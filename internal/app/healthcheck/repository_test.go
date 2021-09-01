package healthcheck

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/internal/dependency/healthcheck"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/linux"
	"github.com/romberli/go-util/middleware/clickhouse"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/go-util/middleware/prometheus"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
)

const (
	// modify the connection information
	defaultDASAddr   = "192.168.10.219:3306"
	defaultDASDBName = "das"
	defaultDBUser    = "root"
	defaultDBPass    = "root"

	defaultPrometheusUser = "admin"
	defaultPrometheusPass = "admin"

	defaultQueryDBName = "pmm"

	defaultOperationID      = 1
	defaultMysqlServerID    = 1
	newResultStatus         = 1
	accuracyReviewStruct    = "accuracyReview"
	newResultAccuracyReview = 1

	defaultVariableName  = "datadir"
	defaultVariableValue = "/mysqldata/mysql3306/data"

	defaultFileSystemNum     = 3
	defaultPrometheusDataNum = 61
)

var (
	operationInfo        = initOperationInfo()
	dasRepo              = initDASRepo()
	applicationMySQLRepo = initApplicationMySQLRepo()
	prometheusRepo       = initPrometheusRepo()
	queryRepo            = initQueryRepo()
	mountPoints, devices = initFileSystems()
)

func initGlobalMySQLPool() error {
	dbAddr := defaultDASAddr
	dbName := defaultDASDBName
	dbUser := defaultDBUser
	dbPass := defaultDBPass
	maxConnections := mysql.DefaultMaxConnections
	initConnections := mysql.DefaultInitConnections
	maxIdleConnections := mysql.DefaultMaxIdleConnections
	maxIdleTime := mysql.DefaultMaxIdleTime
	keepAliveInterval := mysql.DefaultKeepAliveInterval

	config := mysql.NewConfig(dbAddr, dbName, dbUser, dbPass)
	poolConfig := mysql.NewPoolConfigWithConfig(config, maxConnections, initConnections, maxIdleConnections, maxIdleTime, keepAliveInterval)
	log.Debugf("pool config: %v", poolConfig)
	var err error
	global.DASMySQLPool, err = mysql.NewPoolWithPoolConfig(poolConfig)

	return err
}

func initDASRepo() *DASRepo {
	return NewDASRepoWithGlobal()
}

func initOperationInfo() *OperationInfo {
	err := initGlobalMySQLPool()
	if err != nil {
		log.Error(common.CombineMessageWithError("initOperationInfo() failed", err))
		os.Exit(1)
	}

	mysqlServerService := metadata.NewMySQLServerServiceWithDefault()
	err = mysqlServerService.GetByID(defaultMysqlServerID)
	if err != nil {
		log.Error(common.CombineMessageWithError("initOperationInfo() failed", err))
		os.Exit(1)
	}
	mysqlServer := mysqlServerService.GetMySQLServers()[constant.ZeroInt]
	monitorSystem, err := mysqlServer.GetMonitorSystem()
	if err != nil {
		log.Error(common.CombineMessageWithError("initOperationInfo() failed", err))
		os.Exit(1)
	}

	return &OperationInfo{
		operationID:   defaultOperationID,
		mysqlServer:   mysqlServer,
		monitorSystem: monitorSystem,
		startTime:     time.Now().Add(-constant.Week),
		endTime:       time.Now(),
		step:          defaultStep,
	}
}

func initApplicationMySQLRepo() *ApplicationMySQLRepo {
	conn, err := mysql.NewConn(defaultDASAddr, constant.EmptyString, defaultDBUser, defaultDBPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initApplicationMySQLRepo() failed", err))
		os.Exit(1)
	}

	return NewApplicationMySQLRepo(operationInfo, conn)
}

