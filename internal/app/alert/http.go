package alert

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/pingcap/errors"

	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/dependency/alert"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
	"github.com/spf13/viper"
	"github.com/tidwall/pretty"
)

const (
	defaultDialTimeout         = 30 * time.Second
	defaultKeepAlive           = 30 * time.Second
	defaultTLSHandshakeTimeout = 10 * time.Second
	defaultContentType         = "application/json"
)

var (
	_ alert.Sender = (*HTTPSender)(nil)

	defaultTransport = &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		DialContext:         (&net.Dialer{Timeout: defaultDialTimeout, KeepAlive: defaultKeepAlive}).DialContext,
		TLSHandshakeTimeout: defaultTLSHandshakeTimeout,
	}
)

type HTTPSender struct {
	client *http.Client
	config alert.Config
	url    string
}

// NewHTTPSender returns a new alert.Sender
func NewHTTPSender(client *http.Client, cfg alert.Config, url string) alert.Sender {
	return newHTTPSender(client, cfg, url)
}

// NewHTTPSenderWithDefault returns a new alert.Sender with default http client
func NewHTTPSenderWithDefault(cfg alert.Config) alert.Sender {
	client := &http.Client{
		Transport: defaultTransport,
		Timeout:   defaultDialTimeout,
	}
	url := viper.GetString(config.AlertHTTPURLKey)

	return newHTTPSender(client, cfg, url)
}

// newHTTPSender returns a new *HTTPSender
func newHTTPSender(client *http.Client, cfg alert.Config, url string) *HTTPSender {
	return &HTTPSender{
		client: client,
		config: cfg,
		url:    url,
	}
}

// GetClient returns the http client
func (hs *HTTPSender) GetClient() *http.Client {
	return hs.client
}

// GetConfig return the config
func (hs *HTTPSender) GetConfig() alert.Config {
	return hs.config
}

// GetURL returns the http api url
func (hs *HTTPSender) GetURL() string {
	return hs.url
}

// Send sends the email via http api calling
func (hs *HTTPSender) Send() error {
	// get request body
	reqBody, err := hs.buildRequestBody()
	if err != nil {
		return err
	}

	log.Debugf("alert HTTPSender.Send() http body: %s", string(reqBody))
	// call http api
	resp, err := hs.GetClient().Post(hs.GetURL(), defaultContentType, bytes.NewBuffer(reqBody))
	if err != nil {
		return errors.Trace(err)
	}

	return hs.parseResponse(resp)
}

// buildRequestBody builds the http request body, for now, it basically marshals the config
func (hs *HTTPSender) buildRequestBody() ([]byte, error) {
	jsonBytes, err := json.MarshalIndent(hs.GetConfig(), constant.EmptyString, constant.DefaultIndentString)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return pretty.Pretty(jsonBytes), nil
}

// parseResponse parses the http response to find out if sending email completed successfully
func (hs *HTTPSender) parseResponse(resp *http.Response) error {
	// read response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Trace(err)
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("got http error when calling alert http api. status code: %d, message: %s",
			resp.StatusCode, string(respBody))
	}

	return errors.Trace(resp.Body.Close())
}
