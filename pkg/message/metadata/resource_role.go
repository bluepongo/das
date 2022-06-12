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
	DebugMetadataGetResourceRoleAll               = 101601
	DebugMetadataGetResourceRoleByID              = 101602
	DebugMetadataGetResourceRoleByUUID            = 101603
	DebugMetadataGetResourceGroupByResourceRoleID = 101604
	DebugMetadataGetUsersByResourceRoleID         = 101605
	DebugMetadataGetUsersByResourceRoleUUID       = 101606
	DebugMetadataResourceRoleAddUser              = 101607
	DebugMetadataResourceRoleDeleteUser           = 101608
	DebugMetadataAddResourceRole                  = 101609
	DebugMetadataUpdateResourceRole               = 101610
	DebugMetadataDeleteResourceRole               = 101611
	// info
	InfoMetadataGetResourceRoleAll               = 201601
	InfoMetadataGetResourceRoleByID              = 201602
	InfoMetadataGetResourceRoleByUUID            = 201603
	InfoMetadataGetResourceGroupByResourceRoleID = 201604
	InfoMetadataGetUsersByResourceRoleID         = 201605
	InfoMetadataGetUsersByResourceRoleUUID       = 201606
	InfoMetadataResourceRoleAddUser              = 201607
	InfoMetadataResourceRoleDeleteUser           = 201608
	InfoMetadataAddResourceRole                  = 201609
	InfoMetadataUpdateResourceRole               = 201610
	InfoMetadataDeleteResourceRole               = 201611
	// error
	ErrMetadataGetResourceRoleAll               = 401601
	ErrMetadataGetResourceRoleByID              = 401602
	ErrMetadataGetResourceRoleByUUID            = 401603
	ErrMetadataGetResourceGroupByResourceRoleID = 401604
	ErrMetadataGetUsersByResourceRoleID         = 401605
	ErrMetadataGetUsersByResourceRoleUUID       = 401606
	ErrMetadataResourceRoleAddUser              = 401607
	ErrMetadataResourceRoleDeleteUser           = 401608
	ErrMetadataAddResourceRole                  = 401609
	ErrMetadataUpdateResourceRole               = 401610
	ErrMetadataDeleteResourceRole               = 401611
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