func initPrometheusRepo() *PrometheusRepo {
	var config prometheus.Config

	addr := fmt.Sprintf("%s:%s/%s", operationInfo.GetMonitorSystem().GetHostIP(),
		operationInfo.GetMonitorSystem().GetPortNum(), operationInfo.GetMonitorSystem().GetBaseURL())
	switch operationInfo.GetMonitorSystem().GetSystemType() {
	case 1:
		config = prometheus.NewConfig(addr, prometheus.DefaultRoundTripper)
	case 2:
		config = prometheus.NewConfigWithBasicAuth(addr, defaultPrometheusUser, defaultPrometheusPass)
	}

	conn, err := prometheus.NewConnWithConfig(config)
	if err != nil {
		log.Error(common.CombineMessageWithError("initPrometheusRepo() failed", err))
		os.Exit(1)
	}

	return NewPrometheusRepo(operationInfo, conn)
}

func initQueryRepo() healthcheck.QueryRepo {
	var queryRepo healthcheck.QueryRepo

	addr := fmt.Sprintf("%s:%s", operationInfo.GetMonitorSystem().GetHostIP(), operationInfo.GetMonitorSystem().GetPortNumSlow())
	switch operationInfo.GetMonitorSystem().GetSystemType() {
	case 1:
		conn, err := mysql.NewConn(addr, defaultQueryDBName, defaultDBUser, defaultDBPass)
		if err != nil {
			log.Error(common.CombineMessageWithError("initQueryRepo() failed", err))
			os.Exit(1)
		}
		queryRepo = NewMySQLQueryRepo(operationInfo, conn)
	case 2:
		conn, err := clickhouse.NewConnWithDefault(addr, defaultQueryDBName, constant.EmptyString, constant.EmptyString)
		if err != nil {
			log.Error(common.CombineMessageWithError("initQueryRepo() failed", err))
			os.Exit(1)
		}
		queryRepo = NewClickhouseQueryRepo(operationInfo, conn)
	}

	return queryRepo
}

func initFileSystems() ([]string, []string) {
	var (
		systemMountPoints []string
		mysqlMountPoints  []string
		devices           []string
	)

	// get file systems
	fileSystems, err := prometheusRepo.GetFileSystems()
	if err != nil {
		log.Error(common.CombineMessageWithError("initFileSystems() failed", err))
		os.Exit(1)
	}
	// get mount points
	for _, fileSystem := range fileSystems {
		systemMountPoints = append(systemMountPoints, fileSystem.GetMountPoint())
	}
	// get mysql directories
	dirs, err := applicationMySQLRepo.GetMySQLDirs()
	if err != nil {
		log.Error(common.CombineMessageWithError("initFileSystems() failed", err))
		os.Exit(1)
	}

	dirs = append(dirs, constant.RootDir)
	// get mysql mount points and devices
	for _, dir := range dirs {
		mountPoint, err := linux.MatchMountPoint(dir, systemMountPoints)
		if err != nil {
			log.Error(common.CombineMessageWithError("initFileSystems() failed", err))
			os.Exit(1)
		}

		mysqlMountPoints = append(mysqlMountPoints, mountPoint)

		for _, fileSystem := range fileSystems {
			if mountPoint == fileSystem.GetMountPoint() {
				devices = append(devices, fileSystem.GetDevice())
			}
		}
	}

	return mysqlMountPoints, devices
}

func createResult() error {
	hcInfo := NewResultWithDefault(defaultResultOperationID, defaultResultWeightedAverageScore,
		defaultResultDBConfigScore, defaultResultCPUUsageScore, defaultResultIOUtilScore,
		defaultResultDiskCapacityUsageScore, defaultResultConnectionUsageScore, defaultResultAverageActiveSessionPercentsScore,
		defaultResultCacheMissRatioScore, defaultResultTableRowsScore, defaultResultTableSizeScore,
		defaultResultSlowQueryScore, defaultResultAccuracyReview)
	err := dasRepo.SaveResult(hcInfo)

	return err
}

func deleteResultByID(id int) error {
	sql := `delete from t_hc_result where id = ?`
	_, err := dasRepo.Execute(sql, id)
	return err
}

func deleteOperationInfoByID(id int) error {
	sql := `delete from t_hc_operation_info where id = ?`
	_, err := dasRepo.Execute(sql, id)
	return err
}

