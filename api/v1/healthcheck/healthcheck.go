package healthcheck

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/romberli/das/internal/app/healthcheck"
	"github.com/romberli/das/pkg/message"
	msghealth "github.com/romberli/das/pkg/message/healthcheck"
	"github.com/romberli/das/pkg/resp"
	utilhealth "github.com/romberli/das/pkg/util/healthcheck"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
)

const (
	operationIDJSON            = "operation_id"
	reviewJSON                 = "review"
	checkRespMessage           = `{"code": 0, "operation_id: %d", message": "healthcheck started"}`
	checkByHostInfoRespMessage = `{"code": 0, "operation_id: %d", "message": "healthcheck by host info started"}`
	reviewAccuracyRespMessage  = `{"code": 0, "message": "reviewed accuracy"}`
)

// @Tags healthcheck
// @Summary get result by operation id
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": []}"
// @Router /api/v1/healthcheck/result/:id [get]
func GetResultByOperationID(c *gin.Context) {
	// get data
	operationIDStr := c.Param(operationIDJSON)
	if operationIDStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, operationIDJSON)
		return
	}
	operationID, err := strconv.Atoi(operationIDStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, err.Error())
		return
	}
	// init service
	s := healthcheck.NewServiceWithDefault()
	// get entities
	err = s.GetResultByOperationID(operationID)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckGetResultByOperationID, err.Error())
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err.Error())
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msghealth.DebugHealthcheckGetResultByOperationID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msghealth.InfoHealthcheckGetResultByOperationID, operationID)
}

// @Tags healthcheck
// @Summary check health of the database
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": "healthcheck started."}"
// @Router /api/v1/healthcheck/check [post]
func Check(c *gin.Context) {
	var rd *utilhealth.Check
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	startTime, err := time.ParseInLocation(constant.TimeLayoutSecond, rd.GetStartTime(), time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, rd.GetStartTime())
		return
	}
	endTime, err := time.ParseInLocation(constant.TimeLayoutSecond, rd.GetEndTime(), time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, rd.GetEndTime())
		return
	}
	step, err := time.ParseDuration(rd.GetStep())
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeDuration, rd.GetStep())
		return
	}
	// init service
	s := healthcheck.NewServiceWithDefault()
	// check health
	operationID, err := s.Check(rd.GetServerID(), startTime, endTime, step)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckCheck, operationID, err.Error())
		return
	}

	log.Debug(message.NewMessage(msghealth.DebugHealthcheckCheck, operationID).Error())
	resp.ResponseOK(c, fmt.Sprintf(checkRespMessage, operationID), msghealth.InfoHealthcheckCheck, operationID)
}

// @Tags healthcheck
// @Summary check health of the database by host ip and port number
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": ""}"
// @Router /api/v1/healthcheck/check/host-info [post]
func CheckByHostInfo(c *gin.Context) {
	var rd *utilhealth.CheckByHostInfo
	// bind json
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}
	startTime, err := time.ParseInLocation(constant.TimeLayoutSecond, rd.GetStartTime(), time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, rd.GetStartTime())
		return
	}
	endTime, err := time.ParseInLocation(constant.TimeLayoutSecond, rd.GetEndTime(), time.Local)
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeLayout, rd.GetEndTime())
		return
	}
	step, err := time.ParseDuration(rd.GetStep())
	if err != nil {
		resp.ResponseNOK(c, message.ErrNotValidTimeDuration, rd.GetStep())
		return
	}
	// init service
	s := healthcheck.NewServiceWithDefault()
	// get entities
	operationID, err := s.CheckByHostInfo(rd.GetHostIP(), rd.GetPortNum(), startTime, endTime, step)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckCheckByHostInfo, operationID, err.Error())
		return
	}

	log.Debug(message.NewMessage(msghealth.DebugHealthcheckCheckByHostInfo, operationID).Error())
	resp.ResponseOK(c, fmt.Sprintf(checkByHostInfoRespMessage, operationID), msghealth.InfoHealthcheckCheckByHostInfo, operationID)
}

// @Tags healthcheck
// @Summary update accuracy review
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": "{}"
// @Router /api/v1/healthcheck/review [post]
func ReviewAccuracy(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	dataMap := make(map[string]int)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckCheck, err.Error())
		return
	}
	operationID, operationIDExists := dataMap[operationIDJSON]
	if !operationIDExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, operationIDJSON)
		return
	}
	review, reviewExists := dataMap[reviewJSON]
	if !reviewExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, reviewJSON)
		return
	}
	// init service
	s := healthcheck.NewServiceWithDefault()
	// review accuracy
	err = s.ReviewAccuracy(operationID, review)
	if err != nil {
		resp.ResponseNOK(c, msghealth.ErrHealthcheckReviewAccuracy, operationID, err.Error())
		return
	}

	log.Debug(message.NewMessage(msghealth.DebugHealthcheckReviewAccuracy, reviewAccuracyRespMessage).Error())
	resp.ResponseOK(c, reviewAccuracyRespMessage, msghealth.InfoHealthcheckReviewAccuracy)
}
