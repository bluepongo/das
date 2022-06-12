package metadata

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/das/pkg/message"
)

func init() {
	initDebugMySQLClusterMessage()
	initInfoMySQLClusterMessage()
	initErrorMySQLClusterMessage()
}

// Message code
const (
	// debug
	DebugMetadataGetMySQLClusterAll     = 101601
	DebugMetadataGetMySQLClusterByEnv   = 101602
	DebugMetadataGetMySQLClusterByID    = 101603
	DebugMetadataGetMySQLClusterByName  = 101604
	DebugMetadataGetMySQLServers        = 101605
	DebugMetadataGetMasterServers       = 101606
	DebugMetadataGetDBs                 = 101607
	DebugMetadataGetUsers               = 101608
	DebugMetadataMySQLClusterAddUser    = 101609
	DebugMetadataMySQLClusterDeleteUser = 101610
	DebugMetadataGetAppUsers            = 101611
	DebugMetadataGetDBUsers             = 101612
	DebugMetadataGetAllUsers            = 101613
	DebugMetadataAddMySQLCluster        = 101614
	DebugMetadataUpdateMySQLCluster     = 101615
	DebugMetadataDeleteMySQLCluster     = 101616
	// debug
	InfoMetadataGetMySQLClusterAll     = 201601
	InfoMetadataGetMySQLClusterByEnv   = 201602
	InfoMetadataGetMySQLClusterByID    = 201603
	InfoMetadataGetMySQLClusterByName  = 201604
	InfoMetadataGetMySQLServers        = 201605
	InfoMetadataGetMasterServers       = 201606
	InfoMetadataGetDBs                 = 201607
	InfoMetadataGetUsers               = 201608
	InfoMetadataGetAppUsers            = 201609
	InfoMetadataMySQLClusterAddUser    = 101610
	InfoMetadataMySQLClusterDeleteUser = 101611
	InfoMetadataGetDBUsers             = 201612
	InfoMetadataGetAllUsers            = 201613
	InfoMetadataAddMySQLCluster        = 201614
	InfoMetadataUpdateMySQLCluster     = 201615
	InfoMetadataDeleteMySQLCluster     = 201616
	// error
	ErrMetadataGetMySQLClusterAll     = 401601
	ErrMetadataGetMySQLClusterByEnv   = 401602
	ErrMetadataGetMySQLClusterByID    = 401603
	ErrMetadataGetMySQLClusterByName  = 401604
	ErrMetadataGetMySQLServers        = 401605
	ErrMetadataGetMasterServers       = 401606
	ErrMetadataGetDBs                 = 401607
	ErrMetadataGetUsers               = 401608
	ErrMetadataMySQLClusterAddUser    = 101609
	ErrMetadataMySQLClusterDeleteUser = 101610
	ErrMetadataGetAppUsers            = 401611
	ErrMetadataGetDBUsers             = 401612
	ErrMetadataGetAllUsers            = 401613
	ErrMetadataAddMySQLCluster        = 401614
	ErrMetadataUpdateMySQLCluster     = 401615
	ErrMetadataDeleteMySQLCluster     = 401616
)

func initDebugMySQLClusterMessage() {
	message.Messages[DebugMetadataGetMySQLClusterAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterAll,
		"metadata: get all mysql clusters. message: %s")
	message.Messages[DebugMetadataGetMySQLClusterByEnv] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterByEnv,
		"metadata: get mysql cluster by env. message: %s")
	message.Messages[DebugMetadataGetMySQLClusterByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterByID,
		"metadata: get mysql cluster by id. message: %s")
	message.Messages[DebugMetadataGetMySQLClusterByName] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterByName,
		"metadata: get mysql cluster by name. message: %s")
	message.Messages[DebugMetadataGetMySQLServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLServers,
		"metadata: get mysql servers from mysql cluster. message: %s")
	message.Messages[DebugMetadataGetMasterServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMasterServers,
		"metadata: get master servers from mysql cluster. message: %s")
	message.Messages[DebugMetadataGetDBs] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetDBs,
		"metadata: get databases from mysql cluster. message: %s")
	message.Messages[DebugMetadataGetUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetUsers,
		"metadata: get users from mysql cluster. message: %s")
	message.Messages[DebugMetadataMySQLClusterAddUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataMySQLClusterAddUser,
		"metadata: add new user for mysql cluster. message: %s")
	message.Messages[DebugMetadataMySQLClusterDeleteUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataMySQLClusterDeleteUser,
		"metadata: delete users for mysql cluster. message: %s")
	message.Messages[DebugMetadataGetUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetUsers,
		"metadata: get users from mysql cluster. message: %s")
	message.Messages[DebugMetadataGetAppUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetAppUsers,
		"metadata: get app users from mysql cluster. message: %s")
	message.Messages[DebugMetadataGetDBUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetDBUsers,
		"metadata: get database users from mysql cluster. message: %s")
	message.Messages[DebugMetadataGetAllUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetAllUsers,
		"metadata: get all users from mysql cluster. message: %s")
	message.Messages[DebugMetadataAddMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataAddMySQLCluster,
		"metadata: add new mysql cluster. message: %s")
	message.Messages[DebugMetadataUpdateMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataUpdateMySQLCluster,
		"metadata: update mysql cluster. message: %s")
	message.Messages[DebugMetadataDeleteMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataDeleteMySQLCluster,
		"metadata: delete mysql cluster. message: %s")
}

func initInfoMySQLClusterMessage() {
	message.Messages[InfoMetadataGetMySQLClusterAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLClusterAll,
		"metadata: get mysql cluster all completed")
	message.Messages[InfoMetadataGetMySQLClusterByEnv] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLClusterByEnv,
		"metadata: get mysql cluster by env completed. env_id: %d")
	message.Messages[InfoMetadataGetMySQLClusterByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLClusterByID,
		"metadata: get mysql cluster by id completed. id: %d")
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
		"metadata: add new mysql cluster completed. cluster_name: %s, env_id: %d")
	message.Messages[InfoMetadataUpdateMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataUpdateMySQLCluster,
		"metadata: update mysql cluster completed. id: %d")
	message.Messages[InfoMetadataDeleteMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataDeleteMySQLCluster,
		"metadata: delete mysql cluster completed. id: %d")
}

func initErrorMySQLClusterMessage() {
	message.Messages[ErrMetadataGetMySQLClusterAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLClusterAll,
		"metadata: get all mysql cluster failed")
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
