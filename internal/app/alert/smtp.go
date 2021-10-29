package alert

import (
	"strconv"

	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/dependency/alert"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

var (
	_ alert.Sender = (*SMTPSender)(nil)
)

type SMTPSender struct {
	config alert.Config
	url    string
}

// NewSMTPSender returns a new alert.Sender
func NewSMTPSender(cfg alert.Config, url string) alert.Sender {
	return newSMTPSender(cfg, url)
}

// NewSMTPSenderWithDefault returns a new alert.Sender with default SMTP client
func NewSMTPSenderWithDefault(cfg alert.Config) alert.Sender {
	url := viper.GetString(config.AlertSMTPAddrKey)

	return newSMTPSender(cfg, url)
}

// newSMTPSender returns a new *SMTPSender
func newSMTPSender(cfg alert.Config, url string) *SMTPSender {
	return &SMTPSender{
		config: cfg,
		url:    url,
	}
}

// GetConfig returns the config
func (ss *SMTPSender) GetConfig() alert.Config {
	return ss.config
}

// GetAddr returns the addr
func (ss *SMTPSender) GetURL() string {
	return ss.url
}

// Send sends the email via the api calling
func (ss *SMTPSender) Send() error {

	port, _ := strconv.Atoi(ss.GetConfig().Get(smtpPortJSON))

	m := gomail.NewMessage()
	m.SetHeader("From", config.AlertSMTPFromName+"<"+ss.GetConfig().Get(smtpUserJSON)+">")
	m.SetHeader("To", ss.GetConfig().Get(toAddrsJSON))
	m.SetHeader("Cc", ss.GetConfig().Get(ccAddrsJSON))
	m.SetHeader("Subject", ss.GetConfig().Get(subjectJSON))
	m.SetBody("text/html", ss.GetConfig().Get(contentJSON))
	d := gomail.NewDialer(ss.GetURL(), port, ss.GetConfig().Get(smtpUserJSON), ss.GetConfig().Get(smtpPassJSON))

	err := d.DialAndSend(m)
	return err
}
