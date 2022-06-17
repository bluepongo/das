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
	// GetResourceGroupByID get the resource group of the given id from the middleware
	GetResourceGroup() ([]ResourceGroup, error)
	// GetUsers gets the users that own the cluster
	GetUsers() ([]User, error)
	// AddUser add a map of the mysql cluster and user in the middleware
	AddUser(userID int) error
	// DeleteUser delete the map of the mysql cluster and user in the middleware
	DeleteUser(userID int) error
	// GetAppUsers gets the application users of this cluster
	GetAppUsers() ([]User, error)
	// GetDBUsers gets the db users of this cluster
	GetDBUsers() ([]User, error)
	// GetAllUsers gets mysql cluster, application and db users of this cluster
	GetAllUsers() ([]User, error)
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
	// GetResourceGroupByID get the resource group of the given id from the middleware
	GetResourceGroupByID(id int) ([]ResourceGroup, error)
	// GetUsersByID gets the users that the mysql cluster uses
	GetUsersByID(id int) ([]User, error)
	// AddUser add a new map of mysql cluster and user in the middleware
	AddUser(mysqlClusterID, userID int) error
	// DeleteUser delete the map of mysql cluster and user in the middleware
	DeleteUser(mysqlClusterID, userID int) error
	// GetAppUsersByID gets the application users of the given id from the middleware
	GetAppUsersByID(id int) ([]User, error)
	// GetDBUsersByID gets the db users of the given id from the middleware
	GetDBUsersByID(id int) ([]User, error)
	// GetAllUsersByID gets mysql cluster, application and db users of the given id from the middleware
	GetAllUsersByID(id int) ([]User, error)
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
	// GetResourceGroupByID get the resource group of the given id from the middleware
	GetResourceGroupByID(id int) ([]ResourceGroup, error)
	// GetUsersByID gets the user that own the mysql cluster
	GetUsersByID(id int) error
	// AddUser add a new map of mysql cluster and user in the middleware
	AddUser(mysqlClusterID, userID int) error
	// DeleteUser delete the map of mysql cluster and user in the middleware
	DeleteUser(mysqlClusterID, userID int) error
	// GetAppUsersByID gets the application users of the given id
	GetAppUsersByID(id int) error
	// GetDBUsersByID gets the db users of the given id
	GetDBUsersByID(id int) error
	// GetAllUsersByID gets mysql cluster, application and db users of the given id
	GetAllUsersByID(id int) error
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
