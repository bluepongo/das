package query

import (
	"github.com/romberli/das/internal/dependency/query"
	"github.com/romberli/go-util/constant"
)

var _ query.Service = (*Service)(nil)

type Service struct {
	config  *Config
	dasRepo *DASRepo
	queries []query.Query
}

func NewService(config *Config, dasRepo *DASRepo) *Service {
	return newService(config, dasRepo)
}

func NewServiceWithDefault(config *Config) *Service {
	return newService(config, NewDASRepoWithGlobal())
}

func newService(config *Config, dasRepo *DASRepo) *Service {
	return &Service{
		config:  config,
		dasRepo: dasRepo,
	}
}

func (s *Service) GetConfig() *Config {
	return s.config
}

func (s *Service) GetQueries() []query.Query {
	return s.queries
}

func (s *Service) GetByMySQLClusterID(mysqlClusterID int) error {
	return nil
}

func (s *Service) GetByMySQLServerID(mysqlServerID int) error {
	var err error

	querier := NewQuerierWithGlobal(s.GetConfig())
	s.queries, err = querier.GetByMySQLServerID(mysqlServerID)
	if err != nil {
		return err
	}

	return s.Save(constant.DefaultRandomInt, mysqlServerID, constant.DefaultRandomInt, constant.DefaultRandomString)
}

func (s *Service) GetByDBID(dbID int) error {
	return nil
}

func (s *Service) GetBySQLID(mysqlServerID int, sqlID string) error {
	return nil
}

func (s *Service) Save(mysqlClusterID, mysqlServerID, dbID int, sqlID string) error {
	return s.dasRepo.Save(mysqlClusterID, mysqlServerID, dbID, sqlID,
		s.GetConfig().GetStartTime(), s.GetConfig().GetEndTime(), s.GetConfig().GetLimit(), s.GetConfig().GetOffset())
}

func (s *Service) Marshal() ([]byte, error) {
	return nil, nil
}

func (s *Service) MarshalWithFields(fields ...string) ([]byte, error) {
	return nil, nil
}
