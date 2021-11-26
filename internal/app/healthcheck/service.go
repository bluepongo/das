package healthcheck

import (
	"errors"
	"fmt"
	"time"

	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/internal/dependency/healthcheck"
	depmeta "github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/pkg/message"
	msghc "github.com/romberli/das/pkg/message/healthcheck"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/clickhouse"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/go-util/middleware/prometheus"
	"github.com/romberli/log"
	"github.com/spf13/viper"
)

const (
	healthcheckResultStruct        = "Result"
	defaultMonitorClickhouseDBName = "pmm"
	defaultMonitorMySQLDBName      = "pmm"
	defaultSuccessStatus           = 2
	defaultFailedStatus            = 3
)

var _ healthcheck.Service = (*Service)(nil)

// Service of health check
type Service struct {
	healthcheck.DASRepo
	OperationInfo *OperationInfo
	Engine        healthcheck.Engine
	Result        healthcheck.Result `json:"result"`
}

// NewService returns a new *Service
func NewService(repo healthcheck.DASRepo) healthcheck.Service {
	return newService(repo)
}

// NewServiceWithDefault returns a new healthcheck.Service with default repository
func NewServiceWithDefault() healthcheck.Service {
	return newService(NewDASRepoWithGlobal())

}

// newService returns a new *Service
func newService(repo healthcheck.DASRepo) *Service {
	return &Service{
		DASRepo: repo,
		Result:  NewEmptyResult(),
	}
}

// GetDASRepo returns the das repository
func (s *Service) GetDASRepo() healthcheck.DASRepo {
	return s.DASRepo
}

// GetOperationInfo returns the operation information
func (s *Service) GetOperationInfo() healthcheck.OperationInfo {
	return s.OperationInfo
}

// GetEngine returns the healthcheck engine
func (s *Service) GetEngine() healthcheck.Engine {
	return s.Engine
}

// GetResult returns the healthcheck result
func (s *Service) GetResult() healthcheck.Result {
	return s.Result
}

// GetResultByOperationID gets the result of given operation id
func (s *Service) GetResultByOperationID(id int) error {
	var err error

	s.Result, err = s.GetDASRepo().GetResultByOperationID(id)
	if err != nil {
		return err
	}

	return err
}

// Check performs healthcheck on the mysql server with given mysql server id,
// initiating is synchronous, actual running is asynchronous
func (s *Service) Check(mysqlServerID int, startTime, endTime time.Time, step time.Duration) (int, error) {
	return s.check(mysqlServerID, startTime, endTime, step)
}

// CheckByHostInfo performs healthcheck on the mysql server with given mysql server id,
// initiating is synchronous, actual running is asynchronous
func (s *Service) CheckByHostInfo(hostIP string, portNum int, startTime, endTime time.Time, step time.Duration) (int, error) {
	// init mysql server service
	mss := metadata.NewMySQLServerServiceWithDefault()
	// get entities
	err := mss.GetByHostInfo(hostIP, portNum)
	if err != nil {
		return constant.ZeroInt, err
	}
	mysqlServerID := mss.GetMySQLServers()[constant.ZeroInt].Identity()

	return s.check(mysqlServerID, startTime, endTime, step)
}

// check performs healthcheck on the mysql server with given mysql server id,
// initiating is synchronous, actual running is asynchronous
func (s *Service) check(mysqlServerID int, startTime, endTime time.Time, step time.Duration) (int, error) {
	// init
	operationID, err := s.init(mysqlServerID, startTime, endTime, step)
	if err != nil {
		updateErr := s.GetDASRepo().UpdateOperationStatus(operationID, defaultFailedStatus, err.Error())
		if updateErr != nil {
			log.Error(message.NewMessage(msghc.ErrHealthcheckUpdateOperationStatus, updateErr.Error()).Error())
		}

		return operationID, err
	}
	// run asynchronously
	go s.GetEngine().Run()

	return operationID, nil
}

