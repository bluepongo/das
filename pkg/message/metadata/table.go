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

const (
	//debug
	DebugMetadataGetTablesByDBID                 = 100901
	DebugMetadataGetStatisticsByDBIDAndTableName = 100902
	//info
	InfoMetadataGetTablesByDBID                 = 200901
	InfoMetadataGetStatisticsByDBIDAndTableName = 200902
	//error
	ErrMetadataGetTablesByDBID                 = 400901
	ErrMetadataGetStatisticsByDBIDAndTableName = 400902
)

func initDebugTableMessage() {
	message.Messages[DebugMetadataGetTablesByDBID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetTablesByDBID, "metadata: get tables by db id completed. message: %s")
	message.Messages[DebugMetadataGetStatisticsByDBIDAndTableName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetStatisticsByDBIDAndTableName, "metadata: get tables by db id and table name completed. message: %s")
}

func initInfoTableMessage() {
	message.Messages[InfoMetadataGetTablesByDBID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetTablesByDBID, "metadata: get tables by db id completed. db_id: %d")
	message.Messages[InfoMetadataGetStatisticsByDBIDAndTableName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetStatisticsByDBIDAndTableName, "metadata: get tables by db id and table name completed. db_id: %d, table_name: %s")
}

func initErrorTableMessage() {
	message.Messages[ErrMetadataGetTablesByDBID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetTablesByDBID, "metadata: get tables by db id failed. db_id: %d")
	message.Messages[ErrMetadataGetStatisticsByDBIDAndTableName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetStatisticsByDBIDAndTableName, "metadata: get tables by db id and table name failed. db_id: %d, table_name: %s")
}
