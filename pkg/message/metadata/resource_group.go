package metadata

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initDebugResourceGroupMessage()
	initInfoResourceGroupMessage()
	initErrorResourceGroupMessage()
}

const (
	// debug
	DebugMetadataGetResourceGroupAll                  = 102001
	DebugMetadataGetResourceGroupByID                 = 102002
	DebugMetadataGetResourceGroupByGroupUUID          = 102003
	DebugMetadataGetResourceRolesByGroupID            = 102004
	DebugMetadataGetMySQLClustersByGroupID            = 102005
	DebugMetadataGetMySQLServersByGroupID             = 102006
	DebugMetadataGetMiddlewareClustersByGroupID       = 102007
	DebugMetadataGetMiddlewareServersByGroupID        = 102008
	DebugMetadataGetUsersByGroupID                    = 102009
	DebugMetadataGetDASAdminUsersByGroupID            = 102010
	DebugMetadataGetResourceRolesByGroupUUID          = 102011
	DebugMetadataGetMySQLClustersByGroupUUID          = 102012
	DebugMetadataGetMySQLServersByGroupUUID           = 102013
	DebugMetadataGetMiddlewareClustersByGroupUUID     = 102014
	DebugMetadataGetMiddlewareServersByGroupUUID      = 102015
	DebugMetadataGetUsersByGroupUUID                  = 102016
	DebugMetadataGetDASAdminUsersByGroupUUID          = 102017
	DebugMetadataAddResourceGroup                     = 102018
	DebugMetadataUpdateResourceGroup                  = 102019
	DebugMetadataDeleteResourceGroup                  = 102020
	DebugMetadataResourceGroupAddMySQLCluster         = 102021
	DebugMetadataResourceGroupDeleteMySQLCluster      = 102022
	DebugMetadataResourceGroupAddMiddlewareCluster    = 102023
	DebugMetadataResourceGroupDeleteMiddlewareCluster = 102024
	// info
	InfoMetadataGetResourceGroupAll                  = 202001
	InfoMetadataGetResourceGroupByID                 = 202002
	InfoMetadataGetResourceGroupByGroupUUID          = 202003
	InfoMetadataGetResourceRolesByGroupID            = 202004
	InfoMetadataGetMySQLClustersByGroupID            = 202005
	InfoMetadataGetMySQLServersByGroupID             = 202006
	InfoMetadataGetMiddlewareClustersByGroupID       = 202007
	InfoMetadataGetMiddlewareServersByGroupID        = 202008
	InfoMetadataGetUsersByGroupID                    = 202009
	InfoMetadataGetDASAdminUsersByGroupID            = 202010
	InfoMetadataGetResourceRolesByGroupUUID          = 202011
	InfoMetadataGetMySQLClustersByGroupUUID          = 202012
	InfoMetadataGetMySQLServersByGroupUUID           = 202013
	InfoMetadataGetMiddlewareClustersByGroupUUID     = 202014
	InfoMetadataGetMiddlewareServersByGroupUUID      = 202015
	InfoMetadataGetUsersByGroupUUID                  = 202016
	InfoMetadataGetDASAdminUsersByGroupUUID          = 202017
	InfoMetadataAddResourceGroup                     = 202018
	InfoMetadataUpdateResourceGroup                  = 202019
	InfoMetadataDeleteResourceGroup                  = 202020
	InfoMetadataResourceGroupAddMySQLCluster         = 202021
	InfoMetadataResourceGroupDeleteMySQLCluster      = 202022
	InfoMetadataResourceGroupAddMiddlewareCluster    = 202023
	InfoMetadataResourceGroupDeleteMiddlewareCluster = 202024
	// error
	ErrMetadataGetResourceGroupAll                  = 402001
	ErrMetadataGetResourceGroupByID                 = 402002
	ErrMetadataGetResourceGroupByGroupUUID          = 402003
	ErrMetadataGetResourceRolesByGroupID            = 402004
	ErrMetadataGetMySQLClustersByGroupID            = 402005
	ErrMetadataGetMySQLServersByGroupID             = 402006
	ErrMetadataGetMiddlewareClustersByGroupID       = 402007
	ErrMetadataGetMiddlewareServersByGroupID        = 402008
	ErrMetadataGetUsersByGroupID                    = 402009
	ErrMetadataGetDASAdminUsersByGroupID            = 402010
	ErrMetadataGetResourceRolesByGroupUUID          = 402011
	ErrMetadataGetMySQLClustersByGroupUUID          = 402012
	ErrMetadataGetMySQLServersByGroupUUID           = 402013
	ErrMetadataGetMiddlewareClustersByGroupUUID     = 402014
	ErrMetadataGetMiddlewareServersByGroupUUID      = 402015
	ErrMetadataGetUsersByGroupUUID                  = 402016
	ErrMetadataGetDASAdminUsersByGroupUUID          = 402017
	ErrMetadataAddResourceGroup                     = 402018
	ErrMetadataUpdateResourceGroup                  = 402019
	ErrMetadataDeleteResourceGroup                  = 402020
	ErrMetadataResourceGroupAddMySQLCluster         = 402021
	ErrMetadataResourceGroupDeleteMySQLCluster      = 402022
	ErrMetadataResourceGroupAddMiddlewareCluster    = 402023
	ErrMetadataResourceGroupDeleteMiddlewareCluster = 402024
)

