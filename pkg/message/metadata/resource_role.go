package metadata

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/das/pkg/message"
)

func init() {
	initDebugResourceRoleMessage()
	initInfoResourceRoleMessage()
	initErrorResourceRoleMessage()
}

// Message code
const (
	// debug
	DebugMetadataGetResourceRoleAll     = 101601
	DebugMetadataGetResourceRoleByEnv   = 101602
	DebugMetadataGetResourceRoleByID    = 101603
	DebugMetadataGetResourceRoleByName  = 101604
	DebugMetadataGetMySQLServers        = 101605
	DebugMetadataGetMasterServers       = 101606
	DebugMetadataGetDBs                 = 101607
	DebugMetadataGetUsers               = 101608
	DebugMetadataResourceRoleAddUser    = 101609
	DebugMetadataResourceRoleDeleteUser = 101610
	DebugMetadataGetAppUsers            = 101611
	DebugMetadataGetDBUsers             = 101612
	DebugMetadataGetAllUsers            = 101613
	DebugMetadataAddResourceRole        = 101614
	DebugMetadataUpdateResourceRole     = 101615
	DebugMetadataDeleteResourceRole     = 101616
	// debug
	InfoMetadataGetResourceRoleAll     = 201601
	InfoMetadataGetResourceRoleByEnv   = 201602
	InfoMetadataGetResourceRoleByID    = 201603
	InfoMetadataGetResourceRoleByName  = 201604
	InfoMetadataGetMySQLServers        = 201605
	InfoMetadataGetMasterServers       = 201606
	InfoMetadataGetDBs                 = 201607
	InfoMetadataGetUsers               = 201608
	InfoMetadataGetAppUsers            = 201609
	InfoMetadataResourceRoleAddUser    = 101610
	InfoMetadataResourceRoleDeleteUser = 101611
	InfoMetadataGetDBUsers             = 201612
	InfoMetadataGetAllUsers            = 201613
	InfoMetadataAddResourceRole        = 201614
	InfoMetadataUpdateResourceRole     = 201615
	InfoMetadataDeleteResourceRole     = 201616
	// error
	ErrMetadataGetResourceRoleAll     = 401601
	ErrMetadataGetResourceRoleByEnv   = 401602
	ErrMetadataGetResourceRoleByID    = 401603
	ErrMetadataGetResourceRoleByName  = 401604
	ErrMetadataGetMySQLServers        = 401605
	ErrMetadataGetMasterServers       = 401606
	ErrMetadataGetDBs                 = 401607
	ErrMetadataGetUsers               = 401608
	ErrMetadataResourceRoleAddUser    = 101609
	ErrMetadataResourceRoleDeleteUser = 101610
	ErrMetadataGetAppUsers            = 401611
	ErrMetadataGetDBUsers             = 401612
	ErrMetadataGetAllUsers            = 401613
	ErrMetadataAddResourceRole        = 401614
	ErrMetadataUpdateResourceRole     = 401615
	ErrMetadataDeleteResourceRole     = 401616
)

func initDebugResourceRoleMessage() {
	message.Messages[DebugMetadataGetResourceRoleAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetResourceRoleAll,
		"metadata: get all resource roles message: %s")
	message.Messages[DebugMetadataGetResourceRoleByEnv] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetResourceRoleByEnv,
		"metadata: get resource role by env message: %s")
	message.Messages[DebugMetadataGetResourceRoleByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetResourceRoleByID,
		"metadata: get resource role by id message: %s")
	message.Messages[DebugMetadataGetResourceRoleByName] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetResourceRoleByName,
		"metadata: get resource role by name message: %s")
	message.Messages[DebugMetadataGetMySQLServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLServers,
		"metadata: get mysql servers from resource role message: %s")
	message.Messages[DebugMetadataGetMasterServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMasterServers,
		"metadata: get master servers from resource role message: %s")
	message.Messages[DebugMetadataGetDBs] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetDBs,
		"metadata: get databases from resource role message: %s")
	message.Messages[DebugMetadataGetUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetUsers,
		"metadata: get users from resource role message: %s")
	message.Messages[DebugMetadataResourceRoleAddUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataResourceRoleAddUser,
		"metadata: add new user for resource role message: %s")
	message.Messages[DebugMetadataResourceRoleDeleteUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataResourceRoleDeleteUser,
		"metadata: delete users for resource role message: %s")
	message.Messages[DebugMetadataGetUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetUsers,
		"metadata: get users from resource role message: %s")
	message.Messages[DebugMetadataGetAppUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetAppUsers,
		"metadata: get app users from resource role message: %s")
	message.Messages[DebugMetadataGetDBUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetDBUsers,
		"metadata: get database users from resource role message: %s")
	message.Messages[DebugMetadataGetAllUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetAllUsers,
		"metadata: get all users from resource role message: %s")
	message.Messages[DebugMetadataAddResourceRole] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataAddResourceRole,
		"metadata: add new resource role message: %s")
	message.Messages[DebugMetadataUpdateResourceRole] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataUpdateResourceRole,
		"metadata: update resource role message: %s")
	message.Messages[DebugMetadataDeleteResourceRole] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataDeleteResourceRole,
		"metadata: delete resource role message: %s")
}

