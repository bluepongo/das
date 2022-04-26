package health

import (
	"github.com/gin-gonic/gin"
	msghealth "github.com/romberli/das/pkg/message/health"
	"github.com/romberli/das/pkg/resp"
)

const (
	pongString = `{"ping": "pong"}`
)

// @Tags health
// @Summary ping
// @Accept	application/json
// @Param	token body string true "token"
// @Produce application/json
// @Success 200 {string} string "{"code": 200, "data": {"ping": "pong"}}"
// @Router	/api/v1/health/ping [get]
func Ping(c *gin.Context) {
	resp.ResponseOK(c, pongString, msghealth.InfoHealthPing)
}
