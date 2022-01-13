package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type MiddlewareCluster interface {
	// Identity returns the identity
	Identity() int
	// GetClusterName returns the cluster name
	GetClusterName() string
	// GetEnvID returns the env id
	GetEnvID() int
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// GetMiddlewareServers gets the middleware server id list of this cluster
	GetMiddlewareServers() ([]MiddlewareServer, error)
	// GetUsersByMiddlewareClusterID gets user list that own the middleware cluster
	GetUsersByMiddlewareClusterID() ([]User, error)
	// Set sets MiddlewareCluster with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete sets DelFlag to 1
	Delete()
	// MarshalJSON marshals MiddlewareCluster to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified field of the MiddlewareCluster to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type MiddlewareClusterRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// GetAll gets all middleware clusters from the middleware
	GetAll() ([]MiddlewareCluster, error)
	// GetByEnv gets middleware clusters of given env id from the middleware
	GetByEnv(envID int) ([]MiddlewareCluster, error)
	// GetByID gets a middleware cluster by the identity from the middleware
	GetByID(id int) (MiddlewareCluster, error)
	// GetByName gets a middleware cluster of given cluster name from the middle ware
	GetByName(clusterName string) (MiddlewareCluster, error)
	// GetID gets the identity with given cluster name and env id from the middleware
	GetID(clusterName string, envID int) (int, error)
	// GetUsersByMiddlewareClusterID get user list that own the middleware cluster
	GetUsersByMiddlewareClusterID(clusterID int) ([]User, error)
	// MiddlewareClusterAddUser adds a new map of middleware cluster and user in the middleware
	MiddlewareClusterAddUser(middlewareClusterID, userID int) error
	// MiddlewareClusterDeleteUser deletes the map of middleware cluster and user in the middleware
	MiddlewareClusterDeleteUser(middlewareClusterID, userID int) error
	// Create creates a middleware cluster in the middleware
	Create(mc MiddlewareCluster) (MiddlewareCluster, error)
	// Update updates the middleware cluster in the middleware
	Update(mc MiddlewareCluster) error
	// Delete deletes the middleware cluster in the middleware
	Delete(id int) error
}

type MiddlewareClusterService interface {
	// GetMiddlewareClusters returns middleware clusters of the service
	GetMiddlewareClusters() []MiddlewareCluster
	// GetMiddlewareServers returns middleware servers of the service
	GetMiddlewareServers() []MiddlewareServer
	// GetAll gets all middleware clusters from the middleware
	GetAll() error
	// GetByEnv gets middleware clusters of given env id
	GetByEnv(envID int) error
	// GetByID gets a middleware cluster of the given id from the middleware
	GetByID(id int) error
	// GetByName gets a middleware cluster of given cluster name
	GetByName(clusterName string) error
	// GetUsers returns users of the service
	GetUsers() []User
	// GetMiddlewareServersByID gets the middleware servers of given cluster id
	GetMiddlewareServersByID(clusterID int) error
	// GetUsersByMiddlewareClusterID  gets Users that own the middleware cluster
	GetUsersByMiddlewareClusterID(clusterID int) error
	// AddUser adds a new map of middleware cluster and user in the middleware
	AddUser(middlewareClusterID, userID int) error
	// DeleteUser deletes the map of middleware cluster and user in the middleware
	DeleteUser(middlewareClusterID, userID int) error
	// Create creates a middleware cluster in the middleware
	Create(fields map[string]interface{}) error
	// Update gets a middleware cluster of the given id from the middleware,
	// and then updates its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the middleware
	Update(id int, fields map[string]interface{}) error
	// Delete deletes the middleware cluster of given id in the middleware
	Delete(id int) error
	// Marshal marshals MiddlewareClusterService.MiddlewareClusters to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the MiddlewareClusterService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
