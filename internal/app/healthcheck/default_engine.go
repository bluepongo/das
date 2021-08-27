package healthcheck

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/internal/app/sqladvisor"
	"github.com/romberli/das/internal/dependency/healthcheck"
	"github.com/romberli/das/pkg/message"
	msghc "github.com/romberli/das/pkg/message/healthcheck"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/linux"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/go-util/middleware/result"
	"github.com/romberli/log"
)

const (
	defaultDBConfigScore                   = 5
	defaultMinScore                        = 0
	defaultMaxScore                        = 100.0
	defaultHundred                         = 100
	defaultDBConfigItemName                = "db_config"
	defaultCPUUsageItemName                = "cpu_usage"
	defaultIOUtilItemName                  = "io_util"
	defaultDiskCapacityUsageItemName       = "disk_capacity_usage"
	defaultConnectionUsageItemName         = "connection_usage"
	defaultAverageActiveSessionNumItemName = "average_active_session_num"
	defaultCacheMissRatioItemName          = "cache_miss_ratio"
	defaultTableRowsItemName               = "table_rows"
	defaultTableSizeItemName               = "table_size"
	defaultSlowQueryRowsExaminedItemName   = "slow_query_rows_examined"
	defaultSlowQueryTopSQLNum              = 3
	defaultClusterType                     = 1
)

var (
	_ healthcheck.Engine = (*DefaultEngine)(nil)
)

// DefaultEngine work for health check module
type DefaultEngine struct {
	operationInfo        *OperationInfo
	engineConfig         DefaultEngineConfig
	result               *Result
	mountPoints          []string
	devices              []string
	dasRepo              healthcheck.DASRepo
	applicationMySQLRepo healthcheck.ApplicationMySQLRepo
	prometheusRepo       healthcheck.PrometheusRepo
	queryRepo            healthcheck.QueryRepo
}

// NewDefaultEngine returns a new *DefaultEngine
func NewDefaultEngine(operationInfo *OperationInfo,
	dasRepo healthcheck.DASRepo,
	applicationMySQLRepo healthcheck.ApplicationMySQLRepo,
	prometheusRepo healthcheck.PrometheusRepo,
	queryRepo healthcheck.QueryRepo) *DefaultEngine {
	return &DefaultEngine{
		operationInfo:        operationInfo,
		engineConfig:         NewEmptyDefaultEngineConfig(),
		result:               NewEmptyResult(),
		dasRepo:              dasRepo,
		applicationMySQLRepo: applicationMySQLRepo,
		prometheusRepo:       prometheusRepo,
		queryRepo:            queryRepo,
	}
}

// getOperationInfo returns the operation information
func (de *DefaultEngine) getOperationInfo() *OperationInfo {
	return de.operationInfo
}

// getEngineConfig returns the default engine config
func (de *DefaultEngine) getEngineConfig() DefaultEngineConfig {
	return de.engineConfig
}

// getResult returns the result
func (de *DefaultEngine) getResult() *Result {
	return de.result
}

// getMountPoints returns the mount points
func (de *DefaultEngine) getMountPoints() []string {
	return de.mountPoints
}

// getDevices returns the disk devices
func (de *DefaultEngine) getDevices() []string {
	return de.devices
}

// getDASRepo returns the das repository
func (de *DefaultEngine) getDASRepo() healthcheck.DASRepo {
	return de.dasRepo
}

// getApplicationMySQLRepo returns the application mysql repository
func (de *DefaultEngine) getApplicationMySQLRepo() healthcheck.ApplicationMySQLRepo {
	return de.applicationMySQLRepo
}

// getPrometheusRepo returns the prometheus repository
func (de *DefaultEngine) getPrometheusRepo() healthcheck.PrometheusRepo {
	return de.prometheusRepo
}

// getQueryRepo returns the query repository
func (de *DefaultEngine) getQueryRepo() healthcheck.QueryRepo {
	return de.queryRepo
}

// getItemConfig returns *DefaultItemConfig with given item name
func (de *DefaultEngine) getItemConfig(item string) healthcheck.ItemConfig {
	return de.engineConfig.GetItemConfig(item)
}

