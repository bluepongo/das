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
	DebugHealthcheckGetResultByOperationID = 102101
	DebugHealthcheckCheck                  = 102102
	DebugHealthcheckCheckByHostInfo        = 102103
	DebugHealthcheckReviewAccuracy         = 102104
	// info
	InfoHealthcheckGetResultByOperationID = 202101
	InfoHealthcheckCheck                  = 202102
	InfoHealthcheckCheckByHostInfo        = 202103
	InfoHealthcheckReviewAccuracy         = 202104
	// error
	ErrHealthcheckDefaultEngineRun                  = 402101
	ErrHealthcheckGetResultByOperationID            = 402102
	ErrHealthcheckCheck                             = 402103
	ErrHealthcheckCheckByHostInfo                   = 402104
	ErrHealthcheckReviewAccuracy                    = 402105
	ErrHealthcheckCloseConnection                   = 402106
	ErrHealthcheckCreateApplicationMySQLConnection  = 402107
	ErrHealthcheckCreateMonitorMySQLConnection      = 402108
	ErrHealthcheckCreateMonitorClickhouseConnection = 402109
	ErrHealthcheckCreateMonitorPrometheusConnection = 402110
)

func initServiceDebugMessage() {
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
	message.Messages[ErrHealthcheckDefaultEngineRun] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHealthcheckDefaultEngineRun,
		"default engine run failed")
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
