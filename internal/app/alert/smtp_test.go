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
	client := Dial(url)
	as := NewSMTPSender(client, s.config, url)

	asst.Equal(url, as.GetURL(), "Test for NewSMTPSender() failed")
}

func TestSMTP_NewSMTPSenderWithDefault(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupSMTPConfig(toAddrs, ccAddrs, subject, content)

	as := NewSMTPSenderWithDefault(s.config)

	asst.Equal(viper.GetString(config.AlertSMTPURLKey), as.GetURL(), "Test for NewSMTPSenderWithDefault failed")
}

func TestSMTP_GetConfig(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupSMTPConfig(toAddrs, ccAddrs, subject, content)
	client := Dial(url)
	as := newSMTPSender(client, s.config, url)

	asst.Equal(fmt.Sprint(s.config), fmt.Sprint(as.GetConfig()), "Test for GetConfig() Failed")

}

func TestSMTP_GetURL(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupSMTPConfig(toAddrs, ccAddrs, subject, content)
	client := Dial(url)
	as := newSMTPSender(client, s.config, url)
	asst.Equal(url, as.GetURL(), "Test for GetConfig() Failed")

}

func TestSMTP_Send(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupSMTPConfig(toAddrs, ccAddrs, subject, content)
	client := Dial(url)
	as := newSMTPSender(client, s.config, url)
	asst.Equal(nil, as.Send(), "Test SMTP_Send Failed")
}
