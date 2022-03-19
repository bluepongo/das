package healthcheck

import (
	"strings"
	"time"

	"github.com/romberli/das/internal/dependency/healthcheck"
	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/pkg/message"
	msghc "github.com/romberli/das/pkg/message/healthcheck"
	"github.com/romberli/go-util/constant"
)

const (
	dbConfigMaxUserConnection         = "max_user_connection"
	dbConfigSlaveParallelWorkers      = "slave_parallel_workers"
	dbConfigLogBin                    = "log_bin"
	dbConfigBinlogFormat              = "binlog_format"
	dbConfigBinlogRowImage            = "binlog_row_image"
	dbConfigSyncBinlog                = "sync_binlog"
	dbConfigInnodbFlushLogAtTrxCommit = "innodb_flush_log_at_trx_commit"
	dbConfigGTIDMode                  = "gtid_mode"
	dbConfigEnforceGTIDConsistency    = "enforce_gtid_consistency"
	dbConfigSlaveParallelType         = "slave_parallel_type"
	dbConfigMasterInfoRepository      = "master_info_repository"
	dbConfigRelayLogInfoRepository    = "relay_log_info_repository"
	dbConfigReportHost                = "report_host"
	dbConfigReportPort                = "report_port"
	dbConfigInnodbFlushMethod         = "innodb_flush_method"
	dbConfigInnodbMonitorEnable       = "innodb_monitor_enable"
	dbConfigInnodbPrintAllDeadlocks   = "innodb_print_all_deadlocks"
	dbConfigSlowQueryLog              = "slow_query_log"
	dbConfigPerformanceSchema         = "performance_schema"

	dbConfigMaxUserConnectionValid         = 2000
	dbConfigSlaveParallelWorkersValid      = 16
	dbConfigLogBinValid                    = "ON"
	dbConfigBinlogFormatValid              = "ROW"
	dbConfigBinlogRowImageValid            = "FULL"
	dbConfigSyncBinlogValid                = "1"
	dbConfigInnodbFlushLogAtTrxCommitValid = "1"
	dbConfigGTIDModeValid                  = "ON"
	dbConfigEnforceGTIDConsistencyValid    = "ON"
	dbConfigSlaveParallelTypeValid         = "LOGICAL_CLOCK"
	dbConfigMasterInfoRepositoryValid      = "TABLE"
	dbConfigRelayLogInfoRepositoryValid    = "TABLE"
	dbConfigReportHostValid                = constant.EmptyString
	dbConfigReportPortValid                = constant.EmptyString
	dbConfigInnodbFlushMethodValid         = "O_DIRECT"
	dbConfigInnodbMonitorEnableValid       = "ALL"
	dbConfigInnodbPrintAllDeadlocksValid   = "ON"
	dbConfigSlowQueryLogValid              = "ON"
	dbConfigPerformanceSchemaValid         = "ON"
)

var (
	_ healthcheck.OperationInfo  = (*OperationInfo)(nil)
	_ healthcheck.Variable       = (*GlobalVariable)(nil)
	_ healthcheck.Table          = (*Table)(nil)
	_ healthcheck.PrometheusData = (*PrometheusData)(nil)

	dbConfigVariableNames = map[string]string{
		dbConfigLogBin:                    dbConfigLogBinValid,
		dbConfigBinlogFormat:              dbConfigBinlogFormatValid,
		dbConfigBinlogRowImage:            dbConfigBinlogRowImageValid,
		dbConfigSyncBinlog:                dbConfigSyncBinlogValid,
		dbConfigInnodbFlushLogAtTrxCommit: dbConfigInnodbFlushLogAtTrxCommitValid,
		dbConfigGTIDMode:                  dbConfigGTIDModeValid,
		dbConfigEnforceGTIDConsistency:    dbConfigEnforceGTIDConsistencyValid,
		dbConfigSlaveParallelType:         dbConfigSlaveParallelTypeValid,
		dbConfigMasterInfoRepository:      dbConfigMasterInfoRepositoryValid,
		dbConfigRelayLogInfoRepository:    dbConfigRelayLogInfoRepositoryValid,
		dbConfigReportHost:                dbConfigReportHostValid,
		dbConfigReportPort:                dbConfigReportPortValid,
		dbConfigInnodbFlushMethod:         dbConfigInnodbFlushMethodValid,
		// dbConfigInnodbMonitorEnable:       dbConfigInnodbMonitorEnableValid,
		dbConfigInnodbPrintAllDeadlocks: dbConfigInnodbPrintAllDeadlocksValid,
		dbConfigSlowQueryLog:            dbConfigSlowQueryLogValid,
		dbConfigPerformanceSchema:       dbConfigPerformanceSchemaValid,
	}
)

