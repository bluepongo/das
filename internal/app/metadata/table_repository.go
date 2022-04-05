package metadata

import (
	"fmt"

	"github.com/pingcap/errors"
	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/spf13/viper"
)

const (
	msgTypeColumn = "Msg_type"
	msgTextColumn = "Msg_text"
	errorMessage  = "Error"
)

var _ metadata.TableRepo = (*TableRepo)(nil)

// TableRepo implements dependency.TableRepo interface
type TableRepo struct {
	Conn *mysql.Conn
}

// NewTableRepo returns *TableRepo
func NewTableRepo(conn *mysql.Conn) metadata.TableRepo {
	return newTableRepo(conn)
}

// NewTableRepoWithDefault returns *TableRepo with default value
func NewTableRepoWithDefault() metadata.TableRepo {
	return newTableRepo(nil)
}

// newTableRepo returns *TableRepo
func newTableRepo(conn *mysql.Conn) *TableRepo {
	return &TableRepo{conn}
}

// Close closes the mysql connection
func (tr *TableRepo) Close() error {
	if tr.Conn != nil {
		err := tr.Conn.Close()
		if err != nil {
			return err
		}

		tr.Conn = nil
	}

	return nil
}

// InitMySQLConn initialize the mysql connection of the repository
func (tr *TableRepo) InitMySQLConn(hostIP string, portNum int, dbName string) error {
	var err error

	if tr.Conn == nil {
		tr.Conn, err = mysql.NewConn(
			fmt.Sprintf("%s:%d", hostIP, portNum),
			dbName,
			viper.GetString(config.DBApplicationMySQLUserKey),
			viper.GetString(config.DBApplicationMySQLPassKey),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// Execute implements dependency.MySQLServerRepo interface,
// it executes command with arguments on database
func (tr *TableRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	if tr.Conn == nil {
		return nil, errors.New("connection of table repository is nil, please initialize it first")
	}

	return tr.Conn.Execute(command, args...)
}

// GetByDBName gets the table info by db name from middleware
func (tr *TableRepo) GetByDBName(dbName string) ([]metadata.Table, error) {
	sql := `
		SELECT table_schema AS db_name,
			   table_name   AS table_name
		FROM information_schema.tables
		WHERE table_schema = ?;
	`
	log.Debugf("metadata TableRepo.GetByDBName() sql: \n%s\nplaceholders: %s", sql, dbName)
	result, err := tr.Execute(sql, dbName)
	if err != nil {
		return nil, err
	}

	tableList := make([]metadata.Table, result.RowNumber())
	for row := range tableList {
		tableList[row] = NewEmptyTableInfo()
	}
	// map to struct
	err = result.MapToStructSlice(tableList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	return tableList, nil
}

// GetTableStatistics gets table statistics from the middleware
func (tr *TableRepo) GetTableStatistics(dbName, tableName string) ([]metadata.TableStatistic, error) {
	sql := `
		SELECT t.table_schema                        AS db_name,
			t.table_name                             AS table_name,
			t.table_rows                             AS table_rows,
			t.data_length                            AS size,
			TRUNCATE(t.data_length / 1024 / 1024, 3) AS size_mb,
			t.avg_row_length                         AS avg_row_length,
			t.auto_increment                         AS auto_increment,
			t.engine                                 AS engine,
			ccsa.character_set_name                  AS char_set,
			t.table_collation                        AS collation,
			t.create_time                            AS create_time
		FROM information_schema.tables t
		INNER JOIN information_schema.collation_character_set_applicability ccsa
			ON t.table_collation = ccsa.collation_name
		WHERE table_schema = ? AND table_name = ? ;
	`
	log.Debugf("metadata TableRepo.GetTableStatistics() sql: \n%s", sql, dbName, tableName)

	result, err := tr.Execute(sql, dbName, tableName)
	if err != nil {
		return nil, err
	}
	tableStatisticList := make([]metadata.TableStatistic, result.RowNumber())
	for i := range tableStatisticList {
		tableStatisticList[i] = NewEmptyTableStatistic()
	}
	err = result.MapToStructSlice(tableStatisticList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return tableStatisticList, nil
}

// GetIndexStatistics gets index statistics from the middleware
func (tr *TableRepo) GetIndexStatistics(dbName, tableName string) ([]metadata.IndexStatistic, error) {
	sql := `
		SELECT table_schema                     AS db_name,
			table_name                          AS table_name,
			index_name                          AS index_name,
			seq_in_index                        AS sequence,
			column_name                         AS column_name,
			cardinality                         AS cardinality,
			IF(non_unique = 0, 'true', 'false') AS unique,
			IF(nullable = '', 'false', 'true')  AS nullable
		FROM information_schema.statistics
		WHERE table_schema = ?
  		AND table_name = ? ;
	`
	log.Debugf("metadata TableRepo.GetIndexStatistics() sql: \n%s", sql, dbName, tableName)

	result, err := tr.Execute(sql, dbName, tableName)
	if err != nil {
		return nil, err
	}
	indexStatisticList := make([]metadata.IndexStatistic, result.RowNumber())
	for i := range indexStatisticList {
		indexStatisticList[i] = NewEmptyIndexStatistic()
	}
	err = result.MapToStructSlice(indexStatisticList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return indexStatisticList, nil
}

// GetCreateStatement gets the create statement of the table
func (tr *TableRepo) GetCreateStatement(dbName, tableName string) (string, error) {
	sql := fmt.Sprintf(`
		SHOW CREATE TABLE %s.%s;
	`, dbName, tableName)
	log.Debugf("metadata TableRepo.GetCreateStatement() sql: \n%s", sql, dbName, tableName)
	result, err := tr.Execute(sql)
	if err != nil {
		return "", err
	}
	createStatement, err := result.GetStringByName(0, "Create Table")
	if err != nil {
		return "", err
	}

	return string(createStatement), nil
}

// GetStatisticsByDBNameAndTableName gets the full table info by db name and table name from middleware
func (tr *TableRepo) GetStatisticsByDBNameAndTableName(dbName string, tableName string) ([]metadata.TableStatistic, []metadata.IndexStatistic, string, error) {
	tableStatistics, err := tr.GetTableStatistics(dbName, tableName)
	if err != nil {
		return nil, nil, "", err
	}

	indexStatistics, err := tr.GetIndexStatistics(dbName, tableName)
	if err != nil {
		return nil, nil, "", err
	}

	createStatement, err := tr.GetCreateStatement(dbName, tableName)
	if err != nil {
		return nil, nil, "", err
	}

	return tableStatistics, indexStatistics, createStatement, nil
}

// AnalyzeTableByDBNameAndTableName analyzes the table by DB name and table name
func (tr *TableRepo) AnalyzeTableByDBNameAndTableName(dbName, tableName string) error {
	sql := fmt.Sprintf("ANALYZE TABLE %s.%s;", dbName, tableName)
	log.Debugf("metadata TableRepo.AnalyzeTableByDBNameAndTableName() sql: \n%s", sql, dbName, tableName)
	result, err := tr.Execute(sql)
	if err != nil {
		return err
	}

	msgType, err := result.GetStringByName(constant.ZeroInt, msgTypeColumn)
	if err != nil {
		return err
	}
	msgText, err := result.GetStringByName(constant.ZeroInt, msgTextColumn)
	if err != nil {
		return err
	}

	if msgType == errorMessage {
		return errors.New(msgText)
	}

	return nil
}
