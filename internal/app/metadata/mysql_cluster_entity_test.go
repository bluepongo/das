package metadata

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/jinzhu/now"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	defaultMySQLClusterInfoID                   = 1
	defaultMySQLClusterInfoClusterName          = "test"
	defaultMySQLClusterInfoMiddlewareClusterID  = 1
	defaultMySQLClusterInfoMonitorSystemID      = 1
	defaultMySQLClusterInfoOwnerID              = 1
	defaultMySQLClusterInfoEnvID                = 1
	defaultMySQLClusterInfoDelFlag              = 0
	defaultMySQLClusterInfoCreateTimeString     = "2021-01-21 10:00:00.000000"
	defaultMySQLClusterInfoLastUpdateTimeString = "2021-01-21 13:00:00.000000"
	clusterNameJSON                             = "cluster_name"
)

func initNewMySQLClusterInfo() *MySQLClusterInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(defaultMySQLClusterInfoCreateTimeString)
	lastUpdateTime, _ := now.Parse(defaultMySQLClusterInfoLastUpdateTimeString)
	return NewMySQLClusterInfoWithGlobal(
		defaultMySQLClusterInfoID,
		defaultMySQLClusterInfoClusterName,
		defaultMySQLClusterInfoMiddlewareClusterID,
		defaultMySQLClusterInfoMonitorSystemID,
		defaultMySQLClusterInfoOwnerID,
		defaultMySQLClusterInfoEnvID,
		defaultMySQLClusterInfoDelFlag,
		createTime,
		lastUpdateTime)
}

func equalMySQLClusterInfo(a, b *MySQLClusterInfo) bool {
	return a.ID == b.ID &&
		a.ClusterName == b.ClusterName &&
		a.MiddlewareClusterID == b.MiddlewareClusterID &&
		a.MonitorSystemID == b.MonitorSystemID &&
		a.OwnerID == b.OwnerID &&
		a.EnvID == b.EnvID &&
		a.DelFlag == b.DelFlag &&
		a.CreateTime == b.CreateTime &&
		a.LastUpdateTime == b.LastUpdateTime
}

func TestMySQLClusterEntityAll(t *testing.T) {
	TestMySQLClusterInfo_Identity(t)
	TestMySQLClusterInfo_Get(t)
	TestMySQLClusterInfo_GetMySQLServers(t)
	TestMySQLClusterInfo_GetMasterServers(t)
	TestMySQLClusterInfo_Set(t)
	TestMySQLClusterInfo_Delete(t)
	TestMySQLClusterInfo_MarshalJSON(t)
	TestMySQLClusterInfo_MarshalJSONWithFields(t)
}

func TestMySQLClusterInfo_Identity(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	asst.Equal(defaultMySQLClusterInfoID, mysqlClusterInfo.Identity(), "test Identity() failed")
}

func TestMySQLClusterInfo_Get(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	clusterName := mysqlClusterInfo.GetClusterName()
	asst.Equal(mysqlClusterInfo.ClusterName, clusterName, "test GetClusterName() failed")

	middlewareClusterID := mysqlClusterInfo.GetMiddlewareClusterID()
	asst.Equal(mysqlClusterInfo.MiddlewareClusterID, middlewareClusterID, "test GetMiddlewareClusterID() failed")

	monitorSystemID := mysqlClusterInfo.GetMonitorSystemID()
	asst.Equal(mysqlClusterInfo.MonitorSystemID, monitorSystemID, "test GetMonitorSystemID() failed")

	ownerID := mysqlClusterInfo.GetOwnerID()
	asst.Equal(mysqlClusterInfo.OwnerID, ownerID, "test GetOwnerID() failed")

	envID := mysqlClusterInfo.GetEnvID()
	asst.Equal(mysqlClusterInfo.EnvID, envID, "test GetEnvID() failed")

	delFlag := mysqlClusterInfo.GetDelFlag()
	asst.Equal(mysqlClusterInfo.DelFlag, delFlag, "test DelFlag() failed")

	createTime := mysqlClusterInfo.GetCreateTime()
	asst.True(reflect.DeepEqual(mysqlClusterInfo.CreateTime, createTime), "test GetCreateTime() failed")

	lastUpdateTime := mysqlClusterInfo.GetLastUpdateTime()
	asst.True(reflect.DeepEqual(mysqlClusterInfo.LastUpdateTime, lastUpdateTime), "test GetLastUpdateTime() failed")
}

