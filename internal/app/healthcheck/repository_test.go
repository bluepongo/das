package healthcheck

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/internal/dependency/healthcheck"
	depmeta "github.com/romberli/das/internal/dependency/metadata"
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
	testDASMySQLAddr = "192.168.137.11:3306"
	testDASMySQLName = "das"
	testDASMySQLUser = "root"
	testDASMySQLPass = "root"

	testHealthcheckLoginName       = "zhangs"
	testHealthcheckApplicationAddr = "192.168.137.11:3306"
	testHealthcheckApplicationUser = "root"
	testHealthcheckApplicationPass = "root"
	testHealthcheckPrometheusUser  = "admin"
	testHealthcheckPrometheusPass  = "admin"
	testHealthcheckQueryMySQLUser  = "root"
	testHealthcheckQueryMySQLPass  = "root"

	testHealthcheckDBName = "pmm"

	testHealthcheckVariableName  = "datadir"
	testHealthcheckVariableValue = "/data/mysql/mysqld_multi/mysqld3306/data"
	testHealthcheckFileSystemNum = 3
)

var (
	testOperationInfo        *OperationInfo
	testDASRepo              *DASRepo
	testApplicationMySQLRepo *ApplicationMySQLRepo
	testPrometheusRepo       *PrometheusRepo
	testQueryRepo            healthcheck.QueryRepo
	testMountPoints          []string
)

func init() {
	testInitDASMySQLPool()

	testOperationInfo = testInitOperationInfo()
	testDASRepo = newDASRepo(global.DASMySQLPool)
	testApplicationMySQLRepo = testInitApplicationMySQLRepo()
	testPrometheusRepo = testInitPrometheusRepo()
	testQueryRepo = testInitQueryRepo()
	testMountPoints, _ = testInitFileSystems()
}

func testInitOperationInfo() *OperationInfo {
	userService := metadata.NewUserServiceWithDefault()
	err := userService.GetByAccountNameOrEmployeeID(testHealthcheckLoginName)
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitOperationInfo() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
	user := userService.GetUsers()[constant.ZeroInt]
	mysqlServerService := metadata.NewMySQLServerServiceWithDefault()
	err = mysqlServerService.GetByID(testHealthcheckMySQLServerID)
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitOperationInfo() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
	mysqlServer := mysqlServerService.GetMySQLServers()[constant.ZeroInt]
	monitorSystem, err := mysqlServer.GetMonitorSystem()
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitOperationInfo() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
	// get mysql cluster
	mysqlCluster, err := mysqlServer.GetMySQLCluster()
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitOperationInfo() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
	// get dbs
	dbs, err := mysqlCluster.GetDBs()
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitOperationInfo() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
	// get apps
	var apps []depmeta.App
	for _, db := range dbs {
		applications, err := db.GetApps()
		if err != nil {
			log.Error(common.CombineMessageWithError("testInitOperationInfo() failed", err))
			os.Exit(constant.DefaultAbnormalExitCode)
		}
		for _, application := range applications {
			exists, err := common.ElementInSlice(apps, application)
			if err != nil {
				log.Error(common.CombineMessageWithError("testInitOperationInfo() failed", err))
				os.Exit(constant.DefaultAbnormalExitCode)
			}
			if !exists {
				apps = append(apps, applications...)
			}
		}
	}

	return newOperationInfo(testHealthcheckOperationID, user, apps, mysqlServer, monitorSystem, time.Now().Add(-constant.Week), time.Now(), testHealthcheckStep)
}

func testInitDASMySQLPool() {
	var err error

	if global.DASMySQLPool == nil {
		global.DASMySQLPool, err = mysql.NewPoolWithDefault(testDASMySQLAddr, testDASMySQLName, testDASMySQLUser, testDASMySQLPass)
		if err != nil {
			log.Error(common.CombineMessageWithError("testInitDASMySQLPool() failed", err))
			os.Exit(constant.DefaultAbnormalExitCode)
		}
	}
}

