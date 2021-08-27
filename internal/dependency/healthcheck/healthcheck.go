package healthcheck

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type DASRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
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
	// GetDBConfig gets the database configuration
	GetDBConfig(configItems []string) ([]Variable, error)
	// GetMySQLDirs gets the mysql data directory and binlog directory
	GetMySQLDirs() ([]string, error)
	// GetTables gets the tables
	GetLargeTables() ([]Table, error)
}

type PrometheusRepo interface {
	// GetMountPoint gets the mount points from the prometheus
	GetFileSystems(serviceName string) ([]FileSystem, error)
	// CheckCPUUsage gets the cpu usage
	GetCPUUsage(serviceName string) ([]PrometheusData, error)
	// CheckIOUtil gets the io util
	GetIOUtil(serviceName string, devices []string) ([]PrometheusData, error)
	// GetDiskCapacityUsage gets the disk capacity usage
	GetDiskCapacityUsage(serviceName string, mountPoints []string) ([]PrometheusData, error)
	// GetConnectionUsage gets the connection usage
	GetConnectionUsage(serviceName string) ([]PrometheusData, error)
	// GetActiveSessionNum gets the active session number
	GetActiveSessionNum(serviceName string) ([]PrometheusData, error)
	// GetCacheMissRatio gets the cache miss ratio
	GetCacheMissRatio(serviceName string) ([]PrometheusData, error)
}

type QueryRepo interface {
	// Close closes the mysql or clickhouse connection
	Close() error
	// GetSlowQuery gets the slow query
	GetSlowQuery(serviceName string) (middleware.Result, error)
}

type Service interface {
	// GetResult returns the result
	GetResult() Result
	// GetResultByOperationID gets the result by operation id from the middleware
	GetResultByOperationID(id int) error
	// Check checks the server health status
	Check(mysqlServerID int, startTime, endTime time.Time, step time.Duration) error
	// Check checks the server health status
	CheckByHostInfo(hostIP string, portNum int, startTime, endTime time.Time, step time.Duration) error
	// ReviewAccuracy reviews the accuracy of the check
	ReviewAccuracy(id, review int) error
	// MarshalJSON marshals Service to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSON marshals only specified field of the Service to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type Engine interface {
	// Run checks the server health status
	Run()
}
