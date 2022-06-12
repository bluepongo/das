package metadata

import (
	"fmt"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
)

const resourceGroupResourceGroupsStruct = "MySQLClusters"

var _ metadata.ResourceGroupService = (*ResourceGroupService)(nil)

type ResourceGroupService struct {
	ResourceGroupRepo  metadata.ResourceGroupRepo
	ResourceGroups     []metadata.ResourceGroup     `json:"resource_groups"`
	MySQLClusters      []metadata.MySQLCluster      `json:"mysql_clusters"`
	MySQLServers       []metadata.MySQLServer       `json:"mysql_servers"`
	MiddlewareClusters []metadata.MiddlewareCluster `json:"middleware_clusters"`
	MiddlewareServers  []metadata.MiddlewareServer  `json:"middleware_servers"`
	ResourceRoles      []metadata.ResourceRole      `json:"resource_roles"`
	Users              []metadata.User              `json:"users"`
}

// NewResourceGroupService returns a new *ResourceGroupService
func NewResourceGroupService(repo metadata.ResourceGroupRepo) *ResourceGroupService {
	return &ResourceGroupService{
		repo,
		[]metadata.ResourceGroup{},
		[]metadata.MySQLCluster{},
		[]metadata.MySQLServer{},
		[]metadata.MiddlewareCluster{},
		[]metadata.MiddlewareServer{},
		[]metadata.ResourceRole{},
		[]metadata.User{},
	}
}

// NewResourceGroupServiceWithDefault returns a new *ResourceGroupService with default repository
func NewResourceGroupServiceWithDefault() *ResourceGroupService {
	return NewResourceGroupService(NewResourceGroupRepoWithGlobal())
}

// GetResourceGroups returns resource groups of the service
func (rgs *ResourceGroupService) GetResourceGroups() []metadata.ResourceGroup {
	return rgs.ResourceGroups
}

// GetMySQLCluster returns the mysql clusters of the service
func (rgs *ResourceGroupService) GetMySQLClusters() []metadata.MySQLCluster {
	return rgs.MySQLClusters
}

// GetMySQLServers returns the mysql servers of the service
func (rgs *ResourceGroupService) GetMySQLServers() []metadata.MySQLServer {
	return rgs.MySQLServers
}

// GetMiddlewareClusters returns the middleware clusters of the service
func (rgs *ResourceGroupService) GetMiddlewareClusters() []metadata.MiddlewareCluster {
	return rgs.MiddlewareClusters
}

// GetMiddlewareServers returns the middleware servers of the service
func (rgs *ResourceGroupService) GetMiddlewareServers() []metadata.MiddlewareServer {
	return rgs.MiddlewareServers
}

// GetResourceRoles returns the resource roles of the service
func (rgs *ResourceGroupService) GetResourceRoles() []metadata.ResourceRole {
	return rgs.ResourceRoles
}

// GetUsers returns the users of the service
func (rgs *ResourceGroupService) GetUsers() []metadata.User {
	return rgs.Users
}

// GetAll gets all mysql servers from the mysql
func (rgs *ResourceGroupService) GetAll() error {
	var err error

	rgs.ResourceGroups, err = rgs.ResourceGroupRepo.GetAll()

	return err
}

// GetByID gets the resource group of the given id
func (rgs *ResourceGroupService) GetByID(id int) error {
	resourceGroup, err := rgs.ResourceGroupRepo.GetByID(id)
	if err != nil {
		return err
	}

	rgs.ResourceGroups = nil
	rgs.ResourceGroups = append(rgs.ResourceGroups, resourceGroup)

	return nil
}

// GetByGroupUUID gets the resource group by group uuid
func (rgs *ResourceGroupService) GetByGroupUUID(groupUUID string) error {
	resourceGroup, err := rgs.ResourceGroupRepo.GetByGroupUUID(groupUUID)
	if err != nil {
		return err
	}

	rgs.ResourceGroups = nil
	rgs.ResourceGroups = append(rgs.ResourceGroups, resourceGroup)

	return nil
}

// GetResourceRolesByID get all resource roles with given resource group id
func (rgs *ResourceGroupService) GetResourceRolesByID(id int) error {
	var err error

	rgs.ResourceRoles, err = rgs.ResourceGroupRepo.GetResourceRoles(id)

	return err
}

