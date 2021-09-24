package alert

import (
	"github.com/gin-gonic/gin"
)

// @Tags alert
// @Summary send email
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/alert/email [post]
func SendEmail(c *gin.Context) {
	// get data

}
