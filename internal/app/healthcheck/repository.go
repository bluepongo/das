package healthcheck

import (
	"fmt"
	"time"

	"github.com/hashicorp/go-version"
	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/healthcheck"
	"github.com/romberli/das/pkg/message"
	msghc "github.com/romberli/das/pkg/message/healthcheck"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/go-util/middleware/prometheus"
	"github.com/romberli/log"
)

const (
	mysql57           = "5.7"
	performanceSchema = "performance_schema"
	informationSchema = "information_schema"

	deviceLabel     = "device"
	mountPointLabel = "mountpoint"

	DataDirVariable   = "datadir"
	BinlogDirVariable = "log_bin_base"

	MinTableRows = 1000000
)

var (
	_ healthcheck.DASRepo              = (*DASRepo)(nil)
	_ healthcheck.ApplicationMySQLRepo = (*ApplicationMySQLRepo)(nil)
	_ healthcheck.PrometheusRepo       = (*PrometheusRepo)(nil)
)

// DASRepo for health check
type DASRepo struct {
	Database middleware.Pool
}

// NewDASRepo returns *DASRepo with given middleware.Pool
func NewDASRepo(db middleware.Pool) *DASRepo {
	return newDASRepo(db)
}

// NewDASRepoWithGlobal returns *DASRepo with global mysql pool
func NewDASRepoWithGlobal() *DASRepo {
	return NewDASRepo(global.DASMySQLPool)
}

// newDASRepo returns *DASRepo with given middleware.Pool
func newDASRepo(db middleware.Pool) *DASRepo {
	return &DASRepo{Database: db}
}

