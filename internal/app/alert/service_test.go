package alert

import (
	"testing"

	"github.com/romberli/das/config"
	"github.com/romberli/go-util/common"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var testService *Service

func init() {
	testInitDASMySQLPool()
	testInitViper()
	testService = newService(testRepo, NewConfigFromFile())
}

func testInitViper() {
	viper.Set(config.AlertSMTPEnabledKey, true)
	viper.Set(config.AlertSMTPFormatKey, config.AlertSMTPHTMLFormat)
	viper.Set(config.AlertSMTPURLKey, testSMTPURL)
	viper.Set(config.AlertSMTPUserKey, testSMTPUser)
	viper.Set(config.AlertSMTPPassKey, testSMTPPass)
	viper.Set(config.AlertSMTPFromKey, testSMTPFrom)
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

	err := testService.SendEmail(testToAddrs, testCCAddrs, testSubject, testContent)
	asst.Nil(err, common.CombineMessageWithError("test SendEmail() failed", err))
}

func TestAlertService_sendViaSMTP(t *testing.T) {
	asst := assert.New(t)

	testService.setupSMTPConfig(testToAddrs, testCCAddrs, testSubject, testContent)
	err := testService.sendViaSMTP(testToAddrs, testCCAddrs, testSubject, testContent)
	asst.Nil(err, common.CombineMessageWithError("test sendViaSMTP() failed", err))
}

func TestAlertService_sendViaHTTP(t *testing.T) {
	asst := assert.New(t)

	testService.setupHTTPConfig(testToAddrs, testCCAddrs, testContent)
	err := testService.sendViaHTTP(testToAddrs, testCCAddrs, testContent)
	asst.Nil(err, common.CombineMessageWithError("test sendViaHTTP() failed", err))

}

func TestAlertService_saveSMTP(t *testing.T) {
	asst := assert.New(t)

	testService.setupSMTPConfig(testToAddrs, testCCAddrs, testSubject, testContent)
	err := testService.saveSMTP(testToAddrs, testCCAddrs, testSubject, testContent, "test")
	asst.Nil(err, common.CombineMessageWithError("test saveSMTP() failed", err))

}

func TestAlertService_saveHTTP(t *testing.T) {
	asst := assert.New(t)

	testService.setupHTTPConfig(testToAddrs, testCCAddrs, testContent)
	err := testService.saveHTTP(testToAddrs, testCCAddrs, testContent, "test")
	asst.Nil(err, common.CombineMessageWithError("test saveHTTP() failed", err))
}
