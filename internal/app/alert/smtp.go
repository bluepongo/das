package alert

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"

	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/dependency/alert"
	"github.com/romberli/go-util/constant"
	"github.com/spf13/viper"
)

const (
	defaultAlertSMTPFromName    = "DAS"
	defaultAlertSMTPContentText = "text/plain; charset=UTF-8"
	defaultAlertSMTPContentHTML = "text/html; charset=UTF-8"

	smtpHeaderFrom        = "From"
	smtpHeaderTo          = "To"
	smtpHeaderCC          = "Cc"
	smtpHeaderSubject     = "Subject"
	smtpHeaderContentType = "Content-Type"
)

var (
	_ alert.Sender = (*SMTPSender)(nil)
)

type SMTPSender struct {
	client *smtp.Client
	config alert.Config
	url    string
}

// NewSMTPSender returns a new alert.Sender
func NewSMTPSender(client *smtp.Client, cfg alert.Config, url string) alert.Sender {
	return newSMTPSender(client, cfg, url)
}

// NewSMTPSenderWithDefault returns a new alert.Sender with default SMTP client
func NewSMTPSenderWithDefault(cfg alert.Config) (alert.Sender, error) {
	url := viper.GetString(config.AlertSMTPURLKey)
	host, _, _ := net.SplitHostPort(url)
	// init tls connection
	conn, err := tls.Dial(constant.TransportProtocolTCP, url, nil)
	if err != nil {
		return nil, err
	}
	// get smtp client
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return nil, err
	}
	// auth
	err = client.Auth(
		smtp.PlainAuth(
			constant.EmptyString,
			viper.GetString(config.AlertSMTPUserKey),
			viper.GetString(config.AlertSMTPPassKey),
			host,
		),
	)
	if err != nil {
		return nil, err
	}

	return newSMTPSender(client, cfg, url), nil
}

// newSMTPSender returns a new *SMTPSender
func newSMTPSender(client *smtp.Client, cfg alert.Config, url string) *SMTPSender {
	return &SMTPSender{
		client: client,
		config: cfg,
		url:    url,
	}
}

// GetClient returns the smtp client
func (ss *SMTPSender) GetClient() *smtp.Client {
	return ss.client
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
	var ccList []string

	toAddrs := ss.GetConfig().Get(toAddrsJSON)
	if toAddrs == constant.EmptyString {
		return errors.New("to address could not be empty")
	}

	ccAddrs := ss.GetConfig().Get(ccAddrsJSON)
	if ccAddrs != constant.EmptyString {
		ccList = strings.Split(ccAddrs, constant.CommaString)
	}

	return ss.sendMail(
		ss.GetConfig().Get(smtpFromAddrJson),
		strings.Split(toAddrs, constant.CommaString),
		ccList,
		ss.buildMessage(),
	)
}

// sendMail sends mail
func (ss *SMTPSender) sendMail(from string, toList, ccList []string, message []byte) error {
	err := ss.GetClient().Mail(from)
	if err != nil {
		return err
	}
	if len(toList) == constant.ZeroInt {
		return fmt.Errorf("toList could not be empty")
	}
	for _, to := range toList {
		err = ss.GetClient().Rcpt(to)
		if err != nil {
			return err
		}
	}
	ccList = toList
	for _, cc := range ccList {
		err = ss.GetClient().Rcpt(cc)
		if err != nil {
			return err
		}
	}

	w, err := ss.GetClient().Data()
	if err != nil {
		return err
	}

	_, err = w.Write(message)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}

	return ss.GetClient().Quit()
}

// buildMessage return msg
func (ss *SMTPSender) buildMessage() []byte {
	message := constant.EmptyString

	for k, v := range ss.buildHeader() {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	message += fmt.Sprintf("\r\n%s", ss.GetConfig().Get(contentJSON))

	return []byte(message)
}

// buildHeader returns an SMTPSender header
func (ss *SMTPSender) buildHeader() map[string]string {
	header := make(map[string]string)

	header[smtpHeaderFrom] = fmt.Sprintf("%s<%s>", defaultAlertSMTPFromName, ss.GetConfig().Get(smtpFromAddrJson))
	header[smtpHeaderTo] = ss.GetConfig().Get(toAddrsJSON)
	header[smtpHeaderCC] = ss.GetConfig().Get(ccAddrsJSON)
	header[smtpHeaderSubject] = ss.GetConfig().Get(subjectJSON)

	switch viper.GetString(config.AlertSMTPFormatKey) {
	case config.AlertSMTPTextFormat:
		header[smtpHeaderContentType] = defaultAlertSMTPContentText
	case config.AlertSMTPHTMLFormat:
		header[smtpHeaderContentType] = defaultAlertSMTPContentHTML
	}

	return header
}
