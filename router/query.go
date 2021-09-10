package router

import (
	"github.com/gin-gonic/gin"
	"github.com/romberli/das/api/v1/query"
)

// RegisterQuery is the sub-router of das for query
func RegisterQuery(group *gin.RouterGroup) {
	queryGroup := group.Group("/query")
	{
		queryGroup.GET("/cluster/:mysql_cluster_id", query.GetByMySQLClusterID)
		queryGroup.GET("/server/:mysql_server_id", query.GetByMySQLServerID)
		queryGroup.GET("/db/:db_id", query.GetByDBID)
		queryGroup.GET("/:sql_id", query.GetBySQLID)
	}
}
