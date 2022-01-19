package metadata

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/das/pkg/message"
)

func init() {
	initDebugMySQLCLusterMessage()
	initInfoMySQLCLusterMessage()
	initErrorMySQLCLusterMessage()
}

// Message code
const (
	// debug
	DebugMetadataGetMySQLClusterAll     = 100701
	DebugMetadataGetMySQLClusterByEnv   = 100702
	DebugMetadataGetMySQLClusterByID    = 100703
	DebugMetadataGetMySQLClusterByName  = 100704
	DebugMetadataGetMySQLServers        = 100705
	DebugMetadataGetMasterServers       = 100706
	DebugMetadataGetDBs                 = 100707
	DebugMetadataGetUsers               = 100708
	DebugMetadataMySQLClusterAddUser    = 100709
	DebugMetadataMySQLClusterDeleteUser = 100710
	DebugMetadataGetAppUsers            = 100711
	DebugMetadataGetDBUsers             = 100712
	DebugMetadataGetAllUsers            = 100713
	DebugMetadataAddMySQLCluster        = 100714
	DebugMetadataUpdateMySQLCluster     = 100715
	DebugMetadataDeleteMySQLCluster     = 100716
	// debug
	InfoMetadataGetMySQLClusterAll     = 200701
	InfoMetadataGetMySQLClusterByEnv   = 200702
	InfoMetadataGetMySQLClusterByID    = 200703
	InfoMetadataGetMySQLClusterByName  = 200704
	InfoMetadataGetMySQLServers        = 200705
	InfoMetadataGetMasterServers       = 200706
	InfoMetadataGetDBs                 = 200707
	InfoMetadataGetUsers               = 200708
	InfoMetadataGetAppUsers            = 200709
	InfoMetadataMySQLClusterAddUser    = 100710
	InfoMetadataMySQLClusterDeleteUser = 100711
	InfoMetadataGetDBUsers             = 200712
	InfoMetadataGetAllUsers            = 200713
	InfoMetadataAddMySQLCluster        = 200714
	InfoMetadataUpdateMySQLCluster     = 200715
	InfoMetadataDeleteMySQLCluster     = 200716
	// error
	ErrMetadataGetMySQLClusterAll     = 400701
	ErrMetadataGetMySQLClusterByEnv   = 400702
	ErrMetadataGetMySQLClusterByID    = 400703
	ErrMetadataGetMySQLClusterByName  = 400704
	ErrMetadataGetMySQLServers        = 400705
	ErrMetadataGetMasterServers       = 400706
	ErrMetadataGetDBs                 = 400707
	ErrMetadataGetUsers               = 400708
	ErrMetadataMySQLClusterAddUser    = 100709
	ErrMetadataMySQLClusterDeleteUser = 100710
	ErrMetadataGetAppUsers            = 400711
	ErrMetadataGetDBUsers             = 400712
	ErrMetadataGetAllUsers            = 400713
	ErrMetadataAddMySQLCluster        = 400714
	ErrMetadataUpdateMySQLCluster     = 400715
	ErrMetadataDeleteMySQLCluster     = 400716
)

func initDebugMySQLCLusterMessage() {
	message.Messages[DebugMetadataGetMySQLClusterAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterAll,
		"metadata: get all mysql clusters message: %s")
	message.Messages[DebugMetadataGetMySQLClusterByEnv] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterByEnv,
		"metadata: get mysql cluster by env message: %s")
	message.Messages[DebugMetadataGetMySQLClusterByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterByID,
		"metadata: get mysql cluster by id message: %s")
	message.Messages[DebugMetadataGetMySQLClusterByName] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterByName,
		"metadata: get mysql cluster by name message: %s")
	message.Messages[DebugMetadataGetMySQLServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLServers,
		"metadata: get mysql servers from mysql cluster message: %s")
	message.Messages[DebugMetadataGetMasterServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMasterServers,
		"metadata: get master servers from mysql cluster message: %s")
	message.Messages[DebugMetadataGetDBs] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetDBs,
		"metadata: get databases from mysql cluster message: %s")
	message.Messages[DebugMetadataGetUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetUsers,
		"metadata: get users from mysql cluster message: %s")
	message.Messages[DebugMetadataMySQLClusterAddUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataMySQLClusterAddUser,
		"metadata: add new user for mysql cluster message: %s")
	message.Messages[DebugMetadataMySQLClusterDeleteUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataMySQLClusterDeleteUser,
		"metadata: delete users for mysql cluster message: %s")
	message.Messages[DebugMetadataGetUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetUsers,
		"metadata: get users from mysql cluster message: %s")
	message.Messages[DebugMetadataGetAppUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetAppUsers,
		"metadata: get app users from mysql cluster message: %s")
	message.Messages[DebugMetadataGetDBUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetDBUsers,
		"metadata: get database users from mysql cluster message: %s")
	message.Messages[DebugMetadataGetAllUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetAllUsers,
		"metadata: get all users from mysql cluster message: %s")
	message.Messages[DebugMetadataAddMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataAddMySQLCluster,
		"metadata: add new mysql cluster message: %s")
	message.Messages[DebugMetadataUpdateMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataUpdateMySQLCluster,
		"metadata: update mysql cluster message: %s")
	message.Messages[DebugMetadataDeleteMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataDeleteMySQLCluster,
		"metadata: delete mysql cluster message: %s")
}

