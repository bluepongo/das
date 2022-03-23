package query

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type Query interface {
	// GetSQLID returns the sql identity
	GetSQLID() string
	// GetFingerprint returns the fingerprint
	GetFingerprint() string
	// GetExample returns the example
	GetExample() string
	// GetDBName returns the db name
	GetDBName() string
	// GetExecCount returns the execution count
	GetExecCount() int
	// GetTotalExecTime returns the total execution time
	GetTotalExecTime() float64
	// GetAvgExecTime returns the average execution time
	GetAvgExecTime() float64
	// GetRowsExaminedMax returns the maximum row examined
	GetRowsExaminedMax() int
	// SetDBName sets db name to the query
	SetDBName(dbName string)
}

type DASRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// Save saves sql information into the middleware
	Save(mysqlClusterID, mysqlServerID, dbID int, sqlID string, startTime, endTime time.Time, limit, offset int) error
}

type MonitorRepo interface {
	// Close closes the monitor repository
	Close() error
	// GetByServiceNames gets the query slice by the service names of the mysql servers
	GetByServiceNames(serviceNames []string) ([]Query, error)
	// GetByDBName gets the query slice by the service name and db name of the mysql server
	GetByDBName(serviceName, dbName string) ([]Query, error)
	// GetBySQLID gets the query by the service name of the mysql server and sql identity
	GetBySQLID(serviceName, sqlID string) (Query, error)
}

type Service interface {
	// GetQueries returns the query slice
	GetQueries() []Query
	// GetByMySQLClusterID gets the query slice by the mysql cluster identity
	GetByMySQLClusterID(mysqlClusterID int) error
	// GetByMySQLServerID gets the query slice by the mysql server identity
	GetByMySQLServerID(mysqlServerID int) error
	// GetByHostInfo gets the query slice by the mysql server host ip and port number
	GetByHostInfo(hostIP string, portNum int) error
	// GetByDBName gets the query slice by the db identity
	GetByDBID(mysqlServerID, dbID int) error
	// GetBySQLID gets the query by the mysql server identity and the sql identity
	GetBySQLID(mysqlServerID int, sqlID string) error
	// Marshal marshals Service.Queries to json bytes
	Marshal() ([]byte, error)
}
