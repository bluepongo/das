package alert

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
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

	sendEmailRespMessage = `{"message": "sending email completed successfully"}`
)

// @Tags 	alert
// @Summary send email
// @Accept	application/json
// @Param	token 	body string true	"token"
// @Param	config 	body string false	"optional config"
// @Param	toAddrs body string true	"to addrs"
// @Param	ccAddrs body string true 	"cc addrs"
// @Param	content body string true	"to content"
// @Produce application/json
// @Success 200 {string} string "{"message": "sending email completed successfully"}"
// @Router	/api/v1/alert/email [post]
func SendEmail(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, err)
		return
	}
	dataMap := make(map[string]string)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}

	config := alert.NewConfigFromFile()
	configStr, configExists := dataMap[configJSON]
	if configExists {
		err = json.Unmarshal([]byte(configStr), &config)
		if err != nil {
			resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
			return
		}
	}

	toAddrs, toAddrsExists := dataMap[toAddrsJSON]
	if !toAddrsExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, toAddrsJSON)
		return
	}
	ccAddrs := dataMap[ccAddrsJSON]
	subject, contentExists := dataMap[subjectJSON]
	if !contentExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, subjectJSON)
		return
	}
	content, contentExists := dataMap[contentJSON]
	if !contentExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, contentJSON)
		return
	}

	s := alert.NewServiceWithDefault(config)
	err = s.SendEmail(toAddrs, ccAddrs, subject, content)
	if err != nil {
		resp.ResponseNOK(c, msgalert.ErrServiceSendEmail, err, config, toAddrs, ccAddrs, subject)
		return
	}

	log.Debug(message.NewMessage(msgalert.DebugServiceSendEmail, config, toAddrs, ccAddrs, subject, content).Error())
	resp.ResponseOK(c, sendEmailRespMessage, msgalert.InfoServiceSendEmail, config, toAddrs, ccAddrs, subject)
}