func initDebugResourceGroupMessage() {
	message.Messages[DebugMetadataGetResourceGroupAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetResourceGroupAll, "metadata: get all resource groups. message: %s")
	message.Messages[DebugMetadataGetResourceGroupByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetResourceGroupByID, "metadata: get resource group by id. message: %s")
	message.Messages[DebugMetadataGetResourceGroupByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetResourceGroupByGroupUUID, "metadata: get resource group by group uuid. message: %s")
	message.Messages[DebugMetadataGetResourceRolesByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetResourceRolesByGroupID, "metadata: get resource roles by id. message: %s")
	message.Messages[DebugMetadataGetMySQLClustersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMySQLClustersByGroupID, "metadata: get mysql clusters by id. message: %s")
	message.Messages[DebugMetadataGetMySQLServersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMySQLServersByGroupID, "metadata: get mysql servers by id. message: %s")
	message.Messages[DebugMetadataGetMiddlewareClustersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareClustersByGroupID, "metadata: get middleware clusters by id. message: %s")
	message.Messages[DebugMetadataGetMiddlewareServersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareServersByGroupID, "metadata: get middleware servers by id. message: %s")
	message.Messages[DebugMetadataGetUsersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetUsersByGroupID, "metadata: get users by id. message: %s")
	message.Messages[DebugMetadataGetDASAdminUsersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDASAdminUsersByGroupID, "metadata: get das admin users by id. message: %s")
	message.Messages[DebugMetadataGetResourceRolesByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetResourceRolesByGroupUUID, "metadata: get resource roles by resource group uuid. message: %s")
	message.Messages[DebugMetadataGetMySQLClustersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMySQLClustersByGroupUUID, "metadata: get mysql clusters by resource group uuid. message: %s")
	message.Messages[DebugMetadataGetMySQLServersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMySQLServersByGroupUUID, "metadata: get mysql servers by resource group uuid. message: %s")
	message.Messages[DebugMetadataGetMiddlewareClustersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareClustersByGroupUUID, "metadata: get middleware clusters by resource group uuid. message: %s")
	message.Messages[DebugMetadataGetMiddlewareServersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareServersByGroupUUID, "metadata: get middleware servers by resource group uuid. message: %s")
	message.Messages[DebugMetadataGetUsersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetUsersByGroupUUID, "metadata: get users by resource group uuid. message: %s")
	message.Messages[DebugMetadataGetDASAdminUsersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDASAdminUsersByGroupUUID, "metadata: get das admin users by resource group uuid. message: %s")
	message.Messages[DebugMetadataAddResourceGroup] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddResourceGroup, "metadata: add new resource group. message: %s")
	message.Messages[DebugMetadataUpdateResourceGroup] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateResourceGroup, "metadata: update resource group. message: %s")
	message.Messages[DebugMetadataDeleteResourceGroup] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteResourceGroup, "metadata: delete resource group. message: %s")
	message.Messages[DebugMetadataResourceGroupAddMySQLCluster] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataResourceGroupAddMySQLCluster, "metadata: add mysql cluster. message: %s")
	message.Messages[DebugMetadataResourceGroupDeleteMySQLCluster] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataResourceGroupDeleteMySQLCluster, "metadata: delete mysql cluster. message: %s")
	message.Messages[DebugMetadataResourceGroupAddMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataResourceGroupAddMiddlewareCluster, "metadata: add middleware cluster. message: %s")
	message.Messages[DebugMetadataResourceGroupDeleteMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataResourceGroupDeleteMiddlewareCluster, "metadata: delete middleware cluster. message: %s")
}

