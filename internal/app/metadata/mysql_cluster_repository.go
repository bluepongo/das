package metadata

import (
	"fmt"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/log"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/metadata"
)

const (
	ClusterTypeSingle = 1
	ClusterTypeShard  = 2
)

var _ metadata.MySQLClusterRepo = (*MySQLClusterRepo)(nil)

// MySQLClusterRepo implements dependency.MySQLClusterRepo interface
type MySQLClusterRepo struct {
	Database middleware.Pool
}

// NewMySQLClusterRepo returns *MySQLClusterRepo with given middleware.Pool
func NewMySQLClusterRepo(db middleware.Pool) *MySQLClusterRepo {
	return &MySQLClusterRepo{db}
}

// NewMySQLClusterRepoWithGlobal returns *MySQLClusterRepo with global mysql pool
func NewMySQLClusterRepoWithGlobal() *MySQLClusterRepo {
	return NewMySQLClusterRepo(global.DASMySQLPool)
}

// Execute implements metadata.MySQLClusterRepo interface,
// it executes command with arguments on database
func (mcr *MySQLClusterRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := mcr.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata MySQLClusterRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns middleware.PoolConn
func (mcr *MySQLClusterRepo) Transaction() (middleware.Transaction, error) {
	return mcr.Database.Transaction()
}

// GetAll returns all available entities
func (mcr *MySQLClusterRepo) GetAll() ([]metadata.MySQLCluster, error) {
	sql := `
		select id, cluster_name, middleware_cluster_id, monitor_system_id, 
			owner_id, env_id, del_flag, create_time, last_update_time
		from t_meta_mysql_cluster_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata MySQLClusterRepo.GetAll() sql: \n%s", sql)

	result, err := mcr.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*MySQLClusterInfo
	mysqlClusterInfoList := make([]*MySQLClusterInfo, result.RowNumber())
	for i := range mysqlClusterInfoList {
		mysqlClusterInfoList[i] = NewEmptyMySQLClusterInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(mysqlClusterInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []metadata.MySQLCluster
	mysqlClusterList := make([]metadata.MySQLCluster, result.RowNumber())
	for i := range mysqlClusterList {
		mysqlClusterList[i] = mysqlClusterInfoList[i]
	}

	return mysqlClusterList, nil
}

// GetByEnv gets mysql clusters of given env id from the middleware
func (mcr *MySQLClusterRepo) GetByEnv(envID int) ([]metadata.MySQLCluster, error) {
	sql := `
		select id, cluster_name, middleware_cluster_id, monitor_system_id, 
			owner_id, env_id, del_flag, create_time, last_update_time
		from t_meta_mysql_cluster_info
		where del_flag = 0
		and env_id = ?;
	`
	log.Debugf("metadata MySQLClusterRepo.GetByEnv() sql: \n%s\nplaceholders: %d", sql, envID)

	result, err := mcr.Execute(sql, envID)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	mysqlClusterList := make([]metadata.MySQLCluster, resultNum)

	for row := 0; row < resultNum; row++ {
		mysqlClusterList[row] = NewEmptyMySQLClusterInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(mysqlClusterList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return mysqlClusterList, nil
}

// GetByID Select returns an available entity of the given id
func (mcr *MySQLClusterRepo) GetByID(id int) (metadata.MySQLCluster, error) {
	sql := `
		select id, cluster_name, middleware_cluster_id, monitor_system_id, 
			owner_id, env_id, del_flag, create_time, last_update_time
		from t_meta_mysql_cluster_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata MySQLClusterRepo.GetByID() sql: \n%s\nplaceholders: %d", sql, id)

	result, err := mcr.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, fmt.Errorf("metadata MySQLClusterInfo.GetByID(): data does not exists, id: %d", id)
	case 1:
		mysqlClusterInfo := NewEmptyMySQLClusterInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(mysqlClusterInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return mysqlClusterInfo, nil
	default:
		return nil, fmt.Errorf("metadata MySQLClusterInfo.GetByID(): duplicate key exists, id: %d", id)
	}
}

// GetByName gets a mysql cluster of given cluster name from the middle ware
func (mcr *MySQLClusterRepo) GetByName(clusterName string) (metadata.MySQLCluster, error) {
	sql := `
		select id, cluster_name, middleware_cluster_id, monitor_system_id, 
			owner_id, env_id, del_flag, create_time, last_update_time
		from t_meta_mysql_cluster_info where del_flag = 0 and cluster_name = ?;
	`
	log.Debugf("metadata MySQLClusterRepo.GetByName() select sql: %s", sql)
	result, err := mcr.Execute(sql, clusterName)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, fmt.Errorf("metadata MySQLClusterInfo.GetByID(): data does not exists, clusterName: %s", clusterName)
	case 1:
		mysqlClusterInfo := NewEmptyMySQLClusterInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(mysqlClusterInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return mysqlClusterInfo, nil
	default:
		return nil, fmt.Errorf("metadata MySQLClusterInfo.GetByID(): duplicate key exists, clusterName: %s", clusterName)
	}
}

// GetID checks identity of given entity from the middleware
func (mcr *MySQLClusterRepo) GetID(clusterName string) (int, error) {
	sql := `select id from t_meta_mysql_cluster_info where del_flag = 0 and cluster_name = ?;`
	log.Debugf("metadata MySQLClusterRepo.GetID() select sql: %s", sql)
	result, err := mcr.Execute(sql, clusterName)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// GetDBsByID gets the databases of the given id from the middleware
func (mcr *MySQLClusterRepo) GetDBsByID(id int) ([]metadata.DB, error) {
	sql := `
		select id, db_name, cluster_id, cluster_type, owner_id, env_id, del_flag, create_time, last_update_time
		from t_meta_db_info
		where del_flag = 0
		and cluster_id = ?
		and cluster_type = ?;
	`
	log.Debugf("metadata MySQLClusterRepo.GetDBsByID() sql: \n%s\nplaceholders: %d", sql, id, ClusterTypeSingle)

	result, err := mcr.Execute(sql, id, ClusterTypeSingle)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	dbList := make([]metadata.DB, resultNum)

	for row := 0; row < resultNum; row++ {
		dbList[row] = NewEmptyDBInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(dbList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return dbList, nil
}

// GetAppOwnersByID gets the application owners of the given id from the middleware
func (mcr *MySQLClusterRepo) GetAppOwnersByID(id int) ([]metadata.User, error) {
	sql := `
		select user.id, user.user_name, user.department_name, user.employee_id, user.account_name, user.email, user.telephone, user.mobile, user.role, user.del_flag, user.create_time, user.last_update_time 
			from t_meta_user_info as user 
			inner join t_meta_app_info as app
			on user.id = app.owner_id
			inner join t_meta_app_db_map as map
			on app.id = map.app_id
			inner join t_meta_db_info as db
			on db.id = map.db_id
			where user.del_flag = 0 and app.del_flag = 0 and db.del_flag = 0 and map.del_flag = 0
			and db.cluster_id = ?
			and db.cluster_type = ?;
	`
	log.Debugf("metadata MySQLClusterRepo.GetAppOwnersByID() sql: \n%s\nplaceholders: %d", sql, id, ClusterTypeSingle)

	result, err := mcr.Execute(sql, id, ClusterTypeSingle)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	userList := make([]metadata.User, resultNum)

	for row := 0; row < resultNum; row++ {
		userList[row] = NewEmptyUserInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(userList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return userList, nil
}

// GetDBOwnersByID gets the db owners of the given id from the middleware
func (mcr *MySQLClusterRepo) GetDBOwnersByID(id int) ([]metadata.User, error) {
	sql := `
		select user.id, user.user_name, user.department_name, user.employee_id, user.account_name, user.email, user.telephone, user.mobile, user.role, user.del_flag, user.create_time, user.last_update_time 
			from t_meta_user_info as user 
			inner join t_meta_db_info as db
			on user.id = db.owner_id
			where user.del_flag = 0 and db.del_flag = 0
			and db.cluster_id = ?
			and db.cluster_type = ?;
	`
	log.Debugf("metadata MySQLClusterRepo.GetDBOwnersByID() sql: \n%s\nplaceholders: %d", sql, id, ClusterTypeSingle)

	result, err := mcr.Execute(sql, id, ClusterTypeSingle)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	userList := make([]metadata.User, resultNum)

	for row := 0; row < resultNum; row++ {
		userList[row] = NewEmptyUserInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(userList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return userList, nil
}

// GetAllOwnersByID gets both application and db owners of the given id from the middleware
func (mcr *MySQLClusterRepo) GetAllOwnersByID(id int) ([]metadata.User, error) {
	sql := `
		select user.id, user.user_name, user.department_name, user.employee_id, user.account_name, user.email, user.telephone, user.mobile, user.role, user.del_flag, user.create_time, user.last_update_time 
			from t_meta_user_info as user 
			inner join t_meta_app_info as app
			on user.id = app.owner_id
			inner join t_meta_app_db_map as map
			on app.id = map.app_id
			inner join t_meta_db_info as db
			on db.id = map.db_id
			where user.del_flag = 0 and app.del_flag = 0 and db.del_flag = 0 and map.del_flag = 0
			and db.cluster_id = ? 
		union
		select user.id, user.user_name, user.department_name, user.employee_id, user.account_name, user.email, user.telephone, user.mobile, user.role, user.del_flag, user.create_time, user.last_update_time 
			from t_meta_user_info as user 
			inner join t_meta_db_info as db
			on user.id = db.owner_id
			where user.del_flag = 0 and db.del_flag = 0
			and db.cluster_id = ?
			and db.cluster_type = ?;
	`
	log.Debugf("metadata MySQLClusterRepo.GetAppOwnersByID() sql: \n%s\nplaceholders: %d, %d", sql, id, id, ClusterTypeSingle)

	result, err := mcr.Execute(sql, id, id, ClusterTypeSingle)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	userList := make([]metadata.User, resultNum)

	for row := 0; row < resultNum; row++ {
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
func (mcr *MySQLClusterRepo) Create(mysqlCluster metadata.MySQLCluster) (metadata.MySQLCluster, error) {
	sql := `
		insert into t_meta_mysql_cluster_info(cluster_name,middleware_cluster_id,
			 monitor_system_id, owner_id, env_id) 
		values(?,?,?,?,?);`
	log.Debugf("metadata MySQLClusterRepo.Create() insert sql: %s", sql)
	// execute
	_, err := mcr.Execute(sql,
		mysqlCluster.GetClusterName(),
		mysqlCluster.GetMiddlewareClusterID(),
		mysqlCluster.GetMonitorSystemID(),
		mysqlCluster.GetOwnerID(),
		mysqlCluster.GetEnvID())
	if err != nil {
		return nil, err
	}
	// get cluster id
	id, err := mcr.GetID(mysqlCluster.GetClusterName())
	if err != nil {
		return nil, err
	}
	// get entity
	return mcr.GetByID(id)
}

// Update updates data with given entity in the middleware
func (mcr *MySQLClusterRepo) Update(entity metadata.MySQLCluster) error {
	sql := `
		update t_meta_mysql_cluster_info set cluster_name = ?, middleware_cluster_id = ?, 
			monitor_system_id = ?, owner_id = ?, 
			env_id = ?, del_flag = ? 
		where id = ?;`
	log.Debugf("metadata MySQLClusterRepo.Update() update sql: %s", sql)
	mysqlClusterInfo := entity.(*MySQLClusterInfo)
	_, err := mcr.Execute(sql,
		mysqlClusterInfo.ClusterName,
		mysqlClusterInfo.MiddlewareClusterID,
		mysqlClusterInfo.MonitorSystemID,
		mysqlClusterInfo.OwnerID,
		mysqlClusterInfo.EnvID,
		mysqlClusterInfo.DelFlag, mysqlClusterInfo.ID)

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (mcr *MySQLClusterRepo) Delete(id int) error {
	sql := `delete from t_meta_mysql_cluster_info where id = ?;`
	log.Debugf("metadata MySQLCLusterRepo.Delete() delete sql(t_meta_mysql_cluster_info): %s", sql)

	_, err := mcr.Execute(sql, id)
	return err
}
