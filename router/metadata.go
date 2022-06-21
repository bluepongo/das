package router

import (
	"github.com/gin-gonic/gin"

	"github.com/romberli/das/api/v1/metadata"
)

// RegisterMetadata api to gin router
func RegisterMetadata(group *gin.RouterGroup) {
	metadataGroup := group.Group("/metadata")
	{
		// app
		metadataGroup.POST("/app/all", metadata.GetApp)
		metadataGroup.POST("/app/id", metadata.GetAppByID)
		metadataGroup.POST("/app/app-name", metadata.GetAppByName)
		metadataGroup.POST("/app/db", metadata.GetDBsByAppID)
		metadataGroup.POST("/app/user", metadata.GetUsersByAppID)

		metadataGroup.POST("/app/add", metadata.AddApp)
		metadataGroup.POST("/app/update", metadata.UpdateAppByID)
		metadataGroup.POST("/app/delete", metadata.DeleteAppByID)
		metadataGroup.POST("/app/add-db", metadata.AppAddDB)
		metadataGroup.POST("/app/delete-db", metadata.AppDeleteDB)
		metadataGroup.POST("/app/add-user", metadata.AppAddUser)
		metadataGroup.POST("/app/delete-user", metadata.AppDeleteUser)
		// db
		metadataGroup.POST("/db/all", metadata.GetDB)
		metadataGroup.POST("/db/id", metadata.GetDBByID)
		metadataGroup.POST("/db/env", metadata.GetDBByEnv)
		metadataGroup.POST("/db/name-and-cluster-info", metadata.GetDBByNameAndClusterInfo)
		metadataGroup.POST("/db/name-and-host-info", metadata.GetDBByNameAndHostInfo)
		metadataGroup.POST("/db/host-info", metadata.GetDBsByHostInfo)
		metadataGroup.POST("/db/app", metadata.GetAppsByDBID)
		metadataGroup.POST("/db/mysql-cluster", metadata.GetMySQLClusterByDBID)
		metadataGroup.POST("/db/app-user", metadata.GetAppUsersByDBID)
		metadataGroup.POST("/db/db-user", metadata.GetUsersByDBID)
		metadataGroup.POST("/db/all-user", metadata.GetAllUsersByDBID)

		metadataGroup.POST("/db/add", metadata.AddDB)
		metadataGroup.POST("/db/update", metadata.UpdateDBByID)
		metadataGroup.POST("/db/delete", metadata.DeleteDBByID)
		metadataGroup.POST("/db/add-app", metadata.DBAddApp)
		metadataGroup.POST("/db/delete-app", metadata.DBDeleteApp)
		metadataGroup.POST("/db/add-user", metadata.DBAddUser)
		metadataGroup.POST("/db/delete-user", metadata.DBDeleteUser)
		// env
		metadataGroup.POST("/env/all", metadata.GetEnv)
		metadataGroup.POST("/env/id", metadata.GetEnvByID)
		metadataGroup.POST("/env/env-name", metadata.GetEnvByName)

		metadataGroup.POST("/env/add", metadata.AddEnv)
		metadataGroup.POST("/env/update", metadata.UpdateEnvByID)
		metadataGroup.POST("/env/delete", metadata.DeleteEnvByID)
		// middleware cluster
		metadataGroup.POST("/middleware-cluster/all", metadata.GetMiddlewareCluster)
		metadataGroup.POST("/middleware-cluster/env", metadata.GetMiddlewareClusterByEnv)
		metadataGroup.POST("/middleware-cluster/id", metadata.GetMiddlewareClusterByID)
		metadataGroup.POST("/middleware-cluster/cluster-name", metadata.GetMiddlewareClusterByName)
		metadataGroup.POST("/middleware-cluster/middleware-server", metadata.GetMiddlewareServers)
		metadataGroup.POST("/middleware-cluster/user", metadata.GetUsersByMiddlewareClusterID)

		metadataGroup.POST("/middleware-cluster/add", metadata.AddMiddlewareCluster)
		metadataGroup.POST("/middleware-cluster/update", metadata.UpdateMiddlewareClusterByID)
		metadataGroup.POST("/middleware-cluster/delete", metadata.DeleteMiddlewareClusterByID)
		metadataGroup.POST("/middleware-cluster/add-user", metadata.MiddlewareClusterAddUser)
		metadataGroup.POST("/middleware-cluster/delete-user", metadata.MiddlewareClusterDeleteUser)
		// middleware server
		metadataGroup.POST("/middleware-server/all", metadata.GetMiddlewareServer)
		metadataGroup.POST("/middleware-server/id", metadata.GetMiddlewareServerByID)
		metadataGroup.POST("/middleware-server/cluster-id", metadata.GetMiddlewareServerByClusterID)
		metadataGroup.POST("/middleware-server/host-info", metadata.GetMiddlewareServerByHostInfo)

		metadataGroup.POST("/middleware-server/add", metadata.AddMiddlewareServer)
		metadataGroup.POST("/middleware-server/update", metadata.UpdateMiddlewareServerByID)
		metadataGroup.POST("/middleware-server/delete", metadata.DeleteMiddlewareServerByID)
		// monitor system
		metadataGroup.POST("/monitor-system/all", metadata.GetMonitorSystem)
		metadataGroup.POST("/monitor-system/id", metadata.GetMonitorSystemByID)
		metadataGroup.POST("/monitor-system/env", metadata.GetMonitorSystemByEnv)
		metadataGroup.POST("/monitor-system/host-info", metadata.GetMonitorSystemByHostInfo)

		metadataGroup.POST("/monitor-system/add", metadata.AddMonitorSystem)
		metadataGroup.POST("/monitor-system/update", metadata.UpdateMonitorSystemByID)
		metadataGroup.POST("/monitor-system/delete", metadata.DeleteMonitorSystemByID)
		// mysql cluster
		metadataGroup.POST("/mysql-cluster/all", metadata.GetMySQLCluster)
		metadataGroup.POST("/mysql-cluster/id", metadata.GetMySQLClusterByID)
		metadataGroup.POST("/mysql-cluster/env", metadata.GetMySQLClusterByEnv)
		metadataGroup.POST("/mysql-cluster/cluster-name", metadata.GetMySQLClusterByName)
		metadataGroup.POST("/mysql-cluster/mysql-server", metadata.GetMySQLServersByID)
		metadataGroup.POST("/mysql-cluster/master-server", metadata.GetMasterServersByID)
		metadataGroup.POST("/mysql-cluster/db", metadata.GetDBsByMySQLClusterID)
		metadataGroup.POST("/mysql-cluster/user", metadata.GetUsersByMySQLClusterID)
		metadataGroup.POST("/mysql-cluster/app-user", metadata.GetAppUsersByMySQLClusterID)
		metadataGroup.POST("/mysql-cluster/db-user", metadata.GetDBUsersByMySQLClusterID)
		metadataGroup.POST("/mysql-cluster/all-user", metadata.GetAllUsersByMySQLClusterID)
		metadataGroup.POST("/mysql-cluster/resource-group", metadata.GetResourceGroupByMySQLClusterID)

		metadataGroup.POST("/mysql-cluster/add", metadata.AddMySQLCluster)
		metadataGroup.POST("/mysql-cluster/update", metadata.UpdateMySQLClusterByID)
		metadataGroup.POST("/mysql-cluster/delete", metadata.DeleteMySQLClusterByID)
		metadataGroup.POST("/mysql-cluster/add-user", metadata.MySQLClusterAddUser)
		metadataGroup.POST("/mysql-cluster/delete-user", metadata.MySQLClusterDeleteUser)
		// mysql server
		metadataGroup.POST("/mysql-server/all", metadata.GetMySQLServer)
		metadataGroup.POST("/mysql-server/id", metadata.GetMySQLServerByID)
		metadataGroup.POST("/mysql-server/cluster-id", metadata.GetMySQLServerByClusterID)
		metadataGroup.POST("/mysql-server/host-info", metadata.GetMySQLServerByHostInfo)
		metadataGroup.POST("/mysql-server/is-master/host-info", metadata.IsMaster)
		metadataGroup.POST("/mysql-server/mysql-cluster", metadata.GetMySQLClusterByMySQLServerID)

		metadataGroup.POST("/mysql-server/add", metadata.AddMySQLServer)
		metadataGroup.POST("/mysql-server/update", metadata.UpdateMySQLServerByID)
		metadataGroup.POST("/mysql-server/delete", metadata.DeleteMySQLServerByID)
		// resource group
		metadataGroup.POST("/resource-group/all", metadata.GetResourceGroup)
		metadataGroup.POST("/resource-group/id", metadata.GetResourceGroupByID)
		metadataGroup.POST("/resource-group/group-uuid", metadata.GetResourceGroupByGroupUUID)
		metadataGroup.POST("/resource-group/resource-role/id", metadata.GetResourceRolesByGroupID)
		metadataGroup.POST("/resource-group/mysql-cluster/id", metadata.GetMySQLClustersByGroupID)
		metadataGroup.POST("/resource-group/mysql-server/id", metadata.GetMySQLServersByGroupID)
		metadataGroup.POST("/resource-group/middleware-cluster/id", metadata.GetMiddlewareClustersByGroupID)
		metadataGroup.POST("/resource-group/middleware-server/id", metadata.GetMiddlewareServersByGroupID)
		metadataGroup.POST("/resource-group/user/id", metadata.GetUsersByGroupID)
		metadataGroup.POST("/resource-group/das-admin/id", metadata.GetDASAdminUsersByGroupID)
		metadataGroup.POST("/resource-group/resource-role/group-uuid", metadata.GetResourceRolesByGroupUUID)
		metadataGroup.POST("/resource-group/mysql-cluster/group-uuid", metadata.GetMySQLClustersByGroupUUID)
		metadataGroup.POST("/resource-group/mysql-server/group-uuid", metadata.GetMySQLServersByGroupUUID)
		metadataGroup.POST("/resource-group/middleware-cluster/group-uuid", metadata.GetMiddlewareClustersByGroupUUID)
		metadataGroup.POST("/resource-group/middleware-server/group-uuid", metadata.GetMiddlewareServersByGroupUUID)
		metadataGroup.POST("/resource-group/user/group-uuid", metadata.GetUsersByGroupUUID)
		metadataGroup.POST("/resource-group/das-admin/group-uuid", metadata.GetDASAdminUsersByGroupUUID)

		metadataGroup.POST("/resource-group/add", metadata.AddResourceGroup)
		metadataGroup.POST("/resource-group/update", metadata.UpdateResourceGroupByID)
		metadataGroup.POST("/resource-group/delete", metadata.DeleteResourceGroupByID)
		metadataGroup.POST("/resource-group/add-mysql-cluster", metadata.ResourceGroupAddMySQLCluster)
		metadataGroup.POST("/resource-group/delete-mysql-cluster", metadata.ResourceGroupDeleteMySQLCluster)
		metadataGroup.POST("/resource-group/add-middleware-cluster", metadata.ResourceGroupAddMiddlewareCluster)
		metadataGroup.POST("/resource-group/delete-middleware-cluster", metadata.ResourceGroupDeleteMiddlewareCluster)
		// resource role
		metadataGroup.POST("/resource-role/all", metadata.GetResourceRole)
		metadataGroup.POST("/resource-role/id", metadata.GetResourceRoleByID)
		metadataGroup.POST("/resource-role/role-uuid", metadata.GetResourceRoleByUUID)
		metadataGroup.POST("/resource-role/resource-group", metadata.GetResourceGroupByResourceRoleID)
		metadataGroup.POST("/resource-role/user/id", metadata.GetUsersByResourceRoleID)
		metadataGroup.POST("/resource-role/user/role-uuid", metadata.GetUsersByResourceRoleUUID)

		metadataGroup.POST("/resource-role/add", metadata.AddResourceRole)
		metadataGroup.POST("/resource-role/update", metadata.UpdateResourceRoleByID)
		metadataGroup.POST("/resource-role/delete", metadata.DeleteResourceRoleByID)
		metadataGroup.POST("/resource-role/add-user", metadata.ResourceRoleAddUser)
		metadataGroup.POST("/resource-role/delete-user", metadata.ResourceRoleDeleteUser)
		// table
		metadataGroup.POST("/table/db", metadata.GetTablesByDBID)
		metadataGroup.POST("/table/statistic/db", metadata.GetStatisticsByDBIDAndTableName)
		metadataGroup.POST("/table/statistic/host-info-db", metadata.GetStatisticsByHostInfoAndDBNameAndTableName)

		metadataGroup.POST("/table/analyze/db", metadata.AnalyzeTableByDBIDAndTableName)
		metadataGroup.POST("/table/analyze/host-info-db", metadata.AnalyzeTableByHostInfoAndDBNameAndTableName)
		// user
		metadataGroup.POST("/user/all", metadata.GetUser)
		metadataGroup.POST("/user/id", metadata.GetUserByID)
		metadataGroup.POST("/user/user-name", metadata.GetByUserName)
		metadataGroup.POST("/user/employee-id", metadata.GetUserByEmployeeID)
		metadataGroup.POST("/user/account-name", metadata.GetUserByAccountName)
		metadataGroup.POST("/user/login-name", metadata.GetByAccountNameOrEmployeeID)
		metadataGroup.POST("/user/email", metadata.GetUserByEmail)
		metadataGroup.POST("/user/telephone", metadata.GetUserByTelephone)
		metadataGroup.POST("/user/mobile", metadata.GetUserByMobile)
		metadataGroup.POST("/user/app", metadata.GetAppsByUserID)
		metadataGroup.POST("/user/db", metadata.GetDBsByUserID)
		metadataGroup.POST("/user/middleware-cluster", metadata.GetMiddlewareClustersByUserID)
		metadataGroup.POST("/user/mysql-cluster", metadata.GetMySQLClustersByUserID)
		metadataGroup.POST("/user/all-mysql-server", metadata.GetAllMySQLServersByUserID)

		metadataGroup.POST("/user/add", metadata.AddUser)
		metadataGroup.POST("/user/update", metadata.UpdateUserByID)
		metadataGroup.POST("/user/delete", metadata.DeleteUserByID)
	}
}