type OperationInfo struct {
	operationID   int
	apps          []metadata.App
	mysqlServer   metadata.MySQLServer
	monitorSystem metadata.MonitorSystem
	startTime     time.Time
	endTime       time.Time
	step          time.Duration
}

// NewOperationInfo returns a new *OperationInfo
func NewOperationInfo(operationID int, apps []metadata.App, mysqlServer metadata.MySQLServer, MonitorSystem metadata.MonitorSystem, startTime, endTime time.Time, step time.Duration) *OperationInfo {
	return &OperationInfo{
		operationID:   operationID,
		apps:          apps,
		mysqlServer:   mysqlServer,
		monitorSystem: MonitorSystem,
		startTime:     startTime,
		endTime:       endTime,
		step:          step,
	}
}

// GetOperationID returns the operation identity
func (oi *OperationInfo) GetOperationID() int {
	return oi.operationID
}

// GetApps returns the apps
func (oi *OperationInfo) GetApps() []metadata.App {
	return oi.apps
}

// GetMySQLServer returns the mysql server
func (oi *OperationInfo) GetMySQLServer() metadata.MySQLServer {
	return oi.mysqlServer
}

// GetMonitorSystem returns the monitor system
func (oi *OperationInfo) GetMonitorSystem() metadata.MonitorSystem {
	return oi.monitorSystem
}

// GetStartTime returns the start time
func (oi *OperationInfo) GetStartTime() time.Time {
	return oi.startTime
}

// GetEndTime returns the end time
func (oi *OperationInfo) GetEndTime() time.Time {
	return oi.endTime
}

// GetStep returns the step
func (oi *OperationInfo) GetStep() time.Duration {
	return oi.step
}

// GetAppName returns the app name in string, it will concat all the app names with comma
func (oi *OperationInfo) GetAppName() string {
	var name string

	for _, app := range oi.GetApps() {
		name += app.GetAppName() + constant.CommaString
	}

	return strings.Trim(name, constant.CommaString)
}