func TestRepositoryAll(t *testing.T) {
	// das repository
	TestDASRepo_Execute(t)
	TestDASRepo_GetResultByOperationID(t)
	TestDASRepo_IsRunning(t)
	TestDASRepo_InitOperation(t)
	TestDASRepo_UpdateOperationStatus(t)
	TestDASRepo_SaveResult(t)
	TestDASRepo_UpdateAccuracyReviewByOperationID(t)
	// application mysql repository
	TestApplicationMySQLRepo_GetVariables(t)
	TestApplicationMySQLRepo_GetMySQLDirs(t)
	TestApplicationMySQLRepo_GetLargeTables(t)
	// prometheus repository
	TestPrometheusRepo_GetFileSystems(t)
	TestPrometheusRepo_GetCPUUsage(t)
	TestPrometheusRepo_GetIOUtil(t)
	TestPrometheusRepo_GetDiskCapacityUsage(t)
	TestPrometheusRepo_GetConnectionUsage(t)
	TestPrometheusRepo_GetAverageActiveSessionPercents(t)
	TestPrometheusRepo_GetCacheMissRatio(t)
	TestPrometheusRepo_getServiceName(t)
	TestPrometheusRepo_getPMMVersion(t)
	TestPrometheusRepo_execute(t)
	// mysql query repository
	TestMySQLQueryRepo_GetSlowQuery(t)
	TestMySQLQueryRepo_getServiceName(t)
	TestMySQLQueryRepo_getPMMVersion(t)
	// clickhouse query repository
	TestClickhouseQueryRepo_GetSlowQuery(t)
	TestClickhouseQueryRepo_getServiceName(t)
	TestClickhouseQueryRepo_getPMMVersion(t)
}

func TestDASRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := "select 1;"
	result, err := dasRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, r, "test Execute() failed")
}

func TestDASRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_hc_result(operation_id, weighted_average_score, db_config_score, db_config_data,
		db_config_advice, cpu_usage_score, cpu_usage_data, cpu_usage_high, io_util_score,
		io_util_data, io_util_high, disk_capacity_usage_score, disk_capacity_usage_data,
		disk_capacity_usage_high, connection_usage_score, connection_usage_data,
		connection_usage_high, average_active_session_percents_score, average_active_session_percents_data,
		average_active_session_percents_high, cache_miss_ratio_score, cache_miss_ratio_data,
		cache_miss_ratio_high, table_size_score, table_size_data, table_size_high, slow_query_score,
		slow_query_data, slow_query_advice, accuracy_review) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
	?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	tx, err := dasRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql, defaultResultOperationID, defaultResultWeightedAverageScore, defaultResultDBConfigScore,
		defaultResultDBConfigData, defaultResultDBConfigAdvice, defaultResultCPUUsageScore, defaultResultCPUUsageData,
		defaultResultCPUUsageHigh, defaultResultIOUtilScore, defaultResultIOUtilData, defaultResultIOUtilHigh,
		defaultResultDiskCapacityUsageScore, defaultResultDiskCapacityUsageData, defaultResultDiskCapacityUsageHigh,
		defaultResultConnectionUsageScore, defaultResultConnectionUsageData, defaultResultConnectionUsageHigh,
		defaultResultAverageActiveSessionPercentsScore, defaultResultAverageActiveSessionPercentsData, defaultResultAverageActiveSessionPercentsHigh,
		defaultResultCacheMissRatioScore, defaultResultCacheMissRatioData, defaultResultCacheMissRatioHigh,
		defaultResultTableSizeScore, defaultResultTableSizeData, defaultResultTableSizeHigh, defaultResultSlowQueryScore,
		defaultResultSlowQueryData, defaultResultSlowQueryAdvice, defaultResultAccuracyReview)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select operation_id from t_hc_result where operation_id = ?`
	result, err := tx.Execute(sql, defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	operationID, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if operationID != defaultResultOperationID {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entity, err := dasRepo.GetResultByOperationID(defaultResultOperationID)
	if entity != nil {
		asst.Fail("test Transaction() failed")
	}
}

func TestDASRepo_GetResultByOperationID(t *testing.T) {
	asst := assert.New(t)

	err := createResult()
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
	result, err := dasRepo.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
	operationID := result.GetOperationID()
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
	asst.Equal(defaultResultOperationID, operationID, "test GetResultByOperationID() failed")
	// delete
	err = deleteResultByID(result.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
}

func TestDASRepo_IsRunning(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_hc_operation_info(mysql_server_id, start_time, end_time, step) values(?, ?, ?, ?);`
	_, err := dasRepo.Execute(sql, defaultMysqlServerID, time.Now().Add(-constant.Week), time.Now(), defaultStep)
	asst.Nil(err, common.CombineMessageWithError("test IsRunning() failed", err))
	result, err := dasRepo.IsRunning(defaultMysqlServerID)
	asst.Nil(err, common.CombineMessageWithError("test IsRunning() failed", err))
	asst.False(result, "test IsRunning() failed")
	// delete
	sql = `select id from t_hc_operation_info order by id desc limit 0,1`
	resultID, err := dasRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test IsRunning() failed", err))
	id, err := resultID.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test IsRunning() failed", err))
	err = deleteOperationInfoByID(id)
	asst.Nil(err, common.CombineMessageWithError("test IsRunning() failed", err))
}

