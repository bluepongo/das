package alert

import (
	"net/http"
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const testHTTPURL = "http://127.0.0.1:8081"

var testHTTPSender *HTTPSender

func init() {
	testHTTPSender = testInitHTTPSender()
}

func testInitHTTPSender() *HTTPSender {
	client := &http.Client{
		Transport: defaultTransport,
		Timeout:   defaultDialTimeout,
	}
	cfg := NewEmptyConfig()
	cfg.Set(toAddrsJSON, testToAddrs+constant.CommaString+ccAddrsJSON)
	cfg.Set(ccAddrsJSON, testCCAddrs)
	cfg.Set(contentJSON, testContent)

	return newHTTPSender(client, cfg, testHTTPURL)
}

func TestHTTP_ALL(t *testing.T) {
	TestHTTP_Send(t)
}

func TestHTTP_Send(t *testing.T) {
	asst := assert.New(t)

	err := testHTTPSender.Send()
	asst.Nil(err, common.CombineMessageWithError("Test Send() failed", err))
}