func TestMySQLClusterInfo_GetMySQLServers(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	mysqlServers, err := mysqlClusterInfo.GetMySQLServers()
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLServersByID() failed", err))
	asst.Equal(3, len(mysqlServers), "test GetMySQLServersByID() failed")
}

func TestMySQLClusterInfo_GetMasterServers(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	masterServers, err := mysqlClusterInfo.GetMasterServers()
	asst.Nil(err, common.CombineMessageWithError("test GetMasterServers() failed", err))
	asst.Equal(1, masterServers[constant.ZeroInt].Identity(), "test GetMasterServers() failed", err)
}

func TestMySQLClusterInfo_GetDBs(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	dbs, err := mysqlClusterInfo.GetDBs()
	asst.Nil(err, common.CombineMessageWithError("test GetDBs() failed", err))
	asst.Equal(1, dbs[constant.ZeroInt].Identity(), "test GetDBs() failed", err)
}

func TestMySQLClusterInfo_GetAppOwners(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	masterServers, err := mysqlClusterInfo.GetAppOwners()
	asst.Nil(err, common.CombineMessageWithError("test GetAppOwners() failed", err))
	asst.Equal(1, masterServers[constant.ZeroInt].Identity(), "test GetAppOwners() failed", err)
}

func TestMySQLClusterInfo_GetDBOwners(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	masterServers, err := mysqlClusterInfo.GetDBOwners()
	asst.Nil(err, common.CombineMessageWithError("test GetDBOwners() failed", err))
	asst.Equal(1, masterServers[constant.ZeroInt].Identity(), "test GetDBOwners() failed", err)
}

func TestMySQLClusterInfo_GetAllOwners(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	masterServers, err := mysqlClusterInfo.GetAllOwners()
	asst.Nil(err, common.CombineMessageWithError("test GetAllOwners() failed", err))
	asst.Equal(1, masterServers[constant.ZeroInt].Identity(), "test GetAllOwners() failed", err)
}

func TestMySQLClusterInfo_Set(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()

	newClusterName := defaultMySQLClusterInfoClusterName
	err := mysqlClusterInfo.Set(map[string]interface{}{"ClusterName": newClusterName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newClusterName, mysqlClusterInfo.ClusterName, "test Set() failed")

	newMiddlewareClusterID := defaultMySQLClusterInfoMiddlewareClusterID
	err = mysqlClusterInfo.Set(map[string]interface{}{"MiddlewareClusterID": newMiddlewareClusterID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newMiddlewareClusterID, mysqlClusterInfo.MiddlewareClusterID, "test Set() failed")

	newMonitorSystemID := defaultMySQLClusterInfoMonitorSystemID
	err = mysqlClusterInfo.Set(map[string]interface{}{"MonitorSystemID": newMonitorSystemID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newMonitorSystemID, mysqlClusterInfo.MonitorSystemID, "test Set() failed")

	newOwnerID := defaultMySQLClusterInfoOwnerID
	err = mysqlClusterInfo.Set(map[string]interface{}{"OwnerID": newOwnerID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newOwnerID, mysqlClusterInfo.OwnerID, "test Set() failed")

	newEnvID := defaultMySQLClusterInfoEnvID
	err = mysqlClusterInfo.Set(map[string]interface{}{"EnvID": newEnvID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(newEnvID, mysqlClusterInfo.EnvID, "test Set() failed")
}

func TestMySQLClusterInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	mysqlClusterInfo.Delete()
	asst.Equal(1, mysqlClusterInfo.GetDelFlag(), "test Delete() failed")
}

func TestMySQLClusterInfo_MarshalJSON(t *testing.T) {
	var mysqlClusterInfoUnmarshal *MySQLClusterInfo

	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	data, err := mysqlClusterInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = json.Unmarshal(data, &mysqlClusterInfoUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	asst.True(equalMySQLClusterInfo(mysqlClusterInfo, mysqlClusterInfoUnmarshal), "test MarshalJSON() failed")
}

func TestMySQLClusterInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	mysqlClusterInfo := initNewMySQLClusterInfo()
	data, err := mysqlClusterInfo.MarshalJSONWithFields(clusterNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	expect, err := json.Marshal(map[string]interface{}{clusterNameJSON: "test"})
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	asst.Equal(string(expect), string(data), "test MarshalJSONWithFields() failed")
}

