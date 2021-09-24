package router

import (
	"github.com/gin-gonic/gin"
	"github.com/romberli/das/api/v1/alert"
)

// RegisterAlert is the sub-router of das for alert
func RegisterAlert(group *gin.RouterGroup) {
	alertGroup := group.Group("/alert")
	{
		alertGroup.POST("/alert/email", alert.SendEmail)
	}
}
