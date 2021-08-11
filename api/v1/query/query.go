package query

import (
	"github.com/gin-gonic/gin"
)

const (
	idJSON        = "id"
	dbIDJSON      = "db_id"
	startTimeJSON = "start_time"
	endTimeJSON   = "end_time"
	limitJSON     = "limit"
	offsetJSON    = "offset"
)

// @Tags query
// @Summary get all slow queries
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/query [get]
func GetAll(c *gin.Context) {

}

// @Tags query
// @Summary get slow queries by db id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/query/db/:dbID [get]
func GetByDBID(c *gin.Context) {

}

// @Tags query
// @Summary get slow query by query id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/query/:id [get]
func GetByID(c *gin.Context) {

}
