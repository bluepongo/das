package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

// ResourceRole is the entity interface
type ResourceRole interface {
	// Identity returns the identity
	Identity() int
	// GetRoleUUID returns the resource role uuid
	GetRoleUUID() string
	// GetRoleName returns the resource role name
	GetRoleName() string
	// GetResourceGroupID returns the resource group id
	GetResourceGroupID() int
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// GetResourceGroup gets the resource group which this role belongs to
	GetResourceGroup() (ResourceGroup, error)
	// GetUsers gets the users of this resource role
	GetUsers() ([]User, error)
	// Set sets the resource group with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete sets DelFlag to 1
	Delete()
	// AddUser adds a map of the resource role and user
	AddUser(userID int) error
	// DeleteUser deletes the map of the resource role and user
	DeleteUser(userID int) error
	// MarshalJSON marshals ResourceRole to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified field of the ResourceRole to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

// ResourceRoleRepo is the repository interface
type ResourceRoleRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a mysql.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// GetAll gets all resource groups from the middleware
	GetAll() ([]ResourceRole, error)
	// GetByID gets the resource role by the identity from the middleware
	GetByID(id int) (ResourceRole, error)
	// GetID gets the identity with given resource role id from the middleware
	GetID(groupUUID string) (int, error)
	// GetByRoleUUID gets the resource role with given resource role id from the middleware
	GetByRoleUUID(groupUUID string) (ResourceRole, error)
	// GetResourceGroup gets the resource group which this role belongs to with given resource role id from the middleware
	GetResourceGroup(id int) (ResourceGroup, error)
	// GetUsersByID gets the users with given resource group id from the middleware
	GetUsersByID(id int) ([]User, error)
	// Create creates a mysql server in the middleware
	Create(rr ResourceRole) (ResourceRole, error)
	// Update updates the mysql server in the middleware
	Update(rr ResourceRole) error
	// Delete deletes the mysql server from the middleware
	Delete(id int) error
	// AddUser adds a map of the resource role and user from the middleware
	AddUser(roleID int, userID int) error
	// DeleteUser deletes the map of the resource role and user from the middleware
	DeleteUser(roleID int, userID int) error
}

// ResourceRoleService is the service interface
type ResourceRoleService interface {
	// GetResourceRoles returns the resource roles of the service
	GetResourceRoles() []ResourceRole
	// GetResourceGroup returns the resource group of the service
	GetResourceGroup() ResourceGroup
	// GetUsers returns the users of the service
	GetUsers() []User
	// GetAll gets all resource roles
	GetAll() error
	// GetByID gets the resource role of the given id
	GetByID(id int) error
	// GetByRoleUUID gets the resource role by role uuid
	GetByRoleUUID(groupUUID string) (ResourceRole, error)
	// GetResourceGroupByID gets the resource group with given resource role id
	GetResourceGroupByID(id int) (ResourceGroup, error)
	// GetUsersByID gets the users with given resource role id
	GetUsersByID(id int) ([]User, error)
	// GetUsersByRoleUUID gets the users with given resource role uuid
	GetUsersByRoleUUID(id int) ([]User, error)
	// Create creates a mysql server in the mysql
	Create(fields map[string]interface{}) error
	// Update gets a mysql server of the given id from the mysql,
	// and then updates its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the mysql
	Update(id int, fields map[string]interface{}) error
	// Delete deletes the mysql server of given id
	Delete(id int) error
	// AddUser adds a map of the resource role and user
	AddUser(roleID int, userID int) error
	// DeleteUser deletes the map of the resource role and user
	DeleteUser(roleID int, userID int) error
	// Marshal marshals ResourceRoleService.ResourceRoles to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the ResourceRoleService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
