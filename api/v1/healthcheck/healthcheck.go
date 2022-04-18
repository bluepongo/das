package healthcheck

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/app/healthcheck"
	"github.com/romberli/das/pkg/message"
	msghealth "github.com/romberli/das/pkg/message/healthcheck"
	"github.com/romberli/das/pkg/resp"
	utilhealth "github.com/romberli/das/pkg/util/healthcheck"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
	"github.com/spf13/viper"
)

const (
	oneDayHours     = 24
	loginNameJSON   = "login_name"
	operationIDJSON = "operation_id"
	reviewJSON      = "review"

	healthcheckOperationHistoriesStruct = "OperationHistories"

	checkRespMessage           = `{"operation_id": %d, "message": "healthcheck started"}`
	checkByHostInfoRespMessage = `{"operation_id": %d, "message": "healthcheck by host info started"}`
	reviewAccuracyRespMessage  = `{"operation_id": %d, "message": "reviewed accuracy completed"}`
)

// @Tags	healthcheck
// @Summary get result by operation id
// @Accept	application/json
// @Param	login_name body string true "login name"
// @Produce application/json
// @Success 200 {string} string "{"operation_histories":[{"id":30,"mysql_server_id":1,"host_ip":"192.168.137.11","port_num":3306,"start_time":"2022-03-11T19:46:16+08:00","end_time":"2022-03-18T19:46:16+08:00","step":60,"status":2,"message":"healthcheck completed successfully. engine: default, operation_id: 30","del_flag":0,"create_time":"2022-03-18T19:46:16.215941+08:00","last_update_time":"2022-03-18T19:46:17.450918+08:00"}]}"
// @Router	/api/v1/healthcheck/history [get]
func GetOperationHistoriesByLoginName(c *gin.Context) {
	var fields map[string]string

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	err = json.Unmarshal(data, &fields)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	loginName, ok := fields[loginNameJSON]
	if !ok || loginName == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, loginNameJSON)
		return
	}
	// init service
	s := healthcheck.NewServiceWithDefault()
	// get healthcheck histories
	err = s.GetOperationHistoriesByLoginName(loginName)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckGetOperationHistoriesByLoginName, err, loginName)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(healthcheckOperationHistoriesStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msghealth.DebugHealthcheckGetOperationHistoriesByLoginName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msghealth.InfoHealthcheckGetOperationHistoriesByLoginName, loginName)
}

// @Tags	healthcheck
// @Summary get result by operation id
// @Accept	application/json
// @Param	id path int true "operation id"
// @Produce application/json
// @Success 200 {string} string "{"result":{"id":7,"operation_id":16,"host_ip":"192.168.137.11","port_num":3306,"weighted_average_score":99,"db_config_score":90,"db_config_data":"[{"variable_name":"binlog_format","variable_value":"ROW"},{"variable_name":"binlog_row_image","variable_value":"FULL"},{"variable_name":"enforce_gtid_consistency","variable_value":"ON"},{"variable_name":"gtid_mode","variable_value":"ON"},{"variable_name":"innodb_flush_log_at_trx_commit","variable_value":"2"},{"variable_name":"innodb_flush_method","variable_value":"O_DIRECT"},{"variable_name":"innodb_print_all_deadlocks","variable_value":"ON"},{"variable_name":"log_bin","variable_value":"ON"},{"variable_name":"master_info_repository","variable_value":"TABLE"},{"variable_name":"performance_schema","variable_value":"ON"},{"variable_name":"relay_log_info_repository","variable_value":"TABLE"},{"variable_name":"report_host","variable_value":"192.168.137.11"},{"variable_name":"report_port","variable_value":"3306"},{"variable_name":"slave_parallel_type","variable_value":"DATABASE"},{"variable_name":"slow_query_log","variable_value":"ON"},{"variable_name":"sync_binlog","variable_value":"0"}]","db_config_advice":[{"name":"innodb_flush_log_at_trx_commit","value":"2","advice":"1"},{"name":"slave_parallel_type","value":"DATABASE","advice":"LOGICAL_CLOCK"},{"name":"sync_binlog","value":"0","advice":"1"}],"avg_backup_failed_ratio_score":100,"avg_backup_failed_ratio_data":"null","avg_backup_failed_ratio_high":"null","statistics_failed_ratio_score":100,"statistics_failed_ratio_data":"null","statistics_failed_ratio_high":"null","cpu_usage_score":100,"cpu_usage_data":"null","cpu_usage_high":"null","io_util_score":100,"io_util_data":"null","io_util_high":"null","disk_capacity_usage_score":100,"disk_capacity_usage_data":"null","disk_capacity_usage_high":"null","connection_usage_score":100,"connection_usage_data":"null","connection_usage_high":"null","average_active_session_percents_score":100,"average_active_session_percents_data":"null","average_active_session_percents_high":"null","cache_miss_ratio_score":100,"cache_miss_ratio_data":"null","cache_miss_ratio_high":"null","table_rows_score":100,"table_rows_data":"[]","table_rows_high":"null","table_size_score":100,"table_size_data":"[]","table_size_high":"null","slow_query_score":100,"slow_query_data":"[]","slow_query_advice":"","accuracy_review":0,"del_flag":0,"create_time":"2022-03-04 16:32:40.331139 +0800 CST","last_update_time":"2022-03-04 16:32:40.331139 +0800 CST"}"
// @Router	/api/v1/healthcheck/result/:operation_id [get]
func GetResultByOperationID(c *gin.Context) {
	// get data
	operationIDStr := c.Param(operationIDJSON)
	if operationIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, operationIDJSON)
		return
	}
	operationID, err := strconv.Atoi(operationIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := healthcheck.NewServiceWithDefault()
	// get entities
	err = s.GetResultByOperationID(operationID)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckGetResultByOperationID, err)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msghealth.DebugHealthcheckGetResultByOperationID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msghealth.InfoHealthcheckGetResultByOperationID, operationID)
}

