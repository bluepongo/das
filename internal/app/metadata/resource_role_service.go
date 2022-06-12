package metadata

import (
	"fmt"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
)

const resourceRoleResourceRolesStruct = "ResourceRoles"

var _ metadata.ResourceRoleService = (*ResourceRoleService)(nil)

// ResourceRoleService implements Service interface
type ResourceRoleService struct {
	ResourceRoleRepo metadata.ResourceRoleRepo
	ResourceRoles    []metadata.ResourceRole `json:"resource_roles"`
	ResourceGroup    metadata.ResourceGroup  `json:"resource_group"`
	Users            []metadata.User         `json:"users"`
}

// NewResourceRoleService returns a new *ResourceRoleService
func NewResourceRoleService(repo metadata.ResourceRoleRepo) *ResourceRoleService {
	return &ResourceRoleService{
		repo,
		[]metadata.ResourceRole{},
		nil,
		[]metadata.User{},
	}
}

// NewResourceRoleServiceWithDefault returns a new *ResourceRoleService with default repository
func NewResourceRoleServiceWithDefault() *ResourceRoleService {
	return NewResourceRoleService(NewResourceRoleRepoWithGlobal())
}

// GetResourceRoles returns the resource roles of the service
func (rrs *ResourceRoleService) GetResourceRoles() []metadata.ResourceRole {
	return rrs.ResourceRoles
}

// GetResourceGroup returns the resource group of the service
func (rrs *ResourceRoleService) GetResourceGroup() metadata.ResourceGroup {
	return nil
}

// GetUsers returns the users of the service
func (rrs *ResourceRoleService) GetUsers() []metadata.User {
	return rrs.Users
}

// GetAll gets all mysql cluster entities from the middleware
func (rrs *ResourceRoleService) GetAll() error {
	var err error

	rrs.ResourceRoles, err = rrs.ResourceRoleRepo.GetAll()

	return err
}

// GetByID gets a mysql cluster entity that contains the given id from the middleware
func (rrs *ResourceRoleService) GetByID(id int) error {
	resourceRole, err := rrs.ResourceRoleRepo.GetByID(id)
	if err != nil {
		return err
	}

	rrs.ResourceRoles = nil
	rrs.ResourceRoles = append(rrs.ResourceRoles, resourceRole)

	return nil
}

// GetByRoleUUID gets the resource role by role uuid
func (rrs *ResourceRoleService) GetByRoleUUID(groupUUID string) error {
	resourceRole, err := rrs.ResourceRoleRepo.GetByRoleUUID(groupUUID)
	if err != nil {
		return err
	}

	rrs.ResourceRoles = nil
	rrs.ResourceRoles = append(rrs.ResourceRoles, resourceRole)

	return nil
}

// GetResourceGroupByID gets the resource group with given resource role id
func (rrs *ResourceRoleService) GetResourceGroupByID(id int) error {
	var err error

	rrs.ResourceGroup, err = rrs.ResourceRoleRepo.GetResourceGroup(id)
	if err != nil {
		return err
	}

	return nil
}

// GetUsersByID gets the users with given resource role id
func (rrs *ResourceRoleService) GetUsersByID(id int) error {
	var err error

	rrs.Users, err = rrs.ResourceRoleRepo.GetUsersByID(id)

	return err
}

// GetUsersByRoleUUID gets the users with given resource role uuid
func (rrs *ResourceRoleService) GetUsersByRoleUUID(roleUUID string) error {
	id, err := rrs.ResourceRoleRepo.GetID(roleUUID)
	if err != nil {
		return err
	}

	rrs.Users, err = rrs.ResourceRoleRepo.GetUsersByID(id)
	if err != nil {
		return err
	}

	return nil
}

// Create creates a new mysql cluster entity and insert it into the middleware
func (rrs *ResourceRoleService) Create(fields map[string]interface{}) error {
	// generate new map
	_, roleUUIDExists := fields[resourceRoleRoleUUIDStruct]
	_, roleResourceGroupIDExists := fields[resourceRoleResourceGroupIDStruct]

	if !roleUUIDExists || !roleResourceGroupIDExists {
		return message.NewMessage(
			message.ErrFieldNotExists,
			fmt.Sprintf(
				"%s and %s",
				resourceRoleRoleUUIDStruct,
				resourceRoleResourceGroupIDStruct))
	}

	// create a new entity
	resourceRoleInfo, err := NewResourceRoleInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	resourceRole, err := rrs.ResourceRoleRepo.Create(resourceRoleInfo)
	if err != nil {
		return err
	}

	rrs.ResourceRoles = nil
	rrs.ResourceRoles = append(rrs.ResourceRoles, resourceRole)

	return nil
}

// Update gets an mysql cluster entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (rrs *ResourceRoleService) Update(id int, fields map[string]interface{}) error {
	err := rrs.GetByID(id)
	if err != nil {
		return err
	}
	err = rrs.ResourceRoles[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return rrs.ResourceRoleRepo.Update(rrs.ResourceRoles[constant.ZeroInt])
}

// Delete deletes the mysql cluster entity that contains the given id in the middleware
func (rrs *ResourceRoleService) Delete(id int) error {
	err := rrs.GetByID(id)
	if err != nil {
		return err
	}

	return rrs.ResourceRoleRepo.Delete(id)
}

// AddUser adds a new map of mysql cluster and user in the middleware
func (rrs *ResourceRoleService) AddUser(resourceRoleID, userID int) error {
	if err := rrs.ResourceRoleRepo.AddUser(resourceRoleID, userID); err != nil {
		return err
	}

	return rrs.GetUsersByID(resourceRoleID)
}

// DeleteUser deletes the map of mysql cluster and user in the middleware
func (rrs *ResourceRoleService) DeleteUser(resourceRoleID, userID int) error {
	err := rrs.ResourceRoleRepo.DeleteUser(resourceRoleID, userID)
	if err != nil {
		return err
	}

	return rrs.GetUsersByID(resourceRoleID)
}

// Marshal marshals service.Envs
func (rrs *ResourceRoleService) Marshal() ([]byte, error) {
	return rrs.MarshalWithFields(resourceRoleResourceRolesStruct)
}

// MarshalWithFields marshals service.Envs with given fields
func (rrs *ResourceRoleService) MarshalWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(rrs, fields...)
}
