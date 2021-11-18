package healthcheck

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/app/alert"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/internal/app/sqladvisor"
	"github.com/romberli/das/internal/dependency/healthcheck"
	depmeta "github.com/romberli/das/internal/dependency/metadata"
	depquery "github.com/romberli/das/internal/dependency/query"
	"github.com/romberli/das/pkg/message"
	msghc "github.com/romberli/das/pkg/message/healthcheck"
	util "github.com/romberli/das/pkg/util/query"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/linux"
	"github.com/romberli/go-util/middleware/sql/statement"
	"github.com/romberli/log"
	"github.com/spf13/viper"
)

const (
	defaultMinScore                             = 0
	defaultMaxScore                             = 100.0
	defaultHundred                              = 100
	defaultDBConfigItemName                     = "db_config"
	defaultAvgBackupFailedRatioItemName         = "avg_backup_failed_ratio"
	defaultStatisticFailedRatioItemName         = "statistics_failed_ratio"
	defaultCPUUsageItemName                     = "cpu_usage"
	defaultIOUtilItemName                       = "io_util"
	defaultDiskCapacityUsageItemName            = "disk_capacity_usage"
	defaultConnectionUsageItemName              = "connection_usage"
	defaultAverageActiveSessionPercentsItemName = "average_active_session_percents"
	defaultCacheMissRatioItemName               = "cache_miss_ratio"
	defaultTableRowsItemName                    = "table_rows"
	defaultTableSizeItemName                    = "table_size"
	defaultSlowQueryRowsExaminedItemName        = "slow_query_rows_examined"
	defaultSlowQueryTopSQLNum                   = 3
	defaultClusterType                          = 1

	defaultAlertSubject = "das healthcheck result"
)

