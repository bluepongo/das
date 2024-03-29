package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

// ResourceGroup is the entity interface
type ResourceGroup interface {
	// Identity returns the identity
	Identity() int
	// GetGroupUUID returns the resource group uuid
	GetGroupUUID() string
	// GetGroupName returns the resource group name
	GetGroupName() string
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// GetResourceRoles get all resource roles of this resource group
	GetResourceRoles() ([]ResourceRole, error)
	// GetMySQLClusters gets the mysql clusters of this resource group
	GetMySQLClusters() ([]MySQLCluster, error)
	// GetMySQLServers gets the mysql servers of this resource group
	GetMySQLServers() ([]MySQLServer, error)
	// GetMiddlewareClusters gets the mysql clusters of this resource group
	GetMiddlewareClusters() ([]MiddlewareCluster, error)
	// GetMiddlewareServers gets the mysql servers of this resource group
	GetMiddlewareServers() ([]MiddlewareServer, error)
	// GetUsers gets the users of this resource group
	GetUsers() ([]User, error)
	// GetDASAdminUsers gets the das admin users of this resource group
	GetDASAdminUsers() ([]User, error)
	// Set sets the resource group with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete sets DelFlag to 1
	Delete()
	// AddMySQLCluster adds mysql cluster to the resource group
	AddMySQLCluster(mysqlClusterID int) error
	// DeleteMySQLCluster deletes mysql cluster from the resource group
	DeleteMySQLCluster(mysqlClusterID int) error
	// AddMiddlewareCluster adds middleware cluster to the resource group
	AddMiddlewareCluster(mysqlClusterID int) error
	// DeleteMiddlewareCluster deletes middleware cluster from the resource group
	DeleteMiddlewareCluster(middlewareClusterID int) error
	// MarshalJSON marshals ResourceGroup to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified field of the ResourceGroup to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

// ResourceGroupRepo is the repository interface
type ResourceGroupRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a mysql.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// GetAll gets all resource groups from the middleware
	GetAll() ([]ResourceGroup, error)
	// GetByID gets the resource group by the identity from the middleware
	GetByID(id int) (ResourceGroup, error)
	// GetByGroupUUID gets the resource group with given resource group id from the middleware
	GetByGroupUUID(groupUUID string) (ResourceGroup, error)
	// GetID gets the identity with given resource group id from the middleware
	GetID(groupUUID string) (int, error)
	// GetResourceRolesByID get all resource roles with given resource group id from the middleware
	GetResourceRolesByID(id int) ([]ResourceRole, error)
	// GetMySQLClustersByID gets the mysql cluster with given resource group id from the middleware
	GetMySQLClustersByID(id int) ([]MySQLCluster, error)
	// GetMySQLServersByID gets the mysql servers with given resource group uuid from the middleware
	GetMySQLServersByID(id int) ([]MySQLServer, error)
	// GetMiddlewareClustersByID gets the middleware cluster with given resource group id from the middleware
	GetMiddlewareClustersByID(id int) ([]MiddlewareCluster, error)
	// GetMiddlewareServersByID gets the middleware servers with given resource group uuid from the middleware
	GetMiddlewareServersByID(id int) ([]MiddlewareServer, error)
	// GetUsersByID gets the users with given resource group id from the middleware
	GetUsersByID(id int) ([]User, error)
	// GetDASAdminUsersByID gets the das admin users with given resource group id from the middleware
	GetDASAdminUsersByID(id int) ([]User, error)
	// Create creates a mysql server in the middleware
	Create(resourceGroup ResourceGroup) (ResourceGroup, error)
	// Update updates the mysql server in the middleware
	Update(resourceGroup ResourceGroup) error
	// Delete deletes the mysql server from the middleware
	Delete(id int) error
	// AddMySQLCluster adds mysql cluster to the resource group
	AddMySQLCluster(resourceGroupID int, mysqlClusterID int) error
	// DeleteMySQLCluster deletes mysql cluster from the resource group
	DeleteMySQLCluster(resourceGroupID int, mysqlClusterID int) error
	// AddMiddlewareCluster adds middleware cluster to the resource group
	AddMiddlewareCluster(resourceGroupID int, middlewareClusterID int) error
	// DeleteMiddlewareCluster deletes middleware cluster from the resource group
	DeleteMiddlewareCluster(resourceGroupID, middlewareClusterID int) error
}

