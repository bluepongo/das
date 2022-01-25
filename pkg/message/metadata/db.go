package metadata

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/das/pkg/message"
)

func init() {
	initDebugDBMessage()
	initInfoDBMessage()
	initErrorDBMessage()
}

const (
	// debug
	DebugMetadataGetDBAll                  = 100201
	DebugMetadataGetDBByEnv                = 100202
	DebugMetadataGetDBByID                 = 100203
	DebugMetadataGetDBByNameAndClusterInfo = 100204
	DebugMetadataGetAppsByID               = 100205
	DebugMetadataAddDB                     = 100206
	DebugMetadataUpdateDB                  = 100207
	DebugMetadataDeleteDB                  = 100208
	DebugMetadataDBAddApp                  = 100209
	DebugMetadataDBDeleteApp               = 100210
	DebugMetadataDBAddUser                 = 100211
	DebugMetadataDBDeleteUser              = 100212
	// info
	InfoMetadataGetDBAll                  = 200201
	InfoMetadataGetDBByEnv                = 200202
	InfoMetadataGetDBByID                 = 200203
	InfoMetadataGetDBByNameAndClusterInfo = 200204
	InfoMetadataGetAppsByID               = 200205
	InfoMetadataAddDB                     = 200206
	InfoMetadataUpdateDB                  = 200207
	InfoMetadataDeleteDB                  = 200208
	InfoMetadataDBAddApp                  = 200209
	InfoMetadataDBDeleteApp               = 200210
	InfoMetadataDBAddUser                 = 200211
	InfoMetadataDBDeleteUser              = 200212
	// error
	ErrMetadataGetDBAll                  = 400201
	ErrMetadataGetDBByEnv                = 400202
	ErrMetadataGetDBByID                 = 400203
	ErrMetadataGetDBByNameAndClusterInfo = 400205
	ErrMetadataGetAppsByID               = 400204
	ErrMetadataAddDB                     = 400206
	ErrMetadataUpdateDB                  = 400207
	ErrMetadataDeleteDB                  = 400208
	ErrMetadataDBAddApp                  = 400209
	ErrMetadataDBDeleteApp               = 400210
	ErrMetadataDBAddUser                 = 400211
	ErrMetadataDBDeleteUser              = 400212
)

func initDebugDBMessage() {
	message.Messages[DebugMetadataGetDBAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDBAll, "metadata: get all databases completed. message: %s")
	message.Messages[DebugMetadataGetDBByEnv] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDBByEnv, "metadata: get databases by environment completed. message: %s")
	message.Messages[DebugMetadataGetDBByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDBByID, "metadata: get database by id completed. message: %s")
	message.Messages[DebugMetadataGetDBByNameAndClusterInfo] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDBByNameAndClusterInfo, "metadata: get database by name and cluster info completed. message: %s")
	message.Messages[DebugMetadataGetAppsByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAppsByID, "metadata: get app id list completed. message: %s")
	message.Messages[DebugMetadataAddDB] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddDB, "metadata: add new database completed. message: %s")
	message.Messages[DebugMetadataUpdateDB] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateDB, "metadata: update database completed. message: %s")
	message.Messages[DebugMetadataDeleteDB] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteDB, "metadata: delete database completed. message: %s")
	message.Messages[DebugMetadataDBAddApp] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDBAddApp, "metadata: add map of database and app completed. message: %s")
	message.Messages[DebugMetadataDBDeleteApp] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDBDeleteApp, "metadata: delete map of database and app completed. message: %s")
	message.Messages[DebugMetadataDBAddUser] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDBAddUser, "metadata: add map of database and user completed. message: %s")
	message.Messages[DebugMetadataDBDeleteUser] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDBDeleteUser, "metadata: delete map of database and user completed. message: %s")
}

func initInfoDBMessage() {
	message.Messages[InfoMetadataGetDBAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDBAll, "metadata: get database all completed")
	message.Messages[InfoMetadataGetDBByEnv] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDBByEnv, "metadata: get databases by environment completed. env_id: %d")
	message.Messages[InfoMetadataGetDBByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDBByID, "metadata: get database by id completed. id: %d")
	message.Messages[InfoMetadataGetDBByNameAndClusterInfo] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDBByNameAndClusterInfo, "metadata: get database by name and cluster info completed. db_name: %s, cluster_id: %d, cluster_type: %d")
	message.Messages[InfoMetadataGetAppsByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAppsByID, "metadata: get app id list completed. id: %d")
	message.Messages[InfoMetadataAddDB] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddDB, "metadata: add new database completed. db_name: %s, cluster_id: %d, cluster_type: %d, env_id: %d")
	message.Messages[InfoMetadataUpdateDB] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateDB, "metadata: update database completed. id: %d")
	message.Messages[InfoMetadataDeleteDB] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDeleteDB, "metadata: delete database completed. id: %d")
	message.Messages[InfoMetadataDBAddApp] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDBAddApp, "metadata: add map of database and app completed. db_id: %d, app_id: %d")
	message.Messages[InfoMetadataDBDeleteApp] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDBDeleteApp, "metadata: delete map of database and app completed. db_id: %d, app_id: %d")
	message.Messages[InfoMetadataDBAddUser] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDBAddUser, "metadata: add map of database and user completed. db_id: %d, user_id: %d")
	message.Messages[InfoMetadataDBDeleteUser] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDBDeleteUser, "metadata: delete map of database and user completed. db_id: %d, user_id: %d")
}

func initErrorDBMessage() {
	message.Messages[ErrMetadataGetDBAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetDBAll, "metadata: get all databases failed")
	message.Messages[ErrMetadataGetDBByEnv] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetDBByEnv, "metadata: get databases by environment failed. env_id: %d")
	message.Messages[ErrMetadataGetDBByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetDBByID, "metadata: get database by id failed. id: %d")
	message.Messages[ErrMetadataGetDBByNameAndClusterInfo] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetDBByNameAndClusterInfo, "metadata: get database by name and cluster info failed. db_name: %s, cluster_id: %d, cluster_type: %d, env_id: %d")
	message.Messages[ErrMetadataGetAppsByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAppsByID, "metadata: get app id list failed. id: %d")
	message.Messages[ErrMetadataAddDB] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddDB, "metadata: add new databases failed. db_name: %s, cluster_id: %d, cluster_type: %d, env_id: %d")
	message.Messages[ErrMetadataUpdateDB] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateDB, "metadata: update database failed. id: %d")
	message.Messages[ErrMetadataDeleteDB] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteDB, "metadata: delete database failed. id: %d")
	message.Messages[ErrMetadataDBAddApp] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDBAddApp, "metadata: add map of database and app failed. id: %d")
	message.Messages[ErrMetadataDBDeleteApp] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDBDeleteApp, "metadata: delete map of database and app failed. id: %d")
	message.Messages[ErrMetadataDBAddUser] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDBAddUser, "metadata: add map of database and user failed. id: %d")
	message.Messages[ErrMetadataDBDeleteUser] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDBDeleteUser, "metadata: delete map of database and user failed. id: %d")
}
