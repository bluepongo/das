package alert

import (
	"encoding/json"
	"errors"

	"github.com/hashicorp/go-multierror"
	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/dependency/alert"
	"github.com/romberli/go-util/constant"
	"github.com/spf13/viper"
)

const (
	toAddrsJSON = "to_addrs"
	contentJSON = "content"
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
	return nil
}

// sendViaHTTP sends email via http api calling
func (s *Service) sendViaHTTP(toAddrs, ccAddrs, content string) error {
	merr := &multierror.Error{}
	// setup config
	s.setupConfig(toAddrs, ccAddrs, content)
	sender := NewHTTTPSenderWithDefault(s.GetConfig())
	// send email
	err := sender.Send()
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// save result
	err = s.saveHTTP(toAddrs, ccAddrs, content, err.Error())
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	return merr.ErrorOrNil()
}

// setupConfig setups the config
func (s *Service) setupConfig(toAddrs, ccAddrs, content string) {
	toAddrs += constant.CommaString + ccAddrs
	s.GetConfig().Set(toAddrsJSON, toAddrs)
	s.GetConfig().Set(contentJSON, content)
}

// saveHTTP saves the sending results which was done via calling http api to the middleware
func (s *Service) saveHTTP(toAddrs, ccAddrs, content, message string) error {
	cfg, err := json.Marshal(s.GetConfig())
	if err != nil {
		return err
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

// saveSMTP saves the sending results which was done via smtp server to the middleware
func saveSMTP(toAddrs, ccAddrs, content, message string) error {
	return nil
}