// init initiates healthcheck operation and engine
func (s *Service) init(mysqlServerID int, startTime, endTime time.Time, step time.Duration) (int, error) {
	// insert operation message
	operationID, err := s.GetDASRepo().InitOperation(mysqlServerID, startTime, endTime, step)
	if err != nil {
		return operationID, err
	}
	// check if operation with the same mysql server id is still running
	isRunning, err := s.GetDASRepo().IsRunning(mysqlServerID)
	if err != nil {
		return operationID, err
	}
	if isRunning {
		return operationID, fmt.Errorf("healthcheck of mysql server is still running. mysql server id: %d", mysqlServerID)
	}
	mysqlServerService := metadata.NewMySQLServerServiceWithDefault()
	err = mysqlServerService.GetByID(mysqlServerID)
	if err != nil {
		return operationID, err
	}
	// get mysql server
	mysqlServer := mysqlServerService.GetMySQLServers()[constant.ZeroInt]
	// get monitor system
	monitorSystem, err := mysqlServer.GetMonitorSystem()
	if err != nil {
		return operationID, err
	}
	mysqlCluster, err := mysqlServer.GetMySQLCluster()
	if err != nil {
		return operationID, err
	}
	// get dbs
	dbs, err := mysqlCluster.GetDBs()
	if err != nil {
		return operationID, err
	}
	// get apps
	var apps []depmeta.App
	for _, db := range dbs {
		applications, err := db.GetApps()
		if err != nil {
			return operationID, err
		}
		for _, application := range applications {
			exists, err := common.ElementInSlice(apps, application)
			if err != nil {
				return operationID, err
			}
			if !exists {
				apps = append(apps, applications...)
			}
		}
	}
	// init operation information
	s.OperationInfo = NewOperationInfo(operationID, apps, mysqlServer, monitorSystem, startTime, endTime, step)

	// init application mysql connection
	mysqlServerAddr := fmt.Sprintf("%s:%d", mysqlServer.GetHostIP(), mysqlServer.GetPortNum())
	applicationMySQLConn, err := mysql.NewConn(mysqlServerAddr, constant.EmptyString, s.getApplicationMySQLUser(), s.getApplicationMySQLPass())
	if err != nil {
		return operationID, errors.New(
			fmt.Sprintf("create application mysql connection failed. addr: %s, user: %s. error:\n%s",
				mysqlServerAddr, s.getApplicationMySQLUser(), err.Error()))
	}
	// init application mysql repository
	applicationMySQLRepo := NewApplicationMySQLRepo(s.GetOperationInfo(), applicationMySQLConn)

	var (
		prometheusConfig prometheus.Config
		queryRepo        healthcheck.QueryRepo
	)

	prometheusAddr := fmt.Sprintf("%s:%d%s", monitorSystem.GetHostIP(), monitorSystem.GetPortNum(), monitorSystem.GetBaseURL())
	slowQueryAddr := fmt.Sprintf("%s:%d", monitorSystem.GetHostIP(), monitorSystem.GetPortNumSlow())

	switch monitorSystem.GetSystemType() {
	case 1:
		// pmm 1.x
		// init prometheus config
		prometheusConfig = prometheus.NewConfig(prometheusAddr, prometheus.DefaultRoundTripper)
		// init mysql connection
		conn, err := mysql.NewConn(slowQueryAddr, defaultMonitorMySQLDBName, s.getMonitorMySQLUser(), s.getMonitorMySQLPass())
		if err != nil {
			return operationID, errors.New(
				fmt.Sprintf("create monitor mysql connection failed. addr: %s, user: %s. error:\n%s",
					slowQueryAddr, s.getMonitorMySQLUser(), err.Error()))
		}
		queryRepo = NewMySQLQueryRepo(s.GetOperationInfo(), conn)
	case 2:
		// pmm 2.x
		// init prometheus config
		prometheusConfig = prometheus.NewConfigWithBasicAuth(prometheusAddr, s.getMonitorPrometheusUser(), s.getMonitorPrometheusPass())
		// init clickhouse connection
		conn, err := clickhouse.NewConnWithDefault(slowQueryAddr, defaultMonitorClickhouseDBName, s.getMonitorClickhouseUser(), s.getMonitorClickhousePass())
		if err != nil {
			return operationID, errors.New(
				fmt.Sprintf("create monitor clickhouse connection failed. addr: %s, user: %s. error:\n%s",
					slowQueryAddr, s.getMonitorClickhouseUser(), err.Error()))
		}
		queryRepo = NewClickhouseQueryRepo(s.GetOperationInfo(), conn)
	default:
		return operationID, fmt.Errorf("healthcheck: monitor system type should be either 1 or 2, %d is not valid", monitorSystem.GetSystemType())
	}

	prometheusConn, err := prometheus.NewConnWithConfig(prometheusConfig)
	if err != nil {
		return operationID, errors.New(
			fmt.Sprintf("create prometheus connection failed. addr: %s, user: %s. error:\n%s",
				prometheusAddr, s.getMonitorPrometheusUser(), err.Error()))
	}
	prometheusRepo := NewPrometheusRepo(s.GetOperationInfo(), prometheusConn)
	s.Engine = NewDefaultEngine(s.GetOperationInfo(), s.GetDASRepo(), applicationMySQLRepo, prometheusRepo, queryRepo)

	return operationID, nil
}

// getApplicationMySQLUser returns application mysql username
func (s *Service) getApplicationMySQLUser() string {
	return viper.GetString(config.DBApplicationMySQLUserKey)
}

// getApplicationMySQLPass returns application mysql password
func (s *Service) getApplicationMySQLPass() string {
	return viper.GetString(config.DBApplicationMySQLPassKey)
}

// getMonitorPrometheusUser returns prometheus username of monitor system
func (s *Service) getMonitorPrometheusUser() string {
	return viper.GetString(config.DBMonitorPrometheusUserKey)
}

// getMonitorPrometheusPass returns prometheus password of monitor system
func (s *Service) getMonitorPrometheusPass() string {
	return viper.GetString(config.DBMonitorPrometheusPassKey)
}

// getMonitorClickhouseUser returns clickhouse username of monitor system
func (s *Service) getMonitorClickhouseUser() string {
	return viper.GetString(config.DBMonitorClickhouseUserKey)
}

// getMonitorClickhousePass returns clickhouse password of monitor system
func (s *Service) getMonitorClickhousePass() string {
	return viper.GetString(config.DBMonitorClickhousePassKey)
}

// getMonitorMySQLUser returns mysql username of monitor system
func (s *Service) getMonitorMySQLUser() string {
	return viper.GetString(config.DBMonitorMySQLUserKey)
}

// getMonitorMySQLPass returns mysql password of monitor system
func (s *Service) getMonitorMySQLPass() string {
	return viper.GetString(config.DBMonitorMySQLPassKey)
}

// ReviewAccuracy updates accuracy review with given operation id
func (s *Service) ReviewAccuracy(id, review int) error {
	return s.GetDASRepo().UpdateAccuracyReviewByOperationID(id, review)
}

// Marshal marshals Service to json bytes
func (s *Service) Marshal() ([]byte, error) {
	return s.MarshalWithFields(healthcheckResultStruct)
}

// MarshalWithFields marshals only specified fields of the Service to json bytes
func (s *Service) MarshalWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(s, fields...)
}
