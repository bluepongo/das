package alert

import (
	"crypto/tls"
	"fmt"
	"strings"

	"net"
	"net/smtp"

	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/dependency/alert"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
	"github.com/spf13/viper"
)

const (
	defaultAlertSMTPFromName     = "DAS"
	defaultAlertSMTPContentHTML  = "text/html; charset=UTF-8"
	defaultAlertSMTPContentPLAIN = "text/plain; charset=UTF-8"

	headerFromStruct        = "From"
	headerToStruct          = "To"
	headerCcStruct          = "Cc"
	headerSubjectStruct     = "Subject"
	headerContentTypeStruct = "Content-Type"
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
func NewSMTPSenderWithDefault(cfg alert.Config) alert.Sender {
	url := viper.GetString(config.AlertSMTPURLKey)
	client := Dial(url)

	return newSMTPSender(client, cfg, url)
}

// newSMTPSender returns a new *SMTPSender
func newSMTPSender(client *smtp.Client, cfg alert.Config, url string) *SMTPSender {
	return &SMTPSender{
		client: client,
		config: cfg,
		url:    url,
	}
}

// Dial return a smtp client
func Dial(url string) *smtp.Client {
	conn, err := tls.Dial("tcp", url, nil)
	if err != nil {
		log.Errorf("Dialing Error:%s", err.Error())
		return nil
	}

	host, _, _ := net.SplitHostPort(url)
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Errorf("Dialing Error:%s", err.Error())
		return nil
	}
	return client
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

// setHeader set a SMTPSender header
func (ss *SMTPSender) setHeader() map[string]string {
	header := make(map[string]string)
	header[headerFromStruct] = defaultAlertSMTPFromName +
		"<" + ss.GetConfig().Get(smtpUserJSON) + ">"
	header[headerToStruct] = ss.GetConfig().Get(toAddrsJSON)
	header[headerCcStruct] = ss.GetConfig().Get(ccAddrsJSON)
	header[headerSubjectStruct] = ss.GetConfig().Get(subjectJSON)
	if viper.GetBool(config.AlertSMTPHTMLEnabledKey) {
		header[headerContentTypeStruct] = defaultAlertSMTPContentHTML
	}
	if !viper.GetBool(config.AlertSMTPHTMLEnabledKey) {
		header[headerContentTypeStruct] = defaultAlertSMTPContentPLAIN
	}
	return header
}

// makeAuth returns a SMTP Auth
func (ss *SMTPSender) makeAuth() smtp.Auth {
	url := ss.GetURL()
	host, _, _ := net.SplitHostPort(url)
	return smtp.PlainAuth(constant.EmptyString,
		ss.GetConfig().Get(smtpUserJSON),
		ss.GetConfig().Get(smtpPassJSON),
		host)
}

// spiltAddr spilts multi Address
func spiltAddr(addrs string) (addrList []string) {
	return strings.Split(addrs, ";")
}

// makeMessage return msg
func (ss *SMTPSender) makeMessage(header map[string]string) (msg []byte) {
	message := constant.EmptyString
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + ss.GetConfig().Get(contentJSON)
	return []byte(message)
}

// Send sends the email via the api calling
func (ss *SMTPSender) Send() error {
	// init header
	header := ss.setHeader()

	auth := ss.makeAuth()
	return sendMailUsingTLS(ss.GetURL(),
		auth, ss.GetConfig().Get(smtpUserJSON),
		spiltAddr(ss.GetConfig().Get(toAddrsJSON)),
		spiltAddr(ss.GetConfig().Get(ccAddrsJSON)), ss.makeMessage(header))
}

// sendMailUsingTLS sends mail using TLS
func sendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, cc []string, msg []byte) (err error) {
	//create smtp client
	c := Dial(addr)

	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Errorf("Error during AUTH %s", err.Error())
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	for _, addr := range cc {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
