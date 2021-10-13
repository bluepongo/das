package alert

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/dependency/alert"
	"github.com/romberli/go-util/constant"
	"github.com/spf13/viper"
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

// NewHTTTPSender returns a new alert.Sender
func NewHTTTPSender(client *http.Client, cfg Config, url string) alert.Sender {
	return newHTTTPSender(client, cfg, url)
}

// NewHTTTPSenderWithDefault returns a new alert.Sender with default http client
func NewHTTTPSenderWithDefault(cfg alert.Config) alert.Sender {
	client := &http.Client{
		Transport: defaultTransport,
		Timeout:   defaultDialTimeout,
	}
	url := viper.GetString(config.AlertHTTPURLKey)

	return newHTTTPSender(client, cfg, url)
}

// NewHTTTPSender returns a new *HTTPSender
func newHTTTPSender(client *http.Client, cfg alert.Config, url string) *HTTPSender {
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
	// call http api
	resp, err := hs.GetClient().Post(hs.GetURL(), defaultContentType, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	return hs.parseResponse(resp)
}

// buildRequestBody builds the http request body, for now, it basically marshals the config
func (hs *HTTPSender) buildRequestBody() ([]byte, error) {
	return json.Marshal(hs.GetConfig())
}

// parseResponse parses the http response to find out if sending email completed successfully
func (hs *HTTPSender) parseResponse(resp *http.Response) error {
	defer func() {
		_ = resp.Body.Close()
	}()

	// read response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("got http error when calling alert http api. status code:%d, message: %s",
			resp.StatusCode, string(respBody)))
	}
	// unmarshal to a map
	response := NewEmptyResponse()
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return err
	}
	// get value from the map
	if response.GetCode() != constant.ZeroInt {
		return errors.New(fmt.Sprintf("got internal error when calling alert http api. code: %d, message: %s",
			response.GetCode(), response.GetMessage()))
	}

	return nil
}