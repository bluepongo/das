package alert

import (
	"testing"

	"github.com/romberli/das/internal/dependency/alert"
	"github.com/romberli/go-util/common"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
)

const (
	testSMTPURL  = "smtp.163.com:465"
	testSMTPUser = "allinemailtest@163.com"
	testSMTPPass = "LAOMDMZSOMKCZICJ"
	testSMTPFrom = "allinemailtest@163.com"
)

var testSMTPSender alert.Sender

func init() {
	testInitViper()
	testSMTPSender = testInitSMTPSender()
}

func TestSMTP_All(t *testing.T) {
	TestSMTP_Send(t)
}

func testInitSMTPSender() alert.Sender {
	cfg := NewEmptyConfig()
	cfg.Set(toAddrsJSON, testToAddrs)
	cfg.Set(ccAddrsJSON, testCCAddrs)
	cfg.Set(subjectJSON, testSubject)
	cfg.Set(contentJSON, testContent)
	cfg.Set(smtpFromAddrJson, testSMTPFrom)

	sender, err := NewSMTPSenderWithDefault(cfg)
	if err != nil {
		log.Errorf("init smtp sender failed.\n%s", err.Error())
	}

	return sender
}

func TestSMTP_Send(t *testing.T) {
	asst := assert.New(t)

	err := testSMTPSender.Send()
	asst.Nil(err, common.CombineMessageWithError("Test Send() failed", err))
}
