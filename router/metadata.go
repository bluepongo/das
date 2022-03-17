package router

import (
	"github.com/gin-gonic/gin"

	"github.com/romberli/das/api/v1/metadata"
)

func RegisterMetadata(group *gin.RouterGroup) {
	metadataGroup := group.Group("/metadata")
	{
		// app
		metadataGroup.GET("/app", metadata.GetApp)
		metadataGroup.GET("/app/get/:id", metadata.GetAppByID)
		metadataGroup.GET("/app/app-name/:app_name", metadata.GetAppByName)
		metadataGroup.GET("/app/db/:id", metadata.GetDBsByAppID)
		metadataGroup.POST("/app", metadata.AddApp)
		metadataGroup.POST("/app/update/:id", metadata.UpdateAppByID)
		metadataGroup.POST("/app/delete/:id", metadata.DeleteAppByID)
		metadataGroup.POST("/app/add-db/:id", metadata.AppAddDB)
		metadataGroup.POST("/app/delete-db/:id", metadata.AppDeleteDB)
		metadataGroup.GET("/app/user/:id", metadata.GetUsersByAppID)
		metadataGroup.POST("/app/add-user/:id", metadata.AppAddUser)
		metadataGroup.POST("/app/delete-user/:id", metadata.AppDeleteUser)
		// db
		metadataGroup.GET("/db", metadata.GetDB)
		metadataGroup.GET("/db/env/:env_id", metadata.GetDBByEnv)
		metadataGroup.GET("/db/get/:id", metadata.GetDBByID)
		metadataGroup.GET("/db/name-and-cluster-info", metadata.GetDBByNameAndClusterInfo)
		metadataGroup.GET("/db/name-and-host-info", metadata.GetDBByNameAndHostInfo)
		metadataGroup.GET("/db/host-info", metadata.GetDBsByHostInfo)
		metadataGroup.GET("/db/app/:id", metadata.GetAppsByDBID)
		metadataGroup.GET("/db/mysql-cluster/:id", metadata.GetMySQLClusterByDBID)
		metadataGroup.GET("/db/app-user/:id", metadata.GetAppUsersByDBID)
		metadataGroup.GET("/db/db-user/:id", metadata.GetUsersByDBID)
		metadataGroup.GET("/db/all-user/:id", metadata.GetAllUsersByDBID)
		metadataGroup.POST("/db", metadata.AddDB)
		metadataGroup.POST("/db/update/:id", metadata.UpdateDBByID)
		metadataGroup.POST("/db/delete/:id", metadata.DeleteDBByID)
		metadataGroup.POST("/db/add-app/:id", metadata.DBAddApp)
		metadataGroup.POST("/db/delete-app/:id", metadata.DBDeleteApp)
		metadataGroup.POST("/db/add-user/:id", metadata.DBAddUser)
		metadataGroup.POST("/db/delete-user/:id", metadata.DBDeleteUser)
		// env
		metadataGroup.GET("/env", metadata.GetEnv)
		metadataGroup.GET("/env/get/:id", metadata.GetEnvByID)
		metadataGroup.GET("/env/env-name/:env_name", metadata.GetEnvByName)
		metadataGroup.POST("/env", metadata.AddEnv)
		metadataGroup.POST("/env/update/:id", metadata.UpdateEnvByID)
		metadataGroup.POST("/env/delete/:id", metadata.DeleteEnvByID)
		// middleware cluster
		metadataGroup.GET("/middleware-cluster", metadata.GetMiddlewareCluster)
		metadataGroup.GET("/middleware-cluster/env/:env_id", metadata.GetMiddlewareClusterByEnv)
		metadataGroup.GET("/middleware-cluster/get/:id", metadata.GetMiddlewareClusterByID)
		metadataGroup.GET("/middleware-cluster/cluster-name/:cluster_name", metadata.GetMiddlewareClusterByName)
		metadataGroup.GET("/middleware-cluster/middleware-server/:id", metadata.GetMiddlewareServers)
		metadataGroup.GET("/middleware-cluster/user/:id", metadata.GetUsersByMiddlewareClusterID)
		metadataGroup.POST("/middleware-cluster", metadata.AddMiddlewareCluster)
		metadataGroup.POST("/middleware-cluster/update/:id", metadata.UpdateMiddlewareClusterByID)
		metadataGroup.POST("/middleware-cluster/delete/:id", metadata.DeleteMiddlewareClusterByID)
		metadataGroup.POST("/middleware-cluster/add-user/:id", metadata.MiddlewareClusterAddUser)
		metadataGroup.POST("/middleware-cluster/delete-user/:id", metadata.MiddlewareClusterDeleteUser)
		// middleware server
		metadataGroup.GET("/middleware-server", metadata.GetMiddlewareServer)
		metadataGroup.GET("/middleware-server/cluster-id/:cluster_id", metadata.GetMiddlewareServerByClusterID)
		metadataGroup.GET("/middleware-server/get/:id", metadata.GetMiddlewareServerByID)
		metadataGroup.GET("/middleware-server/host-info", metadata.GetMiddlewareServerByHostInfo)
		metadataGroup.POST("/middleware-server", metadata.AddMiddlewareServer)
		metadataGroup.POST("/middleware-server/update/:id", metadata.UpdateMiddlewareServerByID)
		metadataGroup.POST("/middleware-server/delete/:id", metadata.DeleteMiddlewareServerByID)
		// monitor system
		metadataGroup.GET("/monitor-system", metadata.GetMonitorSystem)
		metadataGroup.GET("/monitor-system/env/:env_id", metadata.GetMonitorSystemByEnv)
		metadataGroup.GET("/monitor-system/get/:id", metadata.GetMonitorSystemByID)
		metadataGroup.GET("/monitor-system/host-info", metadata.GetMonitorSystemByHostInfo)
		metadataGroup.POST("/monitor-system", metadata.AddMonitorSystem)
		metadataGroup.POST("/monitor-system/update/:id", metadata.UpdateMonitorSystemByID)
		metadataGroup.POST("/monitor-system/delete/:id", metadata.DeleteMonitorSystemByID)
		// mysql cluster
		metadataGroup.GET("/mysql-cluster", metadata.GetMySQLCluster)
		metadataGroup.GET("/mysql-cluster/env/:env_id", metadata.GetMySQLClusterByEnv)
		metadataGroup.GET("/mysql-cluster/get/:id", metadata.GetMySQLClusterByID)
		metadataGroup.GET("/mysql-cluster/cluster-name/:name", metadata.GetMySQLClusterByName)
		metadataGroup.GET("/mysql-cluster/mysql-server/:id", metadata.GetMySQLServersByID)
		metadataGroup.GET("/mysql-cluster/master-server/:id", metadata.GetMasterServersByID)
		metadataGroup.GET("/mysql-cluster/db/:id", metadata.GetDBsByMySQLClusterID)
		metadataGroup.GET("/mysql-cluster/user/:id", metadata.GetUsersByMySQLClusterID)
		metadataGroup.GET("/mysql-cluster/app-user/:id", metadata.GetAppUsersByMySQLClusterID)
		metadataGroup.GET("/mysql-cluster/db-user/:id", metadata.GetDBUsersByMySQLClusterID)
		metadataGroup.GET("/mysql-cluster/all-user/:id", metadata.GetAllUsersByMySQLClusterID)
		metadataGroup.POST("/mysql-cluster/add-user/:id", metadata.MySQLClusterAddUser)
		metadataGroup.POST("/mysql-cluster/delete-user/:id", metadata.MySQLClusterDeleteUser)
		metadataGroup.POST("/mysql-cluster", metadata.AddMySQLCluster)
		metadataGroup.POST("/mysql-cluster/update/:id", metadata.UpdateMySQLClusterByID)
		metadataGroup.POST("/mysql-cluster/delete/:id", metadata.DeleteMySQLClusterByID)
		// mysql server
		metadataGroup.GET("/mysql-server", metadata.GetMySQLServer)
		metadataGroup.GET("/mysql-server/cluster-id/:cluster_id", metadata.GetMySQLServerByClusterID)
		metadataGroup.GET("/mysql-server/get/:id", metadata.GetMySQLServerByID)
		metadataGroup.GET("/mysql-server/host-info", metadata.GetMySQLServerByHostInfo)
		metadataGroup.GET("/mysql-server/is-master/host-info", metadata.IsMaster)
		metadataGroup.GET("/mysql-server/mysql-cluster/:id", metadata.GetMySQLClusterByMySQLServerID)
		metadataGroup.POST("/mysql-server", metadata.AddMySQLServer)
		metadataGroup.POST("/mysql-server/update/:id", metadata.UpdateMySQLServerByID)
		metadataGroup.POST("/mysql-server/delete/:id", metadata.DeleteMySQLServerByID)
		// user
		metadataGroup.GET("/user", metadata.GetUser)
		metadataGroup.GET("/user/get/:id", metadata.GetUserByID)
		metadataGroup.GET("/user/user-name/:user_name", metadata.GetByUserName)
		metadataGroup.GET("/user/employee-id/:employee_id", metadata.GetUserByEmployeeID)
		metadataGroup.GET("/user/account-name/:account_name", metadata.GetUserByAccountName)
		metadataGroup.GET("/user/login-name/:login_name", metadata.GetByAccountNameOrEmployeeID)
		metadataGroup.GET("/user/email/:email", metadata.GetUserByEmail)
		metadataGroup.GET("/user/telephone/:telephone", metadata.GetUserByTelephone)
		metadataGroup.GET("/user/mobile/:mobile", metadata.GetUserByMobile)
		metadataGroup.GET("/user/app/:id", metadata.GetAppsByUserID)
		metadataGroup.GET("/user/db/:id", metadata.GetDBsByUserID)
		metadataGroup.GET("/user/middleware-cluster/:id", metadata.GetMiddlewareClustersByUserID)
		metadataGroup.GET("/user/mysql-cluster/:id", metadata.GetMySQLClustersByUserID)
		metadataGroup.GET("/user/all-mysql-server/:id", metadata.GetAllMySQLServersByUserID)
		metadataGroup.POST("/user", metadata.AddUser)
		metadataGroup.POST("/user/update/:id", metadata.UpdateUserByID)
		metadataGroup.POST("/user/delete/:id", metadata.DeleteUserByID)
		// table
		metadataGroup.GET("/table/db/:id", metadata.GetTablesByDBID)
		metadataGroup.GET("/table/statistic/db-table", metadata.GetTableByDBIDAndTableName)
		metadataGroup.GET("/table/statistic/host-info-db-table", metadata.GetTableByHostInfoAndDBNameAndTableName)
		// metadataGroup.POST("/table/analyze/db", metadata.AnalyzeTableByDBIDAndTableName)
		// metadataGroup.POST("/table/analyze/host-info", metadata.AnalyzeTableByHostInfoAndDBNameAndTableName)
	}
}