// Run runs healthcheck
func (de *DefaultEngine) Run() {
	defer func() {
		err := de.closeConnections()
		if err != nil {
			log.Error(message.NewMessage(msghc.ErrHealthcheckCloseConnection, err.Error()).Error())
		}
	}()

	// run
	err := de.run()
	if err != nil {
		log.Error(message.NewMessage(msghc.ErrHealthcheckDefaultEngineRun, err.Error()).Error())
		// update status
		updateErr := de.getDASRepo().UpdateOperationStatus(de.operationInfo.operationID, defaultFailedStatus, err.Error())
		if updateErr != nil {
			log.Error(message.NewMessage(msghc.ErrHealthcheckUpdateOperationStatus, updateErr.Error()).Error())
		}
	}

	// update operation status
	msg := fmt.Sprintf("healthcheck completed successfully. engine: default, operation_id: %d", de.operationInfo.operationID)
	updateErr := de.getDASRepo().UpdateOperationStatus(de.operationInfo.operationID, defaultSuccessStatus, msg)
	if updateErr != nil {
		log.Error(message.NewMessage(msghc.ErrHealthcheckUpdateOperationStatus, updateErr.Error()).Error())
	}
}

// run executes the healthcheck
func (de *DefaultEngine) run() error {
	// init MonitorRepo

	// pre run
	err := de.preRun()
	if err != nil {
		return err
	}
	// check db config
	err = de.checkDBConfig()
	if err != nil {
		return err
	}
	// check cpu usage
	err = de.checkCPUUsage()
	if err != nil {
		return err
	}
	// check io util
	err = de.checkIOUtil()
	if err != nil {
		return err
	}
	// check disk capacity usage
	err = de.checkDiskCapacityUsage()
	if err != nil {
		return err
	}
	// check connection usage
	err = de.checkConnectionUsage()
	if err != nil {
		return err
	}
	// check active session number
	err = de.checkActiveSessionNum()
	if err != nil {
		return err
	}
	// check cache miss ratio
	err = de.checkCacheMissRatio()
	if err != nil {
		return err
	}
	// check table size
	err = de.checkTableSize()
	if err != nil {
		return err
	}
	// check slow query
	err = de.checkSlowQuery()
	if err != nil {
		return err
	}
	// summarize
	de.summarize()
	// post run
	return de.postRun()
}

func (de *DefaultEngine) closeConnections() error {
	merr := &multierror.Error{}

	err := de.getApplicationMySQLRepo().Close()
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	err = de.getQueryRepo().Close()
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	return merr.ErrorOrNil()
}

// preRun performs pre-run actions
func (de *DefaultEngine) preRun() error {
	// load engine config
	err := de.loadEngineConfig()
	if err != nil {
		return err
	}
	// get file systems
	fileSystems, err := de.getPrometheusRepo().GetFileSystems(de.getOperationInfo().GetMySQLServer().GetServiceName())
	if err != nil {
		return err
	}
	// get total mount points
	var mountPoints []string
	for _, fileSystem := range fileSystems {
		mountPoints = append(mountPoints, fileSystem.GetMountPoint())
	}
	// get mysql directories
	dirs, err := de.getApplicationMySQLRepo().GetMySQLDirs()
	if err != nil {
		return err
	}
	dirs = append(dirs, constant.RootDir)
	// get mysql mount points and devices
	for _, dir := range dirs {
		mountPoint, err := linux.MatchMountPoint(dir, mountPoints)
		if err != nil {
			return err
		}

		de.mountPoints = append(de.mountPoints, mountPoint)

		for _, fileSystem := range fileSystems {
			if mountPoint == fileSystem.GetMountPoint() {
				de.devices = append(de.devices, fileSystem.GetDevice())
			}
		}
	}
	// init default report host and port
	dbConfigVariableNames[dbConfigReportHost] = de.getOperationInfo().GetMySQLServer().GetHostIP()
	dbConfigVariableNames[dbConfigReportPort] = strconv.Itoa(de.getOperationInfo().GetMySQLServer().GetPortNum())

	return nil
}

