package metadata

import (
	"errors"
	"fmt"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"

	"github.com/romberli/log"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/metadata"
)

var _ metadata.DBRepo = (*DBRepo)(nil)

type DBRepo struct {
	Database middleware.Pool
}

// NewDBRepo returns *DBRepo with given middleware.Pool
func NewDBRepo(db middleware.Pool) *DBRepo {
	return &DBRepo{db}
}

// NewDBRepo returns *DBRepo with global mysql pool
func NewDBRepoWithGlobal() *DBRepo {
	return NewDBRepo(global.DASMySQLPool)
}

// Execute executes given command and placeholders on the middleware
func (dr *DBRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := dr.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata DBRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
func (dr *DBRepo) Transaction() (middleware.Transaction, error) {
	return dr.Database.Transaction()
}

// GetAll gets all databases from the middleware
func (dr *DBRepo) GetAll() ([]metadata.DB, error) {
	sql := `
		select id, db_name, cluster_id, cluster_type, env_id, del_flag, create_time, last_update_time
		from t_meta_db_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata DBRepo.GetAll() sql: \n%s", sql)

	result, err := dr.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*DBInfo
	dbInfoList := make([]*DBInfo, result.RowNumber())
	for i := range dbInfoList {
		dbInfoList[i] = NewEmptyDBInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(dbInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []metadata.DB
	dbList := make([]metadata.DB, result.RowNumber())
	for i := range dbList {
		dbList[i] = dbInfoList[i]
	}

	return dbList, nil
}

// GetByEnv gets databases of given env id from the middleware
func (dr *DBRepo) GetByEnv(envID int) ([]metadata.DB, error) {
	sql := `
		select id, db_name, cluster_id, cluster_type, env_id, del_flag, create_time, last_update_time
		from t_meta_db_info
		where del_flag = 0
		and env_id = ? 
		order by id;
	`
	log.Debugf("metadata DBRepo.GetByEnv sql: \n%s", sql)

	result, err := dr.Execute(sql, envID)
	if err != nil {
		return nil, err
	}
	// init []*DBInfo
	dbInfoList := make([]*DBInfo, result.RowNumber())
	for i := range dbInfoList {
		dbInfoList[i] = NewEmptyDBInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(dbInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []metadata.DB
	dbList := make([]metadata.DB, result.RowNumber())
	for i := range dbList {
		dbList[i] = dbInfoList[i]
	}

	return dbList, nil
}

// GetByID gets a database by the identity from the middleware
func (dr *DBRepo) GetByID(id int) (metadata.DB, error) {
	sql := `
		select id, db_name, cluster_id, cluster_type, env_id, del_flag, create_time, last_update_time
		from t_meta_db_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata DBRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := dr.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata DBInfo.GetByID(): data does not exists, id: %d", id))
	case 1:
		dbInfo := NewEmptyDBInfoWithRepo(dr)
		// map to struct
		err = result.MapToStructByRowIndex(dbInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return dbInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata DBInfo.GetByID(): duplicate key exists, id: %d", id))
	}
}

// GetByNameAndClusterInfo gets a database by the db name and cluster info from the middleware
func (dr *DBRepo) GetByNameAndClusterInfo(name string, clusterID, clusterType int) (metadata.DB, error) {
	sql := `
		select id, db_name, cluster_id, cluster_type, env_id, del_flag, create_time, last_update_time
		from t_meta_db_info
		where del_flag = 0
		and db_name = ?
		and cluster_id = ?
		and cluster_type = ?;
	`
	log.Debugf("metadata DBRepo.GetByNameAndClusterInfo() sql: \n%s\nplaceholders: %s, %d, %d", sql, name, clusterID, clusterType)
	result, err := dr.Execute(sql, name, clusterID, clusterType)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata DBInfo.GetByNameAndClusterInfo(): data does not exists, db name: %s, cluster id: %d, cluster type: %d", name, clusterID, clusterType))
	case 1:
		dbInfo := NewEmptyDBInfoWithRepo(dr)
		// map to struct
		err = result.MapToStructByRowIndex(dbInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return dbInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata DBInfo.GetByNameAndClusterInfo(): duplicate entry exists, db name: %s, cluster id: %d, cluster type: %d", name, clusterID, clusterType))
	}
}

// GetID gets the identity with given database name, cluster id and cluster type from the middleware
func (dr *DBRepo) GetID(dbName string, clusterID int, clusterType int) (int, error) {
	sql := `select id from t_meta_db_info where del_flag = 0 and db_name = ? and cluster_id = ? and cluster_type = ?;`
	log.Debugf("metadata DBRepo.GetID() select sql: %s", sql)
	result, err := dr.Execute(sql, dbName, clusterID, clusterType)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// GetMySQLCLusterByID gets the mysql cluster of the given id from the middleware
func (dr *DBRepo) GetMySQLCLusterByID(id int) (metadata.MySQLCluster, error) {
	sql := `
		select mci.id, mci.cluster_name, mci.middleware_cluster_id, mci.monitor_system_id, 
			mci.env_id, mci.del_flag, mci.create_time, mci.last_update_time
		from t_meta_mysql_cluster_info as mci
		inner join t_meta_db_info as db
		on mci.id = db.cluster_id
		where mci.del_flag = 0 and db.del_flag = 0
		and db.id = ?;
			
	`
	log.Debugf("metadata DBRepo.GetMySQLCLusterByID() sql: \n%s\nplaceholders: %d", sql, id)

	result, err := dr.Execute(sql, id)
	if err != nil {
		return nil, err
	}

	switch result.RowNumber() {
	case 0:
		return nil, fmt.Errorf("metadata DBRepo.GetMySQLCLusterByID(): data does not exists, id: %d", id)
	case 1:
		mysqlClusterInfo := NewEmptyMySQLClusterInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(mysqlClusterInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}
		return mysqlClusterInfo, nil
	default:
		return nil, fmt.Errorf("metadata DBRepo.GetMySQLCLusterByID(): duplicate key exists, id: %d", id)
	}
}

// GetAppsByDBID gets an apps that uses this db
func (dr *DBRepo) GetAppsByDBID(dbID int) ([]metadata.App, error) {
	sql := `
		select app.id, app.app_name, app.level, app.del_flag
			, app.create_time, app.last_update_time
		from t_meta_app_info as app
			inner join t_meta_app_db_map as adm on app.id = adm.app_id
			inner join t_meta_db_info as db on db.id = adm.db_id
		where app.del_flag = 0 
			and adm.del_flag = 0 
			and db.del_flag = 0
			and db.id = ?;
	`
	log.Debugf("metadata DBRepo.GetAppsByDBID() sql: \n%s\nplaceholders: %d", sql, dbID)

	result, err := dr.Execute(sql, dbID)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	appList := make([]metadata.App, resultNum)

	for row := 0; row < resultNum; row++ {
		appList[row] = NewEmptyAppInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(appList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return appList, nil
}

// GetAppUsersByDBID gets the application owners of the given id from the middleware
func (dr *DBRepo) GetAppUsersByDBID(id int) ([]metadata.User, error) {
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
				 inner join t_meta_app_user_map as aum on user.id = aum.user_id
				 inner join t_meta_app_db_map as adm on aum.app_id = adm.app_id
				 inner join t_meta_db_info as db on db.id = adm.db_id
		where user.del_flag = 0
		  and aum.del_flag = 0
		  and db.del_flag = 0
		  and adm.del_flag = 0
		  and db.id = ?;
	`
	log.Debugf("metadata DBRepo.GetAppUsersByDBID() sql: \n%s\nplaceholders: %d, %d", sql, id)

	result, err := dr.Execute(sql, id)
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

// GetUsersByDBID gets the db owners of the given id from the middleware
func (dr *DBRepo) GetUsersByDBID(id int) ([]metadata.User, error) {
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
				 inner join t_meta_db_user_map as dum on user.id = dum.user_id
		where user.del_flag = 0
		  and dum.del_flag = 0
		  and dum.db_id = ?;
	`
	log.Debugf("metadata DBRepo.GetUsersByDBID() sql: \n%s\nplaceholders: %d, %d", sql, id)

	result, err := dr.Execute(sql, id)
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

// GetAllUsersByDBID gets both application and db owners of the given id from the middleware
func (dr *DBRepo) GetAllUsersByDBID(id int) ([]metadata.User, error) {
	sql := `
		select user.id, user.user_name, user.department_name, user.employee_id, user.account_name
			, user.email, user.telephone, user.mobile, user.role, user.del_flag
			, user.create_time, user.last_update_time 
		from t_meta_user_info as user 
			inner join t_meta_app_user_map as aum on user.id = aum.user_id
			inner join t_meta_app_db_map as adm on aum.app_id = adm.app_id
			inner join t_meta_db_info as db on db.id = adm.db_id
		where user.del_flag = 0 
			and aum.del_flag = 0 
			and db.del_flag = 0 
			and adm.del_flag = 0
			and db.id = ? 
		union
		select user.id, user.user_name, user.department_name, user.employee_id, user.account_name
			, user.email, user.telephone, user.mobile, user.role, user.del_flag
			, user.create_time, user.last_update_time 
		from t_meta_user_info as user 
			inner join t_meta_db_user_map as dum on user.id = dum.user_id
		where user.del_flag = 0 
			and dum.del_flag = 0
			and dum.db_id = ?;
	`
	log.Debugf("metadata DBRepo.GetAllUsersByDBID() sql: \n%s\nplaceholders: %d, %d", sql, id, id)

	result, err := dr.Execute(sql, id, id)
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

// Create creates a database in the middleware
func (dr *DBRepo) Create(db metadata.DB) (metadata.DB, error) {
	sql := `insert into t_meta_db_info(db_name, cluster_id, cluster_type, env_id) values(?, ?, ?, ?);`
	log.Debugf("metadata DBRepo.Create() insert sql: %s", sql)

	// execute
	_, err := dr.Execute(sql, db.GetDBName(), db.GetClusterID(), db.GetClusterType(), db.GetEnvID())
	if err != nil {
		return nil, err
	}
	// get id
	id, err := dr.GetID(db.GetDBName(), db.GetClusterID(), db.GetClusterType())
	if err != nil {
		return nil, err
	}
	// get entity
	return dr.GetByID(id)
}

// Update updates the database in the middleware
func (dr *DBRepo) Update(db metadata.DB) error {
	sql := `update t_meta_db_info set db_name = ?, cluster_id = ?, cluster_type = ?, env_id = ?, del_flag = ? where id = ?;`
	log.Debugf("metadata DBRepo.Update() update sql: %s", sql)
	_, err := dr.Execute(sql, db.GetDBName(), db.GetClusterID(), db.GetClusterType(), db.GetEnvID(), db.GetDelFlag(), db.Identity())

	return err
}

// Delete deletes the database in the middleware
func (dr *DBRepo) Delete(id int) error {
	tx, err := dr.Transaction()
	if err != nil {
		return err
	}
	defer func() {
		err = tx.Close()
		if err != nil {
			log.Errorf("metadata DBRepo.Delete(): close database connection failed.\n%s", err.Error())
		}
	}()

	err = tx.Begin()
	if err != nil {
		return err
	}
	sql := `delete from t_meta_db_info where id = ?;`
	log.Debugf("metadata DBRepo.Delete() delete sql(t_meta_db_info): %s", sql)
	_, err = dr.Execute(sql, id)
	if err != nil {
		return err
	}
	sql = `delete from t_meta_app_db_map where db_id = ?;`
	log.Debugf("metadata DBRepo.Delete() delete sql(t_meta_app_db_map): %s", sql)
	_, err = dr.Execute(sql, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// AddApp adds a new map of the app and database in the middleware
func (dr *DBRepo) AddApp(dbID, appID int) error {
	appRepo := NewAppRepoWithGlobal()
	_, err := appRepo.GetByID(appID)
	if err != nil {
		return err
	}
	_, err = dr.GetByID(dbID)
	if err != nil {
		return err
	}
	sql := `insert into t_meta_app_db_map(app_id, db_id) values(?, ?);`
	log.Debugf("metadata DBRepo.AddApp() insert sql: %s", sql)
	_, err = dr.Execute(sql, appID, dbID)

	return err
}

// DeleteApp deletes a map of the app and database in the middleware
func (dr *DBRepo) DeleteApp(dbID, appID int) error {
	appRepo := NewAppRepoWithGlobal()
	_, err := appRepo.GetByID(appID)
	if err != nil {
		return err
	}
	_, err = dr.GetByID(dbID)
	if err != nil {
		return err
	}
	sql := `delete from t_meta_app_db_map where app_id = ? and db_id = ?;`
	log.Debugf("metadata DBRepo.DeleteApp() delete sql: %s", sql)
	_, err = dr.Execute(sql, appID, dbID)

	return err
}

// AddUser adds a new map of app and user in the middleware
func (dr *DBRepo) DBAddUser(dbID, userID int) error {
	userRepo := NewUserRepoWithGlobal()
	_, err := userRepo.GetByID(userID)
	if err != nil {
		return err
	}
	_, err = dr.GetByID(dbID)
	if err != nil {
		return err
	}
	sql := `insert into t_meta_db_user_map(db_id, user_id) values(?, ?);`
	log.Debugf("metadata DBRepo.DBAddUser() insert sql: %s", sql)
	_, err = dr.Execute(sql, dbID, userID)

	return err
}

// DeleteUser delete the map of app and user in the middleware
func (dr *DBRepo) DBDeleteUser(dbID, userID int) error {
	userRepo := NewUserRepoWithGlobal()
	_, err := userRepo.GetByID(userID)
	if err != nil {
		return err
	}
	_, err = dr.GetByID(dbID)
	if err != nil {
		return err
	}
	sql := `delete from t_meta_db_user_map where db_id = ? and user_id = ?;`
	log.Debugf("metadata DBRepo.DBDeleteUser() delete sql: %s", sql)
	_, err = dr.Execute(sql, dbID, userID)

	return err
}