func testInitApplicationMySQLRepo() *ApplicationMySQLRepo {
	conn, err := mysql.NewConn(
		testHealthcheckApplicationAddr,
		constant.EmptyString,
		testHealthcheckApplicationUser,
		testHealthcheckApplicationPass,
	)
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitApplicationMySQLRepo() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}

	return newApplicationMySQLRepo(testOperationInfo, conn)
}

func testInitPrometheusRepo() *PrometheusRepo {
	var cfg prometheus.Config

	addr := fmt.Sprintf("%s:%d%s", testOperationInfo.GetMonitorSystem().GetHostIP(),
		testOperationInfo.GetMonitorSystem().GetPortNum(), testOperationInfo.GetMonitorSystem().GetBaseURL())
	switch testOperationInfo.GetMonitorSystem().GetSystemType() {
	case 1:
		cfg = prometheus.NewConfig(addr, prometheus.DefaultRoundTripper)
	case 2:
		cfg = prometheus.NewConfigWithBasicAuth(addr, testHealthcheckPrometheusUser, testHealthcheckPrometheusPass)
	}

	conn, err := prometheus.NewConnWithConfig(cfg)
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitPrometheusRepo() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}

	return NewPrometheusRepo(testOperationInfo, conn)
}

func testInitQueryRepo() healthcheck.QueryRepo {
	var queryRepo healthcheck.QueryRepo

	addr := fmt.Sprintf("%s:%d", testOperationInfo.GetMonitorSystem().GetHostIP(), testOperationInfo.GetMonitorSystem().GetPortNumSlow())
	switch testOperationInfo.GetMonitorSystem().GetSystemType() {
	case 1:
		conn, err := mysql.NewConn(addr, testHealthcheckDBName, testHealthcheckQueryMySQLUser, testHealthcheckQueryMySQLPass)
		if err != nil {
			log.Error(common.CombineMessageWithError("testInitQueryRepo() failed", err))
			os.Exit(constant.DefaultAbnormalExitCode)
		}
		queryRepo = NewMySQLQueryRepo(testOperationInfo, conn)
	case 2:
		conn, err := clickhouse.NewConnWithDefault(addr, testHealthcheckDBName, constant.EmptyString, constant.EmptyString)
		if err != nil {
			log.Error(common.CombineMessageWithError("testInitQueryRepo() failed", err))
			os.Exit(constant.DefaultAbnormalExitCode)
		}
		queryRepo = NewClickhouseQueryRepo(testOperationInfo, conn)
	}

	return queryRepo
}