type OperationHistory struct {
	ID             int       `middleware:"id" json:"id"`
	MySQLServerID  int       `middleware:"mysql_server_id" json:"mysql_server_id"`
	HostIP         string    `middleware:"host_ip" json:"host_ip"`
	PortNum        int       `middleware:"port_num" json:"port_num"`
	StartTime      time.Time `middleware:"start_time" json:"start_time"`
	EndTime        time.Time `middleware:"end_time" json:"end_time"`
	Step           int       `middleware:"step" json:"step"`
	Status         int       `middleware:"status" json:"status"`
	Message        string    `middleware:"message" json:"message"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

func NewOperationHistory(id, mySQLServerID int, hostIP string, portNum int, startTime, endTime time.Time, step, status int, message string,
	delFlag int, createTime, lastUpdateTime time.Time) healthcheck.OperationHistory {
	return newOperationHistory(id, mySQLServerID, hostIP, portNum, startTime, endTime, step, status, message, delFlag, createTime, lastUpdateTime)
}

func NewEmptyOperationHistory() healthcheck.OperationHistory {
	return &OperationHistory{}
}

func newOperationHistory(id, mysqlServerID int, hostIP string, portNum int, startTime, endTime time.Time, step, status int, message string,
	delFlag int, createTime, lastUpdateTime time.Time) *OperationHistory {
	return &OperationHistory{
		ID:             id,
		MySQLServerID:  mysqlServerID,
		HostIP:         hostIP,
		PortNum:        portNum,
		StartTime:      startTime,
		EndTime:        endTime,
		Step:           step,
		Status:         status,
		Message:        message,
		DelFlag:        delFlag,
		CreateTime:     createTime,
		LastUpdateTime: lastUpdateTime,
	}
}

// GetID returns the identity
func (oh *OperationHistory) GetID() int {
	return oh.ID
}

// GetMySQLServerID returns the mysql server id
func (oh *OperationHistory) GetMySQLServerID() int {
	return oh.MySQLServerID
}

// GetHostIP returns the host ip of mysql server
func (oh *OperationHistory) GetHostIP() string {
	return oh.HostIP
}

// GetPortNum returns the port number of mysql server
func (oh *OperationHistory) GetPortNum() int {
	return oh.PortNum
}

// GetStartTime returns the start time
func (oh *OperationHistory) GetStartTime() time.Time {
	return oh.StartTime
}

// GetEndTime returns the end time
func (oh *OperationHistory) GetEndTime() time.Time {
	return oh.EndTime
}

// GetStep returns the step, the unit is seconds
func (oh *OperationHistory) GetStep() int {
	return oh.Step
}

// GetStatus returns the status
func (oh *OperationHistory) GetStatus() int {
	return oh.Status
}

// GetMessage returns the message
func (oh *OperationHistory) GetMessage() string {
	return oh.Message
}

// GetDelFlag returns the delete flag
func (oh *OperationHistory) GetDelFlag() int {
	return oh.DelFlag
}

// GetCreateTime returns the create time
func (oh *OperationHistory) GetCreateTime() time.Time {
	return oh.CreateTime
}

// GetLastUpdateTime returns the last update time
func (oh *OperationHistory) GetLastUpdateTime() time.Time {
	return oh.LastUpdateTime
}

// DefaultItemConfig include all data for a item
type DefaultItemConfig struct {
	ID                          int       `middleware:"id" json:"id"`
	ItemName                    string    `middleware:"item_name" json:"item_name"`
	ItemWeight                  int       `middleware:"item_weight" json:"item_weight"`
	LowWatermark                float64   `middleware:"low_watermark" json:"low_watermark"`
	HighWatermark               float64   `middleware:"high_watermark" json:"high_watermark"`
	Unit                        float64   `middleware:"unit" json:"unit"`
	ScoreDeductionPerUnitHigh   float64   `middleware:"score_deduction_per_unit_high" json:"score_deduction_per_unit_high"`
	MaxScoreDeductionHigh       float64   `middleware:"max_score_deduction_high" json:"max_score_deduction_high"`
	ScoreDeductionPerUnitMedium float64   `middleware:"score_deduction_per_unit_medium" json:"score_deduction_per_unit_medium"`
	MaxScoreDeductionMedium     float64   `middleware:"max_score_deduction_medium" json:"max_score_deduction_medium"`
	DelFlag                     int       `middleware:"del_flag" json:"del_flag"`
	CreateTime                  time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime              time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewDefaultItemConfig returns new *DefaultItemConfig
func NewDefaultItemConfig(itemName string, itemWeight int, lowWatermark float64, highWatermark float64, unit float64,
	scoreDeductionPerUnitHigh float64, maxScoreDeductionHigh float64, scoreDeductionPerUnitMedium float64, maxScoreDeductionMedium float64) *DefaultItemConfig {
	return &DefaultItemConfig{
		ItemName:                    itemName,
		ItemWeight:                  itemWeight,
		LowWatermark:                lowWatermark,
		HighWatermark:               highWatermark,
		Unit:                        unit,
		ScoreDeductionPerUnitHigh:   scoreDeductionPerUnitHigh,
		MaxScoreDeductionHigh:       maxScoreDeductionHigh,
		ScoreDeductionPerUnitMedium: scoreDeductionPerUnitMedium,
		MaxScoreDeductionMedium:     maxScoreDeductionMedium,
	}
}

// NewEmptyDefaultItemConfig returns a new *DefaultItemConfig
func NewEmptyDefaultItemConfig() *DefaultItemConfig {
	return &DefaultItemConfig{}
}

// GetID returns the identity
func (dic *DefaultItemConfig) GetID() int {
	return dic.ID
}

// GetItemName returns the item name
func (dic *DefaultItemConfig) GetItemName() string {
	return dic.ItemName
}

// GetItemWeight returns the item weight
func (dic *DefaultItemConfig) GetItemWeight() int {
	return dic.ItemWeight
}

// GetLowWatermark returns the low watermark
func (dic *DefaultItemConfig) GetLowWatermark() float64 {
	return dic.LowWatermark
}

// GetHighWatermark returns the high watermark
func (dic *DefaultItemConfig) GetHighWatermark() float64 {
	return dic.HighWatermark
}

// GetUnit returns the unit
func (dic *DefaultItemConfig) GetUnit() float64 {
	return dic.Unit
}

// GetScoreDeductionPerUnitHigh returns the score deduction per unit high
func (dic *DefaultItemConfig) GetScoreDeductionPerUnitHigh() float64 {
	return dic.ScoreDeductionPerUnitHigh
}

// GetMaxScoreDeductionHigh returns the max score deduction high
func (dic *DefaultItemConfig) GetMaxScoreDeductionHigh() float64 {
	return dic.ScoreDeductionPerUnitHigh
}

// GetScoreDeductionPerUnitMedium returns the score deduction per unit medium
func (dic *DefaultItemConfig) GetScoreDeductionPerUnitMedium() float64 {
	return dic.ScoreDeductionPerUnitMedium
}

// GetMaxScoreDeductionMedium returns the max score deduction medium
func (dic *DefaultItemConfig) GetMaxScoreDeductionMedium() float64 {
	return dic.MaxScoreDeductionMedium
}

// GetDelFlag returns the delete flag
func (dic *DefaultItemConfig) GetDelFlag() int {
	return dic.DelFlag
}

// GetCreateTime returns the create time
func (dic *DefaultItemConfig) GetCreateTime() time.Time {
	return dic.CreateTime
}

// GetLastUpdateTime returns the last update time
func (dic *DefaultItemConfig) GetLastUpdateTime() time.Time {
	return dic.LastUpdateTime
}

// DefaultEngineConfig is a map of DefaultItemConfig
type DefaultEngineConfig map[string]*DefaultItemConfig

// NewEmptyDefaultEngineConfig returns a new empty *DefaultItemConfig
func NewEmptyDefaultEngineConfig() DefaultEngineConfig {
	return map[string]*DefaultItemConfig{}
}

// GetItemConfig returns healthcheck.ItemConfig with given item name
func (dec DefaultEngineConfig) GetItemConfig(item string) healthcheck.ItemConfig {
	return dec.getItemConfig(item)
}

// getItemConfig returns *DefaultItemConfig with given item name
func (dec DefaultEngineConfig) getItemConfig(item string) *DefaultItemConfig {
	return dec[item]
}

// Validate validates if engine configuration is valid
func (dec DefaultEngineConfig) Validate() error {
	itemWeightSummary := constant.ZeroInt
	// validate defaultEngineConfig exits items
	if len(dec) == constant.ZeroInt {
		return message.NewMessage(msghc.ErrHealthcheckDefaultEngineEmpty)
	}
	for itemName, defaultItemConfig := range dec {
		// validate item weight
		if defaultItemConfig.ItemWeight > defaultHundred || defaultItemConfig.ItemWeight < constant.ZeroInt {
			return message.NewMessage(msghc.ErrHealthcheckItemWeightItemInvalid, itemName, defaultItemConfig.ItemWeight)
		}
		// validate low watermark
		if defaultItemConfig.LowWatermark < constant.ZeroInt {
			return message.NewMessage(msghc.ErrHealthcheckLowWatermarkItemInvalid, itemName, defaultItemConfig.LowWatermark)
		}
		// validate high watermark
		if defaultItemConfig.HighWatermark < defaultItemConfig.LowWatermark {
			return message.NewMessage(msghc.ErrHealthcheckHighWatermarkItemInvalid, itemName, defaultItemConfig.HighWatermark)
		}
		// validate unit
		if defaultItemConfig.Unit < constant.ZeroInt {
			return message.NewMessage(msghc.ErrHealthcheckUnitItemInvalid, itemName, defaultItemConfig.Unit)
		}
		// validate score deduction per unit high
		if defaultItemConfig.ScoreDeductionPerUnitHigh > defaultHundred || defaultItemConfig.ScoreDeductionPerUnitHigh < constant.ZeroInt || defaultItemConfig.ScoreDeductionPerUnitHigh > defaultItemConfig.MaxScoreDeductionHigh {
			return message.NewMessage(msghc.ErrHealthcheckScoreDeductionPerUnitHighItemInvalid, itemName, defaultItemConfig.ScoreDeductionPerUnitHigh)
		}
		// validate max score deduction high
		if defaultItemConfig.MaxScoreDeductionHigh > defaultHundred || defaultItemConfig.MaxScoreDeductionHigh < constant.ZeroInt {
			return message.NewMessage(msghc.ErrHealthcheckMaxScoreDeductionHighItemInvalid, itemName, defaultItemConfig.MaxScoreDeductionHigh)
		}
		// validate score deduction per unit medium
		if defaultItemConfig.ScoreDeductionPerUnitMedium > defaultHundred || defaultItemConfig.ScoreDeductionPerUnitMedium < constant.ZeroInt || defaultItemConfig.ScoreDeductionPerUnitMedium > defaultItemConfig.MaxScoreDeductionMedium {
			return message.NewMessage(msghc.ErrHealthcheckScoreDeductionPerUnitMediumItemInvalid, itemName, defaultItemConfig.ScoreDeductionPerUnitMedium)
		}
		// validate max score deduction medium
		if defaultItemConfig.MaxScoreDeductionMedium > defaultHundred || defaultItemConfig.MaxScoreDeductionMedium < constant.ZeroInt {
			return message.NewMessage(msghc.ErrHealthcheckMaxScoreDeductionMediumItemInvalid, itemName, defaultItemConfig.MaxScoreDeductionMedium)
		}
		itemWeightSummary += defaultItemConfig.ItemWeight
	}
	// validate item weigh count is 100
	if itemWeightSummary != defaultHundred {
		return message.NewMessage(msghc.ErrHealthcheckItemWeightSummaryInvalid, itemWeightSummary)
	}

	return nil
}

// GlobalVariable encapsulates k-v pairs for global variable
type GlobalVariable struct {
	VariableName  string `middleware:"variable_name" json:"variable_name"`
	VariableValue string `middleware:"variable_value" json:"variable_value"`
}

// NewEmptyGlobalVariable returns a new *GlobalVariables
func NewEmptyGlobalVariable() *GlobalVariable {
	return &GlobalVariable{}
}

// NewGlobalVariable returns a *GlobalVariable
func NewGlobalVariable(name, value string) *GlobalVariable {
	return &GlobalVariable{
		VariableName:  name,
		VariableValue: value,
	}
}

func (gv *GlobalVariable) GetName() string {
	return gv.VariableName
}

func (gv *GlobalVariable) GetValue() string {
	return gv.VariableValue
}

type Variable struct {
	Name   string `middleware:"name" json:"name"`
	Value  string `middleware:"value" json:"value"`
	Advice string `middleware:"advice" json:"advice"`
}

func NewVariable(variableName, currentValue, advice string) *Variable {
	return &Variable{
		Name:   variableName,
		Value:  currentValue,
		Advice: advice,
	}
}

type Table struct {
	TableSchema string  `middleware:"table_schema" json:"table_schema"`
	TableName   string  `middleware:"table_name" json:"table_name"`
	TableRows   int     `middleware:"table_rows" json:"table_rows"`
	TableSize   float64 `middleware:"table_size" json:"table_size"`
}

func NewTable(schema, name string, rows int, size float64) *Table {
	return &Table{
		TableSchema: schema,
		TableName:   name,
		TableRows:   rows,
		TableSize:   size,
	}
}

func NewEmptyTable() *Table {
	return &Table{}
}

func (t *Table) GetSchema() string {
	return t.TableSchema
}

func (t *Table) GetName() string {
	return t.TableName
}

func (t *Table) GetRows() int {
	return t.TableRows
}

func (t *Table) GetSize() float64 {
	return t.TableSize
}

type PrometheusData struct {
	Timestamp string  `middleware:"timestamp" json:"timestamp"`
	Value     float64 `middleware:"value" json:"value"`
}

func NewPrometheusData(ts string, value float64) *PrometheusData {
	return &PrometheusData{
		Timestamp: ts,
		Value:     value,
	}
}

func NewEmptyPrometheusData() *PrometheusData {
	return &PrometheusData{}
}

func (pd *PrometheusData) GetTimestamp() string {
	return pd.Timestamp
}

func (pd *PrometheusData) GetValue() float64 {
	return pd.Value
}

type FileSystem struct {
	MountPoint string `middleware:"mount_point" json:"mount_point"`
	Device     string `middleware:"device" json:"device"`
}

func NewFileSystem(mountPoint, device string) *FileSystem {
	return &FileSystem{
		MountPoint: mountPoint,
		Device:     device,
	}
}

func (fs *FileSystem) GetMountPoint() string {
	return fs.MountPoint
}

func (fs *FileSystem) GetDevice() string {
	return fs.Device
}