func TestDASRepo_InitOperation(t *testing.T) {
	asst := assert.New(t)

	id, err := dasRepo.InitOperation(defaultMysqlServerID, time.Now().Add(-constant.Week), time.Now(), defaultStep)
	asst.Nil(err, common.CombineMessageWithError("test InitOperation() failed", err))
	sql := `select mysql_server_id from t_hc_operation_info where id = ?;`
	result, err := dasRepo.Execute(sql, id)
	asst.Nil(err, common.CombineMessageWithError("test InitOperation() failed", err))
	mysqlServerID, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test InitOperation() failed", err))
	asst.Equal(defaultMysqlServerID, mysqlServerID, "test InitOperation() failed")
	// delete
	err = deleteOperationInfoByID(id)
	asst.Nil(err, common.CombineMessageWithError("test InitOperation() failed", err))
}

func TestDASRepo_UpdateOperationStatus(t *testing.T) {
	asst := assert.New(t)

	id, err := dasRepo.InitOperation(defaultMysqlServerID, time.Now().Add(-constant.Week), time.Now(), defaultStep)
	asst.Nil(err, common.CombineMessageWithError("test UpdateOperationStatus() failed", err))
	err = dasRepo.UpdateOperationStatus(id, newResultStatus, "")
	asst.Nil(err, common.CombineMessageWithError("test UpdateOperationStatus() failed", err))
	sql := `select status from t_hc_operation_info where id = ?;`
	result, err := dasRepo.Execute(sql, id)
	asst.Nil(err, common.CombineMessageWithError("test UpdateOperationStatus() failed", err))
	status, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test UpdateOperationStatus() failed", err))
	asst.Equal(newResultStatus, status, "test UpdateOperationStatus() failed")
	// delete
	err = deleteOperationInfoByID(id)
	asst.Nil(err, common.CombineMessageWithError("test UpdateOperationStatus() failed", err))
}

func TestDASRepo_SaveResult(t *testing.T) {
	asst := assert.New(t)

	err := createResult()
	asst.Nil(err, common.CombineMessageWithError("test SaveResult() failed", err))
	result, err := dasRepo.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test SaveResult() failed", err))
	asst.Equal(defaultResultOperationID, result.GetOperationID(), "test SaveResult() failed")
	// delete
	err = deleteResultByID(result.Identity())
	asst.Nil(err, common.CombineMessageWithError("test SaveResult() failed", err))
}

