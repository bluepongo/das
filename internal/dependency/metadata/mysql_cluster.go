package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

// MySQLCluster is the entity interface
type MySQLCluster interface {
	// Identity returns the identity
	Identity() int
	// GetClusterName returns the env name
	GetClusterName() string
	// GetMiddlewareClusterID returns the middleware cluster id
	GetMiddlewareClusterID() int
	// GetMonitorSystemID returns the monitor system id
	GetMonitorSystemID() int
	// GetOwnerID returns the owner id
	GetOwnerID() int
	// GetEnvID returns the env id
	GetEnvID() int
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// GetMySQLServers gets the mysql servers of this cluster
	GetMySQLServers() ([]MySQLServer, error)
	// GetMasterServers gets the master servers of this cluster
	GetMasterServers() ([]MySQLServer, error)
	// GetDBs gets the databases of this cluster
	GetDBs() ([]DB, error)
	// GetAppOwners gets the application owners of this cluster
	GetAppOwners() ([]User, error)
	// GetDBOwners gets the db owners of this cluster
	GetDBOwners() ([]User, error)
	// GetAllOwners gets both application and db owners of this cluster
	GetAllOwners() ([]User, error)
	// Set sets MySQLCluster with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete sets DelFlag to 1
	Delete()
	// MarshalJSON marshals MySQLCluster to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified field of the MySQLCluster to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

// MySQLClusterRepo is the repository interface
type MySQLClusterRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// GetAll gets all mysql clusters from the middleware
	GetAll() ([]MySQLCluster, error)
	// GetByEnv gets mysql clusters of given env id from the middleware
	GetByEnv(envID int) ([]MySQLCluster, error)
	// GetByID gets a mysql cluster by the identity from the middleware
	GetByID(id int) (MySQLCluster, error)
	// GetByName gets a mysql cluster of given cluster name from the middle ware
	GetByName(clusterName string) (MySQLCluster, error)
	// GetID gets the identity with given cluster name from the middleware
	GetID(clusterName string) (int, error)
	// GetDBsByID gets the databases of the given id from the middleware
	GetDBsByID(id int) ([]DB, error)
	// GetAppOwnersByID gets the application owners of the given id from the middleware
	GetAppOwnersByID(id int) ([]User, error)
	// GetDBOwnersByID gets the db owners of the given id from the middleware
	GetDBOwnersByID(id int) ([]User, error)
	// GetAllOwnersByID gets both application and db owners of the given id from the middleware
	GetAllOwnersByID(id int) ([]User, error)
	// Create creates a mysql cluster in the middleware
	Create(mc MySQLCluster) (MySQLCluster, error)
	// Update updates the mysql cluster in the middleware
	Update(mc MySQLCluster) error
	// Delete deletes the mysql cluster in the middleware
	Delete(id int) error
}

// MySQLClusterService is the service interface
type MySQLClusterService interface {
	// GetMySQLClusters returns the mysql clusters of the service
	GetMySQLClusters() []MySQLCluster
	// GetMySQLServers returns the mysql servers of the service
	GetMySQLServers() []MySQLServer
	// GetDBs returns the dbs of the service
	GetDBs() []DB
	// GetOwners returns the owners of the service
	GetOwners() []User
	// GetAll gets all mysql clusters from the middleware
	GetAll() error
	// GetByEnv gets mysql clusters of given env id
	GetByEnv(envID int) error
	// GetByID gets a mysql cluster of the given id from the middleware
	GetByID(id int) error
	// GetByName gets a mysql cluster of given cluster name
	GetByName(clusterName string) error
	// GetMySQLServersByID gets the mysql servers of given id
	GetMySQLServersByID(id int) error
	// GetMasterServersByID gets the master servers of the given id
	GetMasterServersByID(id int) error
	// GetDBsByID gets the databases of the given id
	GetDBsByID(id int) error
	// GetAppOwnersByID gets the application owners of the given id
	GetAppOwnersByID(id int) error
	// GetDBOwnersByID gets the db owners of the given id
	GetDBOwnersByID(id int) error
	// GetAllOwnersByID gets both application and db owners of the given id
	GetAllOwnersByID(id int) error
	// Create creates a mysql cluster in the middleware
	Create(fields map[string]interface{}) error
	// Update gets a mysql cluster of the given id from the middleware,
	// and then updates its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the middleware
	Update(id int, fields map[string]interface{}) error
	// Delete deletes the mysql cluster of given id in the middleware
	Delete(id int) error
	// Marshal marshals MySQLClusterService.MySQLClusters to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the MySQLClusterService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
