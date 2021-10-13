package alert

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/romberli/das/internal/app/alert"
	"github.com/romberli/das/pkg/message"
	msgalert "github.com/romberli/das/pkg/message/alert"
	"github.com/romberli/das/pkg/resp"
	"github.com/romberli/log"
)

const (
	configJSON  = "config"
	toAddrsJSON = "to_addrs"
	ccAddrsJSON = "cc_addrs"
	subjectJSON = "subject"
	contentJSON = "content"

	sendEmailRespMessage = `{"code": 0, "message": "send email completed successfully"}`
)

// @Tags alert
// @Summary send email
// @Produce  application/json
// @Success 200 {string} string "{"code": 200, "data": {"code": 0, "message": "send email completed successfully"}}"
// @Router /api/v1/alert/email [post]
func SendEmail(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err.Error())
		return
	}
	dataMap := make(map[string]string)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
		return
	}

	config := alert.NewEmptyConfig()
	configStr, configExists := dataMap[configJSON]
	if configExists {
		err = json.Unmarshal([]byte(configStr), &config)
		if err != nil {
			resp.ResponseNOK(c, message.ErrUnmarshalRawData, err.Error())
			return
		}
	}

	toAddrs, toAddrsExists := dataMap[toAddrsJSON]
	if !toAddrsExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, toAddrsJSON)
	}
	ccAddrs := dataMap[ccAddrsJSON]
	subject, contentExists := dataMap[subjectJSON]
	if !contentExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, subjectJSON)
	}
	content, contentExists := dataMap[contentJSON]
	if !contentExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, contentJSON)
	}

	s := alert.NewServiceWithDefault(config)
	err = s.SendEmail(toAddrs, ccAddrs, subject, content)
	if err != nil {
		resp.ResponseNOK(c, msgalert.ErrServiceSendEmail, toAddrs, ccAddrs, subject, content, err.Error())
		return
	}

	log.Debug(message.NewMessage(msgalert.DebugServiceSendEmail, toAddrs, ccAddrs, subject, content).Error())
	resp.ResponseOK(c, sendEmailRespMessage, msgalert.InfoServiceSendEmail, toAddrs, ccAddrs, subject)
}