// GetMySQLClustersByID gets the mysql clusters with given resource group id
func (rgs *ResourceGroupService) GetMySQLClustersByID(id int) error {
	var err error

	rgs.MySQLClusters, err = rgs.ResourceGroupRepo.GetMySQLClustersByID(id)

	return err
}

// GetMySQLServersByID gets the mysql servers with given resource group id
func (rgs *ResourceGroupService) GetMySQLServersByID(id int) error {
	var err error

	rgs.MySQLServers, err = rgs.ResourceGroupRepo.GetMySQLServersByID(id)

	return err
}

// GetMiddlewareClustersByID gets the middleware clusters with given resource group id
func (rgs *ResourceGroupService) GetMiddlewareClustersByID(id int) error {
	var err error

	rgs.MiddlewareClusters, err = rgs.ResourceGroupRepo.GetMiddlewareClustersByID(id)

	return err
}

// GetMiddlewareServersByID gets the middleware servers with given resource group id
func (rgs *ResourceGroupService) GetMiddlewareServersByID(id int) error {
	var err error

	rgs.MiddlewareServers, err = rgs.ResourceGroupRepo.GetMiddlewareServersByID(id)

	return err
}

// GetUsersByID gets the users with given resource group id
func (rgs *ResourceGroupService) GetUsersByID(id int) error {
	var err error

	rgs.Users, err = rgs.ResourceGroupRepo.GetUsersByID(id)

	return err
}

// GetDASAdminUsersByID gets the das admin users with given resource group id
func (rgs *ResourceGroupService) GetDASAdminUsersByID(id int) error {
	var err error

	rgs.Users, err = rgs.ResourceGroupRepo.GetDASAdminUsersByID(id)

	return err
}

// GetResourceRolesByGroupUUID get all resource roles with given resource group uuid
func (rgs *ResourceGroupService) GetResourceRolesByGroupUUID(groupUUID string) error {
	resourceGroup, err := rgs.ResourceGroupRepo.GetByGroupUUID(groupUUID)
	if err != nil {
		return err
	}

	return rgs.GetResourceRolesByID(resourceGroup.Identity())
}

// GetMySQLClustersByGroupUUID gets the mysql clusters with given resource group uuid
func (rgs *ResourceGroupService) GetMySQLClustersByGroupUUID(groupUUID string) error {
	resourceGroup, err := rgs.ResourceGroupRepo.GetByGroupUUID(groupUUID)
	if err != nil {
		return err
	}

	return rgs.GetMySQLClustersByID(resourceGroup.Identity())
}

// GetMySQLServersByGroupUUID gets the mysql servers with given resource group uuid
func (rgs *ResourceGroupService) GetMySQLServersByGroupUUID(groupUUID string) error {
	resourceGroup, err := rgs.ResourceGroupRepo.GetByGroupUUID(groupUUID)
	if err != nil {
		return err
	}

	return rgs.GetMySQLServersByID(resourceGroup.Identity())
}

// GetMiddlewareClustersByGroupUUID gets the middleware clusters with given resource group uuid
func (rgs *ResourceGroupService) GetMiddlewareClustersByGroupUUID(groupUUID string) error {
	resourceGroup, err := rgs.ResourceGroupRepo.GetByGroupUUID(groupUUID)
	if err != nil {
		return err
	}

	return rgs.GetMiddlewareClustersByID(resourceGroup.Identity())
}

// GetMiddlewareServersByGroupUUID gets the middleware servers with given resource group uuid
func (rgs *ResourceGroupService) GetMiddlewareServersByGroupUUID(groupUUID string) error {
	resourceGroup, err := rgs.ResourceGroupRepo.GetByGroupUUID(groupUUID)
	if err != nil {
		return err
	}

	return rgs.GetMiddlewareServersByID(resourceGroup.Identity())
}

// GetUsersByGroupUUID gets the users with given resource group uuid
func (rgs *ResourceGroupService) GetUsersByGroupUUID(groupUUID string) error {
	resourceGroup, err := rgs.ResourceGroupRepo.GetByGroupUUID(groupUUID)
	if err != nil {
		return err
	}

	return rgs.GetUsersByID(resourceGroup.Identity())
}