// @Tags healthcheck
// @Summary check health of the database
// @Accept	application/json
// @Param	server_id	body int	true "mysql server id"
// @Param	start_time	body string true "start time"
// @Param	end_time	body string true "end time"
// @Param	step		body string true "step"
// @Produce application/json
// @Success 200 {string} string "{"operation_id: 16", "message": "healthcheck started"}"
// @Router /api/v1/healthcheck/check [post]
func Check(c *gin.Context) {
	var rd *utilhealth.Check
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	startTime, err := time.ParseInLocation(constant.TimeLayoutSecond, rd.GetStartTime(), time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, errors.Trace(err), rd.GetStartTime())
		return
	}
	endTime, err := time.ParseInLocation(constant.TimeLayoutSecond, rd.GetEndTime(), time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, errors.Trace(err), rd.GetEndTime())
		return
	}

	checkRange := int(endTime.Sub(startTime).Hours() / oneDayHours)
	maxRange := viper.GetInt(config.HealthcheckMaxRangeKey)
	if checkRange > maxRange {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckCheckRange, checkRange, maxRange)
		return
	}

	minStartTime := time.Now().Add(-constant.Day * time.Duration(maxRange))
	if startTime.Before(minStartTime) {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckStartTime, startTime.Format(constant.TimeLayoutSecond), minStartTime.Format(constant.TimeLayoutSecond))
		return
	}

	step, err := time.ParseDuration(rd.GetStep())
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeDuration, errors.Trace(err), rd.GetStep())
		return
	}

	// init service
	s := healthcheck.NewServiceWithDefault()
	// check health
	operationID, err := s.Check(rd.GetServerID(), startTime, endTime, step, rd.GetLoginName())
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckCheck, err, operationID)
		return
	}

	log.Debug(message.NewMessage(msghealth.DebugHealthcheckCheck, operationID).Error())
	resp.ResponseOK(c, fmt.Sprintf(checkRespMessage, operationID), msghealth.InfoHealthcheckCheck, operationID)
}

// @Tags healthcheck
// @Summary check health of the database by host ip and port number
// @Accept	application/json
// @Param	host_ip		body string	true "mysql host ip"
// @Param	port_num	body int	true "mysql port number"
// @Param	start_time	body string true "start time"
// @Param	end_time	body string true "end time"
// @Param	step		body string true "step"
// @Produce application/json
// @Success 200 {string} string "{"operation_id: 18", "message": "healthcheck by host info started"}"
// @Router /api/v1/healthcheck/check/host-info [post]
func CheckByHostInfo(c *gin.Context) {
	var rd *utilhealth.CheckByHostInfo
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	startTime, err := time.ParseInLocation(constant.TimeLayoutSecond, rd.GetStartTime(), time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, errors.Trace(err), rd.GetStartTime())
		return
	}
	endTime, err := time.ParseInLocation(constant.TimeLayoutSecond, rd.GetEndTime(), time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, errors.Trace(err), rd.GetEndTime())
		return
	}

	checkRange := int(endTime.Sub(startTime).Hours() / oneDayHours)
	maxRange := viper.GetInt(config.HealthcheckMaxRangeKey)
	if checkRange > maxRange {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckCheckRange, checkRange, maxRange)
		return
	}

	step, err := time.ParseDuration(rd.GetStep())
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeDuration, errors.Trace(err), rd.GetStep())
		return
	}
	// init service
	s := healthcheck.NewServiceWithDefault()
	// get entities
	operationID, err := s.CheckByHostInfo(rd.GetHostIP(), rd.GetPortNum(), startTime, endTime, step, rd.GetLoginName())
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckCheckByHostInfo, err, operationID)
		return
	}

	log.Debug(message.NewMessage(msghealth.DebugHealthcheckCheckByHostInfo, operationID).Error())
	resp.ResponseOK(c, fmt.Sprintf(checkByHostInfoRespMessage, operationID), msghealth.InfoHealthcheckCheckByHostInfo, operationID)
}

// @Tags healthcheck
// @Summary update accuracy review
// @Accept  application/json
// @Param	operation_id	body int true "operation id"
// @Param	review			body int true "review"
// @Produce application/json
// @Success 200 {string} string "{"operation_id: 1", "message": "reviewed accuracy completed"}"
// @Router /api/v1/healthcheck/review [post]
func ReviewAccuracy(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	dataMap := make(map[string]int)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckCheck, errors.Trace(err))
		return
	}
	operationID, operationIDExists := dataMap[operationIDJSON]
	if !operationIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, operationIDJSON)
		return
	}
	review, reviewExists := dataMap[reviewJSON]
	if !reviewExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, reviewJSON)
		return
	}
	// init service
	s := healthcheck.NewServiceWithDefault()
	// review accuracy
	err = s.ReviewAccuracy(operationID, review)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckReviewAccuracy, err, operationID)
		return
	}

	log.Debug(message.NewMessage(msghealth.DebugHealthcheckReviewAccuracy, reviewAccuracyRespMessage).Error())
	resp.ResponseOK(c, fmt.Sprintf(reviewAccuracyRespMessage, operationID), msghealth.InfoHealthcheckReviewAccuracy, operationID)
}