func TestDASRepo_UpdateAccuracyReviewByOperationID(t *testing.T) {
	asst := assert.New(t)

	err := createResult()
	asst.Nil(err, common.CombineMessageWithError("test UpdateAccuracyReviewByOperationID() failed", err))
	result, err := dasRepo.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test UpdateAccuracyReviewByOperationID() failed", err))
	err = result.Set(map[string]interface{}{accuracyReviewStruct: newResultAccuracyReview})
	asst.Nil(err, common.CombineMessageWithError("test UpdateAccuracyReviewByOperationID() failed", err))
	err = dasRepo.UpdateAccuracyReviewByOperationID(result.GetOperationID(), newResultAccuracyReview)
	asst.Nil(err, common.CombineMessageWithError("test UpdateAccuracyReviewByOperationID() failed", err))
	asst.Equal(newResultAccuracyReview, result.GetAccuracyReview(), "test UpdateAccuracyReviewByOperationID() failed")
	// delete
	err = deleteResultByID(result.Identity())
	asst.Nil(err, common.CombineMessageWithError("test UpdateAccuracyReviewByOperationID() failed", err))
}

func TestApplicationMySQLRepo_GetVariables(t *testing.T) {
	asst := assert.New(t)

	items := []string{defaultVariableName}
	variables, err := applicationMySQLRepo.GetVariables(items)
	asst.Nil(err, common.CombineMessageWithError("test TestApplicationMySQLRepo_GetVariables() failed", err))
	value := variables[constant.ZeroInt].GetValue()
	asst.Equal(defaultVariableValue, value, "test TestApplicationMySQLRepo_GetVariables() failed")
}

func TestApplicationMySQLRepo_GetMySQLDirs(t *testing.T) {
	asst := assert.New(t)

	items := []string{defaultVariableName}
	variables, err := applicationMySQLRepo.GetVariables(items)
	asst.Nil(err, common.CombineMessageWithError("test TestApplicationMySQLRepo_GetMySQLDirs() failed", err))
	value := variables[constant.ZeroInt].GetValue()
	asst.Equal(defaultVariableValue, value, "test TestApplicationMySQLRepo_GetMySQLDirs() failed")
}

func TestApplicationMySQLRepo_GetLargeTables(t *testing.T) {
	asst := assert.New(t)

	tables, err := applicationMySQLRepo.GetLargeTables()
	asst.Nil(err, common.CombineMessageWithError("test TestApplicationMySQLRepo_GetLargeTables() failed", err))
	asst.Equal(constant.ZeroInt, len(tables), "test TestApplicationMySQLRepo_GetLargeTables() failed")
}

func TestPrometheusRepo_GetFileSystems(t *testing.T) {
	asst := assert.New(t)

	fileSystems, err := prometheusRepo.GetFileSystems()
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetFileSystems() failed", err))
	asst.Equal(defaultFileSystemNum, len(fileSystems), "test TestPrometheusRepo_GetFileSystems() failed")
}

func TestPrometheusRepo_GetCPUUsage(t *testing.T) {
	asst := assert.New(t)

	datas, err := prometheusRepo.GetCPUUsage()
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetCPUUsage() failed", err))
	asst.Equal(defaultPrometheusDataNum, len(datas), "test TestPrometheusRepo_GetCPUUsage() failed")
}

func TestPrometheusRepo_GetIOUtil(t *testing.T) {
	asst := assert.New(t)

	datas, err := prometheusRepo.GetIOUtil(devices)
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetIOUtil() failed", err))
	asst.Equal(defaultPrometheusDataNum, len(datas), "test TestPrometheusRepo_GetIOUtil() failed")
}

func TestPrometheusRepo_GetDiskCapacityUsage(t *testing.T) {
	asst := assert.New(t)

	datas, err := prometheusRepo.GetDiskCapacityUsage(mountPoints)
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetDiskCapacityUsage() failed", err))
	asst.Equal(defaultPrometheusDataNum, len(datas), "test TestPrometheusRepo_GetDiskCapacityUsage() failed")
}

func TestPrometheusRepo_GetConnectionUsage(t *testing.T) {
	asst := assert.New(t)

	datas, err := prometheusRepo.GetConnectionUsage()
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetConnectionUsage() failed", err))
	asst.Equal(defaultPrometheusDataNum, len(datas), "test TestPrometheusRepo_GetConnectionUsage() failed")
}

func TestPrometheusRepo_GetAverageActiveSessionPercents(t *testing.T) {
	asst := assert.New(t)

	datas, err := prometheusRepo.GetAverageActiveSessionPercents()
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetAverageActiveSessionPercents() failed", err))
	asst.Equal(defaultPrometheusDataNum, len(datas), "test TestPrometheusRepo_GetAverageActiveSessionPercents() failed")
}