// GetDASAdminUsersByUUID gets the das admin users with given resource group uuid
func (rgs *ResourceGroupService) GetDASAdminUsersByUUID(groupUUID string) error {
	resourceGroup, err := rgs.ResourceGroupRepo.GetByGroupUUID(groupUUID)
	if err != nil {
		return err
	}

	return rgs.GetDASAdminUsersByID(resourceGroup.Identity())
}

// Create creates a mysql server in the mysql
func (rgs *ResourceGroupService) Create(fields map[string]interface{}) error {
	// generate new map
	_, groupUUIDExists := fields[resourceGroupGroupUUIDStruct]
	_, groupNameExists := fields[resourceGroupGroupNameStruct]

	if !groupUUIDExists || !groupNameExists {
		return message.NewMessage(
			message.ErrFieldNotExists,
			fmt.Sprintf(
				"%s and %s",
				resourceGroupGroupUUIDStruct,
				resourceGroupGroupNameStruct))
	}

	// create a new entity
	resourceGroupInfo, err := NewGroupResourceInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	resourceGroup, err := rgs.ResourceGroupRepo.Create(resourceGroupInfo)
	if err != nil {
		return err
	}

	rgs.ResourceGroups = nil
	rgs.ResourceGroups = append(rgs.ResourceGroups, resourceGroup)

	return nil
}

// Update gets a mysql server of the given id from the mysql,
// and then updates its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the mysql
func (rgs *ResourceGroupService) Update(id int, fields map[string]interface{}) error {
	err := rgs.GetByID(id)
	if err != nil {
		return err
	}
	err = rgs.ResourceGroups[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return rgs.ResourceGroupRepo.Update(rgs.ResourceGroups[constant.ZeroInt])
}

// Delete deletes the mysql server of given id
func (rgs *ResourceGroupService) Delete(id int) error {
	err := rgs.GetByID(id)
	if err != nil {
		return err
	}

	return rgs.ResourceGroupRepo.Delete(id)
}

// AddMySQLCluster adds mysql cluster to the resource group
func (rgs *ResourceGroupService) AddMySQLCluster(resourceGroupID int, mysqlClusterID int) error {
	err := rgs.ResourceGroupRepo.AddMySQLCluster(resourceGroupID, mysqlClusterID)
	if err != nil {
		return err
	}

	return rgs.GetMySQLClustersByID(resourceGroupID)
}

// DeleteMySQLCluster deletes mysql cluster from the resource group
func (rgs *ResourceGroupService) DeleteMySQLCluster(resourceGroupID int, mysqlClusterID int) error {
	err := rgs.ResourceGroupRepo.DeleteMySQLCluster(resourceGroupID, mysqlClusterID)
	if err != nil {
		return err
	}

	return rgs.GetMySQLClustersByID(resourceGroupID)
}

// AddMiddlewareCluster adds middleware cluster to the resource group
func (rgs *ResourceGroupService) AddMiddlewareCluster(resourceGroupID int, middlewareClusterID int) error {
	err := rgs.ResourceGroupRepo.AddMiddlewareCluster(resourceGroupID, middlewareClusterID)
	if err != nil {
		return err
	}

	return rgs.GetMiddlewareClustersByID(resourceGroupID)
}

// DeleteMiddlewareCluster deletes middleware cluster from the resource group
func (rgs *ResourceGroupService) DeleteMiddlewareCluster(resourceGroupID int, middlewareClusterID int) error {
	err := rgs.ResourceGroupRepo.DeleteMiddlewareCluster(resourceGroupID, middlewareClusterID)
	if err != nil {
		return err
	}

	return rgs.GetMiddlewareClustersByID(resourceGroupID)
}

// Marshal marshals ResourceGroupService.ResourceGroups to json bytes
func (rgs *ResourceGroupService) Marshal() ([]byte, error) {
	return rgs.MarshalWithFields(resourceGroupResourceGroupsStruct)
}

// MarshalWithFields marshals only specified fields of the ResourceGroupService to json bytes
func (rgs *ResourceGroupService) MarshalWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(rgs, fields...)
}
