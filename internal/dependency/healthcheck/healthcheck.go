package healthcheck

import (
	"time"

	depquery "github.com/romberli/das/internal/dependency/query"
	"github.com/romberli/go-util/middleware"
)

type DASRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// GetOperationHistories gets operation histories from the middleware
	GetHealthCheckHistories(mysqlServerIDList []int, limit int) ([]OperationHistory, error)
	// LoadEngineConfig loads engine config from the middleware
	LoadEngineConfig() (EngineConfig, error)
	// GetResultByOperationID returns the result
	GetResultByOperationID(operationID int) (Result, error)
	// IsRunning returns if the healthcheck of given mysql server is still running
	IsRunning(mysqlServerID int) (bool, error)
	// InitOperation initiates the operation
	InitOperation(mysqlServerID int, startTime, endTime time.Time, step time.Duration) (int, error)
	// UpdateOperationStatus updates operation status
	UpdateOperationStatus(operationID int, status int, message string) error
	// SaveResult saves result into the middleware
	SaveResult(result Result) error
	// UpdateAccuracyReviewByOperationID updates the accuracy review
	UpdateAccuracyReviewByOperationID(operationID int, review int) error
}

type ApplicationMySQLRepo interface {
	// Close closes the mysql connection
	Close() error
	// GetVariables gets the database variables by items
	GetVariables(items []string) ([]Variable, error)
	// GetMySQLDirs gets the mysql data directory and binlog directory
	GetMySQLDirs() ([]string, error)
	// GetLargeTables gets the tables
	GetLargeTables() ([]Table, error)
	// GetDBName gets the db name of given table names
	GetDBName(tableNames []string) (string, error)
}

type PrometheusRepo interface {
	// GetFileSystems gets the file systems from the prometheus
	GetFileSystems() ([]FileSystem, error)
	// GetAvgBackupFailedRatio gets the average backup failed ratio
	GetAvgBackupFailedRatio() ([]PrometheusData, error)
	// GetStatisticFailedRatio gets the statistic failed ratio
	GetStatisticFailedRatio() ([]PrometheusData, error)
	// GetCPUUsage gets the cpu usage
	GetCPUUsage() ([]PrometheusData, error)
	// GetIOUtil gets the io util
	GetIOUtil() ([]PrometheusData, error)
	// GetDiskCapacityUsage gets the disk capacity usage
	GetDiskCapacityUsage(mountPoints []string) ([]PrometheusData, error)
	// GetConnectionUsage gets the connection usage
	GetConnectionUsage() ([]PrometheusData, error)
	// GetAverageActiveSessionPercents gets the average active session percents
	GetAverageActiveSessionPercents() ([]PrometheusData, error)
	// GetCacheMissRatio gets the cache miss ratio
	GetCacheMissRatio() ([]PrometheusData, error)
}

type QueryRepo interface {
	// Close closes the mysql or clickhouse connection
	Close() error
	// GetSlowQuery gets the slow query
	GetSlowQuery() ([]depquery.Query, error)
}

type Service interface {
	// GetDASRepo returns the das repository
	GetDASRepo() DASRepo
	// GetOperationInfo returns the operation information
	GetOperationInfo() OperationInfo
	// GetEngine returns the healthcheck engine
	GetEngine() Engine
	// GetOperationHistories returns the operation histories
	GetOperationHistories() []OperationHistory
	// GetResult returns the result
	GetResult() Result
	// GetOperationHistoriesByLoginName returns the operation histories by login name
	GetOperationHistoriesByLoginName(loginName string) error
	// GetResultByOperationID gets the result by operation id from the middleware
	GetResultByOperationID(id int) error
	// Check checks the server health status
	Check(mysqlServerID int, startTime, endTime time.Time, step time.Duration, loginName string) (int, error)
	// Check checks the server health status
	CheckByHostInfo(hostIP string, portNum int, startTime, endTime time.Time, step time.Duration, loginName string) (int, error)
	// ReviewAccuracy reviews the accuracy of the check
	ReviewAccuracy(id, review int) error
	// Marshal marshals Service to json string
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified field of the Service to json string
	MarshalWithFields(fields ...string) ([]byte, error)
}

type Engine interface {
	// GetOperationInfo returns the operation information
	GetOperationInfo() OperationInfo
	// Run checks the server health status
	Run()
}
