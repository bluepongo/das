package metadata

import (
	"fmt"

	"github.com/pingcap/errors"
	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/log"
)

var _ metadata.ResourceRoleRepo = (*ResourceRoleRepo)(nil)

// ResourceRoleRepo implements dependency.ResourceRoleRepo interface
type ResourceRoleRepo struct {
	Database middleware.Pool
}

// NewResourceRoleRepo returns *ResourceRoleRepo with given middleware.Pool
func NewResourceRoleRepo(db middleware.Pool) *ResourceRoleRepo {
	return &ResourceRoleRepo{db}
}

// NewResourceRoleRepoWithGlobal returns *ResourceRoleRepo with global mysql pool
func NewResourceRoleRepoWithGlobal() *ResourceRoleRepo {
	return NewResourceRoleRepo(global.DASMySQLPool)
}

// Execute executes given command and placeholders on the middleware
// it executes command with arguments on database
func (rrr *ResourceRoleRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := rrr.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata ResourceRoleRepo.Execute(): close database connection failed.\n%+v", err)
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns a mysql.Transaction that could execute multiple commands as a transaction
func (rrr *ResourceRoleRepo) Transaction() (middleware.Transaction, error) {
	return rrr.Database.Transaction()
}

// GetAll gets all resource groups from the middleware
func (rrr *ResourceRoleRepo) GetAll() ([]metadata.ResourceRole, error) {
	sql := `
		select id, role_uuid, role_name, resource_group_id, del_flag, create_time, last_update_time
		from t_meta_resource_role_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata ResourceRoleRepo.GetAll() sql: \n%s", sql)

	result, err := rrr.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*ResourceRoleInfo
	resourceRoleInfoList := make([]*ResourceRoleInfo, result.RowNumber())
	for i := range resourceRoleInfoList {
		resourceRoleInfoList[i] = NewEmptyResourceRoleInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(resourceRoleInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []metadata.ResourceRole
	resourceRoleList := make([]metadata.ResourceRole, result.RowNumber())
	for i := range resourceRoleList {
		resourceRoleList[i] = resourceRoleInfoList[i]
	}

	return resourceRoleList, nil
}

// GetByID gets the resource role by the identity from the middleware
func (rrr *ResourceRoleRepo) GetByID(id int) (metadata.ResourceRole, error) {
	sql := `
		select id, role_uuid, role_name, resource_group_id, del_flag, create_time, last_update_time
		from t_meta_resource_role_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata ResourceRoleRepo.GetByID() sql: \n%s\nplaceholders: %d", sql, id)

	result, err := rrr.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.Trace(fmt.Errorf("metadata ResourceRoleRepo.GetByID(): data does not exists, id: %d", id))
	case 1:
		resourceRoleInfo := NewEmptyResourceRoleInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(resourceRoleInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return resourceRoleInfo, nil
	default:
		return nil, errors.Trace(fmt.Errorf("metadata ResourceRoleRepo.GetByID(): duplicate key exists, id: %d", id))
	}
}

// GetID gets the identity with given resource role id from the middleware
func (rrr *ResourceRoleRepo) GetID(roleUUID string) (int, error) {
	sql := `select id from t_meta_resource_role_info where del_flag = 0 and role_uuid = ?;`
	log.Debugf("metadata ResourceRoleRepo.GetID() select sql: %s", sql)
	result, err := rrr.Execute(sql, roleUUID)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// GetByRoleUUID gets the resource role with given resource role id from the middleware
func (rrr *ResourceRoleRepo) GetByRoleUUID(roleUUID string) (metadata.ResourceRole, error) {
	sql := `
		select id, role_uuid, role_name, resource_group_id, del_flag, create_time, last_update_time
		from t_meta_resource_role_info
		where del_flag = 0 
		and role_uuid = ?;
	`
	log.Debugf("metadata ResourceRoleRepo.GetByName() select sql: %s", sql)
	result, err := rrr.Execute(sql, roleUUID)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.Trace(fmt.Errorf("metadata ResourceRoleRepo.GetByName(): data does not exists, roleUUID: %s", roleUUID))
	case 1:
		resourceRoleInfo := NewEmptyResourceRoleInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(resourceRoleInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return resourceRoleInfo, nil
	default:
		return nil, errors.Trace(fmt.Errorf("metadata ResourceRoleRepo.GetByName(): duplicate key exists, roleUUID: %s", roleUUID))
	}
}

// GetResourceGroup gets the resource group which this role belongs to with given resource role id from the middleware
func (rrr *ResourceRoleRepo) GetResourceGroup(id int) (metadata.ResourceGroup, error) {
	// TODO: select group info by group id from role
	sql := `
		select id, group_uuid, group_name, del_flag, create_time, last_update_time
		from t_meta_resource_group_info as group
				 inner join t_meta_resource_role_info as role
							on group.id = role.resource_group_id
		where role.id = ?
	`

	log.Debugf("metadata ResourceRoleRepo.GetResourceGroup() select sql: %s", sql)
	result, err := rrr.Execute(sql, id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetUsersByID gets the users of the given id from the middleware
func (rrr *ResourceRoleRepo) GetUsersByID(id int) ([]metadata.User, error) {
	sql := `
		select distinct user.id,
						user.user_name,
						user.department_name,
						user.employee_id,
						user.account_name,
						user.email,
						user.telephone,
						user.mobile,
						user.role,
						user.del_flag,
						user.create_time,
						user.last_update_time
		from t_meta_user_info as user
				 inner join t_meta_resource_role_user_map as rum
							on user.id = rum.user_id
				 inner join t_meta_resource_role_info as role
							on role.id = rum.resource_role_id
		where user.del_flag = 0
		  and rum.del_flag = 0
		  and role.del_flag = 0
		  and role.id = ?;
	`
	log.Debugf("metadata ResourceRoleRepo.GetUsersByID() sql: \n%s\nplaceholders: %d", sql, id)

	result, err := rrr.Execute(sql, id)
	if err != nil {
		return nil, err
	}

	userList := make([]metadata.User, result.RowNumber())
	for row := range userList {
		userList[row] = NewEmptyUserInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(userList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return userList, nil
}

// Create creates data with given entity in the middleware
func (rrr *ResourceRoleRepo) Create(resourceRole metadata.ResourceRole) (metadata.ResourceRole, error) {
	sql := `
		insert into t_meta_resource_role_info(role_uuid, role_name, resource_group_id) 
		values(?,?,?);`
	log.Debugf("metadata ResourceRoleRepo.Create() insert sql: %s", sql)
	// execute
	_, err := rrr.Execute(sql,
		resourceRole.GetRoleUUID(),
		resourceRole.GetRoleName(),
		resourceRole.GetResourceGroupID())
	if err != nil {
		return nil, err
	}
	// get role id
	id, err := rrr.GetID(resourceRole.GetRoleUUID())
	if err != nil {
		return nil, err
	}
	// get entity
	return rrr.GetByID(id)
}

// Update updates data with given entity in the middleware
func (rrr *ResourceRoleRepo) Update(entity metadata.ResourceRole) error {
	sql := `
		update t_meta_resource_role_info set role_uuid = ?, role_name = ?, 
			resource_group_id = ?, del_flag = ? 
		where id = ?;`
	log.Debugf("metadata ResourceRoleRepo.Update() update sql: %s", sql)
	resourceRoleInfo := entity.(*ResourceRoleInfo)
	_, err := rrr.Execute(sql,
		resourceRoleInfo.RoleUUID,
		resourceRoleInfo.RoleName,
		resourceRoleInfo.ResourceGroupID,
		resourceRoleInfo.DelFlag, resourceRoleInfo.ID)

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (rrr *ResourceRoleRepo) Delete(id int) error {
	// FIXME: soft delete?
	sql := `delete from t_meta_resource_role_info where id = ?;`
	log.Debugf("metadata ResourceRoleRepo.Delete() delete sql(t_meta_resource_role_info): %s", sql)

	_, err := rrr.Execute(sql, id)
	return err
}

// AddUser adds a map of the resource role and user from the middleware
func (rrr *ResourceRoleRepo) AddUser(roleID int, userID int) error {
	userRepo := NewUserRepoWithGlobal()

	_, err := userRepo.GetByID(userID)
	if err != nil {
		return err
	}
	_, err = rrr.GetByID(roleID)
	if err != nil {
		return err
	}
	sql := `insert into t_meta_resource_role_user_map(resource_role_id, user_id) values(?, ?);`
	log.Debugf("metadata MySQLClusterRepo.AddUser() insert sql: %s", sql)
	_, err = rrr.Execute(sql, roleID, userID)

	return err
}

// DeleteUser deletes the map of the resource role and user from the middleware
func (rrr *ResourceRoleRepo) DeleteUser(roleID int, userID int) error {
	userRepo := NewUserRepoWithGlobal()
	_, err := userRepo.GetByID(userID)
	if err != nil {
		return err
	}
	_, err = rrr.GetByID(roleID)
	if err != nil {
		return err
	}
	sql := `delete from t_meta_resource_role_user_map where resource_role_id = ? and user_id = ?;`
	log.Debugf("metadata AppRepo.DeleteUser() delete sql: %s", sql)
	_, err = rrr.Execute(sql, roleID, userID)

	return err
}
