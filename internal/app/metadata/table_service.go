package metadata

import (
	"github.com/romberli/das/internal/app/privilege"
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

// GetTableStatistics returns the table statistics list
func (ts *TableService) GetTableStatistics() []metadata.TableStatistic {
	return ts.TableStatistics
}

// GetIndexStatistics returns the index statistics list
func (ts *TableService) GetIndexStatistics() []metadata.IndexStatistic {
	return ts.IndexStatistics
}

// GetCreateStatement returns the create statement of table
func (ts *TableService) GetCreateStatement() string {
	return ts.CreateStatement
}

// GetByDBName returns tables info by DB name
func (ts *TableService) GetByDBName(dbName string) error {
	var err error
	ts.Tables, err = ts.TableRepo.GetByDBName(dbName)

	return err
}

// GetStatisticsByHostInfoAndDBNameAndTableName returns the full table info by DB name and table name
func (ts *TableService) GetStatisticsByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, loginName string) error {
	// check privilege
	privilegeService := privilege.NewServiceWithDefault(loginName)
	err := privilegeService.CheckMySQLServerByHostInfo(hostIP, portNum)
	if err != nil {
		return err
	}

	ts.TableStatistics, ts.IndexStatistics, ts.CreateStatement, err = ts.TableRepo.GetStatisticsByDBNameAndTableName(dbName, tableName)
	if err != nil {
		return err
	}

	return nil
}

// AnalyzeTableByHostInfoAndDBNameAndTableName analyzes the table by DBName and TableName
func (ts *TableService) AnalyzeTableByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, loginName string) error {
	// check privilege
	privilegeService := privilege.NewServiceWithDefault(loginName)
	err := privilegeService.CheckMySQLServerByHostInfo(hostIP, portNum)
	if err != nil {
		return err
	}

	return ts.TableRepo.AnalyzeTableByDBNameAndTableName(dbName, tableName)
}

// Marshal marshals TableService.Tables to json bytes
func (ts *TableService) Marshal() ([]byte, error) {
	return ts.MarshalWithFields(tableStruct)
}

// MarshalWithFields marshals only specified fields of the TableService to json bytes
func (ts *TableService) MarshalWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(ts, fields...)
}
