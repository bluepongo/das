package alert

import (
	"encoding/json"
	"strings"

	"github.com/pingcap/errors"

	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/dependency/alert"
	"github.com/romberli/go-multierror"
	"github.com/romberli/go-util/constant"
	"github.com/spf13/viper"
)

const (
	toAddrsJSON          = "to_addrs"
	ccAddrsJSON          = "cc_addrs"
	contentJSON          = "content"
	subjectJSON          = "subject"
	smtpUserJSON         = "user"
	smtpPassJSON         = "pass"
	smtpFromAddrJson     = "from_address"
	defaultPassEncrypted = "****"
)

var _ alert.Service = (*Service)(nil)

type Service struct {
	alert.Repository
	config alert.Config
}

// NewService returns a new alert.Service
func NewService(repository alert.Repository, config alert.Config) alert.Service {
	return newService(repository, config)
}

// NewServiceWithDefault returns a new alert.Service with default repository
func NewServiceWithDefault(config alert.Config) alert.Service {
	return newService(NewRepositoryWithGlobal(), config)
}

// newService returns a new alert.Service
func newService(repository alert.Repository, config alert.Config) *Service {
	return &Service{
		Repository: repository,
		config:     config,
	}
}

// GetRepository returns the repository of the service
func (s *Service) GetRepository() alert.Repository {
	return s.Repository
}

// GetConfig returns the config of the service
func (s *Service) GetConfig() alert.Config {
	return s.config
}

// SendEmail sends the email
func (s *Service) SendEmail(toAddrs, ccAddrs, subject, content string) error {
	smtpEnabled := viper.GetBool(config.AlertSMTPEnabledKey)
	if smtpEnabled {
		return s.sendViaSMTP(toAddrs, ccAddrs, subject, content)
	}

	httpEnabled := viper.GetBool(config.AlertHTTPEnabledKey)
	if httpEnabled {
		return s.sendViaHTTP(toAddrs, ccAddrs, content)
	}

	return errors.New("none of smtp or http is enabled, can not send email")
}

// sendViaSMTP sends email via smtp server
func (s *Service) sendViaSMTP(toAddrs, ccAddrs, subject, content string) error {
	merr := &multierror.Error{}
	var message string
	// setup config
	s.setupConfig(toAddrs, ccAddrs, subject, content)
	sender, err := NewSMTPSenderWithDefault(s.GetConfig())
	if err != nil {
		return err
	}
	// send email
	err = sender.Send()
	if err != nil {
		message = err.Error()
		merr = multierror.Append(merr, err)
	}
	// save result
	err = s.Save(toAddrs, ccAddrs, subject, content, message)
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	return merr.ErrorOrNil()
}

// sendViaHTTP sends email via http api calling
func (s *Service) sendViaHTTP(toAddrs, ccAddrs, content string) error {
	merr := &multierror.Error{}
	var message string
	// setup config
	s.setupConfig(toAddrs, ccAddrs, constant.EmptyString, content)
	sender := NewHTTPSenderWithDefault(s.GetConfig())
	// send email
	err := sender.Send()
	if err != nil {
		message = err.Error()
		merr = multierror.Append(merr, err)
	}
	// save result
	err = s.Save(toAddrs, ccAddrs, constant.EmptyString, content, message)
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	return merr.ErrorOrNil()
}

// setupConfig setups config of HTTP or SMTP
func (s *Service) setupConfig(toAddrs, ccAddrs, subject, content string) {
	smtpEnabled := viper.GetBool(config.AlertSMTPEnabledKey)
	httpEnabled := viper.GetBool(config.AlertHTTPEnabledKey)
	if smtpEnabled {
		s.setupSMTPConfig(toAddrs, ccAddrs, subject, content)
	}
	if !smtpEnabled && httpEnabled {
		s.setupHTTPConfig(toAddrs, ccAddrs, content)
	}
}

// setupSMTPConfig setups the SMTP config
func (s *Service) setupSMTPConfig(toAddrs, ccAddrs, subject, content string) {
	s.GetConfig().Set(smtpUserJSON, viper.GetString(config.AlertSMTPUserKey))
	s.GetConfig().Set(smtpPassJSON, viper.GetString(config.AlertSMTPPassKey))
	s.GetConfig().Set(smtpFromAddrJson, viper.GetString(config.AlertSMTPFromKey))
	s.GetConfig().Set(toAddrsJSON, toAddrs)
	s.GetConfig().Set(ccAddrsJSON, ccAddrs)
	s.GetConfig().Set(subjectJSON, subject)
	s.GetConfig().Set(contentJSON, content)
}

// setupHTTPConfig setups the HTTP config
func (s *Service) setupHTTPConfig(toAddrs, ccAddrs, content string) {
	s.GetConfig().Set(toAddrsJSON, strings.Trim(toAddrs+constant.CommaString+ccAddrs, constant.CommaString))
	s.GetConfig().Set(contentJSON, content)
}

// Save saves the email into the middleware
func (s *Service) Save(toAddrs, ccAddrs, subject, content, message string) error {
	smtpEnabled := viper.GetBool(config.AlertSMTPEnabledKey)
	httpEnabled := viper.GetBool(config.AlertHTTPEnabledKey)
	if smtpEnabled {
		return s.saveSMTP(toAddrs, ccAddrs, subject, content, message)
	}
	if httpEnabled {
		return s.saveHTTP(toAddrs, ccAddrs, content, message)
	}

	return errors.New("none of smtp or http is enabled, can not save the email")
}

// saveSMTP saves the sending results which was done via smtp server to the middleware
func (s *Service) saveSMTP(toAddrs, ccAddrs, subject, content, message string) error {
	// encrypt password with simple string
	s.GetConfig().Set(smtpPassJSON, defaultPassEncrypted)
	cfg, err := json.Marshal(s.GetConfig())
	if err != nil {
		return errors.Trace(err)
	}

	return s.GetRepository().Save(
		viper.GetString(config.AlertSMTPURLKey),
		toAddrs,
		ccAddrs,
		subject,
		content,
		string(cfg),
		message,
	)
}

// saveHTTP saves the sending results which was done via calling http api to the middleware
func (s *Service) saveHTTP(toAddrs, ccAddrs, content, message string) error {
	cfg, err := json.Marshal(s.GetConfig())
	if err != nil {
		return errors.Trace(err)
	}

	return s.GetRepository().Save(
		viper.GetString(config.AlertHTTPURLKey),
		toAddrs,
		ccAddrs,
		constant.DefaultRandomString,
		content,
		string(cfg),
		message,
	)
}
