package alert

import (
	"testing"

	"github.com/romberli/das/config"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	url       = "smtp.163.com:465"
	userAddrs = ""
	toAddrs   = ""
	ccAddrs   = ""
	subject   = "test subject"
	content   = "test content"
	testPass  = ""
)

const (
	appAddr                     = "192.168.10.210:3306"
	appDBName                   = "das"
	appDBUser                   = "root"
	appDBPass                   = "root"
	onlineAppName               = "2"
	newAppName                  = "testApp"
	defaultID                   = 1
	defaultAlertSMTPEnabled     = true
	defaultAlertSMTPHTMLEnabled = true
	defaultAlertHTTPEnabled     = true
)

func initViper() {

	viper.Set(config.AlertSMTPEnabledKey, defaultAlertSMTPEnabled)
	viper.Set(config.AlertSMTPHTMLEnabledKey, defaultAlertSMTPHTMLEnabled)
	viper.Set(config.AlertSMTPURLKey, url)
	//
	viper.Set(config.AlertSMTPUserKey, userAddrs)
	viper.Set(config.AlertSMTPFromKey, toAddrs)
	viper.Set(config.AlertSMTPPassKey, testPass)
}

func initService() (s *Service) {
	initViper()
	config := NewConfigFromFile()
	cr := NewRepositoryWithGlobal()
	s = newService(cr, config)
	pool, _ := mysql.NewPoolWithDefault(appAddr, appDBName, appDBUser, appDBPass)
	s.Repository = NewRepository(pool)
	return

}

func TestAppRepoAll(t *testing.T) {
	TestAlertService_SendEmail(t)
	TestAlertService_sendViaSMTP(t)
	TestAlertService_sendViaHTTP(t)
	TestAlertService_saveSMTP(t)
	TestAlertService_saveHTTP(t)

}

func TestAlertService_SendEmail(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupSMTPConfig(toAddrs, ccAddrs, subject, content)
	err := s.SendEmail(toAddrs, ccAddrs, subject, content)
	asst.Nil(err, common.CombineMessageWithError("test SendEmail() failed", err))
}

func TestAlertService_sendViaSMTP(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupSMTPConfig(toAddrs, ccAddrs, subject, content)
	err := s.sendViaSMTP(toAddrs, ccAddrs, subject, content)
	asst.Nil(err, common.CombineMessageWithError("test sendViaSMTP() failed", err))

}

func TestAlertService_sendViaHTTP(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupHTTPConfig(toAddrs, ccAddrs, content)
	err := s.sendViaHTTP(toAddrs, ccAddrs, content)
	asst.Nil(err, common.CombineMessageWithError("test sendViaHTTP() failed", err))

}

func TestAlertService_saveSMTP(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupSMTPConfig(toAddrs, ccAddrs, subject, content)
	err := s.saveSMTP(toAddrs, ccAddrs, subject, content, "test")
	asst.Nil(err, common.CombineMessageWithError("test saveSMTP() failed", err))

}

func TestAlertService_saveHTTP(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupHTTPConfig(toAddrs, ccAddrs, content)
	err := s.saveHTTP(toAddrs, ccAddrs, content, "test")
	asst.Nil(err, common.CombineMessageWithError("test saveHTTP() failed", err))

}
