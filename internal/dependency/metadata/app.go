package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type App interface {
	// Identity returns the identity
	Identity() int
	// GetSystemName returns the app name
	GetAppName() string
	// GetLevel returns the level
	GetLevel() int
	// GetOwnerID returns the owner id
	GetOwnerID() int
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// GetDBs gets database identity list that the app uses
	GetDBs() ([]DB, error)
	// GetUsers gets user list that own the app
	GetUsers() ([]User, error)
	// Set sets App with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// Delete sets DelFlag to 1
	Delete()
	// AddDB adds a new map of the app and database in the middleware
	AddDB(dbID int) error
	// DeleteDB deletes the map of the app and database in the middleware
	DeleteDB(dbID int) error
	// AddUser adds a new map of the app and user in the middleware
	AddUser(userID int) error
	// DeleteUser deletes the map of the app and user in the middleware
	DeleteUser(userID int) error
	// MarshalJSON marshals App to json bytes
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified fields of the App to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type AppRepo interface {
	// Execute executes command with arguments on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// GetAll gets all apps from the middleware
	GetAll() ([]App, error)
	// GetByID gets an app by the identity from the middleware
	GetByID(id int) (App, error)
	// GetID gets the identity with given app name from the middleware
	GetID(appName string) (int, error)
	// GetAppSystemByName gets the app by name from the middleware
	GetAppByName(appName string) (App, error)
	// GetDBsByAppID gets databases that app uses
	GetDBsByAppID(id int) ([]DB, error)
	// GetUsersByID gets user list that own the app
	GetUsersByAppID(id int) ([]User, error)
	// Create creates an app in the middleware
	Create(appSystem App) (App, error)
	// Update updates the app in the middleware
	Update(appSystem App) error
	// Delete deletes the app in the middleware
	Delete(id int) error
	// AddDB adds a new map of app and database in the middleware
	AddDB(appID, dbID int) error
	// DeleteDB delete the map of app and database in the middleware
	DeleteDB(appID, dbID int) error
	// AddUser adds a new map of app and user in the middleware
	AddUser(appID, userID int) error
	// DeleteUser delete the map of app and user in the middleware
	DeleteUser(appID, userID int) error
}

type AppService interface {
	// GetApps returns apps of the service
	GetApps() []App
	// GetDBs returns dbs of the service
	GetDBs() []DB
	// GetUsers returns users of the service
	GetUsers() []User
	// GetAll gets all apps from the middleware
	GetAll() error
	// GetByID gets an app of the given id from the middleware
	GetByID(id int) error
	// GetAppByName gets App from the middleware by name
	GetAppByName(appName string) error
	// GetDBsByAppID gets databases that the app uses
	GetDBsByAppID(id int) error
	// GetUsersByID gets Users that own the app
	GetUsersByAppID(id int) error
	// Create creates an app in the middleware
	Create(fields map[string]interface{}) error
	// Update gets the app of the given id from the middleware,
	// and then updates its fields that was specified in fields argument,
	// key is the filed name and value is the new field value,
	// it saves the changes to the middleware
	Update(id int, fields map[string]interface{}) error
	// Delete deletes the app of given id in the middleware
	Delete(id int) error
	// AddDB adds a new map of app and database in the middleware
	AddDB(appID, dbID int) error
	// DeleteDB deletes the map of app and database in the middleware
	DeleteDB(appID, dbID int) error
	// AddUser adds a new map of app and user in the middleware
	AddUser(appID, userID int) error
	// DeleteUser deletes the map of app and user in the middleware
	DeleteUser(appID, userID int) error
	// Marshal marshals AppService.Apps to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the AppService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
