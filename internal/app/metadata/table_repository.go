package metadata

import (
	"fmt"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
)

var _ metadata.TableRepo = (*TableRepo)(nil)

// TableRepo implements dependency.TableRepo interface
type TableRepo struct {
	conn *mysql.Conn
}

// NewTableRepo returns *TableRepo with mysql
func NewTableRepo(conn *mysql.Conn) *TableRepo {
	return &TableRepo{conn}
}

// Execute executes command with arguments on database
func (tr *TableRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	defer func() {
		err := tr.conn.Close()
		if err != nil {
			log.Errorf("metadata TableRepo.Execute(): close database connection failed.\n%+v", err)
		}
	}()
	return tr.conn.Execute(command, args...)
}

// GetTableStatistics gets table statistics from the middleware
func (tr *TableRepo) GetTableStatistics(tableSchema, tableName string) ([]metadata.TableStatistic, error) {
	sql := `
		SELECT t.table_schema                        AS table_schema,
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
	log.Debugf("metadata TableRepo.GetTableStatistics() sql: \n%s", sql, tableSchema, tableName)

	result, err := tr.Execute(sql, tableSchema, tableName)
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
func (tr *TableRepo) GetIndexStatistics(tableSchema, tableName string) ([]metadata.IndexStatistic, error) {
	sql := `
		SELECT table_schema                     AS table_schema,
			table_name                          AS table_name,
			index_name                          AS index_name,
			seq_in_index                        AS sequence,
			column_name                         AS column_name,
			cardinality                         AS cardinality,
			IF(non_unique = 0, 'true', 'false') AS non_unique,
			IF(nullable = '', 'false', 'true')  AS nullable
		FROM information_schema.statistics
		WHERE table_schema = ?
  		AND table_name = ? ;
	`
	log.Debugf("metadata TableRepo.GetIndexStatistics() sql: \n%s", sql, tableSchema, tableName)

	result, err := tr.Execute(sql, tableSchema, tableName)
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
func (tr *TableRepo) GetCreateStatement(tableSchema, tableName string) (string, error) {
	sql := fmt.Sprintf(`
		SHOW CREATE TABLE %s.%s;
	`, tableSchema, tableName)
	log.Debugf("metadata TableRepo.GetCreateStatement() sql: \n%s", sql, tableSchema, tableName)
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

// GetByDBName gets the tables info by DBname from middleware
func (tr *TableRepo) GetByDBName(dbName string) ([]metadata.Table, error) {
	// TODO: need to be verified
	sql := `
		SELECT t.table_schema                        AS table_schema,
			t.table_name                             AS table_name,
		FROM information_schema.tables t
		INNER JOIN  
			ON t.table_collation = ccsa.collation_name
		WHERE table_schema = ?;
	`
	log.Debugf("metadata TableRepo.GetByDBName() sql: \n%s\nplaceholders: %s", sql, dbName)
	result, err := tr.Execute(sql)
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

// GetStatisticsByDBNameAndTableName gets the full table info by DB name and table name from middleware
func (tr *TableRepo) GetStatisticsByDBNameAndTableName(dbName, tableName string) ([]metadata.TableStatistic, []metadata.IndexStatistic, string, error) {
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

// AnalyzeTableByDBIDAndTableName analyzes the table by DBID and TableName
func (tr *TableRepo) AnalyzeTableByDBIDAndTableName(dbID int, tableName, userName string) error {
	panic("implement me")
}

// AnalyzeTableByHostInfoAndDBNameAndTableName analyzes the table by host info、DB name and table name
func (tr *TableRepo) AnalyzeTableByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, userName string) error {
	panic("implement me")
}
