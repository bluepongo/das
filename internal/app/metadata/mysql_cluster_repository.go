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

// Cluster type code
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
			log.Errorf("metadata MySQLClusterRepo.Execute(): close database connection failed.\n%+v", err)
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
			env_id, del_flag, create_time, last_update_time
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
			env_id, del_flag, create_time, last_update_time
		from t_meta_mysql_cluster_info
		where del_flag = 0
		and env_id = ?;
	`
	log.Debugf("metadata MySQLClusterRepo.GetByEnv() sql: \n%s\nplaceholders: %d", sql, envID)

	result, err := mcr.Execute(sql, envID)
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

// GetByID Select returns an available entity of the given id
func (mcr *MySQLClusterRepo) GetByID(id int) (metadata.MySQLCluster, error) {
	sql := `
		select id, cluster_name, middleware_cluster_id, monitor_system_id, 
			env_id, del_flag, create_time, last_update_time
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
		return nil, errors.Trace(fmt.Errorf("metadata MySQLClusterRepo.GetByID(): data does not exists, id: %d", id))
	case 1:
		mysqlClusterInfo := NewEmptyMySQLClusterInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(mysqlClusterInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return mysqlClusterInfo, nil
	default:
		return nil, errors.Trace(fmt.Errorf("metadata MySQLClusterRepo.GetByID(): duplicate key exists, id: %d", id))
	}
}

