package query

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initQueryDebugMessage()
	initQueryInfoMessage()
	initQueryErrorMessage()
}

const (
	// debug
	DebugQueryGetByMySQLClusterID = 103001
	DebugQueryGetByMySQLServerID  = 103002
	DebugQueryGetByHostInfo       = 103003
	DebugQueryGetByDBID           = 103004
	DebugQueryGetBySQLID          = 103005
	// info
	InfoQueryGetByMySQLClusterID = 203001
	InfoQueryGetByMySQLServerID  = 203002
	InfoQueryGetByHostInfo       = 203003
	InfoQueryGetByDBID           = 203004
	InfoQueryGetBySQLID          = 203005
	// error
	ErrQueryGetByMySQLClusterID               = 403001
	ErrQueryGetByMySQLServerID                = 403002
	ErrQueryGetByHostInfo                     = 403003
	ErrQueryGetByDBID                         = 403004
	ErrQueryGetBySQLID                        = 403005
	ErrQueryConfigNotValid                    = 403006
	ErrQueryMonitorSystemSystemType           = 403007
	ErrQueryCloseMonitorRepo                  = 403008
	ErrQueryCreateMonitorMysqlConnection      = 403009
	ErrQueryCreateMonitorClickhouseConnection = 403010
)

func initQueryDebugMessage() {
	message.Messages[DebugQueryGetByMySQLClusterID] = config.NewErrMessage(message.DefaultMessageHeader, DebugQueryGetByMySQLClusterID, "get by mysql cluster id. mysql_cluster_id: %d, message: %s")
	message.Messages[DebugQueryGetByMySQLServerID] = config.NewErrMessage(message.DefaultMessageHeader, DebugQueryGetByMySQLServerID, "get by mysql server id. mysql_server_id: %d, message: %s")
	message.Messages[DebugQueryGetByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, DebugQueryGetByHostInfo, "get by mysql server host info. host_ip: %s, port_num: %d, message: %s")
	message.Messages[DebugQueryGetByDBID] = config.NewErrMessage(message.DefaultMessageHeader, DebugQueryGetByDBID, "get by db id. db_id: %d, message: %s")
	message.Messages[DebugQueryGetBySQLID] = config.NewErrMessage(message.DefaultMessageHeader, DebugQueryGetBySQLID, "get by sql id. mysql_server_id: %d, sql_id: %s, message: %s")
}

func initQueryInfoMessage() {
	message.Messages[InfoQueryGetByMySQLClusterID] = config.NewErrMessage(message.DefaultMessageHeader, InfoQueryGetByMySQLClusterID, "get by mysql cluster id completed. mysql_cluster_id: %d")
	message.Messages[InfoQueryGetByMySQLServerID] = config.NewErrMessage(message.DefaultMessageHeader, InfoQueryGetByMySQLServerID, "get by mysql server id completed. mysql_server_id: %d")
	message.Messages[InfoQueryGetByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, DebugQueryGetByHostInfo, "get by mysql server host info completed. host_ip: %s, port_num: %d")
	message.Messages[InfoQueryGetByDBID] = config.NewErrMessage(message.DefaultMessageHeader, InfoQueryGetByDBID, "get by db id completed. db_id: %d.")
	message.Messages[InfoQueryGetBySQLID] = config.NewErrMessage(message.DefaultMessageHeader, InfoQueryGetBySQLID, "get by sql id completed. mysql_server_id: %d, sql_id: %s")
}

func initQueryErrorMessage() {
	message.Messages[ErrQueryGetByMySQLClusterID] = config.NewErrMessage(message.DefaultMessageHeader, ErrQueryGetByMySQLClusterID, "get by mysql cluster id failed. mysql_cluster_id: %d")
	message.Messages[ErrQueryGetByMySQLServerID] = config.NewErrMessage(message.DefaultMessageHeader, ErrQueryGetByMySQLServerID, "get by mysql server id failed. mysql_server_id: %d")
	message.Messages[ErrQueryGetByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, ErrQueryGetByHostInfo, "get by mysql server host info failed. host_ip: %s, port_num: %d")
	message.Messages[ErrQueryGetByDBID] = config.NewErrMessage(message.DefaultMessageHeader, ErrQueryGetByDBID, "get by db id failed. db_id: %d")
	message.Messages[ErrQueryGetBySQLID] = config.NewErrMessage(message.DefaultMessageHeader, ErrQueryGetBySQLID, "get by sql id failed. mysql_server_id: %d, sql_id: %s")
	message.Messages[ErrQueryConfigNotValid] = config.NewErrMessage(message.DefaultMessageHeader, ErrQueryConfigNotValid, "config is not valid. start_time: %s, end_time: %s, limit: %d")
	message.Messages[ErrQueryMonitorSystemSystemType] = config.NewErrMessage(message.DefaultMessageHeader, ErrQueryMonitorSystemSystemType, "monitor system type should be either 1 or 2, %d is not valid")
	message.Messages[ErrQueryCloseMonitorRepo] = config.NewErrMessage(message.DefaultMessageHeader, ErrQueryCloseMonitorRepo, "close monitor repo failed")
	message.Messages[ErrQueryCreateMonitorMysqlConnection] = config.NewErrMessage(message.DefaultMessageHeader, ErrQueryCreateMonitorMysqlConnection, "create monitor mysql connection failed. addr: %s, user: %s")
	message.Messages[ErrQueryCreateMonitorClickhouseConnection] = config.NewErrMessage(message.DefaultMessageHeader, ErrQueryCreateMonitorClickhouseConnection, "create monitor clickhouse connection failed. addr: %s, user: %s")
}
