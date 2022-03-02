package metadata

import (
	"fmt"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
)

const dbDBsStruct = "DBs"

var _ metadata.DBService = (*DBService)(nil)

type DBService struct {
	metadata.DBRepo
	DBs          []metadata.DB         `json:"dbs"`
	MySQLCluster metadata.MySQLCluster `json:"mysql_cluster"`
	Apps         []metadata.App        `json:"apps"`
	Owners       []metadata.User       `json:"owners"`
}

// NewDBService returns a new *DBService
func NewDBService(repo metadata.DBRepo) *DBService {
	return &DBService{DBRepo: repo}
}

// NewDBServiceWithDefault returns a new *DBService with default repository
func NewDBServiceWithDefault() *DBService {
	return NewDBService(NewDBRepoWithGlobal())
}

// GetDBs returns the databases of the service
func (ds *DBService) GetDBs() []metadata.DB {
	return ds.DBs
}

// GetMySQLCluster returns the mysql cluster of the service
func (ds *DBService) GetMySQLCluster() metadata.MySQLCluster {
	return ds.MySQLCluster
}

// GetApps returns the apps of the service
func (ds *DBService) GetApps() []metadata.App {
	return ds.Apps
}

// GetOwners returns the owners of the service
func (ds *DBService) GetOwners() []metadata.User {
	return ds.Owners
}

// GetAll gets all databases from the middleware
func (ds *DBService) GetAll() error {
	var err error

	ds.DBs, err = ds.DBRepo.GetAll()

	return err
}

// GetByEnv gets all databases of given env_id
func (ds *DBService) GetByEnv(envID int) error {
	var err error

	ds.DBs, err = ds.DBRepo.GetByEnv(envID)

	return err
}

// GetByID gets the database of the given id from the middleware
func (ds *DBService) GetByID(id int) error {
	db, err := ds.DBRepo.GetByID(id)
	if err != nil {
		return err
	}

	ds.DBs = nil
	ds.DBs = append(ds.DBs, db)

	return nil
}

// GetDBByNameAndClusterInfo gets the database of the given db name and cluster info from the middleware
func (ds *DBService) GetDBByNameAndClusterInfo(name string, clusterID, clusterType int) error {
	db, err := ds.DBRepo.GetDBByNameAndClusterInfo(name, clusterID, clusterType)
	if err != nil {
		return err
	}

	ds.DBs = nil
	ds.DBs = append(ds.DBs, db)

	return nil
}

// GetMySQLClusterByID gets the cluster of the db
func (ds *DBService) GetMySQLClusterByID(id int) error {
	var err error

	ds.MySQLCluster, err = ds.DBRepo.GetMySQLCLusterByID(id)

	return err
}

// GetAppsByDBID gets an apps that uses this db
func (ds *DBService) GetAppsByDBID(dbID int) error {
	var err error

	ds.Apps, err = ds.DBRepo.GetAppsByDBID(dbID)

	return err
}

// GetAppOwnersByID gets the application owners of the given id
func (ds *DBService) GetAppUsersByDBID(id int) error {
	var err error

	ds.Owners, err = ds.DBRepo.GetAppUsersByDBID(id)

	return err
}

// GetDBOwnersByID gets the db owners of the given id
func (ds *DBService) GetUsersByDBID(id int) error {
	var err error

	ds.Owners, err = ds.DBRepo.GetUsersByDBID(id)

	return err
}

// GetAllOwnersByID gets both application and db owners of the given id
func (ds *DBService) GetAllUsersByDBID(id int) error {
	var err error

	ds.Owners, err = ds.DBRepo.GetAllUsersByDBID(id)

	return err
}

// Create creates a new database in the middleware
func (ds *DBService) Create(fields map[string]interface{}) error {
	// generate new map
	_, dbNameExists := fields[dbDBNameStruct]
	_, clusterIDExists := fields[dbClusterIDStruct]
	_, clusterTypeExists := fields[dbClusterTypeStruct]
	_, envIDExists := fields[dbEnvIDStruct]
	if !dbNameExists || !clusterIDExists || !clusterTypeExists || !envIDExists {
		return message.NewMessage(message.ErrFieldNotExists, fmt.Sprintf("%s and %s and %s and %s",
			dbDBNameStruct, dbClusterIDStruct, dbClusterTypeStruct, dbEnvIDStruct))
	}
	// create a new entity
	dbInfo, err := NewDBInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	db, err := ds.DBRepo.Create(dbInfo)
	if err != nil {
		return err
	}

	ds.DBs = nil
	ds.DBs = append(ds.DBs, db)

	return nil
}

// Update gets a database of the given id from the middleware,
// and then updates its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (ds *DBService) Update(id int, fields map[string]interface{}) error {
	err := ds.GetByID(id)
	if err != nil {
		return err
	}
	err = ds.DBs[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return ds.DBRepo.Update(ds.DBs[constant.ZeroInt])
}

// Delete deletes the database of given id in the middleware
func (ds *DBService) Delete(id int) error {
	err := ds.GetByID(id)
	if err != nil {
		return err
	}

	return ds.DBRepo.Delete(id)
}

// AddApp adds a new map of app and database in the middleware
func (ds *DBService) AddApp(dbID, appID int) error {
	err := ds.DBRepo.AddApp(dbID, appID)
	if err != nil {
		return err
	}

	return ds.GetAppsByDBID(dbID)
}

// DeleteApp deletes the map of app and database in the middleware
func (ds *DBService) DeleteApp(dbID, appID int) error {
	err := ds.DBRepo.DeleteApp(dbID, appID)
	if err != nil {
		return err
	}

	return ds.GetAppsByDBID(dbID)
}

// DBAddUser adds a new map of user and database in the middleware
func (ds *DBService) DBAddUser(dbID, userID int) error {
	err := ds.DBRepo.DBAddUser(dbID, userID)
	if err != nil {
		return err
	}

	return ds.GetUsersByDBID(dbID)
}

// DBDeleteUser deletes the map of user and database in the middleware
func (ds *DBService) DBDeleteUser(dbID, userID int) error {
	err := ds.DBRepo.DBDeleteUser(dbID, userID)
	if err != nil {
		return err
	}

	return ds.GetUsersByDBID(dbID)
}

// Marshal marshals DBService.DBs to json bytes
func (ds *DBService) Marshal() ([]byte, error) {
	return ds.MarshalWithFields(dbDBsStruct)
}

// MarshalWithFields marshals only specified fields of the DBService to json bytes
func (ds *DBService) MarshalWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(ds, fields...)
}
