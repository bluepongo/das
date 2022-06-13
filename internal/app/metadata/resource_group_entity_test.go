package metadata

import (
	"reflect"
	"testing"

	"github.com/jinzhu/now"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testResourceGroupID                   = 1
	testResourceGroupGroupUUID            = "5c6c6d73-eac2-11ec-8183-001c42d502a6"
	testResourceGroupGroupName            = "resource-group-01"
	testResourceGroupDelFlag              = 1
	testResourceGroupCreateTimeString     = "2021-01-21 10:00:00.000000"
	testResourceGroupLastUpdateTimeString = "2021-01-21 10:00:00.000000"
)

var testResourceGroupInfo *ResourceGroupInfo

func init() {
	testInitDASMySQLPool()
	testResourceGroupInfo = testInitNewResourceGroupInfo()
}

func testInitNewResourceGroupInfo() *ResourceGroupInfo {
	now.TimeFormats = append(now.TimeFormats, constant.DefaultTimeLayout)

	createTime, _ := now.Parse(testResourceGroupCreateTimeString)
	lastUpdateTime, _ := now.Parse(testResourceGroupLastUpdateTimeString)
	resourceGroupInfo := NewResourceGroupInfoWithGlobal(
		testResourceGroupID,
		testResourceGroupGroupUUID,
		testResourceGroupGroupName,
		testResourceGroupDelFlag,
		createTime,
		lastUpdateTime,
	)

	return resourceGroupInfo
}

func TestResourceGroupEntity_All(t *testing.T) {
	TestResourceGroupInfo_Get(t)
	TestResourceGroupInfo_GetResourceRoles(t)
	TestResourceGroupInfo_GetMySQLClusters(t)
	TestResourceGroupInfo_GetMySQLServers(t)
	TestResourceGroupInfo_GetMiddlewareClusters(t)
	TestResourceGroupInfo_GetMiddlewareServers(t)
	TestResourceGroupInfo_GetUsers(t)
	TestResourceGroupInfo_GetDASAdminUsers(t)
	TestResourceGroupInfo_Set(t)
	TestResourceGroupInfo_Delete(t)
	TestResourceGroupInfo_AddMySQLCluster(t)
	TestResourceGroupInfo_DeleteMySQLCluster(t)
	TestResourceGroupInfo_AddMiddlewareCluster(t)
	TestResourceGroupInfo_DeleteMiddlewareCluster(t)
	TestResourceGroupInfo_MarshalJSON(t)
	TestResourceGroupInfo_MarshalJSONWithFields(t)
}

func TestResourceGroupInfo_Get(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResourceGroupGroupUUID, testResourceGroupInfo.GetGroupUUID(), "test GetGroupUUID() failed")
	asst.Equal(testResourceGroupGroupName, testResourceGroupInfo.GetGroupName(), "test GetGroupName() failed")
	asst.Equal(testResourceGroupDelFlag, testResourceGroupInfo.GetDelFlag(), "test DelFlag() failed")
	asst.True(reflect.DeepEqual(testResourceGroupInfo.CreateTime, testResourceGroupInfo.GetCreateTime()), "test GetCreateTime() failed")
	asst.True(reflect.DeepEqual(testResourceGroupInfo.LastUpdateTime, testResourceGroupInfo.GetLastUpdateTime()), "test GetLastUpdateTime() failed")
}

func TestResourceGroupInfo_GetResourceRoles(t *testing.T) {
	asst := assert.New(t)

	resourceRoles, err := testResourceGroupInfo.GetResourceRoles()
	asst.Nil(err, common.CombineMessageWithError("test GetResourceRoles() failed", err))
	asst.Equal(2, len(resourceRoles), "test GetResourceRoles() failed", err)
}

func TestResourceGroupInfo_GetMySQLClusters(t *testing.T) {
	asst := assert.New(t)

	mysqlClusters, err := testResourceGroupInfo.GetMySQLClusters()
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLClusters() failed", err))
	asst.Equal(2, len(mysqlClusters), "test GetMySQLClusters() failed")
}

func TestResourceGroupInfo_GetMySQLServers(t *testing.T) {
	asst := assert.New(t)

	mysqlServers, err := testResourceGroupInfo.GetMySQLServers()
	asst.Nil(err, common.CombineMessageWithError("test GetMySQLServers() failed", err))
	asst.Equal(2, len(mysqlServers), "test GetMySQLServers() failed")
}

func TestResourceGroupInfo_GetMiddlewareClusters(t *testing.T) {
	asst := assert.New(t)

	middlewareClusters, err := testResourceGroupInfo.GetMiddlewareClusters()
	asst.Nil(err, common.CombineMessageWithError("test GetMiddlewareClusters() failed", err))
	asst.Equal(1, len(middlewareClusters), "test GetMiddlewareClusters() failed")
}

