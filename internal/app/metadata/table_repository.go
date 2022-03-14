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
			t.table_rows                             AS rows,
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
	tableStatisticInfoList := make([]*TableStatistic, result.RowNumber())
	err = result.MapToStructSlice(tableStatisticInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	tableStatisticList := make([]metadata.TableStatistic, result.RowNumber())
	for i := range tableStatisticList {
		tableStatisticList[i] = tableStatisticInfoList[i]
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
			IF(non_unique = 0, 'true', 'false') AS unique,
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
	indexStatisticInfoList := make([]*IndexStatistic, result.RowNumber())
	err = result.MapToStructSlice(indexStatisticInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	indexStatisticList := make([]metadata.IndexStatistic, result.RowNumber())
	for i := range indexStatisticList {
		indexStatisticList[i] = indexStatisticList[i]
	}

	return indexStatisticList, nil
}

func (tr *TableRepo) GetCreateStatement(tableSchema, tableName string) (string, error) {
	sql := `
		SHOW CREATE TABLE ?.?;
	`
	log.Debugf("metadata TableRepo.GetCreateStatement() sql: \n%s", sql, tableSchema, tableName)
	result, err := tr.Execute(sql, tableSchema, tableName)
	if err != nil {
		return "", err
	}
	fmt.Println(result)

	return "", nil
}

func (tr *TableRepo) AnalyzeTableByDBIDAndTableName(dbID int, tableName, userName string) error {
	panic("implement me")
}

func (tr *TableRepo) AnalyzeTableByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, userName string) error {
	panic("implement me")
}