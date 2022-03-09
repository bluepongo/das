package metadata

import (
	"time"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
)

const (
	dbDBNameStruct      = "DBName"
	dbClusterIDStruct   = "ClusterID"
	dbClusterTypeStruct = "ClusterType"
	dbEnvIDStruct       = "EnvID"
	dbDelFlagStruct     = "DelFlag"
)

var _ metadata.DB = (*DBInfo)(nil)

type DBInfo struct {
	metadata.DBRepo
	ID             int       `middleware:"id" json:"id"`
	DBName         string    `middleware:"db_name" json:"db_name"`
	ClusterID      int       `middleware:"cluster_id" json:"cluster_id"`
	ClusterType    int       `middleware:"cluster_type" json:"cluster_type"`
	EnvID          int       `middleware:"env_id" json:"env_id"`
	DelFlag        int       `middleware:"del_flag" json:"del_flag"`
	CreateTime     time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewDBInfo returns a new *DBInfo
func NewDBInfo(repo *DBRepo, id int, dbName string, clusterID int, clusterType int,
	envID int, delFlag int, createTime time.Time, lastUpdateTime time.Time) *DBInfo {
	return &DBInfo{
		repo,
		id,
		dbName,
		clusterID,
		clusterType,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewDBInfoWithGlobal NewDBInfo returns a new DBInfo with default DBRepo
func NewDBInfoWithGlobal(id int, dbName string, clusterID, clusterType, envID, delFlag int,
	createTime, lastUpdateTime time.Time) *DBInfo {
	return &DBInfo{
		NewDBRepoWithGlobal(),
		id,
		dbName,
		clusterID,
		clusterType,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewEmptyDBInfoWithRepo return a new DBInfo
func NewEmptyDBInfoWithRepo(repo *DBRepo) *DBInfo {
	return &DBInfo{DBRepo: repo}
}

// NewEmptyDBInfoWithGlobal return a new DBInfo
func NewEmptyDBInfoWithGlobal() *DBInfo {
	return NewEmptyDBInfoWithRepo(NewDBRepoWithGlobal())
}

// NewDBInfoWithDefault returns a new *DBInfo with default DBRepo
func NewDBInfoWithDefault(dbName string, clusterID, clusterType, envID int) *DBInfo {
	return &DBInfo{
		DBRepo:      NewDBRepoWithGlobal(),
		DBName:      dbName,
		ClusterID:   clusterID,
		ClusterType: clusterType,
		EnvID:       envID,
	}
}

// NewDBInfoWithMapAndRandom returns a new *DBInfo with given map
func NewDBInfoWithMapAndRandom(fields map[string]interface{}) (*DBInfo, error) {
	di := &DBInfo{}
	err := common.SetValuesWithMapAndRandom(di, fields)
	if err != nil {
		return nil, err
	}

	return di, nil
}

// Identity returns the identity
func (di *DBInfo) Identity() int {
	return di.ID
}

// GetDBName returns the db name
func (di *DBInfo) GetDBName() string {
	return di.DBName
}

// GetClusterID returns the cluster id
func (di *DBInfo) GetClusterID() int {
	return di.ClusterID
}

// GetClusterType returns the cluster type
func (di *DBInfo) GetClusterType() int {
	return di.ClusterType
}

// GetEnvID returns the env id
func (di *DBInfo) GetEnvID() int {
	return di.EnvID
}

// GetDelFlag returns the delete flag
func (di *DBInfo) GetDelFlag() int {
	return di.DelFlag
}

// GetCreateTime returns the create time
func (di *DBInfo) GetCreateTime() time.Time {
	return di.CreateTime
}

// GetLastUpdateTime returns the last update time
func (di *DBInfo) GetLastUpdateTime() time.Time {
	return di.LastUpdateTime
}

// GetApps gets app identity list that uses this db
func (di *DBInfo) GetApps() ([]metadata.App, error) {
	return di.DBRepo.GetAppsByDBID(di.ID)
}

// GetMySQLCluster gets the mysql cluster of this db
func (di *DBInfo) GetMySQLCluster() (metadata.MySQLCluster, error) {
	return di.DBRepo.GetMySQLCLusterByID(di.ID)
}

// GetAppUsers gets the application users of this db
func (di *DBInfo) GetAppUsers() ([]metadata.User, error) {
	return di.DBRepo.GetAppUsersByDBID(di.ID)
}

// GetDBUsers gets the db users of this db
func (di *DBInfo) GetDBUsers() ([]metadata.User, error) {
	return di.DBRepo.GetUsersByDBID(di.ID)
}

// GetAllUsers gets both application and db users of this db
func (di *DBInfo) GetAllUsers() ([]metadata.User, error) {
	return di.DBRepo.GetAllUsersByDBID(di.ID)
}

// Set sets DB with given fields, key is the field name and value is the relevant value of the key
func (di *DBInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(di, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to 1
func (di *DBInfo) Delete() {
	di.DelFlag = 1
}

// AddApp adds a new map of application system and database in the middleware
func (di *DBInfo) AddApp(appID int) error {
	return di.DBRepo.AddApp(di.ID, appID)
}

// DeleteApp delete the map of application system and database in the middleware
func (di *DBInfo) DeleteApp(appID int) error {
	return di.DBRepo.DeleteApp(di.ID, appID)
}

// DBAddUser adds a new map of user and database in the middleware
func (di *DBInfo) DBAddUser(userID int) error {
	return di.DBRepo.DBAddUser(di.ID, userID)
}

// DBDeleteUser delete the map of user and database in the middleware
func (di *DBInfo) DBDeleteUser(userID int) error {
	return di.DBRepo.DBDeleteUser(di.ID, userID)
}

// MarshalJSON marshals DB to json string
func (di *DBInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(di, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only specified field of the DB to json string
func (di *DBInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(di, fields...)
}
