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
		metadataGroup.GET("/app/get", metadata.GetAppByID)
		metadataGroup.GET("/app/app-name", metadata.GetAppByName)
		metadataGroup.GET("/app/db", metadata.GetDBsByAppID)
		metadataGroup.POST("/app", metadata.AddApp)
		metadataGroup.POST("/app/update", metadata.UpdateAppByID)
		metadataGroup.POST("/app/delete", metadata.DeleteAppByID)
		metadataGroup.POST("/app/add-db", metadata.AppAddDB)
		metadataGroup.POST("/app/delete-db", metadata.AppDeleteDB)
		metadataGroup.GET("/app/user", metadata.GetUsersByAppID)
		metadataGroup.POST("/app/add-user", metadata.AppAddUser)
		metadataGroup.POST("/app/delete-user", metadata.AppDeleteUser)
		// db
		metadataGroup.GET("/db", metadata.GetDB)
		metadataGroup.GET("/db/env", metadata.GetDBByEnv)
		metadataGroup.GET("/db/get", metadata.GetDBByID)
		metadataGroup.GET("/db/name-and-cluster-info", metadata.GetDBByNameAndClusterInfo)
		metadataGroup.GET("/db/name-and-host-info", metadata.GetDBByNameAndHostInfo)
		metadataGroup.GET("/db/host-info", metadata.GetDBsByHostInfo)
		metadataGroup.GET("/db/app", metadata.GetAppsByDBID)
		metadataGroup.GET("/db/mysql-cluster", metadata.GetMySQLClusterByDBID)
		metadataGroup.GET("/db/app-user", metadata.GetAppUsersByDBID)
		metadataGroup.GET("/db/db-user", metadata.GetUsersByDBID)
		metadataGroup.GET("/db/all-user", metadata.GetAllUsersByDBID)
		metadataGroup.POST("/db", metadata.AddDB)
		metadataGroup.POST("/db/update", metadata.UpdateDBByID)
		metadataGroup.POST("/db/delete", metadata.DeleteDBByID)
		metadataGroup.POST("/db/add-app", metadata.DBAddApp)
		metadataGroup.POST("/db/delete-app", metadata.DBDeleteApp)
		metadataGroup.POST("/db/add-user", metadata.DBAddUser)
		metadataGroup.POST("/db/delete-user", metadata.DBDeleteUser)
		// env
		metadataGroup.GET("/env", metadata.GetEnv)
		metadataGroup.GET("/env/get", metadata.GetEnvByID)
		metadataGroup.GET("/env/env-name", metadata.GetEnvByName)
		metadataGroup.POST("/env", metadata.AddEnv)
		metadataGroup.POST("/env/update", metadata.UpdateEnvByID)
		metadataGroup.POST("/env/delete", metadata.DeleteEnvByID)
		// middleware cluster
		metadataGroup.GET("/middleware-cluster", metadata.GetMiddlewareCluster)
		metadataGroup.GET("/middleware-cluster/env", metadata.GetMiddlewareClusterByEnv)
		metadataGroup.GET("/middleware-cluster/get", metadata.GetMiddlewareClusterByID)
		metadataGroup.GET("/middleware-cluster/cluster-name", metadata.GetMiddlewareClusterByName)
		metadataGroup.GET("/middleware-cluster/middleware-server", metadata.GetMiddlewareServers)
		metadataGroup.GET("/middleware-cluster/user", metadata.GetUsersByMiddlewareClusterID)
		metadataGroup.POST("/middleware-cluster", metadata.AddMiddlewareCluster)
		metadataGroup.POST("/middleware-cluster/update", metadata.UpdateMiddlewareClusterByID)
		metadataGroup.POST("/middleware-cluster/delete", metadata.DeleteMiddlewareClusterByID)
		metadataGroup.POST("/middleware-cluster/add-user", metadata.MiddlewareClusterAddUser)
		metadataGroup.POST("/middleware-cluster/delete-user", metadata.MiddlewareClusterDeleteUser)
		// middleware server
		metadataGroup.GET("/middleware-server", metadata.GetMiddlewareServer)
		metadataGroup.GET("/middleware-server/cluster-id", metadata.GetMiddlewareServerByClusterID)
		metadataGroup.GET("/middleware-server/get", metadata.GetMiddlewareServerByID)
		metadataGroup.GET("/middleware-server/host-info", metadata.GetMiddlewareServerByHostInfo)
		metadataGroup.POST("/middleware-server", metadata.AddMiddlewareServer)
		metadataGroup.POST("/middleware-server/update", metadata.UpdateMiddlewareServerByID)
		metadataGroup.POST("/middleware-server/delete", metadata.DeleteMiddlewareServerByID)
		// monitor system
		metadataGroup.GET("/monitor-system", metadata.GetMonitorSystem)
		metadataGroup.GET("/monitor-system/env", metadata.GetMonitorSystemByEnv)
		metadataGroup.GET("/monitor-system/get", metadata.GetMonitorSystemByID)
		metadataGroup.GET("/monitor-system/host-info", metadata.GetMonitorSystemByHostInfo)
		metadataGroup.POST("/monitor-system", metadata.AddMonitorSystem)
		metadataGroup.POST("/monitor-system/update", metadata.UpdateMonitorSystemByID)
		metadataGroup.POST("/monitor-system/delete", metadata.DeleteMonitorSystemByID)
		// mysql cluster
		metadataGroup.GET("/mysql-cluster", metadata.GetMySQLCluster)
		metadataGroup.GET("/mysql-cluster/env", metadata.GetMySQLClusterByEnv)
		metadataGroup.GET("/mysql-cluster/get", metadata.GetMySQLClusterByID)
		metadataGroup.GET("/mysql-cluster/cluster-name", metadata.GetMySQLClusterByName)
		metadataGroup.GET("/mysql-cluster/mysql-server", metadata.GetMySQLServersByID)
		metadataGroup.GET("/mysql-cluster/master-server", metadata.GetMasterServersByID)
		metadataGroup.GET("/mysql-cluster/db", metadata.GetDBsByMySQLClusterID)
		metadataGroup.GET("/mysql-cluster/user", metadata.GetUsersByMySQLClusterID)
		metadataGroup.GET("/mysql-cluster/app-user", metadata.GetAppUsersByMySQLClusterID)
		metadataGroup.GET("/mysql-cluster/db-user", metadata.GetDBUsersByMySQLClusterID)
		metadataGroup.GET("/mysql-cluster/all-user", metadata.GetAllUsersByMySQLClusterID)
		metadataGroup.POST("/mysql-cluster/add-user", metadata.MySQLClusterAddUser)
		metadataGroup.POST("/mysql-cluster/delete-user", metadata.MySQLClusterDeleteUser)
		metadataGroup.POST("/mysql-cluster", metadata.AddMySQLCluster)
		metadataGroup.POST("/mysql-cluster/update", metadata.UpdateMySQLClusterByID)
		metadataGroup.POST("/mysql-cluster/delete", metadata.DeleteMySQLClusterByID)
		// mysql server
		metadataGroup.GET("/mysql-server", metadata.GetMySQLServer)
		metadataGroup.GET("/mysql-server/cluster-id", metadata.GetMySQLServerByClusterID)
		metadataGroup.GET("/mysql-server/get", metadata.GetMySQLServerByID)
		metadataGroup.GET("/mysql-server/host-info", metadata.GetMySQLServerByHostInfo)
		metadataGroup.GET("/mysql-server/is-master/host-info", metadata.IsMaster)
		metadataGroup.GET("/mysql-server/mysql-cluster", metadata.GetMySQLClusterByMySQLServerID)
		metadataGroup.POST("/mysql-server", metadata.AddMySQLServer)
		metadataGroup.POST("/mysql-server/update", metadata.UpdateMySQLServerByID)
		metadataGroup.POST("/mysql-server/delete", metadata.DeleteMySQLServerByID)
		// user
		metadataGroup.GET("/user", metadata.GetUser)
		metadataGroup.GET("/user/get", metadata.GetUserByID)
		metadataGroup.GET("/user/user-name", metadata.GetByUserName)
		metadataGroup.GET("/user/employee-id", metadata.GetUserByEmployeeID)
		metadataGroup.GET("/user/account-name", metadata.GetUserByAccountName)
		metadataGroup.GET("/user/login-name", metadata.GetByAccountNameOrEmployeeID)
		metadataGroup.GET("/user/email", metadata.GetUserByEmail)
		metadataGroup.GET("/user/telephone", metadata.GetUserByTelephone)
		metadataGroup.GET("/user/mobile", metadata.GetUserByMobile)
		metadataGroup.GET("/user/app", metadata.GetAppsByUserID)
		metadataGroup.GET("/user/db", metadata.GetDBsByUserID)
		metadataGroup.GET("/user/middleware-cluster", metadata.GetMiddlewareClustersByUserID)
		metadataGroup.GET("/user/mysql-cluster", metadata.GetMySQLClustersByUserID)
		metadataGroup.GET("/user/all-mysql-server", metadata.GetAllMySQLServersByUserID)
		metadataGroup.POST("/user", metadata.AddUser)
		metadataGroup.POST("/user/update", metadata.UpdateUserByID)
		metadataGroup.POST("/user/delete", metadata.DeleteUserByID)
		// table
		metadataGroup.GET("/table/db", metadata.GetTablesByDBID)
		metadataGroup.GET("/table/statistic/db", metadata.GetStatisticsByDBIDAndTableName)
		metadataGroup.GET("/table/statistic/host-info-db", metadata.GetStatisticsByHostInfoAndDBNameAndTableName)
		metadataGroup.POST("/table/analyze/db", metadata.AnalyzeTableByDBIDAndTableName)
		metadataGroup.POST("/table/analyze/host-info-db", metadata.AnalyzeTableByHostInfoAndDBNameAndTableName)
	}
}
