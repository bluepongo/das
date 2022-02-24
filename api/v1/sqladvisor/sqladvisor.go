package sqladvisor

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
	"github.com/romberli/das/internal/app/sqladvisor"
	"github.com/romberli/das/pkg/message"
	msgadvisor "github.com/romberli/das/pkg/message/sqladvisor"
	"github.com/romberli/das/pkg/resp"
	"github.com/romberli/go-util/constant"
)

const (
	sqlTextJSON     = "sql_text"
	fingerprintJSON = "fingerprint"
	sqlIDJSON       = "sql_id"
	dbIDJSON        = "db_id"
)

// @Tags sqladvisor
// @Summary get sql fingerprint
// @Accept  application/json
// @Param	sql_text body string true "sql text"
// @Produce application/json
// @Success 200 {string} string "{"fingerprint": "select * from a","sql_text": "select * from a;"}"
// @Router	/api/v1/sqladvisor/fingerprint/ [get]
func GetFingerprint(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}

	dataMap := make(map[string]string)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}

	sqlText, exists := dataMap[sqlTextJSON]
	if !exists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, sqlTextJSON)
		return
	}
	// init service
	service := sqladvisor.NewServiceWithDefault()
	// get fingerprint
	fingerprint := service.GetFingerprint(sqlText)
	respData := map[string]string{sqlTextJSON: sqlText, fingerprintJSON: fingerprint}
	respMessage, err := json.Marshal(respData)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, errors.Trace(err))
		return
	}

	resp.ResponseOK(c, string(respMessage), msgadvisor.InfoSQLAdvisorGetFingerprint, sqlTextJSON, fingerprint)
}

// @Tags sqladvisor
// @Summary get sql id
// @Accept  application/json
// @Param	sql_text body string true "sql text"
// @Produce  application/json
// @Success 200 {string} string "{"sql_id": "EE56B94E867DC9D5","sql_text": "select * from a;"}"
// @Router /api/v1/sqladvisor/sql-id/ [get]
func GetSQLID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}

	dataMap := make(map[string]string)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}

	sqlText, exists := dataMap[sqlTextJSON]
	if !exists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, sqlTextJSON)
		return
	}
	// init service
	service := sqladvisor.NewServiceWithDefault()
	// get sql id
	sqlID := service.GetSQLID(sqlText)
	respData := map[string]string{sqlTextJSON: sqlText, sqlIDJSON: sqlID}
	respMessage, err := json.Marshal(respData)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, errors.Trace(err))
		return
	}

	resp.ResponseOK(c, string(respMessage), msgadvisor.InfoSQLAdvisorGetSQLID, sqlTextJSON, sqlID)
}

// @Tags sqladvisor
// @Summary get advice
// @Accept  application/json
// @Param	db_id		path int	true "db id"
// @Param	sql_text	body string true "sql text"
// @Produce application/json
// @Success 200 {string} string "{"code": 200, "data": [{"sql_text": "select * from t01", "advice": "xxx"}"
// @Router /api/v1/sqladvisor/advise/:db_id [post]
func Advise(c *gin.Context) {
	// get data
	dbIDStr := c.Param(dbIDJSON)
	if dbIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, dbIDJSON)
		return
	}
	dbID, err := strconv.Atoi(dbIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}

	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}

	dataMap := make(map[string]string)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}

	sqlText, exists := dataMap[sqlTextJSON]
	if !exists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, sqlTextJSON)
		return
	}
	// init service
	service := sqladvisor.NewServiceWithDefault()
	advice, err := service.Advise(dbID, sqlText)
	if err != nil {
		resp.ResponseNOK(c, msgadvisor.ErrSQLAdvisorAdvice, err, dbID, sqlText)
		return
	}

	resp.ResponseOK(c, advice, msgadvisor.InfoSQLAdvisorAdvice, dbID, sqlText, advice)
}
