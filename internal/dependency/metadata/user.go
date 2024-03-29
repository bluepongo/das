package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type User interface {
	// Identity returns the identity
	Identity() int
	// GetUserName returns the username
	GetUserName() string
	// GetDepartmentName returns the department name
	GetDepartmentName() string
	// GetEmployeeID returns the employee id
	GetEmployeeID() string
	// GetAccountName returns the account name
	GetAccountName() string
	// GetEmail returns the email
	GetEmail() string
	// GetEmail returns the telephone
	GetTelephone() string
	// GetMobile returns the mobile
	GetMobile() string
	// GetRole returns the role
	GetRole() int
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetAllMySQLServers gets all mysql servers of this user from the middleware
	GetAllMySQLServers() ([]MySQLServer, error)
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// Set sets User with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete sets DelFlag to 1
	Delete()
	// MarshalJSON marshals User to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified field of the User to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type UserRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// GetAll gets all databases from the middleware
	GetAll() ([]User, error)
	// GetID gets the identity with given account name from the middleware
	GetID(accountName string) (int, error)
	// GetByID gets a user by the identity from the middleware
	GetByID(id int) (User, error)
	// GetByUserName gets users of given username from the middleware
	GetByUserName(userName string) ([]User, error)
	// GetByEmployeeID gets a user of given employee id from the middleware
	GetByEmployeeID(employeeID string) (User, error)
	// GetByAccountName gets a user of given account name from the middleware
	GetByAccountName(accountName string) (User, error)
	// GetByEmail gets a user of given email from the middleware
	GetByEmail(email string) (User, error)
	// GetByTelephone gets a user of given telephone from the middleware
	GetByTelephone(telephone string) (User, error)
	// GetByTelephone gets a user of given mobile from the middleware
	GetByMobile(mobile string) (User, error)
	// GetByAccountNameOrEmployeeID gets a user of given loginName from the middleware
	GetByAccountNameOrEmployeeID(loginName string) (User, error)
	// GetAppsByUserID gets app list that this user owns
	GetAppsByUserID(id int) ([]App, error)
	// GetDBsByUserID gets app list that this user owns
	GetDBsByUserID(id int) ([]DB, error)
	// GetMiddlewareClustersByUserID gets middleware cluster list that this user owns
	GetMiddlewareClustersByUserID(id int) ([]MiddlewareCluster, error)
	// GetMySQLClustersByUserID gets mysql cluster list that this user owns
	GetMySQLClustersByUserID(id int) ([]MySQLCluster, error)
	// GetAllMySQLServersByUserID gets mysql servers list that this user owns
	GetAllMySQLServersByUserID(id int) ([]MySQLServer, error)
	// Create creates a user in the middleware
	Create(db User) (User, error)
	// Update updates a user in the middleware
	Update(db User) error
	// Delete deletes a user in the middleware
	Delete(id int) error
}

type UserService interface {
	// GetUsers returns users of the service
	GetUsers() []User
	// GetApps returns the apps of the service
	GetApps() []App
	// GetDBs returns the dbs of the service
	GetDBs() []DB
	// GetMiddlewareClusters returns the middleware clusters of the service
	GetMiddlewareClusters() []MiddlewareCluster
	// GetMySQLClusters returns the mysql clusters of the service
	GetMySQLClusters() []MySQLCluster
	// GetAll gets all users
	GetAll() error
	// GetByID gets a user by the identity
	GetByID(id int) error
	// GetByUserName gets users of given user name
	GetByUserName(userName string) error
	// GetByEmployeeID gets a user of given employee id
	GetByEmployeeID(employeeID string) error
	// GetByAccountName gets a user of given account name
	GetByAccountName(accountName string) error
	// GetByEmail gets a user of given email
	GetByEmail(email string) error
	// GetByTelephone gets a user of given telephone
	GetByTelephone(telephone string) error
	// GetByTelephone gets a user of given mobile
	GetByMobile(mobile string) error
	// GetByAccountNameOrEmployeeID gets a user of given login name from the middleware
	GetByAccountNameOrEmployeeID(loginName string) error
	// GetAppsByUserID gets apps that this user owns
	GetAppsByUserID(id int) error
	// GetDBsByUserID gets apps that this user owns
	GetDBsByUserID(id int) error
	// GetMiddlewareClustersByUserID gets middleware clusters that this user owns
	GetMiddlewareClustersByUserID(id int) error
	// GetMySQLClustersByUserID gets mysql clusters that this user owns
	GetMySQLClustersByUserID(id int) error
	// GetAllMySQLServersByUserID gets mysql servers list that this user owns
	GetAllMySQLServersByUserID(id int) error
	// Create creates a user in the middleware
	Create(fields map[string]interface{}) error
	// Update gets a user of the given id from the middleware,
	// and then updates its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the middleware
	Update(id int, fields map[string]interface{}) error
	// Delete deletes the user of given id in the middleware
	Delete(id int) error
	// Marshal marshals UserService.Users to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the UserService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