// loadEngineConfig loads engine config
func (de *DefaultEngine) loadEngineConfig() error {
	// load config
	sql := `
		select id, item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high,
		score_deduction_per_unit_medium, max_score_deduction_medium, del_flag, create_time, last_update_time
		from t_hc_default_engine_config
		where del_flag = 0;
	`
	log.Debugf("healthcheck DASRepo.loadEngineConfig() sql: \n%s\n", sql)
	result, err := de.getDASRepo().Execute(sql)
	if err != nil {
		return nil
	}
	// init []*DefaultItemConfig
	defaultEngineConfigList := make([]*DefaultItemConfig, result.RowNumber())
	for i := range defaultEngineConfigList {
		defaultEngineConfigList[i] = NewEmptyDefaultItemConfig()
	}
	// map to struct
	err = result.MapToStructSlice(defaultEngineConfigList, constant.DefaultMiddlewareTag)
	if err != nil {
		return err
	}
	defaultEngine := NewEmptyDefaultEngineConfig()
	for _, defaultEngineConfig := range defaultEngineConfigList {
		itemName := defaultEngineConfig.ItemName
		defaultEngine[itemName] = defaultEngineConfig
	}
	// validate config
	err = defaultEngine.Validate()
	if err == nil {
		return message.NewMessage(msghc.ErrDefaultEngineConfigFormatInValid)
	}
	return nil
}

// checkDBConfig checks database configuration
func (de *DefaultEngine) checkDBConfig() error {
	// load database config
	var configItems []string
	for item := range dbConfigVariableNames {
		configItems = append(configItems, item)
	}

	globalVariables, err := de.getApplicationMySQLRepo().GetDBConfig(configItems)
	if err != nil {
		return err
	}

	dbConfigConfig := de.getItemConfig(defaultDBConfigItemName)

	var (
		dbConfigCount int
		variables     []*Variable
	)

	for _, globalVariable := range globalVariables {
		name := globalVariable.GetName()
		value := globalVariable.GetValue()

		switch name {
		// max_user_connection
		case dbConfigMaxUserConnection:
			maxUserConnection, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			if maxUserConnection < dbConfigMaxUserConnectionValid {
				dbConfigCount++
				variables = append(variables, NewVariable(dbConfigMaxUserConnection, value, strconv.Itoa(dbConfigMaxUserConnectionValid)))
			}
			// others
		case dbConfigLogBin, dbConfigBinlogFormat, dbConfigBinlogRowImage, dbConfigSyncBinlog,
			dbConfigInnodbFlushLogAtTrxCommit, dbConfigGTIDMode, dbConfigEnforceGTIDConsistency,
			dbConfigSlaveParallelType, dbConfigSlaveParallelWorkers, dbConfigMasterInfoRepository,
			dbConfigRelayLogInfoRepository, dbConfigReportHost, dbConfigReportPort, dbConfigInnodbFlushMethod,
			dbConfigInnodbMonitorEnable, dbConfigInnodbPrintAllDeadlocks, dbConfigSlowQueryLog, dbConfigPerformanceSchema:
			if strings.ToUpper(value) != dbConfigVariableNames[name] {
				dbConfigCount++
				variables = append(variables, NewVariable(name, value, dbConfigVariableNames[name]))
			}
		}
	}

	// database config data
	jsonBytesTotal, err := json.Marshal(globalVariables)
	if err != nil {
		return nil
	}
	de.result.DBConfigData = string(jsonBytesTotal)
	// database config advice
	jsonBytesVariables, err := json.Marshal(variables)
	if err != nil {
		return nil
	}
	de.result.DBConfigAdvice = string(jsonBytesVariables)
	// database config score deduction
	dbConfigScoreDeduction := float64(dbConfigCount) * dbConfigConfig.GetScoreDeductionPerUnitHigh()
	if dbConfigScoreDeduction > dbConfigConfig.GetMaxScoreDeductionHigh() {
		dbConfigScoreDeduction = dbConfigConfig.GetMaxScoreDeductionHigh()
	}
	de.result.DBConfigScore = int(defaultMaxScore - dbConfigScoreDeduction)
	if de.result.DBConfigScore < constant.ZeroInt {
		de.result.DBConfigScore = constant.ZeroInt
	}

	return nil
}

// checkCPUUsage checks cpu usage
func (de *DefaultEngine) checkCPUUsage() error {
	// get data
	serviceName := de.getOperationInfo().GetMySQLServer().GetServiceName()
	datas, err := de.getPrometheusRepo().GetCPUUsage(serviceName)
	if err != nil {
		return err
	}
	// parse data
	de.result.CPUUsageScore, de.result.CPUUsageData, de.result.CPUUsageHigh, err = de.parsePrometheusDatas(defaultCPUUsageItemName, datas)
	if err != nil {
		return err
	}

	return nil
}

