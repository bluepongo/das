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
		metadataGroup.GET("/app/:id", metadata.GetAppByID)
		metadataGroup.GET("/app/app-name/:name", metadata.GetAppByName)
		metadataGroup.GET("/app/dbs/:id", metadata.GetDBIDList)
		metadataGroup.POST("/app", metadata.AddApp)
		metadataGroup.POST("/app/:id", metadata.UpdateAppByID)
		metadataGroup.POST("/app/delete/:id", metadata.DeleteAppByID)
		metadataGroup.POST("/app/add-db/:id", metadata.AppAddDB)
		metadataGroup.POST("/app/delete-db/:id", metadata.AppDeleteDB)
		// db
		metadataGroup.GET("/db", metadata.GetDB)
		metadataGroup.GET("/db/env/:env_id", metadata.GetDBByEnv)
		metadataGroup.GET("/db/:id", metadata.GetDBByID)
		metadataGroup.GET("db/apps/:id", metadata.GetAppIDList)
		metadataGroup.POST("/db", metadata.AddDB)
		metadataGroup.POST("/db/:id", metadata.UpdateDBByID)
		metadataGroup.POST("/db/delete/:id", metadata.DeleteDBByID)
		metadataGroup.POST("/db/add-db/:id", metadata.DBAddApp)
		metadataGroup.POST("/db/delete-db/:id", metadata.DBDeleteApp)
		// env
		metadataGroup.GET("/env", metadata.GetEnv)
		metadataGroup.GET("/env/:id", metadata.GetEnvByID)
		metadataGroup.GET("/env/env-name/:name", metadata.GetEnvByName)
		metadataGroup.POST("/env", metadata.AddEnv)
		metadataGroup.POST("/env/:id", metadata.UpdateEnvByID)
		metadataGroup.POST("/env/delete/:id", metadata.DeleteEnvByID)
		// middleware cluster
		metadataGroup.GET("/middleware-cluster", metadata.GetMiddlewareCluster)
		metadataGroup.GET("/middleware-cluster/env/:env_id", metadata.GetMiddlewareClusterByEnv)
		metadataGroup.GET("/middleware-cluster/:id", metadata.GetMiddlewareClusterByID)
		metadataGroup.GET("/middleware-cluster/cluster-name/:name", metadata.GetMiddlewareClusterByName)
		metadataGroup.GET("/middleware-cluster/middleware-server/:id", metadata.GetMiddlewareServerIDList)
		metadataGroup.POST("/middleware-cluster", metadata.AddMiddlewareCluster)
		metadataGroup.POST("/middleware-cluster/:id", metadata.UpdateMiddlewareClusterByID)
		metadataGroup.POST("/middleware-cluster/delete/:id", metadata.DeleteMiddlewareClusterByID)
		// middleware server
		metadataGroup.GET("/middleware-server", metadata.GetMiddlewareServer)
		metadataGroup.GET("/middleware-server/cluster-id/:cluster_id", metadata.GetMiddlewareServerByClusterID)
		metadataGroup.GET("/middleware-server/:id", metadata.GetMiddlewareServerByID)
		metadataGroup.GET("/middleware-server/host-info", metadata.GetMiddlewareServerByHostInfo)
		metadataGroup.POST("/middleware-server", metadata.AddMiddlewareServer)
		metadataGroup.POST("/middleware-server/:id", metadata.UpdateMiddlewareServerByID)
		metadataGroup.POST("/middleware-server/delete/:id", metadata.DeleteMiddlewareClusterByID)
		// monitor system
		metadataGroup.GET("/monitor-system", metadata.GetMonitorSystem)
		metadataGroup.GET("/monitor-system/env/:env_id", metadata.GetMonitorSystemByEnv)
		metadataGroup.GET("/monitor-system/:id", metadata.GetMonitorSystemByID)
		metadataGroup.GET("/monitor-system/host-info", metadata.GetMonitorSystemByHostInfo)
		metadataGroup.POST("/monitor-system", metadata.AddMonitorSystem)
		metadataGroup.POST("/monitor-system/:id", metadata.UpdateMonitorSystemByID)
		metadataGroup.POST("/monitor-system/delete/:id", metadata.DeleteMonitorSystemByID)
		// mysql cluster
		metadataGroup.GET("/mysql-cluster", metadata.GetMySQLCluster)
		metadataGroup.GET("/mysql-cluster/env/:env_id", metadata.GetMySQLClusterByEnv)
		metadataGroup.GET("/mysql-cluster/:id", metadata.GetMySQLClusterByID)
		metadataGroup.GET("/mysql-cluster/cluster-name/:name", metadata.GetMySQLClusterByName)
		metadataGroup.GET("/mysql-cluster/mysql-server/:id", metadata.GetMySQLServerIDList)
		metadataGroup.POST("/mysql-cluster", metadata.AddMySQLCluster)
		metadataGroup.POST("/mysql-cluster/:id", metadata.UpdateMySQLClusterByID)
		metadataGroup.POST("/mysql-cluster/delete/:id", metadata.DeleteMySQLClusterByID)
		// mysql server
		metadataGroup.GET("/mysql-server", metadata.GetMySQLServer)
		metadataGroup.GET("/mysql-server/cluster-id", metadata.GetMySQLServerByClusterID)
		metadataGroup.GET("/mysql-server/:id", metadata.GetMySQLServerByID)
		metadataGroup.GET("/mysql-server/host-info", metadata.GetMySQLServerByHostInfo)
		metadataGroup.POST("/mysql-server", metadata.AddMySQLServer)
		metadataGroup.POST("/mysql-server/:id", metadata.UpdateMySQLServerByID)
		metadataGroup.POST("/mysql-server/delete/:id", metadata.DeleteMySQLServerByID)
		// user
		metadataGroup.GET("/user", metadata.GetUser)
		metadataGroup.GET("/user/user-name/:name", metadata.GetUserByName)
		metadataGroup.GET("/user/:id", metadata.GetUserByID)
		metadataGroup.GET("/user/employee-id/:employee_id", metadata.GetUserByEmployeeID)
		metadataGroup.GET("/user/account-name/:name", metadata.GetUserByAccountName)
		metadataGroup.GET("/user/email/:email", metadata.GetUserByEmail)
		metadataGroup.GET("/user/telephone/:telephone", metadata.GetUserByTelephone)
		metadataGroup.GET("/user/mobile/:mobile", metadata.GetUserByMobile)
		metadataGroup.POST("/user", metadata.AddUser)
		metadataGroup.POST("/user/:id", metadata.UpdateUserByID)
		metadataGroup.POST("/user/delete/:id", metadata.DeleteUserByID)
	}
}
