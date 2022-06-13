package metadata

import (
	"time"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
)

const (
	resourceGroupGroupUUIDStruct = "GroupUUID"
	resourceGroupGroupNameStruct = "GroupName"
)

var _ metadata.ResourceGroup = (*ResourceGroupInfo)(nil)

type ResourceGroupInfo struct {
	ResourceGroupRepo metadata.ResourceGroupRepo
	ID                int       `middleware:"id" json:"id"`
	GroupUUID         string    `middleware:"group_uuid" json:"group_uuid"`
	GroupName         string    `middleware:"group_name" json:"group_name"`
	DelFlag           int       `middleware:"del_flag" json:"del_flag"`
	CreateTime        time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime    time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewResourceGroupInfo returns a new ResourceGroupInfo
func NewResourceGroupInfo(
	repo *ResourceGroupRepo,
	id int,
	groupUUID string,
	groupName string,
	delFlag int,
	createTime, lastUpdateTime time.Time) *ResourceGroupInfo {
	return &ResourceGroupInfo{
		repo,
		id,
		groupUUID,
		groupName,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewResourceGroupInfoWithGlobal returns a new ResourceGroupInfo with default ResourceGroupRepo
func NewResourceGroupInfoWithGlobal(
	id int,
	groupUUID string,
	groupName string,
	delFlag int,
	createTime, lastUpdateTime time.Time) *ResourceGroupInfo {
	return &ResourceGroupInfo{
		NewResourceGroupRepoWithGlobal(),
		id,
		groupUUID,
		groupName,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewEmptyResourceGroupInfoWithGlobal returns a new ResourceGroupInfo with default ResourceGroupRepo
func NewEmptyResourceGroupInfoWithGlobal() *ResourceGroupInfo {
	return &ResourceGroupInfo{ResourceGroupRepo: NewResourceGroupRepoWithGlobal()}
}

// NewResourceGroupInfoWithDefault returns a new *ResourceGroupInfo with default ResourceGroupRepo
func NewResourceGroupInfoWithDefault(groupUUID string, groupName string) *ResourceGroupInfo {
	return &ResourceGroupInfo{
		ResourceGroupRepo: NewResourceGroupRepoWithGlobal(),
		GroupUUID:         groupUUID,
		GroupName:         groupName,
	}
}

// NewGroupResourceInfoWithMapAndRandom returns a new *ResourceGroupInfo with given map
func NewGroupResourceInfoWithMapAndRandom(fields map[string]interface{}) (*ResourceGroupInfo, error) {
	rgi := &ResourceGroupInfo{}
	err := common.SetValuesWithMapAndRandom(rgi, fields)
	if err != nil {
		return nil, err
	}
	return rgi, nil
}

// GetResourceGroupRepo returns the ResourceGroupRepo
func (rgi *ResourceGroupInfo) GetResourceGroupRepo() metadata.ResourceGroupRepo {
	return rgi.ResourceGroupRepo
}

// GetID returns the ID
func (rgi *ResourceGroupInfo) Identity() int {
	return rgi.ID
}

// GetGroupUUID returns the GroupUUID
func (rgi *ResourceGroupInfo) GetGroupUUID() string {
	return rgi.GroupUUID
}

// GetGroupName returns the GroupName
func (rgi *ResourceGroupInfo) GetGroupName() string {
	return rgi.GroupName
}

// GetDelFlag returns the DelFlag
func (rgi *ResourceGroupInfo) GetDelFlag() int {
	return rgi.DelFlag
}

// GetCreateTime returns the CreateTime
func (rgi *ResourceGroupInfo) GetCreateTime() time.Time {
	return rgi.CreateTime
}

// GetLastUpdateTime returns the LastUpdateTime
func (rgi *ResourceGroupInfo) GetLastUpdateTime() time.Time {
	return rgi.LastUpdateTime
}

// GetResourceRoles get all resource roles of this resource group
func (rgi *ResourceGroupInfo) GetResourceRoles() ([]metadata.ResourceRole, error) {
	return rgi.GetResourceGroupRepo().GetResourceRolesByID(rgi.Identity())
}

// GetMySQLClusters gets the mysql clusters of this resource group
func (rgi *ResourceGroupInfo) GetMySQLClusters() ([]metadata.MySQLCluster, error) {
	return rgi.GetResourceGroupRepo().GetMySQLClustersByID(rgi.Identity())
}

// GetMySQLServers gets the mysql servers of this resource group
func (rgi *ResourceGroupInfo) GetMySQLServers() ([]metadata.MySQLServer, error) {
	return rgi.GetResourceGroupRepo().GetMySQLServersByID(rgi.Identity())
}

// GetMiddlewareClusters gets the mysql clusters of this resource group
func (rgi *ResourceGroupInfo) GetMiddlewareClusters() ([]metadata.MiddlewareCluster, error) {
	return rgi.GetResourceGroupRepo().GetMiddlewareClustersByID(rgi.Identity())
}

// GetMiddlewareServers gets the mysql servers of this resource group
func (rgi *ResourceGroupInfo) GetMiddlewareServers() ([]metadata.MiddlewareServer, error) {
	return rgi.GetResourceGroupRepo().GetMiddlewareServersByID(rgi.Identity())
}

// GetUsers gets the users of this resource group
func (rgi *ResourceGroupInfo) GetUsers() ([]metadata.User, error) {
	return rgi.GetResourceGroupRepo().GetUsersByID(rgi.Identity())
}

// GetDASAdminUsers gets the das admin users of this resource group
func (rgi *ResourceGroupInfo) GetDASAdminUsers() ([]metadata.User, error) {
	return rgi.GetResourceGroupRepo().GetDASAdminUsersByID(rgi.Identity())
}

// Set sets the resource group with given fields, key is the field name and value is the relevant value of the key
func (rgi *ResourceGroupInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(rgi, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to 1
func (rgi *ResourceGroupInfo) Delete() {
	rgi.DelFlag = 1
}

// AddMySQLCluster adds mysql cluster to the resource group
func (rgi *ResourceGroupInfo) AddMySQLCluster(mysqlClusterID int) error {
	return rgi.GetResourceGroupRepo().AddMySQLCluster(rgi.Identity(), mysqlClusterID)
}

// DeleteMySQLCluster deletes mysql cluster from the resource group
func (rgi *ResourceGroupInfo) DeleteMySQLCluster(mysqlClusterID int) error {
	return rgi.GetResourceGroupRepo().DeleteMySQLCluster(rgi.Identity(), mysqlClusterID)
}

// AddMiddlewareCluster adds middleware cluster to the resource group
func (rgi *ResourceGroupInfo) AddMiddlewareCluster(mysqlClusterID int) error {
	return rgi.GetResourceGroupRepo().AddMiddlewareCluster(rgi.Identity(), mysqlClusterID)
}

// DeleteMiddlewareCluster deletes middleware cluster from the resource group
func (rgi *ResourceGroupInfo) DeleteMiddlewareCluster(middlewareClusterID int) error {
	return rgi.GetResourceGroupRepo().DeleteMiddlewareCluster(rgi.Identity(), middlewareClusterID)
}

// MarshalJSON marshals ResourceGroup to json string
func (rgi *ResourceGroupInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(rgi, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only specified field of the ResourceGroup to json string
func (rgi *ResourceGroupInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(rgi, fields...)
}