func TestResourceGroupInfo_GetMiddlewareServers(t *testing.T) {
	asst := assert.New(t)

	middlewareServers, err := testResourceGroupInfo.GetMiddlewareServers()
	asst.Nil(err, common.CombineMessageWithError("test GetMiddlewareServers() failed", err))
	asst.Equal(1, len(middlewareServers), "test GetMiddlewareServers() failed")
}

func TestResourceGroupInfo_GetUsers(t *testing.T) {
	asst := assert.New(t)

	users, err := testResourceGroupInfo.GetUsers()
	asst.Nil(err, common.CombineMessageWithError("test GetUsers() failed", err))
	asst.Equal(1, len(users), "test GetUsers() failed", err)
}

func TestResourceGroupInfo_GetDASAdminUsers(t *testing.T) {
	asst := assert.New(t)

	users, err := testResourceGroupInfo.GetDASAdminUsers()
	asst.Nil(err, common.CombineMessageWithError("test GetDASAdminUsers() failed", err))
	asst.Equal(1, len(users), "test GetDASAdminUsers() failed", err)
}

func TestResourceGroupInfo_Set(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupInfo.Set(map[string]interface{}{resourceGroupGroupNameStruct: testResourceGroupUpdateGroupName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testResourceGroupUpdateGroupName, testResourceGroupInfo.GetGroupName(), "test Set() failed")
	err = testResourceGroupInfo.Set(map[string]interface{}{resourceGroupGroupNameStruct: testResourceGroupGroupName})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testResourceGroupGroupName, testResourceGroupInfo.GetGroupName(), "test Set() failed")
}

func TestResourceGroupInfo_Delete(t *testing.T) {
	asst := assert.New(t)

	testResourceGroupInfo.Delete()
	asst.Equal(1, testResourceGroupInfo.GetDelFlag(), "test Delete() failed")
	err := testResourceGroupInfo.Set(map[string]interface{}{resourceGroupDelFlagStruct: constant.ZeroInt})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(constant.ZeroInt, testResourceGroupInfo.GetDelFlag(), "test Set() failed")
}

func TestResourceGroupInfo_AddMySQLCluster(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupInfo.DeleteMySQLCluster(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLCluster() failed", err))
	err = testResourceGroupInfo.AddMySQLCluster(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLCluster() failed", err))
	mysqlClusters, err := testResourceGroupInfo.GetMySQLClusters()
	asst.Nil(err, common.CombineMessageWithError("test AddMySQLCluster() failed", err))
	asst.Equal(2, len(mysqlClusters), "test AddMySQLCluster() failed")

}

func TestResourceGroupInfo_DeleteMySQLCluster(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupInfo.DeleteMySQLCluster(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteMySQLCluster() failed", err))
	err = testResourceGroupInfo.AddMySQLCluster(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteMySQLCluster() failed", err))
	mysqlClusters, err := testResourceGroupInfo.GetMySQLClusters()
	asst.Nil(err, common.CombineMessageWithError("test DeleteMySQLCluster() failed", err))
	asst.Equal(2, len(mysqlClusters), "test DeleteMySQLCluster() failed")
}

func TestResourceGroupInfo_AddMiddlewareCluster(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupInfo.DeleteMiddlewareCluster(testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMiddlewareCluster() failed", err))
	err = testResourceGroupInfo.AddMiddlewareCluster(testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test AddMiddlewareCluster() failed", err))
	middlewareClusters, err := testResourceGroupInfo.GetMiddlewareClusters()
	asst.Nil(err, common.CombineMessageWithError("test AddMiddlewareCluster() failed", err))
	asst.Equal(1, len(middlewareClusters), "test AddMiddlewareCluster() failed")

}

func TestResourceGroupInfo_DeleteMiddlewareCluster(t *testing.T) {
	asst := assert.New(t)

	err := testResourceGroupInfo.DeleteMiddlewareCluster(testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteMiddlewareCluster() failed", err))
	err = testResourceGroupInfo.AddMiddlewareCluster(testMiddlewareClusterClusterID)
	asst.Nil(err, common.CombineMessageWithError("test DeleteMiddlewareCluster() failed", err))
	middlewareClusters, err := testResourceGroupInfo.GetMiddlewareClusters()
	asst.Nil(err, common.CombineMessageWithError("test DeleteMiddlewareCluster() failed", err))
	asst.Equal(1, len(middlewareClusters), "test DeleteMiddlewareCluster() failed")
}

func TestResourceGroupInfo_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testResourceGroupInfo.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	t.Log(string(jsonBytes))
}

func TestResourceGroupInfo_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testResourceGroupInfo.MarshalJSONWithFields(resourceGroupGroupNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	t.Log(string(jsonBytes))
}
