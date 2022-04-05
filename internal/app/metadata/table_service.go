package metadata

import (
	"github.com/romberli/das/internal/app/privilege"
	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
)

const tablesStruct = "Tables"

var _ metadata.TableService = (*TableService)(nil)

// TableService implements dependency.TableService
type TableService struct {
	metadata.TableRepo
	Tables          []metadata.Table          `middleware:"tables" json:"tables"`
	TableStatistics []metadata.TableStatistic `middleware:"table_statistics" json:"table_statistics"`
	IndexStatistics []metadata.IndexStatistic `middleware:"index_statistics" json:"index_statistics"`
	CreateStatement string                    `middleware:"create_statement" json:"create_statement"`
}

// newTableService returns metadata.TableService
func NewTableService(repo metadata.TableRepo) metadata.TableService {
	return newTableService(repo)
}

// newTableService returns metadata.TableService
func NewTableServiceWithDefault() metadata.TableService {
	return newTableService(NewTableRepoWithDefault())
}

// newTableService returns *TableService
func newTableService(repo metadata.TableRepo) *TableService {
	return &TableService{
		repo,
		[]metadata.Table{},
		[]metadata.TableStatistic{},
		[]metadata.IndexStatistic{},
		constant.EmptyString,
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

func (ts *TableService) GetByDBID(dbID int, loginName string) error {
	hostIP, portNum, dbName, err := ts.getHostInfoAndDBNameByDBID(dbID)
	if err != nil {
		return err
	}
	err = ts.init(hostIP, portNum, dbName, loginName)
	if err != nil {
		return err
	}

	defer func() { _ = ts.TableRepo.Close() }()
	ts.Tables, err = ts.TableRepo.GetByDBName(dbName)

	return err
}

// GetStatisticsByDBIDAndTableName gets the table statistics by db id and table name
func (ts *TableService) GetStatisticsByDBIDAndTableName(dbID int, tableName, loginName string) error {
	hostIP, portNum, dbName, err := ts.getHostInfoAndDBNameByDBID(dbID)
	if err != nil {
		return err
	}

	return ts.getStatisticsByHostInfoAndDBNameAndTableName(hostIP, portNum, dbName, tableName, loginName)
}

// GetStatisticsByHostInfoAndDBNameAndTableName gets the table statistics by host info and db name and table name
func (ts *TableService) GetStatisticsByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, loginName string) error {
	return ts.getStatisticsByHostInfoAndDBNameAndTableName(hostIP, portNum, dbName, tableName, loginName)
}

// AnalyzeTableByDBIDAndTableName analyzes the table by db id and table name
func (ts *TableService) AnalyzeTableByDBIDAndTableName(dbID int, tableName, loginName string) error {
	hostIP, portNum, dbName, err := ts.getHostInfoAndDBNameByDBID(dbID)
	if err != nil {
		return err
	}

	return ts.analyzeTableByHostInfoAndDBNameAndTableName(hostIP, portNum, dbName, tableName, loginName)
}

// AnalyzeTableByHostInfoAndDBNameAndTableName analyzes the table by host info and db name and table name
func (ts *TableService) AnalyzeTableByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, loginName string) error {
	return ts.analyzeTableByHostInfoAndDBNameAndTableName(hostIP, portNum, dbName, tableName, loginName)
}

// Marshal marshals TableService.Tables to json bytes
func (ts *TableService) Marshal() ([]byte, error) {
	return ts.MarshalWithFields(tablesStruct)
}

// MarshalWithFields marshals only specified fields of the TableService to json bytes
func (ts *TableService) MarshalWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(ts, fields...)
}

// getHostInfoAndDBNameByDBID gets the host info and db name
func (ts *TableService) getHostInfoAndDBNameByDBID(dbID int) (string, int, string, error) {
	ds := NewDBServiceWithDefault()

	err := ds.GetByID(dbID)
	if err != nil {
		return constant.EmptyString, constant.ZeroInt, constant.EmptyString, err
	}
	err = ds.GetMySQLClusterByID(dbID)
	if err != nil {
		return constant.EmptyString, constant.ZeroInt, constant.EmptyString, err
	}

	masterServers, err := ds.GetMySQLCluster().GetMasterServers()
	if err != nil {
		return constant.EmptyString, constant.ZeroInt, constant.EmptyString, err
	}

	return masterServers[constant.ZeroInt].GetHostIP(), masterServers[constant.ZeroInt].GetPortNum(), ds.GetDBs()[constant.ZeroInt].GetDBName(), nil
}

// init checks the privilege and initialize the repository
func (ts *TableService) init(hostIP string, portNum int, dbName, loginName string) error {
	// check privilege
	privilegeService := privilege.NewServiceWithDefault(loginName)
	err := privilegeService.CheckMySQLServerByHostInfo(hostIP, portNum)
	if err != nil {
		return err
	}

	return ts.TableRepo.InitMySQLConn(hostIP, portNum, dbName)
}

// getStatisticsByHostInfoAndDBNameAndTableName gets the table statistics by host info and db name and table name
func (ts *TableService) getStatisticsByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, loginName string) error {
	err := ts.init(hostIP, portNum, dbName, loginName)
	if err != nil {
		return err
	}

	defer func() { _ = ts.TableRepo.Close() }()
	ts.TableStatistics, ts.IndexStatistics, ts.CreateStatement, err = ts.TableRepo.GetStatisticsByDBNameAndTableName(dbName, tableName)

	return err
}

// analyzeTableByHostInfoAndDBNameAndTableName analyzes the table by host info and db name and table name
func (ts *TableService) analyzeTableByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, loginName string) error {
	err := ts.init(hostIP, portNum, dbName, loginName)
	if err != nil {
		return err
	}

	return ts.TableRepo.AnalyzeTableByDBNameAndTableName(dbName, tableName)
}
