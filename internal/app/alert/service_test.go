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
	// before executing unit test, please modify these test constants appropriately
	testSMTPURL  = "smtp.163.com:465"
	testSMTPUser = "dastest@163.com"
	testSMTPPass = "dastest"
	testSMTPFrom = "dastest@163.com"

	testToAddrs = "dastest@163.com"
	testCCAddrs = "dastest@163.com"
	testSubject = "test subject"
	testContent = "test content"

	testDASAddr   = "127.0.0.1:3306"
	testDASDBName = "das"
	testDASDBUser = "root"
	testDASDBPass = "root"
)

func initViper() {
	viper.Set(config.AlertSMTPEnabledKey, true)
	viper.Set(config.AlertSMTPFormatKey, config.AlertSMTPHTMLFormat)
	viper.Set(config.AlertSMTPURLKey, testSMTPURL)
	viper.Set(config.AlertSMTPUserKey, testSMTPUser)
	viper.Set(config.AlertSMTPPassKey, testSMTPPass)
	viper.Set(config.AlertSMTPFromKey, testSMTPFrom)
}

func initService() (s *Service) {
	initViper()
	cfg := NewConfigFromFile()
	cr := NewRepositoryWithGlobal()
	s = newService(cr, cfg)
	pool, _ := mysql.NewPoolWithDefault(testDASAddr, testDASDBName, testDASDBUser, testDASDBPass)
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
	s.setupSMTPConfig(testToAddrs, testCCAddrs, testSubject, testContent)
	err := s.SendEmail(testToAddrs, testCCAddrs, testSubject, testContent)
	asst.Nil(err, common.CombineMessageWithError("test SendEmail() failed", err))
}

func TestAlertService_sendViaSMTP(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupSMTPConfig(testToAddrs, testCCAddrs, testSubject, testContent)
	err := s.sendViaSMTP(testToAddrs, testCCAddrs, testSubject, testContent)
	asst.Nil(err, common.CombineMessageWithError("test sendViaSMTP() failed", err))

}

func TestAlertService_sendViaHTTP(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupHTTPConfig(testToAddrs, testCCAddrs, testContent)
	err := s.sendViaHTTP(testToAddrs, testCCAddrs, testContent)
	asst.Nil(err, common.CombineMessageWithError("test sendViaHTTP() failed", err))

}

func TestAlertService_saveSMTP(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupSMTPConfig(testToAddrs, testCCAddrs, testSubject, testContent)
	err := s.saveSMTP(testToAddrs, testCCAddrs, testSubject, testContent, "test")
	asst.Nil(err, common.CombineMessageWithError("test saveSMTP() failed", err))

}

func TestAlertService_saveHTTP(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupHTTPConfig(testToAddrs, testCCAddrs, testContent)
	err := s.saveHTTP(testToAddrs, testCCAddrs, testContent, "test")
	asst.Nil(err, common.CombineMessageWithError("test saveHTTP() failed", err))

}
