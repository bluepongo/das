package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type DB interface {
	// Identity returns the identity
	Identity() int
	// GetDBName returns the db name
	GetDBName() string
	// GetClusterID returns the cluster id
	GetClusterID() int
	// GetClusterType returns the cluster type
	GetClusterType() int
	// GetEnvID returns the env id
	GetEnvID() int
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// GetApps gets apps that uses this db
	GetApps() ([]App, error)
	// GetMySQLCluster gets the mysql cluster of this db
	GetMySQLCluster() (MySQLCluster, error)
	// GetAppUsers gets the application users of this db
	GetAppUsers() ([]User, error)
	// GetDBUsers gets the db users of this db
	GetDBUsers() ([]User, error)
	// GetAllUsers gets both application and db users of this db
	GetAllUsers() ([]User, error)
	// Set sets DB with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete sets DelFlag to 1
	Delete()
	// AddApp adds a new map of the app and database in the middleware
	AddApp(appID int) error
	// DeleteApp deletes a new map of the app and database in the middleware
	DeleteApp(appID int) error
	// DBAddUser adds a new map of the user and database in the middleware
	DBAddUser(userID int) error
	// DBDeleteUser deletes a new map of the user and database in the middleware
	DBDeleteUser(userID int) error
	// MarshalJSON marshals DB to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified field of the DB to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type DBRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// GetAll gets all databases from the middleware
	GetAll() ([]DB, error)
	// GetByEnv gets databases of given env id from the middleware
	GetByEnv(envID int) ([]DB, error)
	// GetByID gets a database by the identity from the middleware
	GetByID(id int) (DB, error)
	// GetDBByNameAndClusterInfo gets a database by the db name and cluster info from the middleware
	GetDBByNameAndClusterInfo(name string, clusterID, clusterType int) (DB, error)
	// GetID gets the identity with given database name, cluster id and cluster type from the middleware
	GetID(dbName string, clusterID int, clusterType int) (int, error)
	// GetMySQLCLusterByID gets the mysql cluster of the given id from the middleware
	GetMySQLCLusterByID(id int) (MySQLCluster, error)
	// GetAppsByDBID gets apps that uses this db
	GetAppsByDBID(id int) ([]App, error)
	// GetAppUsersByDBID gets the application users of the given id from the middleware
	GetAppUsersByDBID(id int) ([]User, error)
	// GetUsersByDBID gets the db users of the given id from the middleware
	GetUsersByDBID(id int) ([]User, error)
	// GetAllUsersByDBID gets both application and db users of the given id from the middleware
	GetAllUsersByDBID(id int) ([]User, error)
	// Create creates a database in the middleware
	Create(db DB) (DB, error)
	// Update updates the database in the middleware
	Update(db DB) error
	// Delete deletes the database in the middleware
	Delete(id int) error
	// AddApp adds a new map of the app and database in the middleware
	AddApp(dbID, appID int) error
	// DeleteApp deletes a map of the app and database in the middleware
	DeleteApp(dbID, appID int) error
	// DBAddUser adds a new map of the user and database in the middleware
	DBAddUser(dbID, userID int) error
	// DBDeleteUser deletes a map of the user and database in the middleware
	DBDeleteUser(dbID, userID int) error
}

type DBService interface {
	// GetDBs returns the databases of the service
	GetDBs() []DB
	// GetMySQLCluster returns the mysql cluster of the service
	GetMySQLCluster() MySQLCluster
	// GetApps returns the apps of the service
	GetApps() []App
	// GetUsers returns the users of the service
	GetUsers() []User
	// GetAll gets all databases from the middleware
	GetAll() error
	// GetByEnv gets databases of given env id
	GetByEnv(envID int) error
	// GetByID gets a database of the given id from the middleware
	GetByID(id int) error
	// GetDBByNameAndClusterInfo gets an database of the given db name and cluster info from the middleware
	GetDBByNameAndClusterInfo(name string, clusterID, clusterType int) error
	// GetMySQLClusterByID gets the cluster of the db
	GetMySQLClusterByID(id int) error
	// GetAppsByDBID gets apps that uses this db
	GetAppsByDBID(id int) error
	// GetAppUsersByDBID gets the application users of the given id
	GetAppUsersByDBID(id int) error
	// GetUsersByDBID gets the db users of the given id
	GetUsersByDBID(id int) error
	// GetAllUsersByDBID gets both application and db users of the given id
	GetAllUsersByDBID(id int) error
	// Create creates a database in the middleware
	Create(fields map[string]interface{}) error
	// Update gets a database of the given id from the middleware,
	// and then updates its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the middleware
	Update(id int, fields map[string]interface{}) error
	// Delete deletes the database of given id in the middleware
	Delete(id int) error
	// AddApp adds a new map of app and database in the middleware
	AddApp(dbID, appID int) error
	// DeleteApp deletes the map of app and database in the middleware
	DeleteApp(dbID, appID int) error
	// DBAddUser adds a new map of the user and database in the middleware
	DBAddUser(dbID, userID int) error
	// DBDeleteUser deletes a map of the user and database in the middleware
	DBDeleteUser(dbID, userID int) error
	// Marshal marshals DBService.DBs to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the DBService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