// ResourceGroupService is the service interface
type ResourceGroupService interface {
	// GetResourceGroups returns resource groups of the service
	GetResourceGroups() []ResourceGroup
	// GetMySQLCluster returns the mysql clusters of the service
	GetMySQLClusters() []MySQLCluster
	// GetMySQLServers returns the mysql servers of the service
	GetMySQLServers() []MySQLServer
	// GetMiddlewareClusters returns the middleware clusters of the service
	GetMiddlewareClusters() []MiddlewareCluster
	// GetMiddlewareServers returns the middleware servers of the service
	GetMiddlewareServers() []MiddlewareServer
	// GetResourceRoles returns the resource roles of the service
	GetResourceRoles() []ResourceRole
	// GetUsers returns the users of the service
	GetUsers() []User
	// GetAll gets all mysql servers from the mysql
	GetAll() error
	// GetByID gets the resource group of the given id
	GetByID(id int) error
	// GetByGroupUUID gets the resource group by group uuid
	GetByGroupUUID(groupUUID string) error
	// GetResourceRolesByID get all resource roles with given resource group id
	GetResourceRolesByID(id int) error
	// GetMySQLClustersByID gets the mysql clusters with given resource group id
	GetMySQLClustersByID(id int) error
	// GetMySQLServersByID gets the mysql servers with given resource group id
	GetMySQLServersByID(id int) error
	// GetMiddlewareClustersByID gets the middleware clusters with given resource group id
	GetMiddlewareClustersByID(id int) error
	// GetMiddlewareServersByID gets the middleware servers with given resource group id
	GetMiddlewareServersByID(id int) error
	// GetUsersByID gets the users with given resource group id
	GetUsersByID(id int) error
	// GetDASAdminUsersByID gets the das admin users with given resource group id
	GetDASAdminUsersByID(id int) error
	// GetResourceRolesByGroupUUID get all resource roles with given resource group uuid
	GetResourceRolesByGroupUUID(groupUUID string) error
	// GetMySQLClustersByGroupUUID gets the mysql clusters with given resource group uuid
	GetMySQLClustersByGroupUUID(groupUUID string) error
	// GetMySQLServersByGroupUUID gets the mysql servers with given resource group uuid
	GetMySQLServersByGroupUUID(groupUUID string) error
	// GetMiddlewareClustersByGroupUUID gets the middleware clusters with given resource group uuid
	GetMiddlewareClustersByGroupUUID(groupUUID string) error
	// GetMiddlewareServersByGroupUUID gets the middleware servers with given resource group uuid
	GetMiddlewareServersByGroupUUID(groupUUID string) error
	// GetUsersByGroupUUID gets the users with given resource group uuid
	GetUsersByGroupUUID(groupUUID string) error
	// GetDASAdminUsersByGroupUUID gets the das admin users with given resource group uuid
	GetDASAdminUsersByGroupUUID(groupUUID string) error
	// Create creates a mysql server in the mysql
	Create(fields map[string]interface{}) error
	// Update gets a mysql server of the given id from the mysql,
	// and then updates its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the mysql
	Update(id int, fields map[string]interface{}) error
	// Delete deletes the mysql server of given id
	Delete(id int) error
	// AddMySQLCluster adds mysql cluster to the resource group
	AddMySQLCluster(resourceGroupID int, mysqlClusterID int) error
	// DeleteMySQLCluster deletes mysql cluster from the resource group
	DeleteMySQLCluster(resourceGroupID int, mysqlClusterID int) error
	// AddMiddlewareCluster adds middleware cluster to the resource group
	AddMiddlewareCluster(resourceGroupID int, middlewareClusterID int) error
	// DeleteMiddlewareCluster deletes middleware cluster from the resource group
	DeleteMiddlewareCluster(resourceGroupID int, middlewareClusterID int) error
	// Marshal marshals ResourceGroupService.ResourceGroups to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the ResourceGroupService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
