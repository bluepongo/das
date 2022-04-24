package query

import (
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/internal/dependency/query"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
)

const QueriesStruct = "Queries"

var _ query.Service = (*Service)(nil)

type Service struct {
	config  query.Config
	dasRepo query.DASRepo
	Queries []query.Query `json:"queries"`
}

// NewService returns a new query.Service
func NewService(config query.Config, dasRepo query.DASRepo) query.Service {
	return newService(config, dasRepo)
}

// NewServiceWithDefault returns a new query.Service with default repository
func NewServiceWithDefault(config query.Config) query.Service {
	return newService(config, NewDASRepoWithGlobal())
}

// newService returns a new *Service
func newService(config query.Config, dasRepo query.DASRepo) *Service {
	return &Service{
		config:  config,
		dasRepo: dasRepo,
	}
}

// GetConfig returns the config of query
func (s *Service) GetConfig() query.Config {
	return s.config
}

// GetQueries returns the query slice
func (s *Service) GetQueries() []query.Query {
	return s.Queries
}

// GetByMySQLClusterID gets the query slice by the mysql cluster identity
func (s *Service) GetByMySQLClusterID(mysqlClusterID int) error {
	var err error

	querier := NewQuerierWithGlobal(s.GetConfig())
	s.Queries, err = querier.GetByMySQLClusterID(mysqlClusterID)
	if err != nil {
		return err
	}

	return s.Save(mysqlClusterID, constant.DefaultRandomInt, constant.DefaultRandomInt, constant.DefaultRandomString)
}

// GetByMySQLServerID gets the query slice by the mysql server identity
func (s *Service) GetByMySQLServerID(mysqlServerID int) error {
	var err error

	querier := NewQuerierWithGlobal(s.GetConfig())
	s.Queries, err = querier.GetByMySQLServerID(mysqlServerID)
	if err != nil {
		return err
	}

	return s.Save(constant.DefaultRandomInt, mysqlServerID, constant.DefaultRandomInt, constant.DefaultRandomString)
}

// GetByHostInfo gets the query slice by the mysql server host ip and port number
func (s *Service) GetByHostInfo(hostIP string, portNum int) error {
	mysqlServerService := metadata.NewMySQLServerServiceWithDefault()
	err := mysqlServerService.GetByHostInfo(hostIP, portNum)
	if err != nil {
		return err
	}
	mysqlServerID := mysqlServerService.GetMySQLServers()[constant.ZeroInt].Identity()
	querier := NewQuerierWithGlobal(s.GetConfig())
	s.Queries, err = querier.GetByMySQLServerID(mysqlServerID)
	if err != nil {
		return err
	}

	return s.Save(constant.DefaultRandomInt, mysqlServerID, constant.DefaultRandomInt, constant.DefaultRandomString)
}

// GetByDBID gets the query slice by the db identity
func (s *Service) GetByDBID(dbID int) error {
	var err error

	querier := NewQuerierWithGlobal(s.GetConfig())
	s.Queries, err = querier.GetByDBID(dbID)
	if err != nil {
		return err
	}

	return s.Save(constant.DefaultRandomInt, constant.DefaultRandomInt, dbID, constant.DefaultRandomString)
}

// GetBySQLID gets the query by the mysql server identity and the sql identity
func (s *Service) GetBySQLID(mysqlServerID int, sqlID string) error {
	var err error

	querier := NewQuerierWithGlobal(s.GetConfig())
	s.Queries, err = querier.GetBySQLID(mysqlServerID, sqlID)
	if err != nil {
		return err
	}

	return s.Save(constant.DefaultRandomInt, mysqlServerID, constant.DefaultRandomInt, sqlID)
}

// Save the query info into DAS repo
func (s *Service) Save(mysqlClusterID, mysqlServerID, dbID int, sqlID string) error {
	return s.dasRepo.Save(
		mysqlClusterID,
		mysqlServerID,
		dbID,
		sqlID,
		s.GetConfig().GetStartTime(),
		s.GetConfig().GetEndTime(),
		s.GetConfig().GetLimit(),
		s.GetConfig().GetOffset(),
	)
}

// Marshal marshals Service.Queries to json bytes
func (s *Service) Marshal() ([]byte, error) {
	return common.MarshalStructWithFields(s, QueriesStruct)
}