// checkIOUtil check io util
func (de *DefaultEngine) checkIOUtil() error {
	// get data
	serviceName := de.getOperationInfo().GetMySQLServer().GetServiceName()
	datas, err := de.getPrometheusRepo().GetIOUtil(serviceName, de.getDevices())
	if err != nil {
		return err
	}
	// parse data
	de.result.IOUtilScore, de.result.IOUtilData, de.result.IOUtilHigh, err = de.parsePrometheusDatas(defaultIOUtilItemName, datas)
	if err != nil {
		return err
	}

	return nil
}

// checkDiskCapacityUsage checks disk capacity usage
func (de *DefaultEngine) checkDiskCapacityUsage() error {
	// get data
	serviceName := de.getOperationInfo().GetMySQLServer().GetServiceName()
	datas, err := de.getPrometheusRepo().GetDiskCapacityUsage(serviceName, de.getMountPoints())
	if err != nil {
		return err
	}
	// parse data
	de.result.DiskCapacityUsageScore, de.result.DiskCapacityUsageData, de.result.DiskCapacityUsageHigh, err = de.parsePrometheusDatas(defaultDiskCapacityUsageItemName, datas)
	if err != nil {
		return err
	}

	return nil
}

// checkConnectionUsage checks connection usage
func (de *DefaultEngine) checkConnectionUsage() error {
	// get data
	serviceName := de.getOperationInfo().GetMySQLServer().GetServiceName()
	datas, err := de.getPrometheusRepo().GetConnectionUsage(serviceName)
	if err != nil {
		return err
	}
	// parse data
	de.result.ConnectionUsageScore, de.result.ConnectionUsageData, de.result.ConnectionUsageHigh, err = de.parsePrometheusDatas(defaultConnectionUsageItemName, datas)
	if err != nil {
		return err
	}

	return nil
}

// checkActiveSessionNum check active session number
func (de *DefaultEngine) checkActiveSessionNum() error {
	// get data
	serviceName := de.getOperationInfo().GetMySQLServer().GetServiceName()
	datas, err := de.getPrometheusRepo().GetActiveSessionNum(serviceName)
	// parse data
	de.result.AverageActiveSessionNumScore, de.result.AverageActiveSessionNumData, de.result.AverageActiveSessionNumHigh, err = de.parsePrometheusDatas(defaultAverageActiveSessionNumItemName, datas)
	if err != nil {
		return err
	}

	return nil
}

// checkCacheMissRatio checks cache miss ratio
func (de *DefaultEngine) checkCacheMissRatio() error {
	// get data
	serviceName := de.getOperationInfo().GetMySQLServer().GetServiceName()
	datas, err := de.getPrometheusRepo().GetCacheMissRatio(serviceName)
	// parse data
	de.result.CacheMissRatioScore, de.result.CacheMissRatioData, de.result.CacheMissRatioHigh, err = de.parsePrometheusDatas(defaultCacheMissRatioItemName, datas)
	if err != nil {
		return err
	}

	return nil
}