// GetByName gets a mysql cluster of given cluster name from the middle ware
func (mcr *MySQLClusterRepo) GetByName(clusterName string) (metadata.MySQLCluster, error) {
	sql := `
		select id, cluster_name, middleware_cluster_id, monitor_system_id, 
			env_id, del_flag, create_time, last_update_time
		from t_meta_mysql_cluster_info where del_flag = 0 and cluster_name = ?;
	`
	log.Debugf("metadata MySQLClusterRepo.GetByName() select sql: %s", sql)
	result, err := mcr.Execute(sql, clusterName)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.Trace(fmt.Errorf("metadata MySQLClusterRepo.GetByName(): data does not exists, clusterName: %s", clusterName))
	case 1:
		mysqlClusterInfo := NewEmptyMySQLClusterInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(mysqlClusterInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return mysqlClusterInfo, nil
	default:
		return nil, errors.Trace(fmt.Errorf("metadata MySQLClusterRepo.GetByName(): duplicate key exists, clusterName: %s", clusterName))
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
		select id, db_name, cluster_id, cluster_type, env_id, del_flag, create_time, last_update_time
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

	dbList := make([]metadata.DB, result.RowNumber())
	for row := range dbList {
		dbList[row] = NewEmptyDBInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(dbList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return dbList, nil
}

// GetUsersByID gets the users of the given id from the middleware
func (mcr *MySQLClusterRepo) GetUsersByID(id int) ([]metadata.User, error) {
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
				 inner join t_meta_mysql_cluster_user_map as cum
							on user.id = cum.user_id
				 inner join t_meta_mysql_cluster_info as cluster
							on cluster.id = cum.mysql_cluster_id
		where user.del_flag = 0
		  and cum.del_flag = 0
		  and cluster.del_flag = 0
		  and cluster.id = ?;
	`
	log.Debugf("metadata MySQLClusterRepo.GetUsersByID() sql: \n%s\nplaceholders: %d", sql, id)

	result, err := mcr.Execute(sql, id)
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

// AddUser add a user for mysql cluster in the middleware
func (mcr *MySQLClusterRepo) AddUser(mysqlClusterID, userID int) error {
	userRepo := NewUserRepoWithGlobal()
	_, err := userRepo.GetByID(userID)
	if err != nil {
		return err
	}
	_, err = mcr.GetByID(mysqlClusterID)
	if err != nil {
		return err
	}
	sql := `insert into t_meta_mysql_cluster_user_map(mysql_cluster_id, user_id) values(?, ?);`
	log.Debugf("metadata MySQLClusterRepo.AddUser() insert sql: %s", sql)
	_, err = mcr.Execute(sql, mysqlClusterID, userID)

	return err
}

// DeleteUser delete a user for mysql cluster in the middleware
func (mcr *MySQLClusterRepo) DeleteUser(mysqlClusterID, userID int) error {
	userRepo := NewUserRepoWithGlobal()
	_, err := userRepo.GetByID(userID)
	if err != nil {
		return err
	}
	_, err = mcr.GetByID(mysqlClusterID)
	if err != nil {
		return err
	}
	sql := `delete from t_meta_mysql_cluster_user_map where mysql_cluster_id = ? and user_id = ?;`
	log.Debugf("metadata AppRepo.DeleteUser() delete sql: %s", sql)
	_, err = mcr.Execute(sql, mysqlClusterID, userID)

	return err
}

// GetAppUsersByID gets the application users of the given id from the middleware
func (mcr *MySQLClusterRepo) GetAppUsersByID(id int) ([]metadata.User, error) {
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
				 inner join t_meta_app_user_map as aum
							on user.id = aum.user_id
				 inner join t_meta_app_db_map as adm
							on aum.app_id = adm.app_id
				 inner join t_meta_db_info as db
							on db.id = adm.db_id
		where user.del_flag = 0
		  and aum.del_flag = 0
		  and db.del_flag = 0
		  and adm.del_flag = 0
		  and db.cluster_id = ?
		  and db.cluster_type = ?;
	`
	log.Debugf("metadata MySQLClusterRepo.GetAppUsersByID() sql: \n%s\nplaceholders: %d, %d", sql, id, ClusterTypeSingle)

	result, err := mcr.Execute(sql, id, ClusterTypeSingle)
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

// GetDBUsersByID gets the db users of the given id from the middleware
func (mcr *MySQLClusterRepo) GetDBUsersByID(id int) ([]metadata.User, error) {
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
				 inner join t_meta_db_user_map as dum
							on user.id = dum.user_id
				 inner join t_meta_db_info as db
							on dum.db_id = db.id
		where user.del_flag = 0
		  and dum.del_flag = 0
		  and db.del_flag = 0
		  and db.cluster_id = ?
		  and db.cluster_type = ?;
	`
	log.Debugf("metadata MySQLClusterRepo.GetDBUsersByID() sql: \n%s\nplaceholders: %d", sql, id, ClusterTypeSingle)

	result, err := mcr.Execute(sql, id, ClusterTypeSingle)
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

// GetAllUsersByID gets both application and db users of the given id from the middleware
func (mcr *MySQLClusterRepo) GetAllUsersByID(id int) ([]metadata.User, error) {
	sql := `
		select user.id,
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
				 inner join t_meta_app_user_map as aum
							on user.id = aum.user_id
				 inner join t_meta_app_db_map as map
							on aum.app_id = map.app_id
				 inner join t_meta_db_info as db
							on db.id = map.db_id
		where user.del_flag = 0
		  and aum.del_flag = 0
		  and db.del_flag = 0
		  and map.del_flag = 0
		  and db.cluster_id = ?
		  and db.cluster_type = ?
		union
		select user.id,
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
				 inner join t_meta_db_user_map as dum
							on user.id = dum.user_id
				 inner join t_meta_db_info as db
							on dum.db_id = db.id
		where user.del_flag = 0
		  and dum.del_flag = 0
		  and db.del_flag = 0
		  and db.cluster_id = ?
		  and db.cluster_type = ?
		union
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
				 inner join t_meta_mysql_cluster_user_map as cum
							on user.id = cum.user_id
				 inner join t_meta_mysql_cluster_info as cluster
							on cluster.id = cum.mysql_cluster_id
		where user.del_flag = 0
		  and cum.del_flag = 0
		  and cluster.del_flag = 0
		  and cluster.id = ?;
	`
	log.Debugf("metadata MySQLClusterRepo.GetAppUsersByID() sql: \n%s\nplaceholders: %d, %d, %d, %d", sql, id, ClusterTypeSingle, id, ClusterTypeSingle)

	result, err := mcr.Execute(sql, id, ClusterTypeSingle, id, ClusterTypeSingle, id)
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
func (mcr *MySQLClusterRepo) Create(mysqlCluster metadata.MySQLCluster) (metadata.MySQLCluster, error) {
	sql := `
		insert into t_meta_mysql_cluster_info(cluster_name,middleware_cluster_id,
			 monitor_system_id, env_id) 
		values(?,?,?,?);`
	log.Debugf("metadata MySQLClusterRepo.Create() insert sql: %s", sql)
	// execute
	_, err := mcr.Execute(sql,
		mysqlCluster.GetClusterName(),
		mysqlCluster.GetMiddlewareClusterID(),
		mysqlCluster.GetMonitorSystemID(),
		// mysqlCluster.GetOwnerID(),
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
			monitor_system_id = ?, env_id = ?, del_flag = ? 
		where id = ?;`
	log.Debugf("metadata MySQLClusterRepo.Update() update sql: %s", sql)
	mysqlClusterInfo := entity.(*MySQLClusterInfo)
	_, err := mcr.Execute(sql,
		mysqlClusterInfo.ClusterName,
		mysqlClusterInfo.MiddlewareClusterID,
		mysqlClusterInfo.MonitorSystemID,
		// mysqlClusterInfo.OwnerID,
		mysqlClusterInfo.EnvID,
		mysqlClusterInfo.DelFlag, mysqlClusterInfo.ID)

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (mcr *MySQLClusterRepo) Delete(id int) error {
	sql := `delete from t_meta_mysql_cluster_info where id = ?;`
	log.Debugf("metadata MySQLClusterRepo.Delete() delete sql(t_meta_mysql_cluster_info): %s", sql)

	_, err := mcr.Execute(sql, id)
	return err
}
