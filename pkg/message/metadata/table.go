package metadata

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initDebugTableMessage()
	initInfoTableMessage()
	initErrorTableMessage()
}

// Message Code for metadata-table
const (
	// debug
	DebugMetadataGetTablesByDBID                              = 101801
	DebugMetadataGetStatisticsByDBIDAndTableName              = 101802
	DebugMetadataGetStatisticsByHostInfoAndDBNameAndTableName = 101803
	DebugMetadataAnalyzeTableByDBIDAndTableName               = 101804
	DebugMetadataAnalyzeTableByHostInfoAndDBNameAndTableName  = 101805
	// info
	InfoMetadataGetTablesByDBID                              = 201801
	InfoMetadataGetStatisticsByDBIDAndTableName              = 201802
	InfoMetadataGetStatisticsByHostInfoAndDBNameAndTableName = 201803
	InfoMetadataAnalyzeTableByDBIDAndTableName               = 201804
	InfoMetadataAnalyzeTableByHostInfoAndDBNameAndTableName  = 201805
	// error
	ErrMetadataGetTablesByDBID                              = 401801
	ErrMetadataGetStatisticsByDBIDAndTableName              = 401802
	ErrMetadataGetStatisticsByHostInfoAndDBNameAndTableName = 401803
	ErrMetadataAnalyzeTableByDBIDAndTableName               = 401804
	ErrMetadataAnalyzeTableByHostInfoAndDBNameAndTableName  = 401805
	ErrMetadataTableCreateApplicationMySQLConn              = 401806
)

func initDebugTableMessage() {
	message.Messages[DebugMetadataGetTablesByDBID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetTablesByDBID, "metadata: get tables by db id completed. message: %s")
	message.Messages[DebugMetadataGetStatisticsByDBIDAndTableName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetStatisticsByDBIDAndTableName, "metadata: get statistics by db id and table name completed. message: %s")
	message.Messages[DebugMetadataGetStatisticsByHostInfoAndDBNameAndTableName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetStatisticsByHostInfoAndDBNameAndTableName, "metadata: get statistics by host info and db name and table name completed. message: %s")
	message.Messages[DebugMetadataAnalyzeTableByDBIDAndTableName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAnalyzeTableByDBIDAndTableName, "metadata: analyze table by db id and table name completed. message: %s")
	message.Messages[DebugMetadataAnalyzeTableByHostInfoAndDBNameAndTableName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAnalyzeTableByHostInfoAndDBNameAndTableName, "metadata: analyze table by host info and db name and table name completed. message: %s")
}

func initInfoTableMessage() {
	message.Messages[InfoMetadataGetTablesByDBID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetTablesByDBID, "metadata: get tables by db id completed. db_id: %d, login_name: %s")
	message.Messages[InfoMetadataGetStatisticsByDBIDAndTableName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetStatisticsByDBIDAndTableName, "metadata: get tables by db id and table name completed. db_id: %d, table_name: %s, login_name: %s")
	message.Messages[InfoMetadataGetStatisticsByHostInfoAndDBNameAndTableName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetStatisticsByHostInfoAndDBNameAndTableName, "metadata: get tables by host info and db name and table name completed. host_ip: %s, port_num: %d, db_name: %s, table_name: %s, login_name %s")
	message.Messages[InfoMetadataAnalyzeTableByDBIDAndTableName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAnalyzeTableByDBIDAndTableName, "metadata: analyze table by db id and table name completed. db_id: %d, table_name: %s, login_name: %s")
	message.Messages[InfoMetadataAnalyzeTableByHostInfoAndDBNameAndTableName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAnalyzeTableByHostInfoAndDBNameAndTableName, "metadata: analyze table by host info and db name and table name completed. host_ip: %s, port_num: %d, db_name: %s, table_name: %s, login_name: %s")
}

func initErrorTableMessage() {
	message.Messages[ErrMetadataGetTablesByDBID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetTablesByDBID, "metadata: get tables by db id failed. db_id: %d, login_name: %s")
	message.Messages[ErrMetadataGetStatisticsByDBIDAndTableName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetStatisticsByDBIDAndTableName, "metadata: get statistics by db id and table name failed. db_id: %d, table_name: % s, login_name:%s")
	message.Messages[ErrMetadataGetStatisticsByHostInfoAndDBNameAndTableName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetStatisticsByHostInfoAndDBNameAndTableName, "metadata: get statistics by host info and db name and table name failed. host_ip: %s, port_num: %d, db_name: %s, table_name: %s, login_name: %s")
	message.Messages[ErrMetadataAnalyzeTableByDBIDAndTableName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAnalyzeTableByDBIDAndTableName, "metadata: analyze table by db id and table name failed. db_id: %d, table_name: %s, login_name: %s")
	message.Messages[ErrMetadataAnalyzeTableByHostInfoAndDBNameAndTableName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAnalyzeTableByHostInfoAndDBNameAndTableName, "metadata: analyze table by host info and db name and table name failed. host_ip: %s, port_num: %d, db_name: %s, table_name: %s, login_name: %s")
	message.Messages[ErrMetadataTableCreateApplicationMySQLConn] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataTableCreateApplicationMySQLConn, "metadata: create application mysql connection failed. db id: %d")
}
