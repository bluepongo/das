package alert

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/stretchr/testify/assert"
)

const (
	message = "mestest"
	cfg     = "{\"pass\": \"****\"}"
)

func TestRepoALL(t *testing.T) {
	TestRepository_Save(t)
	TestRepository_Execute(t)
}

func TestRepository_Execute(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupSMTPConfig(testToAddrs, testCCAddrs, testSubject, testContent)
	sr := s.Repository
	sql := `
	insert into t_alert_operation_info(url, to_addrs, cc_addrs, subject, content, config, message)
	values(?, ?, ?, ?, ?, ?, ?);
`
	_, err := sr.Execute(sql, testSMTPURL, testToAddrs, testCCAddrs, testSubject, testContent, cfg, message)

	asst.Nil(err, common.CombineMessageWithError("test Save() failed", err))

}

func TestRepository_Save(t *testing.T) {
	asst := assert.New(t)

	s := initService()
	s.setupSMTPConfig(testToAddrs, testCCAddrs, testSubject, testContent)
	sr := s.Repository
	err := sr.Save(testSMTPURL, testToAddrs, testCCAddrs, testSubject, testContent, cfg, message)
	asst.Nil(err, common.CombineMessageWithError("test Save() failed", err))

	s.setupHTTPConfig(testToAddrs, testCCAddrs, testContent)
	err = sr.Save(testHTTPURL, testToAddrs, testCCAddrs, testSubject, testContent, cfg, message)

	asst.Nil(err, common.CombineMessageWithError("test Save() failed", err))
}
