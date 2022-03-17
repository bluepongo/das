package metadata

import (
	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
)

const tableStruct = "Tables"

var _ metadata.TableService = (*TableService)(nil)

type TableService struct {
	metadata.TableRepo
	Tables          []metadata.Table          `middleware:"tables" json:"tables"`
	TableStatistics []metadata.TableStatistic `middleware:"table_statistics" json:"table_statistics"`
	IndexStatistics []metadata.IndexStatistic `middleware:"index_statistics" json:"index_statistics"`
	CreateStatement string                    `middleware:"create_statement" json:"create_statement"`
}

func NewTableService(repo metadata.TableRepo) *TableService {
	return &TableService{
		repo,
		[]metadata.Table{},
		[]metadata.TableStatistic{},
		[]metadata.IndexStatistic{},
		"",
	}
}

// GetTables returns the tables list
func (ts *TableService) GetTables() []metadata.Table {
	return ts.Tables
}

// GetTableByDBIDAndTableName returns the table info by DBID and table name
func (ts *TableService) GetTableByDBIDAndTableName(dbID int, tableName string) error {
	// TODO: complete it
	return nil
}

// GetTableByHostInfoAndDBNameAndTableName returns the table info by host info、 DB name and table name
func (ts *TableService) GetTableByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName string) error {
	// TODO: complete it
	return nil
}

// AnalyzeTableByDBIDAndTableName analyzes the table by DBID and TableName
func (ts *TableService) AnalyzeTableByDBIDAndTableName(dbID int, tableName, accountName string) error {
	// TODO: complete it
	return nil
}

// AnalyzeTableByHostInfoAndDBNameAndTableName analyzes the table by host info、DB name and table name
func (ts *TableService) AnalyzeTableByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, accountName string) error {
	// TODO: complete it
	return nil
}

// Marshal marshals TableService.Tables to json bytes
func (ts *TableService) Marshal() ([]byte, error) {
	return ts.MarshalWithFields(tableStruct)
}

// MarshalWithFields marshals only specified fields of the TableService to json bytes
func (ts *TableService) MarshalWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(ts, fields...)
}
