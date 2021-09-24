package alert

import (
	"github.com/romberli/das/internal/dependency/alert"
)

var _ alert.Service = (*Service)(nil)

type Service struct {
	alert.Repository
}

func NewService(repository alert.Repository) *Service {
	return newService(repository)
}

func NewServiceWithDefault() *Service {
	return newService(NewRepositoryWithGlobal())
}

func newService(repository alert.Repository) *Service {
	return &Service{repository}
}

func (s *Service) SendEmail(url string, toAddr, ccAddr []string, content string) error {
	return nil
}

func (s *Service) save(url string, toAddr, ccAddr []string, content string) error {
	return nil
}
