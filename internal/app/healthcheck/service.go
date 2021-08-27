package healthcheck

import (
	"fmt"
	"time"

	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/internal/dependency/healthcheck"
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
	resultStruct                   = "Result"
	defaultStep                    = time.Minute
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

// GetResult returns the healthcheck result
func (s *Service) GetResult() healthcheck.Result {
	return s.Result
}

// GetResultByOperationID gets the result of given operation id
func (s *Service) GetResultByOperationID(id int) error {
	var err error

	s.Result, err = s.DASRepo.GetResultByOperationID(id)
	if err != nil {
		return err
	}

	return err
}

// Check performs healthcheck on the mysql server with given mysql server id,
// initiating is synchronous, actual running is asynchronous
func (s *Service) Check(mysqlServerID int, startTime, endTime time.Time, step time.Duration) error {
	return s.check(mysqlServerID, startTime, endTime, step)
}

// CheckByHostInfo performs healthcheck on the mysql server with given mysql server id,
// initiating is synchronous, actual running is asynchronous
func (s *Service) CheckByHostInfo(hostIP string, portNum int, startTime, endTime time.Time, step time.Duration) error {
	// init mysql server service
	mss := metadata.NewMySQLServerServiceWithDefault()
	// get entities
	err := mss.GetByHostInfo(hostIP, portNum)
	if err != nil {
		return err
	}
	mysqlServerID := mss.MySQLServers[0].Identity()
	return s.check(mysqlServerID, startTime, endTime, step)
}

// check performs healthcheck on the mysql server with given mysql server id,
// initiating is synchronous, actual running is asynchronous
func (s *Service) check(mysqlServerID int, startTime, endTime time.Time, step time.Duration) error {
	// init
	err := s.init(mysqlServerID, startTime, endTime, step)
	if err != nil {
		updateErr := s.DASRepo.UpdateOperationStatus(s.OperationInfo.operationID, defaultFailedStatus, err.Error())
		if updateErr != nil {
			log.Error(message.NewMessage(msghc.ErrHealthcheckUpdateOperationStatus, updateErr.Error()).Error())
		}

		return err
	}
	// run asynchronously
	go s.Engine.Run()

	return nil
}

// init initiates healthcheck operation and engine
func (s *Service) init(mysqlServerID int, startTime, endTime time.Time, step time.Duration) error {
	// check if operation with the same mysql server id is still running
	isRunning, err := s.DASRepo.IsRunning(mysqlServerID)
	if err != nil {
		return err
	}
	if isRunning {
		return fmt.Errorf("healthcheck of mysql server is still running. mysql server id: %d", mysqlServerID)
	}
	// insert operation message
	id, err := s.DASRepo.InitOperation(mysqlServerID, startTime, endTime, step)
	if err != nil {
		return err
	}
	// get operation info
	// init application mysql connection
	mysqlServerService := metadata.NewMySQLServerServiceWithDefault()
	err = mysqlServerService.GetByID(mysqlServerID)
	if err != nil {
		return err
	}
	mysqlServer := mysqlServerService.GetMySQLServers()[constant.ZeroInt]
	mysqlServerAddr := fmt.Sprintf("%s:%d", mysqlServer.GetHostIP(), mysqlServer.GetPortNum())
	applicationMySQLConn, err := mysql.NewConn(mysqlServerAddr, constant.EmptyString, s.getApplicationMySQLUser(), s.getApplicationMySQLPass())
	if err != nil {
		return err
	}
	// get monitor system info
	monitorSystem, err := mysqlServer.GetMonitorSystem()
	if err != nil {
		return err
	}

	var (
		monitorPrometheusConn *prometheus.Conn
		monitorClickhouseConn *clickhouse.Conn
		monitorMySQLConn      *mysql.Conn
	)

	monitorSystemType := monitorSystem.GetSystemType()
	switch monitorSystemType {
	case 1:
		// pmm 1.x
		// init prometheus connection
		prometheusAddr := fmt.Sprintf("%s:%d", monitorSystem.GetHostIP(), monitorSystem.GetPortNum())
		prometheusConfig := prometheus.NewConfig(prometheusAddr, prometheus.DefaultRoundTripper)
		monitorPrometheusConn, err = prometheus.NewConnWithConfig(prometheusConfig)
		if err != nil {
			return err
		}
		// init mysql connection
		mysqlAddr := fmt.Sprintf("%s:%d", monitorSystem.GetHostIP(), monitorSystem.GetPortNumSlow())
		monitorMySQLConn, err = mysql.NewConn(mysqlAddr, defaultMonitorMySQLDBName, s.getMonitorMySQLUser(), s.getMonitorMySQLPass())
		if err != nil {
			return err
		}
	case 2:
		// pmm 2.x
		// init prometheus connection
		prometheusAddr := fmt.Sprintf("%s:%d%s", monitorSystem.GetHostIP(), monitorSystem.GetPortNum(), monitorSystem.GetBaseURL())
		prometheusConfig := prometheus.NewConfigWithBasicAuth(prometheusAddr, s.getMonitorPrometheusUser(), s.getMonitorPrometheusPass())
		monitorPrometheusConn, err = prometheus.NewConnWithConfig(prometheusConfig)
		if err != nil {
			return err
		}
		// init clickhouse connection
		clickhouseAddr := fmt.Sprintf("%s:%d", monitorSystem.GetHostIP(), monitorSystem.GetPortNumSlow())
		monitorClickhouseConn, err = clickhouse.NewConnWithDefault(clickhouseAddr, defaultMonitorClickhouseDBName, s.getMonitorClickhouseUser(), s.getMonitorClickhousePass())
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("healthcheck: monitor system type should be either 1 or 2, %d is not valid", monitorSystemType)
	}

	s.OperationInfo = NewOperationInfo(id, mysqlServer, monitorSystem, startTime, endTime, step)
	s.Engine = NewDefaultEngine(s.DASRepo, s.OperationInfo, applicationMySQLConn, monitorPrometheusConn, monitorClickhouseConn, monitorMySQLConn)

	return nil
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
	return s.DASRepo.UpdateAccuracyReviewByOperationID(id, review)
}

// MarshalJSON marshals Service to json bytes
func (s *Service) MarshalJSON() ([]byte, error) {
	return s.MarshalJSONWithFields(resultStruct)
}

// MarshalJSONWithFields marshals only specified fields of the Service to json bytes
func (s *Service) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(s.Result, fields...)
}