// checkTableSize checks table size by checking rows
func (de *DefaultEngine) checkTableSize() error {
	// check table rows
	// get data
	sql := `
		select TABLE_SCHEMA,TABLE_NAME,TABLE_ROWS,(DATA_LENGTH+INDEX_LENGTH)/1024/1024/1024
		as TABLE_SIZE from TABLES
		where TABLE_TYPE='BASE TABLE';
	`
	log.Debugf("healthcheck DASRepo.checkTableSize() sql: \n%s\n", sql)
	result, err := de.monitorMySQLConn.Execute(sql)
	if err != nil {
		return err
	}

	// analyze result
	length := result.RowNumber()
	if length == constant.ZeroInt {
		return nil
	}

	tableRowsConfig := de.getItemConfig(defaultTableRowsItemName)

	var (
		tableRows            float64
		tableRowsHighSum     float64
		tableRowsHighCount   int
		tableRowsMediumSum   float64
		tableRowsMediumCount int

		tableRowsHigh [][]driver.Value
	)

	for i, rowData := range result.Rows.Values {
		tableRows, err = result.GetFloat(i, constant.ZeroInt)
		if err != nil {
			return err
		}

		switch {
		case tableRows >= tableRowsConfig.GetHighWatermark():
			tableRowsHigh = append(tableRowsHigh, rowData)
			tableRowsHighSum += tableRows
			tableRowsHighCount++
		case tableRows >= tableRowsConfig.GetLowWatermark():
			tableRowsMediumSum += tableRows
			tableRowsMediumCount++
		}
	}

	// table rows data
	jsonBytesTotal, err := json.Marshal(result.Rows.Values)
	if err != nil {
		return nil
	}
	de.result.TableSizeData = string(jsonBytesTotal)
	// table rows high
	jsonBytesHigh, err := json.Marshal(tableRowsHigh)
	if err != nil {
		return nil
	}
	de.result.TableSizeHigh = string(jsonBytesHigh)

	// table rows high score deduction
	tableRowsScoreDeductionHigh := (tableRowsHighSum/float64(tableRowsHighCount) - tableRowsConfig.GetHighWatermark()) / tableRowsConfig.GetUnit() * tableRowsConfig.GetScoreDeductionPerUnitHigh()
	if tableRowsScoreDeductionHigh > tableRowsConfig.GetMaxScoreDeductionHigh() {
		tableRowsScoreDeductionHigh = tableRowsConfig.GetMaxScoreDeductionHigh()
	}
	// table rows medium score deduction
	tableRowsScoreDeductionMedium := (tableRowsMediumSum/float64(tableRowsMediumCount) - tableRowsConfig.GetLowWatermark()) / tableRowsConfig.GetUnit() * tableRowsConfig.GetScoreDeductionPerUnitMedium()
	if tableRowsScoreDeductionMedium > tableRowsConfig.GetMaxScoreDeductionMedium() {
		tableRowsScoreDeductionMedium = tableRowsConfig.GetMaxScoreDeductionMedium()
	}
	// table rows score
	de.result.TableSizeScore = int(defaultMaxScore - tableRowsScoreDeductionHigh - tableRowsScoreDeductionMedium)
	if de.result.TableSizeScore < constant.ZeroInt {
		de.result.TableSizeScore = constant.ZeroInt
	}

	return nil
}

