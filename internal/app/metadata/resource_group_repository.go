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

const defaultDASAdminRoleName = "das_admin"

var _ metadata.ResourceGroupRepo = (*ResourceGroupRepo)(nil)

type ResourceGroupRepo struct {
	Database middleware.Pool
}

// NewResourceGroupRepo returns *ResourceGroupRepo with given middleware.Pool
func NewResourceGroupRepo(db middleware.Pool) *ResourceGroupRepo {
	return &ResourceGroupRepo{db}
}

// NewResourceGroupRepoWithGlobal returns *ResourceGroupRepo with global mysql pool
func NewResourceGroupRepoWithGlobal() *ResourceGroupRepo {
	return NewResourceGroupRepo(global.DASMySQLPool)
}

// Execute implements metadata.ResourceGroupRepo interface,
// it executes command with arguments on database
func (rgr *ResourceGroupRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := rgr.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata ResourceGroupRepo.Execute(): close database connection failed.\n%+v", err)
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns middleware.PoolConn
func (rgr *ResourceGroupRepo) Transaction() (middleware.Transaction, error) {
	return rgr.Database.Transaction()
}

// GetAll returns all available entities
func (rgr *ResourceGroupRepo) GetAll() ([]metadata.ResourceGroup, error) {
	sql := `
        select id, group_uuid, group_name, del_flag, create_time, last_update_time
        from t_meta_resource_group_info
        where del_flag = 0
        order by id;
    `
	log.Debugf("metadata ResourceGroupRepo.GetAll() sql: \n%s", sql)

	result, err := rgr.Execute(sql)
	if err != nil {
		return nil, err
	}

	// init []metadata.ResourceGroup
	resourceGroupList := make([]metadata.ResourceGroup, result.RowNumber())
	for i := range resourceGroupList {
		resourceGroupList[i] = NewEmptyResourceGroupInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(resourceGroupList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return resourceGroupList, nil
}

// GetByID Select returns an available entity of the given id
func (rgr *ResourceGroupRepo) GetByID(id int) (metadata.ResourceGroup, error) {
	sql := `
        select id, group_uuid, group_name, del_flag, create_time, last_update_time
        from t_meta_resource_group_info
        where del_flag = 0
        and id = ?;
    `
	log.Debugf("metadata ResourceGroupRepo.GetByID() sql: \n%s\nplaceholders: %d", sql, id)

	result, err := rgr.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.Trace(fmt.Errorf("metadata ResourceGroupRepo.GetByID(): data does not exists, id: %d", id))
	case 1:
		resourceGroupInfo := NewEmptyResourceGroupInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(resourceGroupInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return resourceGroupInfo, nil
	default:
		return nil, errors.Trace(fmt.Errorf("metadata ResourceGroupRepo.GetByID(): duplicate key exists, id: %d", id))
	}
}

// GetByGroupUUID gets the resource group with given resource group id from the middleware
func (rgr *ResourceGroupRepo) GetByGroupUUID(groupUUID string) (metadata.ResourceGroup, error) {
	sql := `
        select id, group_uuid, group_name, del_flag, create_time, last_update_time
        from t_meta_resource_group_info
        where del_flag = 0
        and group_uuid = ?;
    `
	log.Debugf("metadata ResourceGroupRepo.GetByGroupUUID() select sql: %s", sql)
	result, err := rgr.Execute(sql, groupUUID)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.Trace(fmt.Errorf("metadata ResourceGroupRepo.GetByGroupUUID(): data does not exists, groupUUID: %s", groupUUID))
	case 1:
		resourceGroupInfo := NewEmptyResourceGroupInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(resourceGroupInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return resourceGroupInfo, nil
	default:
		return nil, errors.Trace(fmt.Errorf("metadata ResourceGroupRepo.GetByGroupUUID(): duplicate key exists, groupUUID: %s", groupUUID))
	}
}

// GetID gets the identity with given resource group id from the middleware
func (rgr *ResourceGroupRepo) GetID(groupUUID string) (int, error) {
	sql := `select id from t_meta_resource_group_info where del_flag = 0 and group_uuid = ?;`
	log.Debugf("metadata ResourceGroupRepo.GetID() select sql: %s", sql)
	result, err := rgr.Execute(sql, groupUUID)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// GetResourceRolesByID get all resource roles with given resource group id from the middleware
func (rgr *ResourceGroupRepo) GetResourceRolesByID(id int) ([]metadata.ResourceRole, error) {
	sql := `
        select rri.id,
			   rri.role_uuid,
			   rri.role_name,
			   rri.resource_group_id,
			   rri.del_flag,
			   rri.create_time,
			   rri.last_update_time
		from t_meta_resource_group_info rgi
			inner join t_meta_resource_role_info rri on rgi.id = rri.resource_group_id
		where rgi.del_flag = 0
		  and rri.del_flag = 0
		  and rgi.id = ?;
    `
	log.Debugf("metadata ResourceGroupRepo.GetResourceRoles() sql: \n%s\nplaceholders: %d", sql, id)

	result, err := rgr.Execute(sql, id)
	if err != nil {
		return nil, err
	}

	resourceRoleList := make([]metadata.ResourceRole, result.RowNumber())
	for row := range resourceRoleList {
		resourceRoleList[row] = NewEmptyResourceRoleInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(resourceRoleList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return resourceRoleList, nil
}

// GetMySQLClustersByID gets the mysql cluster with given resource group id from the middleware
func (rgr *ResourceGroupRepo) GetMySQLClustersByID(id int) ([]metadata.MySQLCluster, error) {
	sql := `
		select mci.id,
			   mci.cluster_name,
			   mci.middleware_cluster_id,
			   mci.monitor_system_id,
			   mci.env_id,
			   mci.del_flag,
			   mci.create_time,
			   mci.last_update_time
		from t_meta_mysql_cluster_info mci
			inner join t_meta_mysql_cluster_resource_group_map mcrgm on mci.id = mcrgm.mysql_cluster_id
			inner join t_meta_resource_group_info rgi on mcrgm.resource_group_id = rgi.id
		where mci.del_flag = 0
		  and mcrgm.del_flag = 0
		  and rgi.del_flag = 0
		  and rgi.id = ?
    `
	log.Debugf("metadata ResourceGroupRepo.GetMySQLClustersByID() sql: \n%s\nplaceholders: %d", sql, id)

	result, err := rgr.Execute(sql, id)
	if err != nil {
		return nil, err
	}

	mysqlClusterList := make([]metadata.MySQLCluster, result.RowNumber())
	for row := range mysqlClusterList {
		mysqlClusterList[row] = NewEmptyMySQLClusterInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(mysqlClusterList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return mysqlClusterList, nil
}

// GetMySQLServersByID gets the mysql servers with given resource group uuid from the middleware
func (rgr *ResourceGroupRepo) GetMySQLServersByID(id int) ([]metadata.MySQLServer, error) {
	sql := `
		select msi.id,
			   msi.cluster_id,
			   msi.server_name,
			   msi.service_name,
			   msi.host_ip,
			   msi.port_num,
			   msi.deployment_type,
			   msi.version,
			   msi.del_flag,
			   msi.create_time,
			   msi.last_update_time
		from t_meta_mysql_server_info msi
			inner join t_meta_mysql_cluster_info mci on msi.cluster_id = mci.id
			inner join t_meta_mysql_cluster_resource_group_map mcrgm on mci.id = mcrgm.mysql_cluster_id
			inner join t_meta_resource_group_info rgi on mcrgm.resource_group_id = rgi.id
		where mci.del_flag = 0
		  and mcrgm.del_flag = 0
		  and rgi.del_flag = 0
		  and rgi.id = ?
    `
	log.Debugf("metadata ResourceGroupRepo.GetMySQLServersByID() sql: \n%s\nplaceholders: %d", sql, id)

	result, err := rgr.Execute(sql, id)
	if err != nil {
		return nil, err
	}

	mysqlServerList := make([]metadata.MySQLServer, result.RowNumber())
	for row := range mysqlServerList {
		mysqlServerList[row] = NewEmptyMySQLServerInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(mysqlServerList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return mysqlServerList, nil
}

// GetMiddlewareClustersByID gets the middleware cluster with given resource group id from the middleware
func (rgr *ResourceGroupRepo) GetMiddlewareClustersByID(id int) ([]metadata.MiddlewareCluster, error) {
	sql := `
		select mci.id,
			   mci.cluster_name,
			   mci.env_id,
			   mci.del_flag,
			   mci.create_time,
			   mci.last_update_time
		from t_meta_middleware_cluster_info mci
			inner join t_meta_middleware_cluster_resource_group_map mcrgm on mci.id = mcrgm.middleware_cluster_id
			inner join t_meta_resource_group_info rgi on mcrgm.resource_group_id = rgi.id
		where mci.del_flag = 0
		  and mcrgm.del_flag = 0
		  and rgi.del_flag = 0
		  and rgi.id = ?
    `
	log.Debugf("metadata ResourceGroupRepo.GetMiddlewareClustersByID() sql: \n%s\nplaceholders: %d", sql, id)

	result, err := rgr.Execute(sql, id)
	if err != nil {
		return nil, err
	}

	middlewareClusterList := make([]metadata.MiddlewareCluster, result.RowNumber())
	for row := range middlewareClusterList {
		middlewareClusterList[row] = NewEmptyMiddlewareClusterInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(middlewareClusterList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return middlewareClusterList, nil
}

// GetMiddlewareServersByID gets the middleware servers with given resource group uuid from the middleware
func (rgr *ResourceGroupRepo) GetMiddlewareServersByID(id int) ([]metadata.MiddlewareServer, error) {
	sql := `
		select msi.id,
			   msi.cluster_id,
			   msi.server_name,
			   msi.middleware_role,
			   msi.host_ip,
			   msi.port_num,
			   msi.del_flag,
			   msi.create_time,
			   msi.last_update_time
		from t_meta_middleware_server_info msi
			inner join t_meta_middleware_cluster_info mci on msi.cluster_id = mci.id
			inner join t_meta_middleware_cluster_resource_group_map mcrgm on mci.id = mcrgm.middleware_cluster_id
			inner join t_meta_resource_group_info rgi on mcrgm.resource_group_id = rgi.id
		where mci.del_flag = 0
		  and mcrgm.del_flag = 0
		  and rgi.del_flag = 0
		  and rgi.id = ?
    `
	log.Debugf("metadata ResourceGroupRepo.GetMySQLServersByID() sql: \n%s\nplaceholders: %d", sql, id)

	result, err := rgr.Execute(sql, id)
	if err != nil {
		return nil, err
	}

	middlewareServerList := make([]metadata.MiddlewareServer, result.RowNumber())
	for row := range middlewareServerList {
		middlewareServerList[row] = NewEmptyMiddlewareServerInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(middlewareServerList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return middlewareServerList, nil
}

// GetUsersByID gets the users of the given id from the middleware
func (rgr *ResourceGroupRepo) GetUsersByID(id int) ([]metadata.User, error) {
	sql := `
		select distinct ui.id,
						ui.user_name,
						ui.department_name,
						ui.employee_id,
						ui.account_name,
						ui.email,
						ui.telephone,
						ui.mobile,
						ui.role,
						ui.del_flag,
						ui.create_time,
						ui.last_update_time
		from t_meta_resource_group_info rgi
			inner join t_meta_resource_role_info rri on rgi.id = rri.resource_group_id
			inner join t_meta_resource_role_user_map rrum on rri.id = rrum.resource_role_id
			inner join t_meta_user_info ui on rrum.user_id = ui.id
		where rgi.del_flag = 0
		  and rri.del_flag = 0
		  and rrum.del_flag = 0
		  and ui.del_flag = 0
		  and rgi.id = ?;
    `
	log.Debugf("metadata ResourceGroupRepo.GetUsersByID() sql: \n%s\nplaceholders: %d", sql, id)

	result, err := rgr.Execute(sql, id)
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

// GetUsersByID gets the users of the given id from the middleware
func (rgr *ResourceGroupRepo) GetDASAdminUsersByID(id int) ([]metadata.User, error) {
	sql := `
		select distinct ui.id,
						ui.user_name,
						ui.department_name,
						ui.employee_id,
						ui.account_name,
						ui.email,
						ui.telephone,
						ui.mobile,
						ui.role,
						ui.del_flag,
						ui.create_time,
						ui.last_update_time
		from t_meta_resource_group_info rgi
			inner join t_meta_resource_role_info rri on rgi.id = rri.resource_group_id
			inner join t_meta_resource_role_user_map rrum on rri.id = rrum.resource_role_id
			inner join t_meta_user_info ui on rrum.user_id = ui.id
		where rgi.del_flag = 0
		  and rri.del_flag = 0
		  and rrum.del_flag = 0
		  and ui.del_flag = 0
		  and rgi.id = ?
		  and rri.role_name = ?;
    `
	log.Debugf("metadata ResourceGroupRepo.GetDASAdminUsersByID() sql: \n%s\nplaceholders: %d, %s", sql, id, defaultDASAdminRoleName)

	result, err := rgr.Execute(sql, id, defaultDASAdminRoleName)
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
func (rgr *ResourceGroupRepo) Create(rg metadata.ResourceGroup) (metadata.ResourceGroup, error) {
	sql := `insert into t_meta_resource_group_info(group_uuid, group_name) values(?,?);`
	log.Debugf("metadata ResourceGroupRepo.Create() insert sql: \n%s\nplaceholders: %s, %s", sql, rg.GetGroupUUID(), rg.GetGroupName())
	// execute
	_, err := rgr.Execute(sql, rg.GetGroupUUID(), rg.GetGroupName())
	if err != nil {
		return nil, err
	}

	// get entity
	return rgr.GetByGroupUUID(rg.GetGroupUUID())
}

// Update updates data with given entity in the middleware
func (rgr *ResourceGroupRepo) Update(rg metadata.ResourceGroup) error {
	sql := `update t_meta_resource_group_info set group_uuid = ?, group_name = ?, del_flag = ? where id = ?;`
	log.Debugf("metadata ResourceGroupRepo.Update() update sql: %s", sql)
	_, err := rgr.Execute(sql,
		rg.GetGroupUUID(),
		rg.GetGroupName(),
		rg.GetDelFlag(),
		rg.Identity(),
	)

	return err
}

// Delete deletes data in the middleware, it will also delete concerning map data
func (rgr *ResourceGroupRepo) Delete(id int) error {
	// delete concerning resource roles
	resourceRoles, err := rgr.GetResourceRolesByID(id)
	if err != nil {
		return err
	}
	rrr := NewResourceRoleRepoWithGlobal()
	for _, resourceRole := range resourceRoles {
		err = rrr.Delete(resourceRole.Identity())
		if err != nil {
			return err
		}
	}

	// delete mysql clusters from the map
	sql := `delete from t_meta_mysql_cluster_resource_group_map where resource_group_id = ?`
	log.Debugf("metadata ResourceGroupRepo.Delete() delete sql(t_meta_mysql_cluster_resource_group_map): %s", sql)
	_, err = rgr.Execute(sql, id)
	if err != nil {
		return err
	}

	// delete resource group
	sql = `delete from t_meta_resource_group_info where id = ?;`
	log.Debugf("metadata ResourceGroupRepo.Delete() delete sql(t_meta_resource_group_info): %s", sql)

	_, err = rgr.Execute(sql, id)

	return err
}

// AddMySQLCluster adds mysql cluster to the resource group
func (rgr *ResourceGroupRepo) AddMySQLCluster(resourceGroupID int, mysqlClusterID int) error {
	_, err := rgr.GetByID(resourceGroupID)
	if err != nil {
		return err
	}

	mysqlClusterRepo := NewMySQLClusterRepoWithGlobal()
	_, err = mysqlClusterRepo.GetByID(mysqlClusterID)
	if err != nil {
		return err
	}

	sql := `
		insert into t_meta_mysql_cluster_resource_group_map(mysql_cluster_id, resource_group_id)
		values(?, ?);
	`
	log.Debugf("metadata ResourceGroupRepo.AddMySQLCluster() insert sql: %s", sql)
	_, err = rgr.Execute(sql, resourceGroupID, mysqlClusterID)

	return err
}

// DeleteMySQLCluster deletes mysql cluster from the resource group
func (rgr *ResourceGroupRepo) DeleteMySQLCluster(resourceGroupID, mysqlClusterID int) error {
	_, err := rgr.GetByID(resourceGroupID)
	if err != nil {
		return err
	}

	sql := `delete from t_meta_mysql_cluster_resource_group_map where resource_group_id = ? and mysql_cluster_id = ?;`
	log.Debugf("metadata ResourceGroupRepo.DeleteMySQLCluster() delete sql: %s", sql)
	_, err = rgr.Execute(sql, resourceGroupID, mysqlClusterID)

	return err
}

// AddMiddlewareCluster adds middleware cluster to the resource group
func (rgr *ResourceGroupRepo) AddMiddlewareCluster(resourceGroupID int, middlewareClusterID int) error {
	_, err := rgr.GetByID(resourceGroupID)
	if err != nil {
		return err
	}

	middlewareClusterRepo := NewMiddlewareClusterRepoWithGlobal()
	_, err = middlewareClusterRepo.GetByID(middlewareClusterID)
	if err != nil {
		return err
	}

	sql := `
		insert into t_meta_middleware_cluster_resource_group_map(middleware_cluster_id, resource_group_id)
		values(?, ?);
	`
	log.Debugf("metadata ResourceGroupRepo.AddMiddlewareCluster() insert sql: %s", sql)
	_, err = rgr.Execute(sql, resourceGroupID, middlewareClusterID)

	return err
}

// DeleteMiddlewareCluster deletes middleware cluster from the resource group
func (rgr *ResourceGroupRepo) DeleteMiddlewareCluster(resourceGroupID, middlewareClusterID int) error {
	_, err := rgr.GetByID(resourceGroupID)
	if err != nil {
		return err
	}

	sql := `delete from t_meta_middleware_cluster_resource_group_map where resource_group_id = ? and middleware_cluster_id = ?;`
	log.Debugf("metadata ResourceGroupRepo.DeleteMiddlewareCluster() delete sql: %s", sql)
	_, err = rgr.Execute(sql, resourceGroupID, middlewareClusterID)

	return err
}