var (
	_ healthcheck.Engine = (*DefaultEngine)(nil)

	ignoreDBList = []string{"information_schema", "performance_schema", "mysql", "test", "sys"}
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
		result:               NewEmptyResultWithOperationID(operationInfo.GetOperationID()),
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
			return
		}
		return
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
	// check mysql average backup failed ratio
	err = de.CheckAvgBackupFailedRatio()
	if err != nil {
		return err
	}
	// check mysql statistic failed ratio
	err = de.CheckStatisticFailedRatio()
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
	// check average active session percents
	err = de.checkAverageActiveSessionPercents()
	if err != nil {
		return err
	}
	// check cache miss ratio
	err = de.checkCacheMissRatio()
	if err != nil {
		return err
	}
	// check table rows
	err = de.checkTableRows()
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
	fileSystems, err := de.getPrometheusRepo().GetFileSystems()
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

		if !common.StringInSlice(de.mountPoints, mountPoint) {
			de.mountPoints = append(de.mountPoints, mountPoint)
		}

		for _, fileSystem := range fileSystems {
			device := fileSystem.GetDevice()
			if mountPoint == fileSystem.GetMountPoint() && !common.StringInSlice(de.devices, device) {
				de.devices = append(de.devices, device)
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

	for _, defaultEngineConfig := range defaultEngineConfigList {
		de.engineConfig[defaultEngineConfig.ItemName] = defaultEngineConfig
	}
	// validate config
	return de.engineConfig.Validate()
}

// checkDBConfig checks database configuration
func (de *DefaultEngine) checkDBConfig() error {
	// load database config
	var configItems []string
	for item := range dbConfigVariableNames {
		configItems = append(configItems, item)
	}

	globalVariables, err := de.getApplicationMySQLRepo().GetVariables(configItems)
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
			// slave_parallel_workers
		case dbConfigSlaveParallelWorkers:
			workers, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			if workers != dbConfigSlaveParallelWorkersValid {
				dbConfigCount++
				variables = append(variables, NewVariable(dbConfigSlaveParallelWorkers, value, strconv.Itoa(dbConfigSlaveParallelWorkersValid)))
			}
			// others
		case dbConfigLogBin, dbConfigBinlogFormat, dbConfigBinlogRowImage, dbConfigSyncBinlog,
			dbConfigInnodbFlushLogAtTrxCommit, dbConfigGTIDMode, dbConfigEnforceGTIDConsistency,
			dbConfigSlaveParallelType, dbConfigMasterInfoRepository, dbConfigRelayLogInfoRepository,
			dbConfigReportHost, dbConfigReportPort, dbConfigInnodbFlushMethod, dbConfigInnodbMonitorEnable,
			dbConfigInnodbPrintAllDeadlocks, dbConfigSlowQueryLog, dbConfigPerformanceSchema:
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

// CheckAvgBackupFailedRatio checks the mysql backup information
func (de *DefaultEngine) CheckAvgBackupFailedRatio() error {
	// get data
	datas, err := de.getPrometheusRepo().GetAvgBackupFailedRatio()
	if err != nil {
		return err
	}
	// parse data
	de.result.AvgBackupFailedRatioScore, de.result.AvgBackupFailedRatioData, de.result.AvgBackupFailedRatioHigh, err = de.parsePrometheusDatas(defaultAvgBackupFailedRatioItemName, datas)
	if err != nil {
		return err
	}

	return nil
}

// CheckStatisticFailedRatio checks the statistic of mysql
func (de *DefaultEngine) CheckStatisticFailedRatio() error {
	// get data
	datas, err := de.getPrometheusRepo().GetCPUUsage()
	if err != nil {
		return err
	}
	// parse data
	de.result.StatisticFailedRatioScore, de.result.StatisticFailedRatioData, de.result.StatisticFailedRatioHigh, err = de.parsePrometheusDatas(defaultStatisticFailedRatioItemName, datas)
	if err != nil {
		return err
	}

	return nil
}

// checkCPUUsage checks cpu usage
func (de *DefaultEngine) checkCPUUsage() error {
	// get data
	datas, err := de.getPrometheusRepo().GetCPUUsage()
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
	datas, err := de.getPrometheusRepo().GetIOUtil()
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
	datas, err := de.getPrometheusRepo().GetDiskCapacityUsage(de.getMountPoints())
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
	datas, err := de.getPrometheusRepo().GetConnectionUsage()
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

// checkAverageActiveSessionPercents check active session number
func (de *DefaultEngine) checkAverageActiveSessionPercents() error {
	// get data
	datas, err := de.getPrometheusRepo().GetAverageActiveSessionPercents()
	if err != nil {
		return err
	}
	// parse data
	de.result.AverageActiveSessionPercentsScore, de.result.AverageActiveSessionPercentsData, de.result.AverageActiveSessionPercentsHigh, err = de.parsePrometheusDatas(defaultAverageActiveSessionPercentsItemName, datas)
	if err != nil {
		return err
	}

	return nil
}

// checkCacheMissRatio checks cache miss ratio
func (de *DefaultEngine) checkCacheMissRatio() error {
	// get data
	datas, err := de.getPrometheusRepo().GetCacheMissRatio()
	if err != nil {
		return err
	}
	// parse data
	de.result.CacheMissRatioScore, de.result.CacheMissRatioData, de.result.CacheMissRatioHigh, err = de.parsePrometheusDatas(defaultCacheMissRatioItemName, datas)
	if err != nil {
		return err
	}

	return nil
}

// checkTableSize checks table rows
func (de *DefaultEngine) checkTableRows() error {
	// get tables
	tables, err := de.getApplicationMySQLRepo().GetLargeTables()
	if err != nil {
		return err
	}

	tableRowsConfig := de.getItemConfig(defaultTableRowsItemName)

	var (
		tableRowsHighSum              int
		tableRowsHighCount            int
		tableRowsMediumSum            int
		tableRowsMediumCount          int
		tableRowsScoreDeductionHigh   float64
		tableRowsScoreDeductionMedium float64

		tableRowsHigh []healthcheck.Table
	)

	for _, table := range tables {
		switch {
		case float64(table.GetRows()) >= tableRowsConfig.GetHighWatermark():
			tableRowsHigh = append(tableRowsHigh, table)
			tableRowsHighSum += table.GetRows()
			tableRowsHighCount++
		case float64(table.GetRows()) >= tableRowsConfig.GetLowWatermark():
			tableRowsMediumSum += table.GetRows()
			tableRowsMediumCount++
		}
	}

	// table rows data
	jsonBytesTotal, err := json.Marshal(tables)
	if err != nil {
		return nil
	}
	de.result.TableRowsData = string(jsonBytesTotal)
	// table rows high
	jsonBytesHigh, err := json.Marshal(tableRowsHigh)
	if err != nil {
		return nil
	}
	de.result.TableRowsHigh = string(jsonBytesHigh)

	// table rows high score deduction
	if tableRowsHighCount > constant.ZeroInt {
		tableRowsScoreDeductionHigh = (float64(tableRowsHighSum)/float64(tableRowsHighCount) - tableRowsConfig.GetHighWatermark()) / tableRowsConfig.GetUnit() * tableRowsConfig.GetScoreDeductionPerUnitHigh()
		if tableRowsScoreDeductionHigh > tableRowsConfig.GetMaxScoreDeductionHigh() {
			tableRowsScoreDeductionHigh = tableRowsConfig.GetMaxScoreDeductionHigh()
		}
	}
	// table rows medium score deduction
	if tableRowsMediumCount > constant.ZeroInt {
		tableRowsScoreDeductionMedium = (float64(tableRowsMediumSum)/float64(tableRowsMediumCount) - tableRowsConfig.GetLowWatermark()) / tableRowsConfig.GetUnit() * tableRowsConfig.GetScoreDeductionPerUnitMedium()
		if tableRowsScoreDeductionMedium > tableRowsConfig.GetMaxScoreDeductionMedium() {
			tableRowsScoreDeductionMedium = tableRowsConfig.GetMaxScoreDeductionMedium()
		}
	}
	// table rows score
	de.result.TableRowsScore = int(defaultMaxScore - tableRowsScoreDeductionHigh - tableRowsScoreDeductionMedium)
	if de.result.TableRowsScore < constant.ZeroInt {
		de.result.TableRowsScore = constant.ZeroInt
	}

	return nil
}

// checkTableSize checks table sizes
func (de *DefaultEngine) checkTableSize() error {
	// get tables
	tables, err := de.getApplicationMySQLRepo().GetLargeTables()
	if err != nil {
		return err
	}

	tableSizeConfig := de.getItemConfig(defaultTableSizeItemName)

	var (
		tableSizeHighSum              float64
		tableSizeHighCount            int
		tableSizeMediumSum            float64
		tableSizeMediumCount          int
		tableSizeScoreDeductionHigh   float64
		tableSizeScoreDeductionMedium float64

		tableSizeHigh []healthcheck.Table
	)

	for _, table := range tables {
		switch {
		case table.GetSize() >= tableSizeConfig.GetHighWatermark():
			tableSizeHigh = append(tableSizeHigh, table)
			tableSizeHighSum += table.GetSize()
			tableSizeHighCount++
		case table.GetSize() >= tableSizeConfig.GetLowWatermark():
			tableSizeMediumSum += table.GetSize()
			tableSizeMediumCount++
		}
	}

	// table size data
	jsonBytesTotal, err := json.Marshal(tables)
	if err != nil {
		return nil
	}
	de.result.TableSizeData = string(jsonBytesTotal)
	// table size high
	jsonBytesHigh, err := json.Marshal(tableSizeHigh)
	if err != nil {
		return nil
	}
	de.result.TableSizeHigh = string(jsonBytesHigh)

	// table size high score deduction
	if tableSizeHighCount > constant.ZeroInt {
		tableSizeScoreDeductionHigh = (tableSizeHighSum/float64(tableSizeHighCount) - tableSizeConfig.GetHighWatermark()) / tableSizeConfig.GetUnit() * tableSizeConfig.GetScoreDeductionPerUnitHigh()
		if tableSizeScoreDeductionHigh > tableSizeConfig.GetMaxScoreDeductionHigh() {
			tableSizeScoreDeductionHigh = tableSizeConfig.GetMaxScoreDeductionHigh()
		}
	}
	// table size medium score deduction
	if tableSizeMediumCount > constant.ZeroInt {
		tableSizeScoreDeductionMedium = (tableSizeMediumSum/float64(tableSizeMediumCount) - tableSizeConfig.GetLowWatermark()) / tableSizeConfig.GetUnit() * tableSizeConfig.GetScoreDeductionPerUnitMedium()
		if tableSizeScoreDeductionMedium > tableSizeConfig.GetMaxScoreDeductionMedium() {
			tableSizeScoreDeductionMedium = tableSizeConfig.GetMaxScoreDeductionMedium()
		}
	}
	// table size score
	de.result.TableSizeScore = int(defaultMaxScore - tableSizeScoreDeductionHigh - tableSizeScoreDeductionMedium)
	if de.result.TableSizeScore < constant.ZeroInt {
		de.result.TableSizeScore = constant.ZeroInt
	}

	return nil
}

// checkSlowQuery checks slow query
func (de *DefaultEngine) checkSlowQuery() error {
	// check slow query execution time
	slowQueries, err := de.getQueryRepo().GetSlowQuery()
	if err != nil {
		return err
	}

	var (
		i                                int
		slowQueryRowsExaminedHighSum     int
		slowQueryRowsExaminedHighCount   int
		slowQueryRowsExaminedMediumSum   int
		slowQueryRowsExaminedMediumCount int

		slowQueryRowsExaminedHighScore   float64
		slowQueryRowsExaminedMediumScore float64

		topSQLList []depquery.Query
	)

	// slow query data
	jsonBytesRowsExamined, err := json.Marshal(slowQueries)
	if err != nil {
		return err
	}
	de.result.SlowQueryData = string(jsonBytesRowsExamined)

	slowQueryRowsExaminedConfig := de.getItemConfig(defaultSlowQueryRowsExaminedItemName)

	for _, slowQuery := range slowQueries {
		if statement.GetType(slowQuery.GetExample()) == statement.Unknown {
			continue
		}
		if i < defaultSlowQueryTopSQLNum {
			dbName, err := util.GetDBName(slowQuery.GetExample())
			if err != nil {
				return err
			}
			if !common.StringInSlice(ignoreDBList, dbName) {
				slowQuery.SetDBName(dbName)
				topSQLList = append(topSQLList, slowQuery)
				i++
			}
		}
		if slowQuery.GetRowsExaminedMax() >= int(slowQueryRowsExaminedConfig.GetHighWatermark()) {
			// slow query rows examined high
			slowQueryRowsExaminedHighSum += slowQuery.GetRowsExaminedMax()
			slowQueryRowsExaminedHighCount++
			continue
		}
		if slowQuery.GetRowsExaminedMax() >= int(slowQueryRowsExaminedConfig.GetLowWatermark()) {
			// slow query rows examined medium
			slowQueryRowsExaminedMediumSum += slowQuery.GetRowsExaminedMax()
			slowQueryRowsExaminedMediumCount++
		}
	}
	// slow query rows examined high score
	if slowQueryRowsExaminedHighCount > constant.ZeroInt {
		slowQueryRowsExaminedHighScore = (float64(slowQueryRowsExaminedHighSum)/float64(slowQueryRowsExaminedHighCount) - slowQueryRowsExaminedConfig.GetHighWatermark()) / slowQueryRowsExaminedConfig.GetUnit() * slowQueryRowsExaminedConfig.GetScoreDeductionPerUnitHigh()
		if slowQueryRowsExaminedHighScore > slowQueryRowsExaminedConfig.GetMaxScoreDeductionHigh() {
			slowQueryRowsExaminedHighScore = slowQueryRowsExaminedConfig.GetMaxScoreDeductionHigh()
		}
	}
	// slow query rows examined medium score
	if slowQueryRowsExaminedMediumCount > constant.ZeroInt {
		slowQueryRowsExaminedMediumScore = (float64(slowQueryRowsExaminedMediumSum)/float64(slowQueryRowsExaminedMediumCount) - slowQueryRowsExaminedConfig.GetLowWatermark()) / slowQueryRowsExaminedConfig.GetUnit() * slowQueryRowsExaminedConfig.GetScoreDeductionPerUnitMedium()
		if slowQueryRowsExaminedMediumScore > slowQueryRowsExaminedConfig.GetMaxScoreDeductionMedium() {
			slowQueryRowsExaminedMediumScore = slowQueryRowsExaminedConfig.GetMaxScoreDeductionMedium()
		}
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
		var advice string

		// get db info
		if sql.GetDBName() != constant.EmptyString {
			err = dbService.GetByNameAndClusterInfo(sql.GetDBName(), clusterID, defaultClusterType)
			if err != nil {
				return err
			}
			// get db id
			dbID := dbService.GetDBs()[constant.ZeroInt].Identity()
			// init sql advisor service
			advisorService := sqladvisor.NewServiceWithDefault()
			// get advice
			advice, err = advisorService.Advise(dbID, sql.GetExample())
			if err != nil {
				return err
			}
		} else {
			jsonBytes, err := json.Marshal(sql)
			if err != nil {
				return err
			}
			advice = string(jsonBytes)
		}

		de.result.SlowQueryAdvice += advice + constant.CommaString
	}

	de.result.SlowQueryAdvice = strings.Trim(de.result.SlowQueryAdvice, constant.CommaString)

	return nil
}

// summarize summarizes all item scores with weight
func (de *DefaultEngine) summarize() {
	de.result.WeightedAverageScore = (de.result.DBConfigScore*de.getItemConfig(defaultDBConfigItemName).GetItemWeight() +
		de.result.AvgBackupFailedRatioScore*de.getItemConfig(defaultAvgBackupFailedRatioItemName).GetItemWeight() +
		de.result.StatisticFailedRatioScore*de.getItemConfig(defaultStatisticFailedRatioItemName).GetItemWeight() +
		de.result.CPUUsageScore*de.getItemConfig(defaultCPUUsageItemName).GetItemWeight() +
		de.result.IOUtilScore*de.getItemConfig(defaultIOUtilItemName).GetItemWeight() +
		de.result.DiskCapacityUsageScore*de.getItemConfig(defaultDiskCapacityUsageItemName).GetItemWeight() +
		de.result.ConnectionUsageScore*de.getItemConfig(defaultConnectionUsageItemName).GetItemWeight() +
		de.result.AverageActiveSessionPercentsScore*de.getItemConfig(defaultAverageActiveSessionPercentsItemName).GetItemWeight() +
		de.result.CacheMissRatioScore*de.getItemConfig(defaultCacheMissRatioItemName).GetItemWeight() +
		de.result.TableRowsScore*de.getItemConfig(defaultTableRowsItemName).GetItemWeight() +
		de.result.TableSizeScore*de.getItemConfig(defaultTableSizeItemName).GetItemWeight() +
		de.result.SlowQueryScore*(de.getItemConfig(defaultSlowQueryRowsExaminedItemName).GetItemWeight())) /
		constant.MaxPercentage

	if de.result.WeightedAverageScore < defaultMinScore {
		de.result.WeightedAverageScore = defaultMinScore
	}
}

// postRun performs post-run actions, for now, it ony saves healthcheck result to the middleware
func (de *DefaultEngine) postRun() error {
	// save result
	err := de.getDASRepo().SaveResult(de.result) // send email
	if err != nil {
		return err
	}

	err = de.sendEmail()
	if err != nil {
		return err
	}

	return nil
}

// parsePrometheusDatas parses prometheus datas
func (de *DefaultEngine) parsePrometheusDatas(item string, datas []healthcheck.PrometheusData) (int, string, string, error) {
	cfg := de.getItemConfig(item)

	var (
		highSum              float64
		highCount            int
		mediumSum            float64
		mediumCount          int
		scoreDeductionHigh   float64
		scoreDeductionMedium float64

		highDatas []healthcheck.PrometheusData
	)
	// parse monitor data
	for _, data := range datas {
		switch {
		case data.GetValue() >= cfg.GetHighWatermark():
			highDatas = append(highDatas, data)
			highSum += data.GetValue()
			highCount++
		case data.GetValue() >= cfg.GetLowWatermark():
			mediumSum += data.GetValue()
			mediumCount++
		}
	}

	// high score deduction
	if highCount > constant.ZeroInt {
		scoreDeductionHigh = (highSum/float64(highCount) - cfg.GetHighWatermark()) / cfg.GetUnit() * cfg.GetScoreDeductionPerUnitHigh()
		if scoreDeductionHigh > cfg.GetMaxScoreDeductionHigh() {
			scoreDeductionHigh = cfg.GetMaxScoreDeductionHigh()
		}
	}
	// medium score deduction
	if mediumCount > constant.ZeroInt {
		scoreDeductionMedium = (mediumSum/float64(mediumCount) - cfg.GetLowWatermark()) / cfg.GetUnit() * cfg.GetScoreDeductionPerUnitMedium()
		if scoreDeductionMedium > cfg.GetMaxScoreDeductionMedium() {
			scoreDeductionMedium = cfg.GetMaxScoreDeductionMedium()
		}
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

func (de *DefaultEngine) sendEmail() error {
	toAddrs, err := de.getToAddrs()
	if err != nil {
		return err
	}

	result, err := de.getDASRepo().GetResultByOperationID(de.getOperationInfo().GetOperationID())
	if err != nil {
		return err
	}
	alertService := alert.NewServiceWithDefault(alert.NewConfigFromFile())

	return alertService.SendEmail(toAddrs, constant.EmptyString, defaultAlertSubject, result.String())
}

// getToAddrs gets to addrs that will send email to
func (de *DefaultEngine) getToAddrs() (string, error) {
	mysqlCluster, err := de.getOperationInfo().GetMySQLServer().GetMySQLCluster()
	if err != nil {
		return constant.EmptyString, err
	}

	var (
		owners  []depmeta.User
		toAddrs string
	)
	switch viper.GetString(config.HealthcheckAlertOwnerTypeKey) {
	case config.HealthcheckAlertOwnerTypeApp:
		owners, err = mysqlCluster.GetAppOwners()
	case config.HealthcheckAlertOwnerTypeDB:
		owners, err = mysqlCluster.GetDBOwners()
	case config.HealthcheckAlertOwnerTypeAll:
		owners, err = mysqlCluster.GetAllOwners()
	}
	if err != nil {
		return constant.EmptyString, err
	}

	for _, owner := range owners {
		toAddrs += owner.GetEmail() + constant.CommaString
	}

	return strings.Trim(toAddrs, constant.CommaString), nil
}
