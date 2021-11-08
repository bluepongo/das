package alert

import (
	"testing"

	"github.com/romberli/das/internal/dependency/alert"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
)

func TestSMTP_ALL(t *testing.T) {
	TestSMTP_Send(t)
}

func initSMTPSender() alert.Sender {
	initViper()

	cfg := NewEmptyConfig()
	cfg.Set(toAddrsJSON, testToAddrs)
	cfg.Set(ccAddrsJSON, testCCAddrs)
	cfg.Set(subjectJSON, testSubject)
	cfg.Set(contentJSON, testContent)
	cfg.Set(smtpFromAddrJson, testSMTPFrom)
	s, err := NewSMTPSenderWithDefault(cfg)
	if err != nil {
		log.Errorf("init smtp sender failed.\n%s", err.Error())
	}

	return s
}

func TestSMTP_Send(t *testing.T) {
	asst := assert.New(t)

	ss := initSMTPSender()
	err := ss.Send()
	asst.Equal(nil, err, "test Send() failed")
}