func initInfoMySQLCLusterMessage() {
	message.Messages[InfoMetadataGetMySQLClusterAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLClusterAll,
		"metadata: get mysql cluster all completed")
	message.Messages[InfoMetadataGetMySQLClusterByEnv] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLClusterByEnv,
		"metadata: get mysql cluster by env completed. env_id: %s")
	message.Messages[InfoMetadataGetMySQLClusterByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLClusterByID,
		"metadata: get mysql cluster by id completed. id: %s")
	message.Messages[InfoMetadataGetMySQLClusterByName] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLClusterByName,
		"metadata: get mysql cluster by name completed. cluster_name: %s")
	message.Messages[InfoMetadataGetMySQLServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLServers,
		"metadata: get mysql servers from mysql cluster completed. id: %d")
	message.Messages[InfoMetadataGetMasterServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMasterServers,
		"metadata: get master servers from mysql cluster completed. id: %d")
	message.Messages[InfoMetadataGetDBs] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetDBs,
		"metadata: get databases from mysql cluster completed. id: %d")
	message.Messages[InfoMetadataGetUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetUsers,
		"metadata: get users from mysql cluster completed. id: %d")
	message.Messages[InfoMetadataMySQLClusterAddUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataMySQLClusterAddUser,
		"metadata: add user for mysql cluster completed. id: %d, user_id: %d")
	message.Messages[InfoMetadataMySQLClusterDeleteUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataMySQLClusterDeleteUser,
		"metadata: delete user for mysql cluster completed. id: %d, user_id: %d")
	message.Messages[InfoMetadataGetAppUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetAppUsers,
		"metadata: get app users from mysql cluster completed. id: %d")
	message.Messages[InfoMetadataGetDBUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetDBUsers,
		"metadata: get database users from mysql cluster completed. id: %d")
	message.Messages[InfoMetadataGetAllUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetAllUsers,
		"metadata: get all users from mysql cluster completed. id: %d")
	message.Messages[InfoMetadataAddMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataAddMySQLCluster,
		"metadata: add new mysql cluster completed. cluster_name: %s, env_id: %s")
	message.Messages[InfoMetadataUpdateMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataUpdateMySQLCluster,
		"metadata: update mysql cluster completed. id: %s")
	message.Messages[InfoMetadataDeleteMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataDeleteMySQLCluster,
		"metadata: delete mysql cluster completed. id: %s")
}

func initErrorMySQLCLusterMessage() {
	message.Messages[ErrMetadataGetMySQLClusterAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLClusterAll,
		"metadata: get all mysql cluster failed.")
	message.Messages[ErrMetadataGetMySQLClusterByEnv] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLClusterByEnv,
		"metadata: get mysql cluster by env failed. env_id: %d")
	message.Messages[ErrMetadataGetMySQLClusterByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLClusterByID,
		"metadata: get mysql cluster by id failed. id: %d")
	message.Messages[ErrMetadataGetMySQLClusterByName] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLClusterByName,
		"metadata: get mysql cluster by name failed. cluster_name: %s")
	message.Messages[ErrMetadataGetMySQLServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLServers,
		"metadata: get mysql servers from mysql cluster failed. id: %d")
	message.Messages[ErrMetadataGetMasterServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMasterServers,
		"metadata: get master servers from mysql cluster failed. id: %d")
	message.Messages[ErrMetadataGetDBs] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetDBs,
		"metadata: get databases from mysql cluster failed. id: %d")
	message.Messages[ErrMetadataGetUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetUsers,
		"metadata: get users from mysql cluster failed. id: %d")
	message.Messages[ErrMetadataMySQLClusterAddUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataMySQLClusterAddUser,
		"metadata: add user for mysql cluster failed. id: %d, user_id: %d")
	message.Messages[ErrMetadataMySQLClusterDeleteUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataMySQLClusterDeleteUser,
		"metadata: delete user for mysql cluster failed. id: %d, user_id: %d")
	message.Messages[ErrMetadataGetAppUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetAppUsers,
		"metadata: get app users from mysql cluster failed. id: %d")
	message.Messages[ErrMetadataGetDBUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetDBUsers,
		"metadata: get database users from mysql cluster failed. id: %d")
	message.Messages[ErrMetadataGetAllUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetAllUsers,
		"metadata: get all users from mysql cluster failed. id: %d")
	message.Messages[ErrMetadataAddMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataAddMySQLCluster,
		"metadata: add new mysql cluster failed. cluster_name: %s, env_id: %d")
	message.Messages[ErrMetadataUpdateMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataUpdateMySQLCluster,
		"metadata: update mysql cluster failed. id: %d")
	message.Messages[ErrMetadataDeleteMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataDeleteMySQLCluster,
		"metadata: delete mysql cluster failed. id: %d")
}
