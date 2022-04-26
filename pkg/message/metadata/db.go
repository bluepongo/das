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
	DebugMetadataGetDBAll                  = 101101
	DebugMetadataGetDBByEnv                = 101102
	DebugMetadataGetDBByID                 = 101103
	DebugMetadataGetDBByNameAndClusterInfo = 101104
	DebugMetadataGetDBByNameAndHostInfo    = 101105
	DebugMetadataGetDBsByHostInfo          = 101106
	DebugMetadataGetAppsByDBID             = 101107
	DebugMetadataGetMySQLClusterByDBID     = 101108
	DebugMetadataGetAppUsersByDBID         = 101109
	DebugMetadataGetUsersByDBID            = 101110
	DebugMetadataGetAllUsersByDBID         = 101111
	DebugMetadataAddDB                     = 101112
	DebugMetadataUpdateDB                  = 101113
	DebugMetadataDeleteDB                  = 101114
	DebugMetadataDBAddApp                  = 101115
	DebugMetadataDBDeleteApp               = 101116
	DebugMetadataDBAddUser                 = 101117
	DebugMetadataDBDeleteUser              = 101118
	// info
	InfoMetadataGetDBAll                  = 201101
	InfoMetadataGetDBByEnv                = 201102
	InfoMetadataGetDBByID                 = 201103
	InfoMetadataGetDBByNameAndClusterInfo = 201104
	InfoMetadataGetDBByNameAndHostInfo    = 201105
	InfoMetadataGetDBsByHostInfo          = 201106
	InfoMetadataGetAppsByDBID             = 201107
	InfoMetadataGetMySQLClusterByDBID     = 201108
	InfoMetadataGetAppUsersByDBID         = 201109
	InfoMetadataGetUsersByDBID            = 201110
	InfoMetadataGetAllUsersByDBID         = 201111
	InfoMetadataAddDB                     = 201112
	InfoMetadataUpdateDB                  = 201113
	InfoMetadataDeleteDB                  = 201114
	InfoMetadataDBAddApp                  = 201115
	InfoMetadataDBDeleteApp               = 201116
	InfoMetadataDBAddUser                 = 201117
	InfoMetadataDBDeleteUser              = 201118
	// error
	ErrMetadataGetDBAll                  = 401101
	ErrMetadataGetDBByEnv                = 401102
	ErrMetadataGetDBByID                 = 401103
	ErrMetadataGetDBByNameAndClusterInfo = 401104
	ErrMetadataGetDBByNameAndHostInfo    = 401105
	ErrMetadataGetDBsByHostInfo          = 401106
	ErrMetadataGetAppsByDBID             = 401107
	ErrMetadataGetMySQLClusterByDBID     = 401108
	ErrMetadataGetAppUsersByDBID         = 401109
	ErrMetadataGetUsersByDBID            = 401110
	ErrMetadataGetAllUsersByDBID         = 401111
	ErrMetadataAddDB                     = 401112
	ErrMetadataUpdateDB                  = 401113
	ErrMetadataDeleteDB                  = 401114
	ErrMetadataDBAddApp                  = 401115
	ErrMetadataDBDeleteApp               = 401116
	ErrMetadataDBAddUser                 = 401117
	ErrMetadataDBDeleteUser              = 401118
)

func initDebugDBMessage() {
	message.Messages[DebugMetadataGetDBAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDBAll, "metadata: get all databases. message: %s")
	message.Messages[DebugMetadataGetDBByEnv] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDBByEnv, "metadata: get databases by environment. message: %s")
	message.Messages[DebugMetadataGetDBByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDBByID, "metadata: get database by id. message: %s")
	message.Messages[DebugMetadataGetDBByNameAndClusterInfo] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDBByNameAndClusterInfo, "metadata: get database by name and cluster info. message: %s")
	message.Messages[DebugMetadataGetDBByNameAndHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDBByNameAndHostInfo, "metadata: get database by name and host info. message: %s")
	message.Messages[DebugMetadataGetDBsByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDBsByHostInfo, "metadata: get databases by host info. message: %s")
	message.Messages[DebugMetadataGetAppsByDBID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAppsByDBID, "metadata: get app id list. message: %s")
	message.Messages[DebugMetadataGetMySQLClusterByDBID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMySQLClusterByDBID, "metadata: get mysql cluster by id. message: %s")
	message.Messages[DebugMetadataGetAppUsersByDBID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAppUsersByDBID, "metadata: get app users by id. message: %s")
	message.Messages[DebugMetadataGetUsersByDBID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetUsersByDBID, "metadata: get users by id. message: %s")
	message.Messages[DebugMetadataGetAllUsersByDBID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAllUsersByDBID, "metadata: get all users by id. message: %s")
	message.Messages[DebugMetadataAddDB] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddDB, "metadata: add new database. message: %s")
	message.Messages[DebugMetadataUpdateDB] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateDB, "metadata: update database. message: %s")
	message.Messages[DebugMetadataDeleteDB] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteDB, "metadata: delete database. message: %s")
	message.Messages[DebugMetadataDBAddApp] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDBAddApp, "metadata: add map of database and app. message: %s")
	message.Messages[DebugMetadataDBDeleteApp] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDBDeleteApp, "metadata: delete map of database and app. message: %s")
	message.Messages[DebugMetadataDBAddUser] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDBAddUser, "metadata: add map of database and user. message: %s")
	message.Messages[DebugMetadataDBDeleteUser] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDBDeleteUser, "metadata: delete map of database and user. message: %s")
}

