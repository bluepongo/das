package alert

import (
	"net/http"
	"testing"

	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/dependency/alert"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	_ alert.Sender = (*HTTPSender)(nil)
)

func makeClient() (client *http.Client) {
	client = &http.Client{
		Transport: defaultTransport,
		Timeout:   defaultDialTimeout,
	}
	return
}

func TestHTTPALL(t *testing.T) {
	TestHTTP_NewHTTPSender(t)
	TestHTTP_NewHTTPSenderWithDefault(t)
	TestHTTP_GetClient(t)
	TestHTTP_GetConfig(t)
	TestHTTP_GetURL(t)
	TestHTTP_Send(t)
}

func TestHTTP_NewHTTPSender(t *testing.T) {
	asst := assert.New(t)
	s := initService()
	s.setupHTTPConfig(toAddrs, ccAddrs, content)

	as := NewHTTPSender(makeClient(), s.GetConfig(), url)

	asst.Equal(url, as.GetURL(), "Test NewHTTPSender() failed")
}

func TestHTTP_NewHTTPSenderWithDefault(t *testing.T) {
	asst := assert.New(t)
	s := initService()
	s.setupHTTPConfig(toAddrs, ccAddrs, content)

	as := NewHTTPSenderWithDefault(s.GetConfig())

	asst.Equal(viper.GetString(config.AlertHTTPURLKey), as.GetURL(), "Test NewHTTTPSenderWithDefault() failed")
}

func TestHTTP_GetClient(t *testing.T) {
	asst := assert.New(t)
	s := initService()
	s.setupHTTPConfig(toAddrs, ccAddrs, content)

	as := newHTTPSender(makeClient(), s.GetConfig(), url)
	asst.Equal(makeClient(), as.GetClient(), "Test GetClient() failed")
}

func TestHTTP_GetConfig(t *testing.T) {
	asst := assert.New(t)
	s := initService()
	s.setupHTTPConfig(toAddrs, ccAddrs, content)

	as := newHTTPSender(makeClient(), s.GetConfig(), url)
	asst.Equal(s.config, as.GetConfig(), "Test GetConfig() failed")
}
func TestHTTP_GetURL(t *testing.T) {
	asst := assert.New(t)
	s := initService()
	s.setupHTTPConfig(toAddrs, ccAddrs, content)

	as := newHTTPSender(makeClient(), s.GetConfig(), url)
	asst.Equal(url, as.GetURL(), "Test GetURL() failed")
}
func TestHTTP_Send(t *testing.T) {
	asst := assert.New(t)
	s := initService()
	s.setupHTTPConfig(toAddrs, ccAddrs, content)

	as := newHTTPSender(makeClient(), s.GetConfig(), url)
	asst.Equal(nil, as.Send(), "Test HTTP_Send() failed")
}
