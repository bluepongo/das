package router

import (
	"github.com/gin-gonic/gin"
	"github.com/romberli/das/api/v1/query"
)

// RegisterQuery is the sub-router of das for query
func RegisterQuery(group *gin.RouterGroup) {
	queryGroup := group.Group("/query")
	{
		queryGroup.GET("/cluster", query.GetByMySQLClusterID)
		queryGroup.GET("/server", query.GetByMySQLServerID)
		queryGroup.GET("/host-info", query.GetByHostInfo)
		queryGroup.GET("/db", query.GetByDBID)
		queryGroup.GET("/sql", query.GetBySQLID)
	}
}
