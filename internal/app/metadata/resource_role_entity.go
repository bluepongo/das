package metadata

import (
	"time"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
)

const (
	resourceRoleRoleUUIDStruct        = "RoleUUID"
	resourceRoleRoleNameStruct        = "RoleName"
	resourceRoleResourceGroupIDStruct = "ResourceGroupID"
	resourceRoleDelFlagStruct         = "DelFlag"
)

var _ metadata.ResourceRole = (*ResourceRoleInfo)(nil)

// ResourceRoleInfo is the entity interface
type ResourceRoleInfo struct {
	ResourceRoleRepo metadata.ResourceRoleRepo
	ID               int       `middleware:"id" json:"id"`
	RoleUUID         string    `middleware:"role_uuid" json:"role_uuid"`
	RoleName         string    `middleware:"role_name" json:"role_name"`
	ResourceGroupID  int       `middleware:"resource_group_id" json:"resource_group_id"`
	DelFlag          int       `middleware:"del_flag" json:"del_flag"`
	CreateTime       time.Time `middleware:"create_time" json:"create_time"`
	LastUpdateTime   time.Time `middleware:"last_update_time" json:"last_update_time"`
}

// NewResourceRoleInfo returns a new ResourceRoleInfo
func NewResourceRoleInfo(repo *ResourceRoleRepo,
	id int,
	roleUUID string,
	roleName string,
	resourceGroupID int,
	delFlag int,
	createTime, lastUpdateTime time.Time) *ResourceRoleInfo {
	return &ResourceRoleInfo{
		repo,
		id,
		roleUUID,
		roleName,
		resourceGroupID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewResourceRoleInfoWithGlobal returns a new ResourceRoleInfo with default ResourceRoleRepo
func NewResourceRoleInfoWithGlobal(
	id int,
	roleUUID string,
	roleName string,
	resourceGroupID int,
	delFlag int,
	createTime, lastUpdateTime time.Time) *ResourceRoleInfo {
	return &ResourceRoleInfo{
		NewResourceRoleRepoWithGlobal(),
		id,
		roleUUID,
		roleName,
		resourceGroupID,
		delFlag,
		createTime,
		lastUpdateTime,
	}
}

// NewEmptyResourceRoleInfoWithGlobal returns a new ResourceRoleInfo with default ResourceRoleRepo
func NewEmptyResourceRoleInfoWithGlobal() *ResourceRoleInfo {
	return &ResourceRoleInfo{ResourceRoleRepo: NewResourceRoleRepoWithGlobal()}
}

// NewResourceRoleInfoWithDefault returns a new ResourceRoleInfo with default ResourceRoleRepo
func NewResourceRoleInfoWithDefault(
	roleUUID string,
	resourceGroupID int) *ResourceRoleInfo {
	return &ResourceRoleInfo{
		ResourceRoleRepo: NewResourceRoleRepoWithGlobal(),
		RoleUUID:         roleUUID,
		ResourceGroupID:  resourceGroupID,
	}
}

// NewResourceRoleInfoWithMapAndRandom returns a new *ResourceRoleInfo with given map
func NewResourceRoleInfoWithMapAndRandom(fields map[string]interface{}) (*ResourceRoleInfo, error) {
	rri := &ResourceRoleInfo{}
	err := common.SetValuesWithMapAndRandom(rri, fields)
	if err != nil {
		return nil, err
	}
	return rri, nil
}

// Identity cluster returns ID of mysql cluster
func (rri *ResourceRoleInfo) Identity() int {
	return rri.ID
}

// GetRoleUUID returns the role uuid
func (rri *ResourceRoleInfo) GetRoleUUID() string {
	return rri.RoleUUID
}

// GetRoleName returns the role name
func (rri *ResourceRoleInfo) GetRoleName() string {
	return rri.RoleName
}

// GetResourceGroupID returns the resource group id
func (rri *ResourceRoleInfo) GetResourceGroupID() int {
	return rri.ResourceGroupID
}

// GetDelFlag returns the delete flag
func (rri *ResourceRoleInfo) GetDelFlag() int {
	return rri.DelFlag
}

// GetCreateTime returns created time
func (rri *ResourceRoleInfo) GetCreateTime() time.Time {
	return rri.CreateTime
}

// GetLastUpdateTime returns last updated time
func (rri *ResourceRoleInfo) GetLastUpdateTime() time.Time {
	return rri.LastUpdateTime
}

// GetResourceGroup gets the resource group which this role belongs to
func (rri *ResourceRoleInfo) GetResourceGroup() (metadata.ResourceGroup, error) {
	return rri.ResourceRoleRepo.GetResourceGroup(rri.Identity())
}

// GetUsers gets the users of this resource role
func (rri *ResourceRoleInfo) GetUsers() ([]metadata.User, error) {
	return rri.ResourceRoleRepo.GetUsersByID(rri.Identity())
}

// Set sets the resource group with given fields, key is the field name and value is the relevant value of the key
func (rri *ResourceRoleInfo) Set(fields map[string]interface{}) error {
	for fieldName, fieldValue := range fields {
		err := common.SetValueOfStruct(rri, fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete sets DelFlag to 1
func (rri *ResourceRoleInfo) Delete() {
	rri.DelFlag = 1
}

// AddUser adds a map of the resource role and user
func (rri *ResourceRoleInfo) AddUser(userID int) error {
	return rri.ResourceRoleRepo.AddUser(rri.Identity(), userID)
}

// DeleteUser deletes the map of the resource role and user
func (rri *ResourceRoleInfo) DeleteUser(userID int) error {
	return rri.ResourceRoleRepo.DeleteUser(rri.Identity(), userID)
}

// MarshalJSON marshals ResourceRole to json string
func (rri *ResourceRoleInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(rri, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only specified field of the ResourceRole to json string
func (rri *ResourceRoleInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(rri, fields...)
}
