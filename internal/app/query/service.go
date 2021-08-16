package query

import (
	"github.com/romberli/das/internal/dependency/query"
)

var _ query.Service = (*Service)(nil)

type Service struct {
	Config  *Config
	Queries []query.Query
}

func NewService(config *Config) *Service {
	return &Service{
		Config: config,
	}
}

func (s *Service) GetQueries() []query.Query {
	return s.Queries
}

func (s *Service) GetAll() error {
	return nil
}

func (s *Service) GetByMySQLServerID() error {
	return nil
}

func (s *Service) GetByDBID() error {
	return nil
}

func (s *Service) GetByID() error {
	return nil
}

func (s *Service) Marshal() ([]byte, error) {
	return nil, nil
}

func (s *Service) MarshalWithFields(fields ...string) ([]byte, error) {
	return nil, nil
}