func initInfoResourceRoleMessage() {
	message.Messages[InfoMetadataGetResourceRoleAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetResourceRoleAll,
		"metadata: get resource role all completed")
	message.Messages[InfoMetadataGetResourceRoleByEnv] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetResourceRoleByEnv,
		"metadata: get resource role by env completed. env_id: %s")
	message.Messages[InfoMetadataGetResourceRoleByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetResourceRoleByID,
		"metadata: get resource role by id completed. id: %s")
	message.Messages[InfoMetadataGetResourceRoleByName] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetResourceRoleByName,
		"metadata: get resource role by name completed. cluster_name: %s")
	message.Messages[InfoMetadataGetMySQLServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLServers,
		"metadata: get mysql servers from resource role completed. id: %d")
	message.Messages[InfoMetadataGetMasterServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMasterServers,
		"metadata: get master servers from resource role completed. id: %d")
	message.Messages[InfoMetadataGetDBs] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetDBs,
		"metadata: get databases from resource role completed. id: %d")
	message.Messages[InfoMetadataGetUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetUsers,
		"metadata: get users from resource role completed. id: %d")
	message.Messages[InfoMetadataResourceRoleAddUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataResourceRoleAddUser,
		"metadata: add user for resource role completed. id: %d, user_id: %d")
	message.Messages[InfoMetadataResourceRoleDeleteUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataResourceRoleDeleteUser,
		"metadata: delete user for resource role completed. id: %d, user_id: %d")
	message.Messages[InfoMetadataGetAppUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetAppUsers,
		"metadata: get app users from resource role completed. id: %d")
	message.Messages[InfoMetadataGetDBUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetDBUsers,
		"metadata: get database users from resource role completed. id: %d")
	message.Messages[InfoMetadataGetAllUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetAllUsers,
		"metadata: get all users from resource role completed. id: %d")
	message.Messages[InfoMetadataAddResourceRole] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataAddResourceRole,
		"metadata: add new resource role completed. cluster_name: %s, env_id: %s")
	message.Messages[InfoMetadataUpdateResourceRole] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataUpdateResourceRole,
		"metadata: update resource role completed. id: %s")
	message.Messages[InfoMetadataDeleteResourceRole] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataDeleteResourceRole,
		"metadata: delete resource role completed. id: %s")
}

func initErrorResourceRoleMessage() {
	message.Messages[ErrMetadataGetResourceRoleAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetResourceRoleAll,
		"metadata: get all resource role failed")
	message.Messages[ErrMetadataGetResourceRoleByEnv] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetResourceRoleByEnv,
		"metadata: get resource role by env failed. env_id: %d")
	message.Messages[ErrMetadataGetResourceRoleByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetResourceRoleByID,
		"metadata: get resource role by id failed. id: %d")
	message.Messages[ErrMetadataGetResourceRoleByName] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetResourceRoleByName,
		"metadata: get resource role by name failed. cluster_name: %s")
	message.Messages[ErrMetadataGetMySQLServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLServers,
		"metadata: get mysql servers from resource role failed. id: %d")
	message.Messages[ErrMetadataGetMasterServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMasterServers,
		"metadata: get master servers from resource role failed. id: %d")
	message.Messages[ErrMetadataGetDBs] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetDBs,
		"metadata: get databases from resource role failed. id: %d")
	message.Messages[ErrMetadataGetUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetUsers,
		"metadata: get users from resource role failed. id: %d")
	message.Messages[ErrMetadataResourceRoleAddUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataResourceRoleAddUser,
		"metadata: add user for resource role failed. id: %d, user_id: %d")
	message.Messages[ErrMetadataResourceRoleDeleteUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataResourceRoleDeleteUser,
		"metadata: delete user for resource role failed. id: %d, user_id: %d")
	message.Messages[ErrMetadataGetAppUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetAppUsers,
		"metadata: get app users from resource role failed. id: %d")
	message.Messages[ErrMetadataGetDBUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetDBUsers,
		"metadata: get database users from resource role failed. id: %d")
	message.Messages[ErrMetadataGetAllUsers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetAllUsers,
		"metadata: get all users from resource role failed. id: %d")
	message.Messages[ErrMetadataAddResourceRole] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataAddResourceRole,
		"metadata: add new resource role failed. cluster_name: %s, env_id: %d")
	message.Messages[ErrMetadataUpdateResourceRole] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataUpdateResourceRole,
		"metadata: update resource role failed. id: %d")
	message.Messages[ErrMetadataDeleteResourceRole] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataDeleteResourceRole,
		"metadata: delete resource role failed. id: %d")
}
