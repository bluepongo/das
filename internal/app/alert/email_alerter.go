package alert

import (
	"github.com/romberli/das/internal/dependency/alert"
)

const (
	defaultEmailURL = ""
)

var _ alert.EmailAlerter = (*EmailAlerter)(nil)

type EmailAlerter struct {
	url     string
	toAddr  []string
	ccAddr  []string
	content string
}

func NewEmailAlerter(url string, toAddr, ccAddr []string, content string) *EmailAlerter {
	return newEmailAlerter(url, toAddr, ccAddr, content)
}

func NewEmailAlerterWithDefault(toAddr, ccAddr []string, content string) *EmailAlerter {
	return newEmailAlerter(defaultEmailURL, toAddr, ccAddr, content)
}

func newEmailAlerter(url string, toAddr, ccAddr []string, content string) *EmailAlerter {
	return &EmailAlerter{
		url:     url,
		toAddr:  toAddr,
		ccAddr:  ccAddr,
		content: content,
	}
}

func (ea *EmailAlerter) GetURL() string {
	return ea.url
}

func (ea *EmailAlerter) GetToAddr() []string {
	return ea.toAddr
}

func (ea *EmailAlerter) GetCCAddr() []string {
	return ea.ccAddr
}

func (ea *EmailAlerter) GetContent() string {
	return ea.content
}

func (ea *EmailAlerter) Send() error {
	return nil
}
