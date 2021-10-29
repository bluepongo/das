package alert

import (
	"fmt"
	"testing"

	"github.com/romberli/das/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestSMTP_ALL(t *testing.T) {
	TestSMTP_NewSMTPSender(t)
	TestSMTP_NewSMTPSenderWithDefault(t)
	TestSMTP_GetConfig(t)
	TestSMTP_GetURL(t)
	TestSMTP_Send(t)

}

func TestSMTP_NewSMTPSender(t *testing.T) {
	asst := assert.New(t)
	s := initService()
	s.setupSMTPConfig(toAddrs, ccAddrs, subject, content)

	as := NewSMTPSender(s.config, url)

	asst.Equal(url, as.GetURL(), "Test for NewSMTPSender() failed")
}

func TestSMTP_NewSMTPSenderWithDefault(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupSMTPConfig(toAddrs, ccAddrs, subject, content)

	as := NewSMTPSenderWithDefault(s.config)

	asst.Equal(viper.GetString(config.AlertSMTPAddrKey), as.GetURL(), "Test for NewSMTPSenderWithDefault failed")
}

func TestSMTP_GetConfig(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupSMTPConfig(toAddrs, ccAddrs, subject, content)
	as := newSMTPSender(s.config, url)

	asst.Equal(fmt.Sprintln(s.config), fmt.Sprintln(as.GetConfig()), "Test for GetConfig() Failed")

}

func TestSMTP_GetURL(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupSMTPConfig(toAddrs, ccAddrs, subject, content)
	as := newSMTPSender(s.config, url)
	asst.Equal(url, fmt.Sprint(as.GetURL()), "Test for GetConfig() Failed")

}

func TestSMTP_Send(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupSMTPConfig(toAddrs, ccAddrs, subject, content)
	as := newSMTPSender(s.config, url)
	asst.Equal(nil, as.Send(), "Test SMTP_Send Failed")
}