func initInfoResourceGroupMessage() {
	message.Messages[InfoMetadataGetResourceGroupAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetResourceGroupAll, "metadata: get all resource groups completed")
	message.Messages[InfoMetadataGetResourceGroupByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetResourceGroupByID, "metadata: get resource group by id completed. id: %d")
	message.Messages[InfoMetadataGetResourceGroupByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetResourceGroupByGroupUUID, "metadata: get resource group by group uuid completed. group_uuid: %s")
	message.Messages[InfoMetadataGetResourceRolesByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetResourceRolesByGroupID, "metadata: get resource roles by id completed. id: %d")
	message.Messages[InfoMetadataGetMySQLClustersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMySQLClustersByGroupID, "metadata: get mysql clusters by id completed. id: %d")
	message.Messages[InfoMetadataGetMySQLServersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMySQLServersByGroupID, "metadata: get mysql servers by id completed. id: %d")
	message.Messages[InfoMetadataGetMiddlewareClustersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareClustersByGroupID, "metadata: get middleware clusters by id completed. id: %d")
	message.Messages[InfoMetadataGetMiddlewareServersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareServersByGroupID, "metadata: get middleware servers by id completed. id: %d")
	message.Messages[InfoMetadataGetUsersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetUsersByGroupID, "metadata: get users by id completed. id: %d")
	message.Messages[InfoMetadataGetDASAdminUsersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDASAdminUsersByGroupID, "metadata: get das admin users by id completed. id: %d")
	message.Messages[InfoMetadataGetResourceRolesByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetResourceRolesByGroupUUID, "metadata: get resource roles by resource group uuid completed. group_uuid: %d")
	message.Messages[InfoMetadataGetMySQLClustersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMySQLClustersByGroupUUID, "metadata: get mysql clusters by resource group uuid completed. group_uuid: %d")
	message.Messages[InfoMetadataGetMySQLServersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMySQLServersByGroupUUID, "metadata: get mysql servers by resource group uuid completed. group_uuid: %d")
	message.Messages[InfoMetadataGetMiddlewareClustersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareClustersByGroupUUID, "metadata: get middleware clusters by resource group uuid completed. group_uuid: %d")
	message.Messages[InfoMetadataGetMiddlewareServersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareServersByGroupUUID, "metadata: get middleware servers by resource group uuid completed. group_uuid: %d")
	message.Messages[InfoMetadataGetUsersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetUsersByGroupUUID, "metadata: get users by resource group uuid completed. group_uuid: %d")
	message.Messages[InfoMetadataGetDASAdminUsersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDASAdminUsersByGroupUUID, "metadata: get das admin users by resource group uuid completed. group_uuid: %d")
	message.Messages[InfoMetadataAddResourceGroup] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddResourceGroup, "metadata: add new resource group completed. group_uuid: %s, group_name: %s")
	message.Messages[InfoMetadataUpdateResourceGroup] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateResourceGroup, "metadata: update resource group completed. id: %d")
	message.Messages[InfoMetadataDeleteResourceGroup] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDeleteResourceGroup, "metadata: delete resource group completed. id: %d")
	message.Messages[InfoMetadataResourceGroupAddMySQLCluster] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataResourceGroupAddMySQLCluster, "metadata: add mysql cluster completed. resource_group_id: %d, mysql_cluster_id: %d")
	message.Messages[InfoMetadataResourceGroupDeleteMySQLCluster] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataResourceGroupDeleteMySQLCluster, "metadata: delete mysql cluster completed. resource_group_id: %d, mysql_cluster_id: %d")
	message.Messages[InfoMetadataResourceGroupAddMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataResourceGroupAddMiddlewareCluster, "metadata: add middleware cluster completed. resource_group_id: %d, middleware_cluster_id: %d")
	message.Messages[InfoMetadataResourceGroupDeleteMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataResourceGroupDeleteMiddlewareCluster, "metadata: delete middleware cluster completed. resource_group_id: %d, middleware_cluster_id: %d")
}

