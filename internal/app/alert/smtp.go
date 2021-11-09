package alert

import (
	"crypto/tls"
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

	headerFrom        = "From"
	headerTo          = "To"
	headerCc          = "Cc"
	headerSubject     = "Subject"
	headerContentType = "Content-Type"

	crlfString = "\r\n"
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
	return ss.sendMail(
		ss.GetConfig().Get(smtpFromAddrJson),
		strings.Split(ss.GetConfig().Get(toAddrsJSON), constant.CommaString),
		strings.Split(ss.GetConfig().Get(ccAddrsJSON), constant.CommaString),
		ss.buildMessage(),
	)
}

// sendMail sends mail
func (ss *SMTPSender) sendMail(from string, toList []string, ccList []string, message []byte) error {
	err := ss.GetClient().Mail(from)
	if err != nil {
		return err
	}
	if len(toList) == 0 {
		return fmt.Errorf("Email toList cant be null")
	}
	for _, to := range toList {
		if err = ss.GetClient().Rcpt(to); err != nil {
			return err
		}
	}
	if len(ccList) > 0 {
		for _, cc := range ccList {
			if err = ss.GetClient().Rcpt(cc); err != nil {
				return err
			}
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

	// message += crlfString + ss.GetConfig().Get(contentJSON)
	message += fmt.Sprintf("\r\n%s", ss.GetConfig().Get(contentJSON))

	return []byte(message)
}

// buildHeader returns an SMTPSender header
func (ss *SMTPSender) buildHeader() map[string]string {
	header := make(map[string]string)

	header[headerFrom] = fmt.Sprintf("%s<%s>", defaultAlertSMTPFromName, ss.GetConfig().Get(smtpFromAddrJson))
	header[headerTo] = ss.GetConfig().Get(toAddrsJSON)
	header[headerCc] = ss.GetConfig().Get(ccAddrsJSON)
	header[headerSubject] = ss.GetConfig().Get(subjectJSON)

	switch viper.GetString(config.AlertSMTPFormatKey) {
	case config.AlertSMTPTextFormat:
		header[headerContentType] = defaultAlertSMTPContentText
	case config.AlertSMTPHTMLFormat:
		header[headerContentType] = defaultAlertSMTPContentHTML
	}

	return header
}