func testInitFileSystems() ([]string, []string) {
	var (
		systemMountPoints []string
		mysqlMountPoints  []string
		devices           []string
	)

	// get file systems
	fileSystems, err := testPrometheusRepo.GetFileSystems()
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitFileSystems() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
	// get mount points
	for _, fileSystem := range fileSystems {
		systemMountPoints = append(systemMountPoints, fileSystem.GetMountPoint())
	}
	// get mysql directories
	dirs, err := testApplicationMySQLRepo.GetMySQLDirs()
	if err != nil {
		log.Error(common.CombineMessageWithError("testInitFileSystems() failed", err))
		os.Exit(constant.DefaultAbnormalExitCode)
	}

	dirs = append(dirs, constant.RootDir)
	// get mysql mount points and devices
	for _, dir := range dirs {
		mountPoint, err := linux.MatchMountPoint(dir, systemMountPoints)
		if err != nil {
			log.Error(common.CombineMessageWithError("testInitFileSystems() failed", err))
			os.Exit(constant.DefaultAbnormalExitCode)
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

func testDeleteOperationInfoByID(id int) error {
	sql := `delete from t_hc_operation_history where id = ?`
	_, err := testDASRepo.Execute(sql, id)

	return err
}

func testDeleteResultByID(id int) error {
	sql := `delete from t_hc_result where id = ?`
	_, err := testDASRepo.Execute(sql, id)

	return err
}

func TestRepository_All(t *testing.T) {
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
	TestPrometheusRepo_GetAvgBackupFailedRatio(t)
	TestPrometheusRepo_GetStatisticFailedRatio(t)
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
	result, err := testDASRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
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
	tx, err := testDASRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(
		sql,
		testResultNewOperationID,
		testResultWeightedAverageScore,
		testResultDBConfigScore,
		testResultDBConfigData,
		testResultDBConfigAdvice,
		testResultAvgBackupFailedRatioScore,
		testResultAvgBackupFailedRatioData,
		testResultAvgBackupFailedRatioHigh,
		testResultStatisticFailedRatioScore,
		testResultStatisticFailedRatioData,
		testResultStatisticFailedRatioHigh,
		testResultCPUUsageScore,
		testResultCPUUsageData,
		testResultCPUUsageHigh,
		testResultIOUtilScore,
		testResultIOUtilData,
		testResultIOUtilHigh,
		testResultDiskCapacityUsageScore,
		testResultDiskCapacityUsageData,
		testResultDiskCapacityUsageHigh,
		testResultConnectionUsageScore,
		testResultConnectionUsageData,
		testResultConnectionUsageHigh,
		testResultAverageActiveSessionPercentsScore,
		testResultAverageActiveSessionPercentsData,
		testResultAverageActiveSessionPercentsHigh,
		testResultCacheMissRatioScore,
		testResultCacheMissRatioData,
		testResultCacheMissRatioHigh,
		testResultTableSizeScore,
		testResultTableSizeData,
		testResultTableSizeHigh,
		testResultSlowQueryScore,
		testResultSlowQueryData,
		testResultSlowQueryAdvice,
		testResultAccuracyReview,
	)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select operation_id from t_hc_result where operation_id = ?`
	result, err := tx.Execute(sql, testResultNewOperationID)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	operationID, err := result.GetInt(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if operationID != testResultNewOperationID {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	_, err = testDASRepo.GetResultByOperationID(testResultNewOperationID)
	asst.NotNil(err, common.CombineMessageWithError("test Transaction() failed", err))
}

func TestDASRepo_GetResultByOperationID(t *testing.T) {
	asst := assert.New(t)

	err := testDASRepo.SaveResult(testResult)
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
	result, err := testDASRepo.GetResultByOperationID(testResult.GetOperationID())
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
	asst.Equal(testResultOperationID, result.GetOperationID())
	// delete
	err = testDeleteResultByID(result.Identity())
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
}

func TestDASRepo_IsRunning(t *testing.T) {
	asst := assert.New(t)

	id, err := testDASRepo.InitOperation(
		testOperationInfo.GetUser().Identity(),
		testHealthcheckMySQLServerID,
		time.Now().Add(-constant.Week),
		time.Now(),
		testHealthcheckStep,
	)
	asst.Nil(err, common.CombineMessageWithError("test IsRunning() failed", err))
	isRunning, err := testDASRepo.IsRunning(testHealthcheckMySQLServerID)
	asst.Nil(err, common.CombineMessageWithError("test IsRunning() failed", err))
	asst.False(isRunning, "test IsRunning() failed")
	// delete
	err = testDeleteOperationInfoByID(id)
	asst.Nil(err, common.CombineMessageWithError("test IsRunning() failed", err))
}

func TestDASRepo_InitOperation(t *testing.T) {
	asst := assert.New(t)

	id, err := testDASRepo.InitOperation(
		testOperationInfo.GetUser().Identity(),
		testHealthcheckMySQLServerID,
		time.Now().Add(-constant.Week),
		time.Now(),
		testHealthcheckStep,
	)
	asst.Nil(err, common.CombineMessageWithError("test InitOperation() failed", err))
	// delete
	err = testDeleteOperationInfoByID(id)
	asst.Nil(err, common.CombineMessageWithError("test InitOperation() failed", err))
}

func TestDASRepo_UpdateOperationStatus(t *testing.T) {
	asst := assert.New(t)

	id, err := testDASRepo.InitOperation(
		testOperationInfo.GetUser().Identity(),
		testHealthcheckMySQLServerID,
		time.Now().Add(-constant.Week),
		time.Now(),
		testHealthcheckStep,
	)
	asst.Nil(err, common.CombineMessageWithError("test UpdateOperationStatus() failed", err))
	err = testDASRepo.UpdateOperationStatus(id, testHealthcheckResultUpdateStatus, constant.EmptyString)
	asst.Nil(err, common.CombineMessageWithError("test UpdateOperationStatus() failed", err))
	sql := `select status from t_hc_operation_history where id = ?;`
	r, err := testDASRepo.Execute(sql, id)
	asst.Nil(err, common.CombineMessageWithError("test UpdateOperationStatus() failed", err))
	status, err := r.GetInt(constant.ZeroInt, constant.ZeroInt)
	asst.Nil(err, common.CombineMessageWithError("test UpdateOperationStatus() failed", err))
	asst.Equal(testHealthcheckResultUpdateStatus, status, "test UpdateOperationStatus() failed")
	// delete
	err = testDeleteOperationInfoByID(id)
	asst.Nil(err, common.CombineMessageWithError("test UpdateOperationStatus() failed", err))
}

func TestDASRepo_SaveResult(t *testing.T) {
	asst := assert.New(t)

	err := testDASRepo.SaveResult(testResult)
	asst.Nil(err, common.CombineMessageWithError("test SaveResult() failed", err))
	r, err := testDASRepo.GetResultByOperationID(testResult.GetOperationID())
	asst.Nil(err, common.CombineMessageWithError("test SaveResult() failed", err))
	asst.Equal(testResultOperationID, r.GetOperationID(), "test SaveResult() failed")
	// delete
	err = testDeleteResultByID(r.Identity())
	asst.Nil(err, common.CombineMessageWithError("test SaveResult() failed", err))
}

func TestDASRepo_UpdateAccuracyReviewByOperationID(t *testing.T) {
	asst := assert.New(t)

	err := testDASRepo.SaveResult(testResult)
	asst.Nil(err, common.CombineMessageWithError("test UpdateAccuracyReviewByOperationID() failed", err))
	err = testDASRepo.UpdateAccuracyReviewByOperationID(testResult.GetOperationID(), testHealthcheckResultAccuracyReview)
	asst.Nil(err, common.CombineMessageWithError("test UpdateAccuracyReviewByOperationID() failed", err))
	r, err := testDASRepo.GetResultByOperationID(testResult.GetOperationID())
	asst.Nil(err, common.CombineMessageWithError("test UpdateAccuracyReviewByOperationID() failed", err))
	asst.Equal(testHealthcheckResultAccuracyReview, r.GetAccuracyReview(), "test UpdateAccuracyReviewByOperationID() failed")
	// delete
	err = testDeleteResultByID(r.Identity())
	asst.Nil(err, common.CombineMessageWithError("test UpdateAccuracyReviewByOperationID() failed", err))
}

func TestApplicationMySQLRepo_GetVariables(t *testing.T) {
	asst := assert.New(t)

	variables, err := testApplicationMySQLRepo.GetVariables([]string{testHealthcheckVariableName})
	asst.Nil(err, common.CombineMessageWithError("test TestApplicationMySQLRepo_GetVariables() failed", err))
	asst.Equal(
		strings.TrimRight(testHealthcheckVariableValue, constant.SlashString),
		strings.TrimRight(variables[constant.ZeroInt].GetValue(), constant.SlashString),
		"test TestApplicationMySQLRepo_GetVariables() failed",
	)
}

func TestApplicationMySQLRepo_GetMySQLDirs(t *testing.T) {
	asst := assert.New(t)

	variables, err := testApplicationMySQLRepo.GetVariables([]string{testHealthcheckVariableName})
	asst.Nil(err, common.CombineMessageWithError("test TestApplicationMySQLRepo_GetMySQLDirs() failed", err))
	value := variables[constant.ZeroInt].GetValue()
	asst.Equal(
		strings.TrimRight(testHealthcheckVariableValue, constant.SlashString),
		strings.TrimRight(value, constant.SlashString),
		"test TestApplicationMySQLRepo_GetMySQLDirs() failed",
	)
}

func TestApplicationMySQLRepo_GetLargeTables(t *testing.T) {
	asst := assert.New(t)

	tables, err := testApplicationMySQLRepo.GetLargeTables()
	asst.Nil(err, common.CombineMessageWithError("test TestApplicationMySQLRepo_GetLargeTables() failed", err))
	asst.Equal(constant.ZeroInt, len(tables), "test TestApplicationMySQLRepo_GetLargeTables() failed")
}

func TestPrometheusRepo_GetFileSystems(t *testing.T) {
	asst := assert.New(t)

	fileSystems, err := testPrometheusRepo.GetFileSystems()
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetFileSystems() failed", err))
	asst.Equal(testHealthcheckFileSystemNum, len(fileSystems), "test TestPrometheusRepo_GetFileSystems() failed")
}

func TestPrometheusRepo_GetAvgBackupFailedRatio(t *testing.T) {
	asst := assert.New(t)

	datas, err := testPrometheusRepo.GetAvgBackupFailedRatio()
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetAvgBackupFailedRatio() failed", err))
	asst.GreaterOrEqual(len(datas), constant.ZeroInt, "test TestPrometheusRepo_GetAvgBackupFailedRatio() failed")
}

func TestPrometheusRepo_GetStatisticFailedRatio(t *testing.T) {
	asst := assert.New(t)

	datas, err := testPrometheusRepo.GetStatisticFailedRatio()
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetStatisticFailedRatio() failed", err))
	asst.GreaterOrEqual(len(datas), constant.ZeroInt, "test TestPrometheusRepo_GetStatisticFailedRatio() failed")
}

func TestPrometheusRepo_GetCPUUsage(t *testing.T) {
	asst := assert.New(t)

	datas, err := testPrometheusRepo.GetCPUUsage()
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetCPUUsage() failed", err))
	asst.GreaterOrEqual(len(datas), constant.ZeroInt, "test TestPrometheusRepo_GetCPUUsage() failed")
}

func TestPrometheusRepo_GetIOUtil(t *testing.T) {
	asst := assert.New(t)

	datas, err := testPrometheusRepo.GetIOUtil()
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetIOUtil() failed", err))
	asst.GreaterOrEqual(len(datas), constant.ZeroInt, "test TestPrometheusRepo_GetIOUtil() failed")
}

func TestPrometheusRepo_GetDiskCapacityUsage(t *testing.T) {
	asst := assert.New(t)

	datas, err := testPrometheusRepo.GetDiskCapacityUsage(testMountPoints)
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetDiskCapacityUsage() failed", err))
	asst.GreaterOrEqual(len(datas), constant.ZeroInt, "test TestPrometheusRepo_GetDiskCapacityUsage() failed")
}

func TestPrometheusRepo_GetConnectionUsage(t *testing.T) {
	asst := assert.New(t)

	datas, err := testPrometheusRepo.GetConnectionUsage()
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetConnectionUsage() failed", err))
	asst.GreaterOrEqual(len(datas), constant.ZeroInt, "test TestPrometheusRepo_GetConnectionUsage() failed")
}

func TestPrometheusRepo_GetAverageActiveSessionPercents(t *testing.T) {
	asst := assert.New(t)

	datas, err := testPrometheusRepo.GetAverageActiveSessionPercents()
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetAverageActiveSessionPercents() failed", err))
	asst.GreaterOrEqual(len(datas), constant.ZeroInt, "test TestPrometheusRepo_GetAverageActiveSessionPercents() failed")
}

func TestPrometheusRepo_GetCacheMissRatio(t *testing.T) {
	asst := assert.New(t)

	datas, err := testPrometheusRepo.GetCacheMissRatio()
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_GetCacheMissRatio() failed", err))
	asst.GreaterOrEqual(len(datas), constant.ZeroInt, "test TestPrometheusRepo_GetCacheMissRatio() failed")
}

func TestPrometheusRepo_getServiceName(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testOperationInfo.GetMySQLServer().GetServiceName(), testPrometheusRepo.getServiceName(),
		"test TestPrometheusRepo_getServiceName() failed")
}

func TestPrometheusRepo_getPMMVersion(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testOperationInfo.GetMonitorSystem().GetSystemType(), testPrometheusRepo.getPMMVersion(),
		"test TestPrometheusRepo_getPMMVersion() failed")
}

func TestPrometheusRepo_execute(t *testing.T) {
	asst := assert.New(t)

	var prometheusQuery string

	// prepare query
	switch testPrometheusRepo.getPMMVersion() {
	case 1:
		// pmm 1.x
		prometheusQuery = PrometheusCPUUsageV1
	case 2:
		// pmm 2.x
		prometheusQuery = PrometheusCPUUsageV2
	}

	nodeName := testPrometheusRepo.getNodeName()
	prometheusQuery = fmt.Sprintf(prometheusQuery, nodeName, nodeName, nodeName, nodeName, nodeName, nodeName)

	datas, err := testPrometheusRepo.execute(prometheusQuery)
	asst.Nil(err, common.CombineMessageWithError("test TestPrometheusRepo_execute() failed", err))
	asst.GreaterOrEqual(len(datas), constant.ZeroInt, "test TestPrometheusRepo_execute() failed")
}

func TestMySQLQueryRepo_GetSlowQuery(t *testing.T) {
	asst := assert.New(t)

	if testOperationInfo.GetMonitorSystem().GetSystemType() == 1 {
		queries, err := testQueryRepo.(*MySQLQueryRepo).GetSlowQuery()
		asst.Nil(err, common.CombineMessageWithError("test TestMySQLQueryRepo_GetSlowQuery() failed", err))
		asst.LessOrEqual(constant.ZeroInt, len(queries), "test TestMySQLQueryRepo_GetSlowQuery() failed")
	}
}

func TestMySQLQueryRepo_getServiceName(t *testing.T) {
	asst := assert.New(t)

	if testOperationInfo.GetMonitorSystem().GetSystemType() == 1 {
		asst.Equal(testOperationInfo.GetMySQLServer().GetServiceName(), testQueryRepo.(*MySQLQueryRepo).getServiceName(),
			"test TestMySQLQueryRepo_getServiceName() failed")
	}
}

func TestMySQLQueryRepo_getPMMVersion(t *testing.T) {
	asst := assert.New(t)

	if testOperationInfo.GetMonitorSystem().GetSystemType() == 1 {
		asst.Equal(testOperationInfo.GetMySQLServer().GetServiceName(), testQueryRepo.(*MySQLQueryRepo).getPMMVersion(),
			"test TestMySQLQueryRepo_getPMMVersion() failed")
	}
}

func TestClickhouseQueryRepo_GetSlowQuery(t *testing.T) {
	asst := assert.New(t)

	if testOperationInfo.GetMonitorSystem().GetSystemType() == 2 {
		queries, err := testQueryRepo.(*ClickhouseQueryRepo).GetSlowQuery()
		asst.Nil(err, common.CombineMessageWithError("test TestClickhouseQueryRepo_GetSlowQuery() failed", err))
		asst.LessOrEqual(constant.ZeroInt, len(queries), "test TestClickhouseQueryRepo_GetSlowQuery() failed")
	}
}

func TestClickhouseQueryRepo_getServiceName(t *testing.T) {
	asst := assert.New(t)

	if testOperationInfo.GetMonitorSystem().GetSystemType() == 2 {
		asst.Equal(testOperationInfo.GetMySQLServer().GetServiceName(), testQueryRepo.(*ClickhouseQueryRepo).getServiceName(),
			"test TestClickhouseQueryRepo_getServiceName() failed")
	}
}

func TestClickhouseQueryRepo_getPMMVersion(t *testing.T) {
	asst := assert.New(t)

	if testOperationInfo.GetMonitorSystem().GetSystemType() == 2 {
		asst.Equal(testOperationInfo.GetMonitorSystem().GetSystemType(), testQueryRepo.(*ClickhouseQueryRepo).getPMMVersion(),
			"test TestClickhouseQueryRepo_getServiceName() failed")
	}
}
