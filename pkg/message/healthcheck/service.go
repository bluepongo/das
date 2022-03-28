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
	DebugHealthcheckGetOperationHistoriesByLoginName = 102101
	DebugHealthcheckGetResultByOperationID           = 102102
	DebugHealthcheckCheck                            = 102103
	DebugHealthcheckCheckByHostInfo                  = 102103
	DebugHealthcheckReviewAccuracy                   = 102104
	// info
	InfoHealthcheckGetOperationHistoriesByLoginName = 202101
	InfoHealthcheckGetResultByOperationID           = 202102
	InfoHealthcheckCheck                            = 202103
	InfoHealthcheckCheckByHostInfo                  = 202103
	InfoHealthcheckReviewAccuracy                   = 202104
	// error
	ErrHealthcheckCheckRange                        = 402101
	ErrHealthcheckStartTime                         = 402102
	ErrHealthcheckDefaultEngineRun                  = 402103
	ErrHealthcheckGetOperationHistoriesByLoginName  = 402104
	ErrHealthcheckGetResultByOperationID            = 402105
	ErrHealthcheckCheck                             = 402106
	ErrHealthcheckCheckByHostInfo                   = 402107
	ErrHealthcheckReviewAccuracy                    = 402108
	ErrHealthcheckCloseConnection                   = 402109
	ErrHealthcheckCreateApplicationMySQLConnection  = 402110
	ErrHealthcheckCreateMonitorMySQLConnection      = 402111
	ErrHealthcheckCreateMonitorClickhouseConnection = 402112
	ErrHealthcheckCreateMonitorPrometheusConnection = 402113
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
