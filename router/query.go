package router

import (
	"github.com/gin-gonic/gin"
	"github.com/romberli/das/api/v1/query"
)

// RegisterQuery is the sub-router of das for query
func RegisterQuery(group *gin.RouterGroup) {
	queryGroup := group.Group("/query")
	{
		queryGroup.POST("/cluster", query.GetByMySQLClusterID)
		queryGroup.POST("/server", query.GetByMySQLServerID)
		queryGroup.POST("/host-info", query.GetByHostInfo)
		queryGroup.POST("/db", query.GetByDBID)
		queryGroup.POST("/sql", query.GetBySQLID)
	}
}