func TestPrometheusRepo_GetCacheMissRatio(t *testing.T) {
	asst := assert.New(t)

	datas, err := prometheusRepo.GetCacheMissRatio()
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetCacheMissRatio() failed", err))
	asst.Equal(defaultPrometheusDataNum, len(datas), "test TestPrometheusRepo_GetCacheMissRatio() failed")
}

func TestPrometheusRepo_getServiceName(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(operationInfo.GetMySQLServer().GetServiceName(), prometheusRepo.getServiceName(), "test TestPrometheusRepo_getServiceName() failed")
}

func TestPrometheusRepo_getPMMVersion(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(operationInfo.GetMonitorSystem().GetSystemType(), prometheusRepo.getPMMVersion(), "test TestPrometheusRepo_getPMMVersion() failed")
}

func TestPrometheusRepo_execute(t *testing.T) {
	asst := assert.New(t)

	datas, err := prometheusRepo.execute(PrometheusCPUUsageV2)
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_execute() failed", err))
	asst.Equal(defaultPrometheusDataNum, len(datas), "test TestPrometheusRepo_execute() failed")
}

func TestMySQLQueryRepo_GetSlowQuery(t *testing.T) {
	asst := assert.New(t)

	if operationInfo.GetMonitorSystem().GetSystemType() == 1 {
		queries, err := queryRepo.(*MySQLQueryRepo).GetSlowQuery()
		asst.Nil(err, common.CombineMessageWithError("test TestMySQLQueryRepo_GetSlowQuery() failed", err))
		asst.LessOrEqual(constant.ZeroInt, len(queries), "test TestMySQLQueryRepo_GetSlowQuery() failed")
	}
}

func TestMySQLQueryRepo_getServiceName(t *testing.T) {
	asst := assert.New(t)

	if operationInfo.GetMonitorSystem().GetSystemType() == 1 {
		asst.Equal(operationInfo.GetMySQLServer().GetServiceName(), queryRepo.(*MySQLQueryRepo).getServiceName(), "test TestMySQLQueryRepo_getServiceName() failed")
	}
}

func TestMySQLQueryRepo_getPMMVersion(t *testing.T) {
	asst := assert.New(t)

	if operationInfo.GetMonitorSystem().GetSystemType() == 1 {
		asst.Equal(operationInfo.GetMySQLServer().GetServiceName(), queryRepo.(*MySQLQueryRepo).getPMMVersion(), "test TestMySQLQueryRepo_getPMMVersion() failed")
	}
}

func TestClickhouseQueryRepo_GetSlowQuery(t *testing.T) {
	asst := assert.New(t)

	if operationInfo.GetMonitorSystem().GetSystemType() == 1 {
		queries, err := queryRepo.(*ClickhouseQueryRepo).GetSlowQuery()
		asst.Nil(err, common.CombineMessageWithError("test TestClickhouseQueryRepo_GetSlowQuery() failed", err))
		asst.LessOrEqual(constant.ZeroInt, len(queries), "test TestClickhouseQueryRepo_GetSlowQuery() failed")
	}
}

func TestClickhouseQueryRepo_getServiceName(t *testing.T) {
	asst := assert.New(t)

	if operationInfo.GetMonitorSystem().GetSystemType() == 1 {
		asst.Equal(operationInfo.GetMySQLServer().GetServiceName(), queryRepo.(*ClickhouseQueryRepo).getServiceName(), "test TestClickhouseQueryRepo_getServiceName() failed")
	}
}

func TestClickhouseQueryRepo_getPMMVersion(t *testing.T) {
	asst := assert.New(t)

	if operationInfo.GetMonitorSystem().GetSystemType() == 1 {
		asst.Equal(operationInfo.GetMySQLServer().GetServiceName(), queryRepo.(*ClickhouseQueryRepo).getPMMVersion(), "test TestClickhouseQueryRepo_getServiceName() failed")
	}
}
