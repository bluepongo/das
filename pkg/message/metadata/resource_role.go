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
	DebugMetadataGetResourceRoleAll               = 102201
	DebugMetadataGetResourceRoleByID              = 102202
	DebugMetadataGetResourceRoleByUUID            = 102203
	DebugMetadataGetResourceGroupByResourceRoleID = 102204
	DebugMetadataGetUsersByResourceRoleID         = 102205
	DebugMetadataGetUsersByResourceRoleUUID       = 102206
	DebugMetadataResourceRoleAddUser              = 102207
	DebugMetadataResourceRoleDeleteUser           = 102208
	DebugMetadataAddResourceRole                  = 102209
	DebugMetadataUpdateResourceRole               = 102210
	DebugMetadataDeleteResourceRole               = 102211
	// info
	InfoMetadataGetResourceRoleAll               = 202201
	InfoMetadataGetResourceRoleByID              = 202202
	InfoMetadataGetResourceRoleByUUID            = 202203
	InfoMetadataGetResourceGroupByResourceRoleID = 202204
	InfoMetadataGetUsersByResourceRoleID         = 202205
	InfoMetadataGetUsersByResourceRoleUUID       = 202206
	InfoMetadataResourceRoleAddUser              = 202207
	InfoMetadataResourceRoleDeleteUser           = 202208
	InfoMetadataAddResourceRole                  = 202209
	InfoMetadataUpdateResourceRole               = 202210
	InfoMetadataDeleteResourceRole               = 202211
	// error
	ErrMetadataGetResourceRoleAll               = 402201
	ErrMetadataGetResourceRoleByID              = 402202
	ErrMetadataGetResourceRoleByUUID            = 402203
	ErrMetadataGetResourceGroupByResourceRoleID = 402204
	ErrMetadataGetUsersByResourceRoleID         = 402205
	ErrMetadataGetUsersByResourceRoleUUID       = 402206
	ErrMetadataResourceRoleAddUser              = 402207
	ErrMetadataResourceRoleDeleteUser           = 402208
	ErrMetadataAddResourceRole                  = 402209
	ErrMetadataUpdateResourceRole               = 402210
	ErrMetadataDeleteResourceRole               = 402211
)

func initDebugResourceRoleMessage() {
	message.Messages[DebugMetadataGetResourceRoleAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetResourceRoleAll,
		"metadata: get all resource roles message: %s")
	message.Messages[DebugMetadataGetResourceRoleByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetResourceRoleByID,
		"metadata: get resource role by id message: %s")
	message.Messages[DebugMetadataGetResourceRoleByUUID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetResourceRoleByUUID,
		"metadata: get resource role by uuid message: %s")
	message.Messages[DebugMetadataGetResourceGroupByResourceRoleID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetResourceGroupByResourceRoleID,
		"metadata: get resource group by resource role id message: %s")
	message.Messages[DebugMetadataGetUsersByResourceRoleID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetUsersByResourceRoleID,
		"metadata: get users by resource role id message: %s")
	message.Messages[DebugMetadataGetResourceRoleByUUID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetResourceRoleByUUID,
		"metadata: get users by resource role uuid message: %s")
	message.Messages[DebugMetadataResourceRoleAddUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataResourceRoleAddUser,
		"metadata: add new user for resource role message: %s")
	message.Messages[DebugMetadataResourceRoleDeleteUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataResourceRoleDeleteUser,
		"metadata: delete users for resource role message: %s")
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
	message.Messages[InfoMetadataGetResourceRoleByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetResourceRoleByID,
		"metadata: get resource role by id completed. id: %s")
	message.Messages[InfoMetadataGetResourceRoleByUUID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetResourceRoleByUUID,
		"metadata: get resource role by uuid completed. role_uuid: %s")
	message.Messages[InfoMetadataGetResourceGroupByResourceRoleID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetResourceGroupByResourceRoleID,
		"metadata: get resource group by resource role id completed. id: %d")
	message.Messages[InfoMetadataGetUsersByResourceRoleID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetUsersByResourceRoleID,
		"metadata: get users by resource role id completed. id: %d")
	message.Messages[InfoMetadataGetUsersByResourceRoleUUID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetUsersByResourceRoleUUID,
		"metadata: get users by resource role uuid completed. uuid: %s")
	message.Messages[InfoMetadataResourceRoleAddUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataResourceRoleAddUser,
		"metadata: add user for resource role completed. id: %d, user_id: %d")
	message.Messages[InfoMetadataResourceRoleDeleteUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataResourceRoleDeleteUser,
		"metadata: delete user for resource role completed. id: %d, user_id: %d")
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
	message.Messages[ErrMetadataGetResourceRoleByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetResourceRoleByID,
		"metadata: get resource role by id failed. id: %d")
	message.Messages[ErrMetadataGetResourceRoleByUUID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetResourceRoleByUUID,
		"metadata: get resource role by uuid failed. role_uuid: %s")
	message.Messages[ErrMetadataGetResourceGroupByResourceRoleID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetResourceGroupByResourceRoleID,
		"metadata: get resource group by resource role id failed. id: %d")
	message.Messages[ErrMetadataGetUsersByResourceRoleID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetUsersByResourceRoleID,
		"metadata: get users by resource role id failed. id: %d")
	message.Messages[ErrMetadataGetUsersByResourceRoleUUID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetUsersByResourceRoleUUID,
		"metadata: get users by resource role uuid failed. role_uuid: %s")
	message.Messages[ErrMetadataResourceRoleAddUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataResourceRoleAddUser,
		"metadata: add user for resource role failed. id: %d, user_id: %d")
	message.Messages[ErrMetadataResourceRoleDeleteUser] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataResourceRoleDeleteUser,
		"metadata: delete user for resource role failed. id: %d, user_id: %d")
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
