package sqladvisor

import (
	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/dependency/sqladvisor"
	"github.com/romberli/go-multierror"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
	"github.com/spf13/viper"
)

var _ sqladvisor.Service = (*Service)(nil)

type Service struct {
	sqladvisor.Repository
	Advisor sqladvisor.Advisor
	Advice  string `json:"advice"`
	Message string `json:"message"`
}

// NewService returns a new *Service
func NewService(soarBin, configFile string) sqladvisor.Service {
	return newService(soarBin, configFile)
}

// NewServiceWithDefault returns a new *Service with default value
func NewServiceWithDefault() sqladvisor.Service {
	soarBin := viper.GetString(config.SQLAdvisorSoarBinKey)
	configFile := viper.GetString(config.SQLAdvisorSoarConfigKey)

	return newService(soarBin, configFile)
}

// newService returns a new *Service
func newService(soarBin, configFile string) *Service {
	return &Service{
		Repository: NewRepositoryWithGlobal(),
		Advisor:    NewDefaultAdvisor(soarBin, configFile),
	}
}

// GetFingerprint returns the fingerprint of the sql text
func (s *Service) GetFingerprint(sqlText string) string {
	return s.Advisor.GetFingerprint(sqlText)
}

// GetSQLID returns the identity of the sql text
func (s *Service) GetSQLID(sqlText string) string {
	return s.Advisor.GetSQLID(sqlText)
}

// Advise parses the sql text and returns the tuning advice,
// note that only the first sql statement in the sql text will be advised
func (s *Service) Advise(dbID int, sqlText string) (string, error) {
	merr := &multierror.Error{}

	sqlList, err := s.Advisor.GetParser().Split(sqlText)
	if err != nil {
		return constant.EmptyString, err
	}

	advice, msg, err := s.Advisor.Advise(dbID, sqlList[constant.ZeroInt])
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	if msg != constant.EmptyString {
		log.Infof("advisor message: %s", msg)
	}

	err = s.Repository.Save(dbID, sqlText, advice, msg)
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	return advice, merr.ErrorOrNil()
}