// checkSlowQuery checks slow query
func (de *DefaultEngine) checkSlowQuery() error {
	// check slow query execution time
	var (
		sql    string
		result middleware.Result
		err    error
	)

	serviceName := de.operationInfo.mysqlServer.GetServiceName()
	slowQueryRowsExaminedConfig := de.getItemConfig(defaultSlowQueryRowsExaminedItemName)
	pmmVersion := de.getPMMVersion()

	switch pmmVersion {
	case 1:
		sql = `
			select qc.checksum as sql_id,
				   qc.fingerprint,
				   qe.query    as example,
				   qe.db       as db_name,
				   m.exec_count,
				   m.total_exec_time,
				   m.avg_exec_time,
				   m.rows_examined_max
			from (
					 select qcm.query_class_id,
							sum(qcm.query_count)                                        as exec_count,
							truncate(sum(qcm.query_time_sum), 2)                        as total_exec_time,
							truncate(sum(qcm.query_time_sum) / sum(qcm.query_count), 2) as avg_exec_time,
							qcm.rows_examined_max
					 from query_class_metrics qcm
							  inner join instances i on qcm.instance_id = i.instance_id
					 where i.name = ?
					   and qcm.start_ts >= ?
					   and qcm.start_ts < ?
					   and qcm.rows_examined_max >= ?
					 group by query_class_id
					 order by rows_examined_max desc) m
					 inner join query_examples qe on m.query_class_id = qe.query_class_id
					 inner join query_classes qc on m.query_class_id = qc.query_class_id
			;
		`
		result, err = de.monitorMySQLConn.Execute(sql, serviceName, de.operationInfo.startTime, de.operationInfo.endTime, slowQueryRowsExaminedConfig.GetLowWatermark())
	case 2:
		sql = `
			select queryid                                                       as sql_id,
				   fingerprint,
				   (select example from metrics where queryid = queryid limit 1) as example,
				   database                                                      as db_name,
				   sum(num_queries)                                              as exec_count,
				   truncate(sum(m_query_time_sum), 2)                            as total_exec_time,
				   truncate(sum(m_query_time_sum) / sum(num_queries), 2)         as avg_exec_time,
				   max(m_rows_examined_max)                                      as rows_examined_max
			from metrics
			where service_type = 'mysql'
			  and service_name = ?
			  and period_start >= ?
			  and period_start < ?
			  and m_rows_examined_max >= ?
			group by queryid, fingerprint
			order by rows_examined_max desc;
		`
		result, err = de.monitorClickhouseConn.Execute(sql, serviceName, de.operationInfo.startTime, de.operationInfo.endTime, slowQueryRowsExaminedConfig.GetLowWatermark())
	default:
		return message.NewMessage(msghc.ErrPmmVersionInvalid, pmmVersion)
	}
	if err != nil {
		return err
	}

	var (
		topSQLList                       []*SlowQuery
		slowQueryRowsExaminedHighSum     int
		slowQueryRowsExaminedHighCount   int
		slowQueryRowsExaminedMediumSum   int
		slowQueryRowsExaminedMediumCount int
	)

	slowQueries := make([]*SlowQuery, result.RowNumber())
	err = result.MapToStructSlice(slowQueries, constant.DefaultMiddlewareTag)
	if err != nil {
		return err
	}
	// slow query data
	jsonBytesRowsExamined, err := json.Marshal(slowQueries)
	if err != nil {
		return err
	}
	de.result.SlowQueryData = string(jsonBytesRowsExamined)

	for i, slowQuery := range slowQueries {
		if i < defaultSlowQueryTopSQLNum {
			topSQLList = append(topSQLList, slowQuery)
		}

		if slowQuery.RowsExaminedMax >= int(slowQueryRowsExaminedConfig.GetHighWatermark()) {
			// slow query rows examined high
			slowQueryRowsExaminedHighSum += slowQuery.RowsExaminedMax
			slowQueryRowsExaminedHighCount++
			continue
		}
		// slow query rows examined medium
		slowQueryRowsExaminedMediumSum += slowQuery.RowsExaminedMax
		slowQueryRowsExaminedMediumCount++
	}
	// slow query rows examined high score
	slowQueryRowsExaminedHighScore := (float64(slowQueryRowsExaminedHighSum)/float64(slowQueryRowsExaminedHighCount) - slowQueryRowsExaminedConfig.GetHighWatermark()) / slowQueryRowsExaminedConfig.GetUnit() * slowQueryRowsExaminedConfig.GetScoreDeductionPerUnitHigh()
	if slowQueryRowsExaminedHighScore > slowQueryRowsExaminedConfig.GetMaxScoreDeductionHigh() {
		slowQueryRowsExaminedHighScore = slowQueryRowsExaminedConfig.GetMaxScoreDeductionHigh()
	}
	// slow query rows examined medium score
	slowQueryRowsExaminedMediumScore := (float64(slowQueryRowsExaminedMediumSum)/float64(slowQueryRowsExaminedMediumCount) - slowQueryRowsExaminedConfig.GetLowWatermark()) / slowQueryRowsExaminedConfig.GetUnit() * slowQueryRowsExaminedConfig.GetScoreDeductionPerUnitMedium()
	if slowQueryRowsExaminedMediumScore > slowQueryRowsExaminedConfig.GetMaxScoreDeductionMedium() {
		slowQueryRowsExaminedMediumScore = slowQueryRowsExaminedConfig.GetMaxScoreDeductionMedium()
	}
	// slow query score
	de.result.SlowQueryScore = int(defaultMaxScore - slowQueryRowsExaminedHighScore - slowQueryRowsExaminedMediumScore)
	if de.result.SlowQueryScore < defaultMinScore {
		de.result.SlowQueryScore = defaultMinScore
	}

	// sql tuning
	clusterID := de.operationInfo.mysqlServer.GetClusterID()
	// init db service
	dbService := metadata.NewDBServiceWithDefault()
	for _, sql := range topSQLList {
		// get db info
		err = dbService.GetByNameAndClusterInfo(sql.DBName, clusterID, defaultClusterType)
		if err != nil {
			return err
		}
		if len(dbService.GetDBs()) == constant.ZeroInt {
			return errors.New(fmt.Sprintf("could not find db info. db_name: %s, cluster_id: %d, cluster_type: %d",
				sql.DBName, clusterID, defaultClusterType))
		}
		// get db id
		dbID := dbService.GetDBs()[constant.ZeroInt].Identity()
		// init sql advisor service
		advisorService := sqladvisor.NewServiceWithDefault()
		// get advice
		advice, err := advisorService.Advise(dbID, sql.Example)
		if err != nil {
			return err
		}

		de.result.SlowQueryAdvice += advice + constant.CommaString
	}

	strings.Trim(de.result.SlowQueryAdvice, constant.CommaString)

	return nil
}