func initInfoDBMessage() {
	message.Messages[InfoMetadataGetDBAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDBAll, "metadata: get database all completed")
	message.Messages[InfoMetadataGetDBByEnv] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDBByEnv, "metadata: get databases by environment completed. env_id: %d")
	message.Messages[InfoMetadataGetDBByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDBByID, "metadata: get database by id completed. id: %d")
	message.Messages[InfoMetadataGetDBByNameAndClusterInfo] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDBByNameAndClusterInfo, "metadata: get database by name and cluster info completed. db_name: %s, cluster_id: %d, cluster_type: %d")
	message.Messages[InfoMetadataGetDBByNameAndHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDBByNameAndHostInfo, "metadata: get database by name and host info completed. db_name: %s, host_ip: %s, port_num: %d")
	message.Messages[InfoMetadataGetDBsByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDBsByHostInfo, "metadata: get databases by host info completed. host_ip: %s, port_num: %d")
	message.Messages[InfoMetadataGetAppsByDBID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAppsByDBID, "metadata: get app id list completed. id: %d")
	message.Messages[InfoMetadataGetMySQLClusterByDBID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMySQLClusterByDBID, "metadata: get mysql cluster by id completed. id: %d")
	message.Messages[InfoMetadataGetAppUsersByDBID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAppUsersByDBID, "metadata: get app users by id completed. id: %d")
	message.Messages[InfoMetadataGetUsersByDBID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetUsersByDBID, "metadata: get users by id completed. id: %d")
	message.Messages[InfoMetadataGetAllUsersByDBID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAllUsersByDBID, "metadata: get all users by id completed. id: %d")
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
	message.Messages[ErrMetadataGetDBByNameAndHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetDBByNameAndHostInfo, "metadata: get database by name and host info failed. db_name: %s, host_ip: %s, port_num: %d")
	message.Messages[ErrMetadataGetDBsByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetDBsByHostInfo, "metadata: get databases by host info failed. host_ip: %s, port_num: %d")
	message.Messages[ErrMetadataGetAppsByDBID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAppsByDBID, "metadata: get app id list failed. id: %d")
	message.Messages[ErrMetadataGetMySQLClusterByDBID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMySQLClusterByDBID, "metadata: get mysql cluster by id failed. id: %d")
	message.Messages[ErrMetadataGetAppUsersByDBID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAppUsersByDBID, "metadata: get app users by id failed. id: %d")
	message.Messages[ErrMetadataGetUsersByDBID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetUsersByDBID, "metadata: get users by id failed. id: %d")
	message.Messages[ErrMetadataGetAllUsersByDBID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAllUsersByDBID, "metadata: get all users by id failed. id: %d")
	message.Messages[ErrMetadataAddDB] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddDB, "metadata: add new databases failed. db_name: %s, cluster_id: %d, cluster_type: %d, env_id: %d")
	message.Messages[ErrMetadataUpdateDB] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateDB, "metadata: update database failed. id: %d")
	message.Messages[ErrMetadataDeleteDB] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteDB, "metadata: delete database failed. id: %d")
	message.Messages[ErrMetadataDBAddApp] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDBAddApp, "metadata: add map of database and app failed. id: %d")
	message.Messages[ErrMetadataDBDeleteApp] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDBDeleteApp, "metadata: delete map of database and app failed. id: %d")
	message.Messages[ErrMetadataDBAddUser] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDBAddUser, "metadata: add map of database and user failed. id: %d")
	message.Messages[ErrMetadataDBDeleteUser] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDBDeleteUser, "metadata: delete map of database and user failed. id: %d")
}
