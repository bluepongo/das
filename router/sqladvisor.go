package router

import (
	"github.com/gin-gonic/gin"
	"github.com/romberli/das/api/v1/sqladvisor"
)

func RegisterSQLAdvisor(group *gin.RouterGroup) {
	sqladvisorGroup := group.Group("/sqladvisor")
	{
		sqladvisorGroup.POST("/fingerprint", sqladvisor.GetFingerprint)
		sqladvisorGroup.POST("/sql-id", sqladvisor.GetSQLID)
		sqladvisorGroup.POST("/advise", sqladvisor.Advise)
	}
}
