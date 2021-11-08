package alert

import (
	"net/http"
	"testing"

	"github.com/romberli/das/internal/dependency/alert"
	"github.com/stretchr/testify/assert"
)

const (
	testHTTPURL = "http://127.0.0.1:8081"
)

var (
	_ alert.Sender = (*HTTPSender)(nil)
)

func initHTTPClient() *http.Client {
	return &http.Client{
		Transport: defaultTransport,
		Timeout:   defaultDialTimeout,
	}
}

func TestHTTPALL(t *testing.T) {
	TestHTTP_Send(t)
}

func TestHTTP_Send(t *testing.T) {
	asst := assert.New(t)
	s := initService()
	s.setupHTTPConfig(testToAddrs, testCCAddrs, testContent)

	as := newHTTPSender(initHTTPClient(), s.GetConfig(), testHTTPURL)
	asst.Equal(nil, as.Send(), "Test HTTP_Send() failed")
}
