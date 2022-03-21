package metadata

import (
	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
)

const tableStruct = "Tables"

var _ metadata.TableService = (*TableService)(nil)

// TableService implements dependency.TableService
type TableService struct {
	metadata.TableRepo
	Tables          []metadata.Table          `middleware:"tables" json:"tables"`
	TableStatistics []metadata.TableStatistic `middleware:"table_statistics" json:"table_statistics"`
	IndexStatistics []metadata.IndexStatistic `middleware:"index_statistics" json:"index_statistics"`
	CreateStatement string                    `middleware:"create_statement" json:"create_statement"`
}

// NewTableService return *TableService
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

// GetByDBName returns tables info by DB name
func (ts *TableService) GetByDBName(dbName string) error {
	var err error
	ts.Tables, err = ts.TableRepo.GetByDBName(dbName)

	return err
}

// GetStatisticsByDBNameAndTableName returns the full table info by DB name and table name
func (ts *TableService) GetStatisticsByDBNameAndTableName(dbName, tableName string) error {
	var err error
	ts.TableStatistics, ts.IndexStatistics, ts.CreateStatement, err = ts.TableRepo.GetStatisticsByDBNameAndTableName(dbName, tableName)
	if err != nil {
		return err
	}

	return nil
}

// AnalyzeTableByDBIDAndTableName analyzes the table by DBID and TableName
func (ts *TableService) AnalyzeTableByDBIDAndTableName(dbID int, tableName, accountName string) error {
	// TODO: service AnalyzeTableByDBIDAndTableName
	return nil
}

// AnalyzeTableByHostInfoAndDBNameAndTableName analyzes the table by host info、DB name and table name
func (ts *TableService) AnalyzeTableByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, accountName string) error {
	// TODO: service AnalyzeTableByHostInfoAndDBNameAndTableName
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
