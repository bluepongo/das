package metadata

import (
	"time"

	"github.com/pingcap/errors"
	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
)

const (
	mysqlClusterClusterNameStruct         = "ClusterName"
	mysqlCLusterMiddlewareClusterIDStruct = "MiddlewareClusterID"
	mysqlClusterMonitorSystemIDStruct     = "MonitorSystemID"
	mysqlClusterEnvIDStruct               = "EnvID"
	mysqlClusterDelFlagStruct             = "DelFlag"
)

var _ metadata.MySQLCluster = (*MySQLClusterInfo)(nil)

// MySQLClusterInfo is a struct map to table in the database
type MySQLClusterInfo struct {
	MySQLClusterRepo    metadata.MySQLClusterRepo
	ID                  int       `middleware:"id" json:"id"`
	ClusterName         string    `middleware:"cluster_name" json:"cluster_name"`
	MiddlewareClusterID int       `middleware:"middleware_cluster_id" json:"middleware_cluster_id"`
	MonitorSystemID     int       `middleware:"monitor_system_id" json:"monitor_system_id"`
	EnvID               int       `middleware:"env_id" json:"env_id"`
	DelFlag             int       `middleware:"del_flag" json:"del_flag"`
	CreateTime          time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime      time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewMySQLClusterInfo returns a new MySQLClusterInfo
func NewMySQLClusterInfo(repo *MySQLClusterRepo,
	id int,
	clusterName string,
	middlewareClusterID int,
	monitorSystemID int,
	// ownerID int,
	envID int,
	delFlag int,
	createTime, lastUpdateTime time.Time) *MySQLClusterInfo {
	return &MySQLClusterInfo{
		repo,
		id,
		clusterName,
		middlewareClusterID,
		monitorSystemID,
		// ownerID,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewMySQLClusterInfoWithGlobal returns a new MySQLClusterInfo with default MySQLClusterRepo
func NewMySQLClusterInfoWithGlobal(
	id int,
	clusterName string,
	middlewareClusterID int,
	monitorSystemID int,
	// ownerID int,
	envID int,
	delFlag int,
	createTime, lastUpdateTime time.Time) *MySQLClusterInfo {
	return &MySQLClusterInfo{
		NewMySQLClusterRepoWithGlobal(),
		id,
		clusterName,
		middlewareClusterID,
		monitorSystemID,
		// ownerID,
		envID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewEmptyMySQLClusterInfoWithGlobal returns a new MySQLClusterInfo with default MySQLClusterRepo
func NewEmptyMySQLClusterInfoWithGlobal() *MySQLClusterInfo {
	return &MySQLClusterInfo{MySQLClusterRepo: NewMySQLClusterRepoWithGlobal()}
}

// NewMySQLClusterInfoWithDefault returns a new MySQLClusterInfo with default MySQLClusterRepo
func NewMySQLClusterInfoWithDefault(
	clusterName string,
	envID int) *MySQLClusterInfo {
	return &MySQLClusterInfo{
		MySQLClusterRepo:    NewMySQLClusterRepoWithGlobal(),
		ClusterName:         clusterName,
		MiddlewareClusterID: constant.DefaultRandomInt,
		MonitorSystemID:     constant.DefaultRandomInt,
		EnvID:               envID,
	}
}

// NewMySQLClusterInfoWithMapAndRandom returns a new *MySQLClusterInfo with given map
func NewMySQLClusterInfoWithMapAndRandom(fields map[string]interface{}) (*MySQLClusterInfo, error) {
	mci := &MySQLClusterInfo{}
	err := common.SetValuesWithMapAndRandom(mci, fields)
	if err != nil {
		return nil, err
	}
	return mci, nil
}

// Identity cluster returns ID of mysql cluster
func (mci *MySQLClusterInfo) Identity() int {
	return mci.ID
}

// GetClusterName returns the env name
func (mci *MySQLClusterInfo) GetClusterName() string {
	return mci.ClusterName
}

// GetMiddlewareClusterID returns the middleware cluster id
func (mci *MySQLClusterInfo) GetMiddlewareClusterID() int {
	return mci.MiddlewareClusterID
}

// GetMonitorSystemID returns the monitor system id
func (mci *MySQLClusterInfo) GetMonitorSystemID() int {
	return mci.MonitorSystemID
}

// GetEnvID returns the env id
func (mci *MySQLClusterInfo) GetEnvID() int {
	return mci.EnvID
}

// GetDelFlag returns the delete flag
func (mci *MySQLClusterInfo) GetDelFlag() int {
	return mci.DelFlag
}

// GetCreateTime returns created time of mysql cluster
func (mci *MySQLClusterInfo) GetCreateTime() time.Time {
	return mci.CreateTime
}

// GetLastUpdateTime returns last updated time of mysql cluster
func (mci *MySQLClusterInfo) GetLastUpdateTime() time.Time {
	return mci.LastUpdateTime
}

// GetMySQLServers gets the mysql servers of this cluster
func (mci *MySQLClusterInfo) GetMySQLServers() ([]metadata.MySQLServer, error) {
	mysqlServerService := NewMySQLServerServiceWithDefault()
	err := mysqlServerService.GetByClusterID(mci.Identity())

	return mysqlServerService.GetMySQLServers(), err
}

// GetMasterServers gets the master servers of this cluster
func (mci *MySQLClusterInfo) GetMasterServers() ([]metadata.MySQLServer, error) {
	var masterServers []metadata.MySQLServer

	mysqlServers, err := mci.GetMySQLServers()
	if err != nil {
		return nil, err
	}

	for _, mysqlServer := range mysqlServers {
		isMaster, err := mysqlServer.IsMaster()
		if err != nil {
			log.Errorf("failed to connect server %s:%s. %s", mysqlServer.GetHostIP(), mysqlServer.GetPortNum(), err.Error())
		}

		if isMaster {
			masterServers = append(masterServers, mysqlServer)
		}
	}

	if len(masterServers) == constant.ZeroInt {
		return nil, errors.Errorf("no master exists, something goes wrong. mysql cluster id: %d", mci.Identity())
	}

	return masterServers, nil
}

// GetDBs gets the databases of this cluster
func (mci *MySQLClusterInfo) GetDBs() ([]metadata.DB, error) {
	return mci.MySQLClusterRepo.GetDBsByID(mci.Identity())
}

// GetResourceGroupByID get the resource group of the given id from the middleware
func (mci *MySQLClusterInfo) GetResourceGroup() ([]metadata.ResourceGroup, error) {
	// todo: implement
	return nil, nil
}

// GetUsers gets the users of this cluster
func (mci *MySQLClusterInfo) GetUsers() ([]metadata.User, error) {
	return mci.MySQLClusterRepo.GetUsersByID(mci.Identity())
}

// AddUser add a user for mysql cluster in the middleware
func (mci *MySQLClusterInfo) AddUser(userID int) error {
	return mci.MySQLClusterRepo.AddUser(mci.Identity(), userID)
}

// DeleteUser delete a user for mysql cluster in the middleware
func (mci *MySQLClusterInfo) DeleteUser(userID int) error {
	return mci.MySQLClusterRepo.DeleteUser(mci.Identity(), userID)
}

// DeleteUser delete a user for mysql cluster in the middleware

// GetAppUsers gets the application users of this cluster
func (mci *MySQLClusterInfo) GetAppUsers() ([]metadata.User, error) {
	return mci.MySQLClusterRepo.GetAppUsersByID(mci.Identity())
}

// GetDBUsers gets the db users of this cluster
func (mci *MySQLClusterInfo) GetDBUsers() ([]metadata.User, error) {
	return mci.MySQLClusterRepo.GetDBUsersByID(mci.Identity())
}

// GetAllUsers gets both application and db users of this cluster
func (mci *MySQLClusterInfo) GetAllUsers() ([]metadata.User, error) {
	return mci.MySQLClusterRepo.GetAllUsersByID(mci.Identity())
}

// Set sets mysql cluster with given fields, key is the field name and value is the relevant value of the key
func (mci *MySQLClusterInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(mci, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to true, need to use Save to write to the middleware
func (mci *MySQLClusterInfo) Delete() {
	mci.DelFlag = 1
}

// MarshalJSON marshals mysql cluster to json string, it only marshals fields that has default tag
func (mci *MySQLClusterInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(mci, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only with specified fields of mysql cluster to json string
func (mci *MySQLClusterInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(mci, fields...)
}