func initErrorResourceGroupMessage() {
	message.Messages[ErrMetadataGetResourceGroupAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetResourceGroupAll, "metadata: get all resource groups failed.")
	message.Messages[ErrMetadataGetResourceGroupByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetResourceGroupByID, "metadata: get resource group by id failed. id: %d")
	message.Messages[ErrMetadataGetResourceGroupByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetResourceGroupByGroupUUID, "metadata: get resource group by group uuid failed. group_uuid: %s")
	message.Messages[ErrMetadataGetResourceRolesByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetResourceRolesByGroupID, "metadata: get resource roles by id failed. id: %d")
	message.Messages[ErrMetadataGetMySQLClustersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMySQLClustersByGroupID, "metadata: get mysql clusters by id failed. id: %d")
	message.Messages[ErrMetadataGetMySQLServersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMySQLServersByGroupID, "metadata: get mysql servers by id failed. id: %d")
	message.Messages[ErrMetadataGetMiddlewareClustersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareClustersByGroupID, "metadata: get middleware clusters by id failed. id: %d")
	message.Messages[ErrMetadataGetMiddlewareServersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareServersByGroupID, "metadata: get middleware servers by id failed. id: %d")
	message.Messages[ErrMetadataGetUsersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetUsersByGroupID, "metadata: get users by id failed. id: %d")
	message.Messages[ErrMetadataGetDASAdminUsersByGroupID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetDASAdminUsersByGroupID, "metadata: get das admin users by id failed. id: %d")
	message.Messages[ErrMetadataGetResourceRolesByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetResourceRolesByGroupUUID, "metadata: get resource roles by resource group uuid failed. group_uuid: %d")
	message.Messages[ErrMetadataGetMySQLClustersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMySQLClustersByGroupUUID, "metadata: get mysql clusters by resource group uuid failed. group_uuid: %d")
	message.Messages[ErrMetadataGetMySQLServersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMySQLServersByGroupUUID, "metadata: get mysql servers by resource group uuid failed. group_uuid: %d")
	message.Messages[ErrMetadataGetMiddlewareClustersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareClustersByGroupUUID, "metadata: get middleware clusters by resource group uuid failed. group_uuid: %d")
	message.Messages[ErrMetadataGetMiddlewareServersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareServersByGroupUUID, "metadata: get middleware servers by resource group uuid failed. group_uuid: %d")
	message.Messages[ErrMetadataGetUsersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetUsersByGroupUUID, "metadata: get users by resource group uuid failed. group_uuid: %d")
	message.Messages[ErrMetadataGetDASAdminUsersByGroupUUID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetDASAdminUsersByGroupUUID, "metadata: get das admin users by resource group uuid failed. group_uuid: %d")
	message.Messages[ErrMetadataAddResourceGroup] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddResourceGroup, "metadata: add new resource group failed. group_uuid: %s, group_name: %s")
	message.Messages[ErrMetadataUpdateResourceGroup] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateResourceGroup, "metadata: update resource group failed. id: %d")
	message.Messages[ErrMetadataDeleteResourceGroup] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteResourceGroup, "metadata: delete resource group failed. id: %d")
	message.Messages[ErrMetadataResourceGroupAddMySQLCluster] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataResourceGroupAddMySQLCluster, "metadata: add mysql cluster failed. resource_group_id: %d, mysql_cluster_id: %d")
	message.Messages[ErrMetadataResourceGroupDeleteMySQLCluster] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataResourceGroupDeleteMySQLCluster, "metadata: delete mysql cluster failed. resource_group_id: %d, mysql_cluster_id: %d")
	message.Messages[ErrMetadataResourceGroupAddMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataResourceGroupAddMiddlewareCluster, "metadata: add middleware cluster failed. resource_group_id: %d, middleware_cluster_id: %d")
	message.Messages[ErrMetadataResourceGroupDeleteMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataResourceGroupDeleteMiddlewareCluster, "metadata: delete middleware cluster failed. resource_group_id: %d, middleware_cluster_id: %d")
}
