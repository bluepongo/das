package router

import (
	"github.com/gin-gonic/gin"
	"github.com/romberli/das/api/v1/healthcheck"
)

// RegisterHealthcheck is the sub-router of das for healthcheck
func RegisterHealthcheck(group *gin.RouterGroup) {
	healthcheckGroup := group.Group("/healthcheck")
	{
		healthcheckGroup.POST("/history", healthcheck.GetOperationHistoriesByLoginName)
		healthcheckGroup.POST("/result", healthcheck.GetResultByOperationID)
		healthcheckGroup.POST("/check", healthcheck.Check)
		healthcheckGroup.POST("/check/host-info", healthcheck.CheckByHostInfo)
		healthcheckGroup.POST("/review", healthcheck.ReviewAccuracy)
	}
}