// summarize summarizes all item scores with weight
func (de *DefaultEngine) summarize() {
	de.result.WeightedAverageScore = (de.result.DBConfigScore*de.getItemConfig(defaultDBConfigItemName).GetItemWeight() +
		de.result.CPUUsageScore*de.getItemConfig(defaultCPUUsageItemName).GetItemWeight() +
		de.result.IOUtilScore*de.getItemConfig(defaultIOUtilItemName).GetItemWeight() +
		de.result.DiskCapacityUsageScore*de.getItemConfig(defaultDiskCapacityUsageItemName).GetItemWeight() +
		de.result.ConnectionUsageScore*de.getItemConfig(defaultConnectionUsageItemName).GetItemWeight() +
		de.result.AverageActiveSessionNumScore*de.getItemConfig(defaultAverageActiveSessionNumItemName).GetItemWeight() +
		de.result.CacheMissRatioScore*de.getItemConfig(defaultCacheMissRatioItemName).GetItemWeight() +
		de.result.TableSizeScore*(de.getItemConfig(defaultTableRowsItemName).GetItemWeight()+de.getItemConfig(defaultTableSizeItemName).GetItemWeight()) +
		de.result.SlowQueryScore*(de.getItemConfig(defaultSlowQueryRowsExaminedItemName).GetItemWeight())) /
		constant.MaxPercentage

	if de.result.WeightedAverageScore < defaultMinScore {
		de.result.WeightedAverageScore = defaultMinScore
	}
}

// postRun performs post-run actions, for now, it ony saves healthcheck result to the middleware
func (de *DefaultEngine) postRun() error {
	// save result
	return de.getDASRepo().SaveResult(de.result)
}

func (de *DefaultEngine) parsePrometheusDatas(item string, datas []healthcheck.PrometheusData) (int, string, string, error) {
	config := de.getItemConfig(item)

	var (
		highSum     float64
		highCount   int
		mediumSum   float64
		mediumCount int

		highDatas []healthcheck.PrometheusData
	)
	// parse monitor data
	for _, data := range datas {
		switch {
		case data.GetValue() >= config.GetHighWatermark():
			highDatas = append(highDatas, data)
			highSum += data.GetValue()
			highCount++
		case data.GetValue() >= config.GetLowWatermark():
			mediumSum += data.GetValue()
			mediumCount++
		}
	}

	// high score deduction
	scoreDeductionHigh := (highSum/float64(highCount) - config.GetHighWatermark()) / config.GetUnit() * config.GetScoreDeductionPerUnitHigh()
	if scoreDeductionHigh > config.GetMaxScoreDeductionHigh() {
		scoreDeductionHigh = config.GetMaxScoreDeductionHigh()
	}
	// medium score deduction
	scoreDeductionMedium := (mediumSum/float64(mediumCount) - config.GetLowWatermark()) / config.GetUnit() * config.GetScoreDeductionPerUnitMedium()
	if scoreDeductionMedium > config.GetMaxScoreDeductionMedium() {
		scoreDeductionMedium = config.GetMaxScoreDeductionMedium()
	}
	// calculate score
	score := int(defaultMaxScore - scoreDeductionHigh - scoreDeductionMedium)
	if score < constant.ZeroInt {
		score = constant.ZeroInt
	}

	jsonBytesTotal, err := json.Marshal(datas)
	if err != nil {
		return constant.ZeroInt, constant.EmptyString, constant.EmptyString, err
	}
	jsonBytesHigh, err := json.Marshal(highDatas)
	if err != nil {
		return constant.ZeroInt, constant.EmptyString, constant.EmptyString, err
	}

	return score, string(jsonBytesTotal), string(jsonBytesHigh), nil
}
