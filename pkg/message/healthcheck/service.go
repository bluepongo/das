package healthcheck

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initServiceDebugMessage()
	initServiceInfoMessage()
	initServiceErrorMessage()
}

const (
	// debug
	DebugHealthcheckGetOperationHistoriesByLoginName = 103101
	DebugHealthcheckGetResultByOperationID           = 103102
	DebugHealthcheckCheck                            = 103103
	DebugHealthcheckCheckByHostInfo                  = 103103
	DebugHealthcheckReviewAccuracy                   = 103104
	// info
	InfoHealthcheckGetOperationHistoriesByLoginName = 203101
	InfoHealthcheckGetResultByOperationID           = 203102
	InfoHealthcheckCheck                            = 203103
	InfoHealthcheckCheckByHostInfo                  = 203103
	InfoHealthcheckReviewAccuracy                   = 203104
	// error
	ErrHealthcheckCheckRange                        = 403101
	ErrHealthcheckStartTime                         = 403102
	ErrHealthcheckDefaultEngineRun                  = 403103
	ErrHealthcheckGetOperationHistoriesByLoginName  = 403104
	ErrHealthcheckGetResultByOperationID            = 403105
	ErrHealthcheckCheck                             = 403106
	ErrHealthcheckCheckByHostInfo                   = 403107
	ErrHealthcheckReviewAccuracy                    = 403108
	ErrHealthcheckCloseConnection                   = 403109
	ErrHealthcheckCreateApplicationMySQLConnection  = 403110
	ErrHealthcheckCreateMonitorMySQLConnection      = 403111
	ErrHealthcheckCreateMonitorClickhouseConnection = 403112
	ErrHealthcheckCreateMonitorPrometheusConnection = 403113
)

func initServiceDebugMessage() {
	message.Messages[DebugHealthcheckGetOperationHistoriesByLoginName] = config.NewErrMessage(
		message.DefaultMessageHeader, DebugHealthcheckGetOperationHistoriesByLoginName,
		"healthcheck: get operation histories by login name completed. message: %s")
	message.Messages[DebugHealthcheckGetResultByOperationID] = config.NewErrMessage(
		message.DefaultMessageHeader, DebugHealthcheckGetResultByOperationID,
		"healthcheck: get result by operation id completed. message: %s")
	message.Messages[DebugHealthcheckCheck] = config.NewErrMessage(
		message.DefaultMessageHeader, DebugHealthcheckCheck,
		"healthcheck: check started. operation id: %d")
	message.Messages[DebugHealthcheckCheckByHostInfo] = config.NewErrMessage(
		message.DefaultMessageHeader, DebugHealthcheckCheckByHostInfo,
		"healthcheck: check by host info started. operation id: %d")
	message.Messages[DebugHealthcheckReviewAccuracy] = config.NewErrMessage(
		message.DefaultMessageHeader, DebugHealthcheckReviewAccuracy,
		"healthcheck: review accuracy message: %s")
}

func initServiceInfoMessage() {
	message.Messages[InfoHealthcheckGetOperationHistoriesByLoginName] = config.NewErrMessage(
		message.DefaultMessageHeader, InfoHealthcheckGetOperationHistoriesByLoginName,
		"healthcheck: get operation histories by login name completed. login name: %s")
	message.Messages[InfoHealthcheckGetResultByOperationID] = config.NewErrMessage(
		message.DefaultMessageHeader, InfoHealthcheckGetResultByOperationID,
		"healthcheck: get result by operation id completed. operation_id: %d")
	message.Messages[InfoHealthcheckCheck] = config.NewErrMessage(
		message.DefaultMessageHeader, InfoHealthcheckCheck,
		"healthcheck: check started. operation id: %d")
	message.Messages[InfoHealthcheckCheckByHostInfo] = config.NewErrMessage(
		message.DefaultMessageHeader, InfoHealthcheckCheckByHostInfo,
		"healthcheck: check by host info started. operation id: %d")
	message.Messages[InfoHealthcheckReviewAccuracy] = config.NewErrMessage(
		message.DefaultMessageHeader, InfoHealthcheckReviewAccuracy,
		"healthcheck: review accuracy completed. operation id: %d")
}

func initServiceErrorMessage() {
	message.Messages[ErrHealthcheckCheckRange] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckCheckRange,
		"check range is larger than the maximum allowed range. check range: %d, allowed range: %d")
	message.Messages[ErrHealthcheckStartTime] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckStartTime,
		"start time is older than the minimum allowed time. start time: %s, minimum allowed time: %s")
	message.Messages[ErrHealthcheckDefaultEngineRun] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckDefaultEngineRun,
		"default engine run failed")
	message.Messages[ErrHealthcheckGetOperationHistoriesByLoginName] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckGetOperationHistoriesByLoginName,
		"healthcheck: get operation histories by login name failed. login name: %s")
	message.Messages[ErrHealthcheckGetResultByOperationID] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckGetResultByOperationID,
		"healthcheck: get result by operation id failed. operation id: %d")
	message.Messages[ErrHealthcheckCheck] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckCheck,
		"healthcheck: check failed. operation id: %d")
	message.Messages[ErrHealthcheckCheckByHostInfo] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckCheckByHostInfo,
		"healthcheck: check by host info failed. operation id: %d")
	message.Messages[ErrHealthcheckReviewAccuracy] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckReviewAccuracy,
		"healthcheck: review accuracy failed. operation id: %d")
	message.Messages[ErrHealthcheckCloseConnection] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckCloseConnection,
		"healthcheck: close middleware connection failed")
	message.Messages[ErrHealthcheckCreateApplicationMySQLConnection] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckCreateApplicationMySQLConnection,
		"create application mysql connection failed. addr: %s, user: %s")
	message.Messages[ErrHealthcheckCreateMonitorMySQLConnection] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckCreateMonitorMySQLConnection,
		"create monitor mysql connection failed. addr: %s, user: %s")
	message.Messages[ErrHealthcheckCreateMonitorClickhouseConnection] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckCreateMonitorClickhouseConnection,
		"create monitor clickhouse connection failed. addr: %s, user: %s")
	message.Messages[ErrHealthcheckCreateMonitorPrometheusConnection] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckCreateMonitorPrometheusConnection,
		"create prometheus connection failed. addr: %s, user: %s")
}