// Execute executes given command and placeholders on the middleware
func (dr *DASRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := dr.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("healthcheck DASRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
func (dr *DASRepo) Transaction() (middleware.Transaction, error) {
	return dr.Database.Transaction()
}

// LoadEngineConfig loads engine config from the middleware
func (dr *DASRepo) LoadEngineConfig() (healthcheck.EngineConfig, error) {
	// load config
	sql := `
		select id, item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high,
		score_deduction_per_unit_medium, max_score_deduction_medium, del_flag, create_time, last_update_time
		from t_hc_default_engine_config
		where del_flag = 0;
	`
	log.Debugf("healthcheck DASRepo.loadEngineConfig() sql: \n%s\n", sql)
	result, err := dr.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*DefaultItemConfig
	defaultItemConfigs := make([]*DefaultItemConfig, result.RowNumber())
	for i := range defaultItemConfigs {
		defaultItemConfigs[i] = NewEmptyDefaultItemConfig()
	}
	// map to struct
	err = result.MapToStructSlice(defaultItemConfigs, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	defaultEngineConfig := NewEmptyDefaultEngineConfig()
	for _, defaultItemConfig := range defaultItemConfigs {
		defaultEngineConfig[defaultItemConfig.ItemName] = defaultItemConfig
	}

	return defaultEngineConfig, nil
}

// GetResultByOperationID gets a Result by the operationID from the middleware
func (dr *DASRepo) GetResultByOperationID(operationID int) (healthcheck.Result, error) {
	sql := `
		select id, operation_id, weighted_average_score, db_config_score, db_config_data, 
		db_config_advice, cpu_usage_score, cpu_usage_data, cpu_usage_high, io_util_score,
		io_util_data, io_util_high, disk_capacity_usage_score, disk_capacity_usage_data, 
		disk_capacity_usage_high, connection_usage_score, connection_usage_data, 
		connection_usage_high, average_active_session_num_score, average_active_session_num_data,
		average_active_session_num_high, cache_miss_ratio_score, cache_miss_ratio_data, 
		cache_miss_ratio_high, table_size_score, table_size_data, table_size_high, slow_query_score,
		slow_query_data, slow_query_advice, accuracy_review, del_flag, create_time, last_update_time
		from t_hc_result
		where del_flag = 0
		and operation_id = ? 
		order by id;
	`
	log.Debugf("healthCheck DASRepo.GetResultByOperationID select sql: \n%s\nplaceholders: %s", sql, operationID)

	result, err := dr.Execute(sql, operationID)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, fmt.Errorf("healthCheck DASRepo.GetResultByOperationID(): data does not exists, operation_id: %d", operationID)
	case 1:
		hcInfo := NewEmptyResultWithRepo(dr)
		// map to struct
		err = result.MapToStructByRowIndex(hcInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return hcInfo, nil
	default:
		return nil, fmt.Errorf("healthCheck DASRepo.GetResultByOperationID(): duplicate key exists, operation_id: %d", operationID)
	}
}

// IsRunning gets status by the mysqlServerID from the middleware
func (dr *DASRepo) IsRunning(mysqlServerID int) (bool, error) {
	sql := `select count(1) from t_hc_operation_info where del_flag = 0 and mysql_server_id = ? and status = 1;`
	log.Debugf("healthCheck DASRepo.IsRunning() select sql: \n%s\nplaceholders: %s", sql, mysqlServerID)

	result, err := dr.Execute(sql, mysqlServerID)
	if err != nil {
		return false, err
	}
	count, _ := result.GetInt(constant.ZeroInt, constant.ZeroInt)

	return count != 0, nil
}

// InitOperation creates a operationInfo in the middleware
func (dr *DASRepo) InitOperation(mysqlServerID int, startTime, endTime time.Time, step time.Duration) (int, error) {
	startTimeStr := startTime.Format(constant.TimeLayoutSecond)
	endTimeStr := endTime.Format(constant.TimeLayoutSecond)
	stepInt := int(step.Seconds())

	sql := `insert into t_hc_operation_info(mysql_server_id, start_time, end_time, step) values(?, ?, ?, ?);`
	log.Debugf("healthCheck DASRepo.InitOperation() insert sql: \n%s\nplaceholders: %s, %s, %s, %s", sql, mysqlServerID, startTimeStr, endTimeStr, stepInt)

	_, err := dr.Execute(sql, mysqlServerID, startTimeStr, endTimeStr, stepInt)
	if err != nil {
		return constant.ZeroInt, err
	}

	sql = `
		select id from t_hc_operation_info where del_flag = 0 and 
		mysql_server_id = ? and start_time = ? and end_time = ? and step = ?;
	`
	log.Debugf("healthCheck DASRepo.InitOperation() select sql: \n%s\nplaceholders: %s, %s, %s, %s", sql, mysqlServerID, startTimeStr, endTimeStr, stepInt)

	result, err := dr.Execute(sql, mysqlServerID, startTimeStr, endTimeStr, stepInt)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// UpdateOperationStatus updates the status and message by the operationID in the middleware
func (dr *DASRepo) UpdateOperationStatus(operationID int, status int, message string) error {
	sql := `update t_hc_operation_info set status = ?, message = ? where id = ?;`
	log.Debugf("healthCheck DASRepo.UpdateOperationStatus() update sql: \n%s\nplaceholders: %s, %s, %s", sql, operationID, status, message)
	_, err := dr.Execute(sql, status, message, operationID)

	return err
}

// SaveResult saves the result in the middleware
func (dr *DASRepo) SaveResult(result healthcheck.Result) error {
	sql := `insert into t_hc_result(operation_id, weighted_average_score, db_config_score, db_config_data, 
		db_config_advice, cpu_usage_score, cpu_usage_data, cpu_usage_high, io_util_score,
		io_util_data, io_util_high, disk_capacity_usage_score, disk_capacity_usage_data, 
		disk_capacity_usage_high, connection_usage_score, connection_usage_data, 
		connection_usage_high, average_active_session_num_score, average_active_session_num_data,
		average_active_session_num_high, cache_miss_ratio_score, cache_miss_ratio_data, 
		cache_miss_ratio_high, table_size_score, table_size_data, table_size_high, slow_query_score,
		slow_query_data, slow_query_advice, accuracy_review) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	log.Debugf("healthCheck DASRepo.SaveResult() insert sql: \n%s\nplaceholders: %s, %s, %s, %s, %s, "+
		"%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
		sql, result.GetOperationID(), result.GetWeightedAverageScore(), result.GetDBConfigScore(), result.GetDBConfigData(),
		result.GetDBConfigAdvice(), result.GetCPUUsageScore(), result.GetCPUUsageData(), result.GetCPUUsageHigh(),
		result.GetIOUtilScore(), result.GetIOUtilData(), result.GetIOUtilHigh(), result.GetDiskCapacityUsageScore(),
		result.GetDiskCapacityUsageData(), result.GetDiskCapacityUsageHigh(), result.GetConnectionUsageScore(),
		result.GetConnectionUsageData(), result.GetConnectionUsageHigh(), result.GetAverageActiveSessionNumScore(),
		result.GetAverageActiveSessionNumData(), result.GetAverageActiveSessionNumHigh(), result.GetCacheMissRatioScore(),
		result.GetCacheMissRatioData(), result.GetCacheMissRatioHigh(), result.GetTableSizeScore(), result.GetTableSizeData(),
		result.GetTableSizeHigh(), result.GetSlowQueryScore(), result.GetSlowQueryData(), result.GetSlowQueryAdvice(),
		result.GetAccuracyReview())

	// execute
	_, err := dr.Execute(sql, result.GetOperationID(), result.GetWeightedAverageScore(), result.GetDBConfigScore(),
		result.GetDBConfigData(), result.GetDBConfigAdvice(), result.GetCPUUsageScore(), result.GetCPUUsageData(),
		result.GetCPUUsageHigh(), result.GetIOUtilScore(), result.GetIOUtilData(), result.GetIOUtilHigh(),
		result.GetDiskCapacityUsageScore(), result.GetDiskCapacityUsageData(), result.GetDiskCapacityUsageHigh(),
		result.GetConnectionUsageScore(), result.GetConnectionUsageData(), result.GetConnectionUsageHigh(),
		result.GetAverageActiveSessionNumScore(), result.GetAverageActiveSessionNumData(), result.GetAverageActiveSessionNumHigh(),
		result.GetCacheMissRatioScore(), result.GetCacheMissRatioData(), result.GetCacheMissRatioHigh(),
		result.GetTableSizeScore(), result.GetTableSizeData(), result.GetTableSizeHigh(), result.GetSlowQueryScore(),
		result.GetSlowQueryData(), result.GetSlowQueryAdvice(), result.GetAccuracyReview())

	return err
}

// UpdateAccuracyReviewByOperationID updates the accuracyReview by the operationID in the middleware
func (dr *DASRepo) UpdateAccuracyReviewByOperationID(operationID int, review int) error {
	sql := `update t_hc_result set accuracy_review = ? where operation_id = ?;`
	log.Debugf("healthCheck DASRepo.UpdateAccuracyReviewByOperationID() update sql: \n%s\nplaceholders: %s, %s", sql, operationID, review)

	_, err := dr.Execute(sql, review, operationID)
	return err
}

// loadEngineConfig loads engine config from the middleware
func (dr *DASRepo) loadEngineConfig() (DefaultEngineConfig, error) {
	// load config
	sql := `
		select id, item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high,
		score_deduction_per_unit_medium, max_score_deduction_medium, del_flag, create_time, last_update_time
		from t_hc_default_engine_config
		where del_flag = 0;
	`
	log.Debugf("healthcheck DASRepo.loadEngineConfig() sql: \n%s\n", sql)
	result, err := dr.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*DefaultItemConfig
	defaultItemConfigs := make([]*DefaultItemConfig, result.RowNumber())
	for i := range defaultItemConfigs {
		defaultItemConfigs[i] = NewEmptyDefaultItemConfig()
	}
	// map to struct
	err = result.MapToStructSlice(defaultItemConfigs, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	defaultEngineConfig := NewEmptyDefaultEngineConfig()
	for _, defaultItemConfig := range defaultItemConfigs {
		defaultEngineConfig[defaultItemConfig.ItemName] = defaultItemConfig
	}

	return defaultEngineConfig, nil
}

type ApplicationMySQLRepo struct {
	operationInfo *OperationInfo
	conn          *mysql.Conn
}

// NewApplicationMySQLRepo returns a new *ApplicationMySQLRepo
func NewApplicationMySQLRepo(operationInfo *OperationInfo, conn *mysql.Conn) *ApplicationMySQLRepo {
	return &ApplicationMySQLRepo{
		operationInfo: operationInfo,
		conn:          conn,
	}
}

// GetOperationInfo returns the operation information
func (amr *ApplicationMySQLRepo) GetOperationInfo() *OperationInfo {
	return amr.operationInfo
}

// Close closes the application mysql connection
func (amr *ApplicationMySQLRepo) Close() error {
	return amr.conn.Close()
}

// GetDBConfig gets db config with given items
func (amr *ApplicationMySQLRepo) GetDBConfig(configItems []string) ([]healthcheck.Variable, error) {
	variables, err := amr.getDBConfig(configItems)
	if err != nil {
		return nil, err
	}

	config := make([]healthcheck.Variable, len(variables))
	for i, variable := range variables {
		config[i] = variable
	}

	return config, err
}

// GetMySQLDirs gets the mysql directories
func (amr *ApplicationMySQLRepo) GetMySQLDirs() ([]string, error) {
	config, err := amr.getDBConfig([]string{DataDirVariable, BinlogDirVariable})
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, variable := range config {
		dirs = append(dirs, variable.GetValue())
	}

	return dirs, nil
}

// GetLargeTables gets the large tables
func (amr *ApplicationMySQLRepo) GetLargeTables() ([]healthcheck.Table, error) {
	tableList, err := amr.getLargeTables()
	if err != nil {
		return nil, err
	}

	tables := make([]healthcheck.Table, len(tableList))
	for i, table := range tableList {
		tables[i] = table
	}

	return tables, nil
}

// getDBConfig gets db config with given items
func (amr *ApplicationMySQLRepo) getDBConfig(configItems []string) ([]*GlobalVariable, error) {
	// prepare args
	interfaces, err := common.ConvertInterfaceToSliceInterface(configItems)
	if err != nil {
		return nil, err
	}
	items, err := middleware.ConvertSliceToString(interfaces...)
	if err != nil {
		return nil, err
	}
	// check mysql version
	mysqlVersion, err := version.NewVersion(amr.GetOperationInfo().GetMySQLServer().GetVersion())
	if err != nil {
		return nil, err
	}
	defaultVersion, err := version.NewVersion(mysql57)
	if err != nil {
		return nil, err
	}
	// prepare sql
	sql := fmt.Sprintf(applicationMySQLDBConfig, performanceSchema, items)
	if mysqlVersion.LessThan(defaultVersion) {
		sql = fmt.Sprintf(applicationMySQLDBConfig, informationSchema, items)
	}
	// get result
	result, err := amr.conn.Execute(sql)
	if err != nil {
		return nil, err
	}
	variables := make([]*GlobalVariable, result.RowNumber())
	for i := range variables {
		variables[i] = NewEmptyGlobalVariable()
	}
	err = result.MapToStructSlice(variables, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return variables, nil
}

// GetLargeTables gets the large tables
func (amr *ApplicationMySQLRepo) getLargeTables() ([]*Table, error) {
	result, err := amr.conn.Execute(applicationMySQLTableSize, MinTableRows)
	if err != nil {
		return nil, err
	}
	tables := make([]*Table, result.RowNumber())
	for i := range tables {
		tables[i] = NewEmptyTable()
	}
	err = result.MapToStructSlice(tables, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return tables, nil
}

type PrometheusRepo struct {
	operationInfo *OperationInfo
	conn          *prometheus.Conn
}

// NewPrometheusRepo returns a new *PrometheusRepo
func NewPrometheusRepo(operationInfo *OperationInfo, conn *prometheus.Conn) *PrometheusRepo {
	return &PrometheusRepo{
		operationInfo: operationInfo,
		conn:          conn,
	}
}

// GetOperationInfo returns the operation information
func (pr *PrometheusRepo) GetOperationInfo() *OperationInfo {
	return pr.operationInfo
}

// GetFileSystems gets the file systems from the prometheus
func (pr *PrometheusRepo) GetFileSystems(serviceName string) ([]healthcheck.FileSystem, error) {
	fileSystems, err := pr.getFileSystems(serviceName)
	if err != nil {
		return nil, err
	}

	fileSystemInterfaces := make([]healthcheck.FileSystem, len(fileSystems))
	for i, fileSystem := range fileSystems {
		fileSystemInterfaces[i] = fileSystem
	}

	return fileSystemInterfaces, nil
}

// GetCPUUsage gets the cpu usage
func (pr *PrometheusRepo) GetCPUUsage(serviceName string) ([]healthcheck.PrometheusData, error) {
	prometheusDatas, err := pr.getCPUUsage(serviceName)
	if err != nil {
		return nil, err
	}

	return pr.convertToSliceInterface(prometheusDatas), nil
}

// GetIOUtil gets the io util
func (pr *PrometheusRepo) GetIOUtil(serviceName string, devices []string) ([]healthcheck.PrometheusData, error) {
	prometheusDatas, err := pr.getIOUtil(serviceName, devices)
	if err != nil {
		return nil, err
	}

	return pr.convertToSliceInterface(prometheusDatas), nil
}

// GetDiskCapacityUsage gets the disk capacity usage
func (pr *PrometheusRepo) GetDiskCapacityUsage(serviceName string, mountPoints []string) ([]healthcheck.PrometheusData, error) {
	prometheusDatas, err := pr.getDiskCapacityUsage(serviceName, mountPoints)
	if err != nil {
		return nil, err
	}

	return pr.convertToSliceInterface(prometheusDatas), nil
}

// GetConnectionUsage gets the connection usage
func (pr *PrometheusRepo) GetConnectionUsage(serviceName string) ([]healthcheck.PrometheusData, error) {
	return nil, nil
}

// GetActiveSessionNum gets the active session number
func (pr *PrometheusRepo) GetActiveSessionNum(serviceName string) ([]healthcheck.PrometheusData, error) {
	return nil, nil
}

// GetCacheMissRatio gets the cache miss ratio
func (pr *PrometheusRepo) GetCacheMissRatio(serviceName string) ([]healthcheck.PrometheusData, error) {
	return nil, nil
}

// getPMMVersion return the pmm version
func (pr *PrometheusRepo) getServiceName() string {
	return pr.GetOperationInfo().GetMySQLServer().GetServiceName()
}

// getPMMVersion return the pmm version
func (pr *PrometheusRepo) getPMMVersion() int {
	return pr.GetOperationInfo().GetMonitorSystem().GetSystemType()
}

// convertToSliceInterface converts []*PrometheusData to []healthcheck.PrometheusData
func (pr *PrometheusRepo) convertToSliceInterface(prometheusDatas []*PrometheusData) []healthcheck.PrometheusData {
	datas := make([]healthcheck.PrometheusData, len(prometheusDatas))
	for i, prometheusData := range prometheusDatas {
		datas[i] = prometheusData
	}

	return datas
}

// getFileSystems gets the file system from the prometheus
func (pr *PrometheusRepo) getFileSystems(serviceName string) ([]*FileSystem, error) {
	var query string

	// prepare query
	switch pr.getPMMVersion() {
	case 1:
		// pmm 1.x
		query = PrometheusFileSystemV1
	case 2:
		// pmm 2.x
		query = PrometheusFileSystemV2
	default:
		return nil, message.NewMessage(msghc.ErrPmmVersionInvalid)
	}

	query = fmt.Sprintf(query, serviceName)
	log.Debugf("healthcheck PrometheusRepo.getFileSystems() query: \n%s\n", query)
	// get data
	result, err := pr.conn.Execute(query)
	if err != nil {
		return nil, err
	}
	// parse result
	vector, err := result.Raw.GetVector()
	if err != nil {
		return nil, err
	}

	var fileSystems []*FileSystem
	for _, sample := range vector {
		fileSystems = append(fileSystems, NewFileSystem(string(sample.Metric[mountPointLabel]), string(sample.Metric[deviceLabel])))
	}

	return fileSystems, nil
}

// getCPUUsage gets the cpu usage
func (pr *PrometheusRepo) getCPUUsage(serviceName string) ([]*PrometheusData, error) {
	var query string

	switch pr.getPMMVersion() {
	case 1:
		// pmm 1.x
		query = PrometheusCPUUsageV1
	case 2:
		// pmm 2.x
		query = PrometheusCPUUsageV2
	default:
		return nil, message.NewMessage(msghc.ErrPmmVersionInvalid)
	}

	query = fmt.Sprintf(query, serviceName, serviceName, serviceName, serviceName, serviceName, serviceName)
	log.Debugf("healthcheck PrometheusRepo.checkCPUUsage() query: \n%s\n", query)

	return pr.execute(query)
}

// getIOUtil gets the io util
func (pr *PrometheusRepo) getIOUtil(serviceName string, devices []string) ([]*PrometheusData, error) {
	var query string

	devs := common.ConvertStringSliceToString(devices, constant.VerticalBarString)

	// prepare query
	switch pr.getPMMVersion() {
	case 1:
		// pmm 1.x
		query = fmt.Sprintf(PrometheusIOUtilV1, devs, serviceName, devs, serviceName)
	case 2:
		// pmm 2.x
		query = fmt.Sprintf(PrometheusIOUtilV2, devs, serviceName, devs, serviceName, devs, serviceName, devs, serviceName)
	default:
		return nil, message.NewMessage(msghc.ErrPmmVersionInvalid)
	}

	log.Debugf("healthcheck PrometheusRepo.getIOUtil() query: \n%s\n", query)
	// get data
	return pr.execute(query)
}

// getDiskCapacityUsage gets the disk capacity usage
func (pr *PrometheusRepo) getDiskCapacityUsage(serviceName string, mountPoints []string) ([]*PrometheusData, error) {
	var query string

	mps := common.ConvertStringSliceToString(mountPoints, constant.VerticalBarString)

	// prepare query
	switch pr.getPMMVersion() {
	case 1:
		// pmm 1.x
		query = fmt.Sprintf(PrometheusDiskCapacityV1, serviceName, mps, serviceName, mps)
	case 2:
		// pmm 2.x
		query = fmt.Sprintf(PrometheusDiskCapacityV2, serviceName, mps, serviceName, mps, serviceName, mps, serviceName, mps)
	default:
		return nil, message.NewMessage(msghc.ErrPmmVersionInvalid)
	}

	log.Debugf("healthcheck PrometheusRepo.getDiskCapacityUsage() query: \n%s\n", query)
	// get data
	return pr.execute(query)
}

// getConnectionUsage gets the connection usage
func (pr *PrometheusRepo) getConnectionUsage(serviceName string) ([]*PrometheusData, error) {
	return nil, nil
}

// getActiveSessionNum gets the active session number
func (pr *PrometheusRepo) getActiveSessionNum(serviceName string) ([]*PrometheusData, error) {
	return nil, nil
}

// getCacheMissRatio gets the cache miss ratio
func (pr *PrometheusRepo) getCacheMissRatio(serviceName string) ([]*PrometheusData, error) {
	return nil, nil
}

// execute executes the given query
func (pr *PrometheusRepo) execute(query string) ([]*PrometheusData, error) {
	// execute query
	result, err := pr.conn.Execute(query, pr.GetOperationInfo().GetStartTime(), pr.GetOperationInfo().GetEndTime(), pr.GetOperationInfo().GetStep())
	if err != nil {
		return nil, err
	}
	// parse result
	var datas []*PrometheusData

	matrix, err := result.Raw.GetMatrix()
	if err != nil {
		return nil, err
	}
	for _, sampleStream := range matrix {
		for _, samplePair := range sampleStream.Values {
			datas = append(datas, NewPrometheusData(samplePair.Timestamp.String(), float64(samplePair.Value)))
		}
	}

	return datas, nil
}
