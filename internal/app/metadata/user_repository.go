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

var _ metadata.UserRepo = (*UserRepo)(nil)

// UserRepo struct
type UserRepo struct {
	Database middleware.Pool
}

// NewUserRepo returns *UserRepo with given middleware.Pool
func NewUserRepo(db middleware.Pool) *UserRepo {
	return &UserRepo{db}
}

// NewUserRepoWithGlobal returns *UserRepo with global mysql pool
func NewUserRepoWithGlobal() *UserRepo {
	return NewUserRepo(global.DASMySQLPool)
}

// Execute implements dependency.UserRepo interface,
// it executes command with arguments on database
// Execute executes given command and placeholders on the middleware
func (ur *UserRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := ur.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata UserRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// GetByName gets users of given username from the middleware
func (ur *UserRepo) GetByName(userName string) ([]metadata.User, error) {
	sql := `
	select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
	from t_meta_user_info
	where del_flag = 0
	and user_name = ?;
`
	log.Debugf("metadata UserRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, userName)

	result, err := ur.Execute(sql, userName)
	if err != nil {
		return nil, err
	}
	// init []*UserInfo
	userInfoList := make([]*UserInfo, result.RowNumber())
	for i := range userInfoList {
		userInfoList[i] = NewEmptyUserInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(userInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []metadata.User
	userList := make([]metadata.User, result.RowNumber())
	for i := range userList {
		userList[i] = userInfoList[i]
	}

	return userList, nil

}

// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
func (ur *UserRepo) Transaction() (middleware.Transaction, error) {
	return ur.Database.Transaction()
}

// GetAll gets all databases from the middleware
func (ur *UserRepo) GetAll() ([]metadata.User, error) {
	sql := `
	select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
	from t_meta_user_info
	where del_flag = 0
	order by id;
	`
	log.Debugf("metadata UserRepo.GetAll() sql: \n%s", sql)

	result, err := ur.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*UserInfo
	userInfoList := make([]*UserInfo, result.RowNumber())
	for i := range userInfoList {
		userInfoList[i] = NewEmptyUserInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(userInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []metadata.User
	userList := make([]metadata.User, result.RowNumber())
	for i := range userList {
		userList[i] = userInfoList[i]
	}

	return userList, nil
}

// GetByTelephone gets a user of given mobile from the middleware
func (ur *UserRepo) GetByMobile(mobile string) (metadata.User, error) {
	sql := `
		select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
		from t_meta_user_info
		where del_flag = 0
		and mobile = ?;
`
	log.Debugf("metadata UserRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, mobile)

	result, err := ur.Execute(sql, mobile)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByMobile(): data does not exists, id: %s", mobile))
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByMobile(): duplicate key exists, id: %s", mobile))
	}
}

// GetByTelephone gets a user of given telephone from the middleware
func (ur *UserRepo) GetByTelephone(telephone string) (metadata.User, error) {
	sql := `
		select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
		from t_meta_user_info
		where del_flag = 0
		and telephone = ?;
`
	log.Debugf("metadata UserRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, telephone)

	result, err := ur.Execute(sql, telephone)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByTelephone(): data does not exists, id: %s", telephone))
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByTelephone(): duplicate key exists, id: %s", telephone))
	}
}

// GetByID gets a user by the identity from the middleware
func (ur *UserRepo) GetByID(id int) (metadata.User, error) {
	sql := `
		select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
		from t_meta_user_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata UserRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := ur.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByID(): data does not exists, id: %d", id))
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByID(): duplicate key exists, id: %d", id))
	}
}

// GetByAccountName gets a user of given account name from the middleware
func (ur *UserRepo) GetByAccountName(accountName string) (metadata.User, error) {
	sql := `
	select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
	from t_meta_user_info
	where del_flag = 0
	and account_name = ?;
`
	log.Debugf("metadata UserRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, accountName)

	result, err := ur.Execute(sql, accountName)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByAccountName(): data does not exists, id: %s", accountName))
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByAccountName(): duplicate key exists, id: %s", accountName))
	}
}

// GetByEmail gets a user of given email from the middleware
func (ur *UserRepo) GetByEmail(email string) (metadata.User, error) {
	sql := `
	select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
	from t_meta_user_info
	where del_flag = 0
	and email = ?;
`
	log.Debugf("metadata UserRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, email)

	result, err := ur.Execute(sql, email)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByEmail(): data does not exists, id: %s", email))
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByEmail(): duplicate key exists, id: %s", email))
	}
}

// GetID gets the identity with given accountName from the middleware
func (ur *UserRepo) GetID(accountName string) (int, error) {
	sql := `select id from t_meta_user_info where del_flag = 0 and account_name = ?;`
	log.Debugf("metadata UserRepo.GetID() select sql: %s", sql)
	result, err := ur.Execute(sql, accountName)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// GetByEmployeeID gets a user of given employee id from the middleware
func (ur *UserRepo) GetByEmployeeID(employeeID string) (metadata.User, error) {
	sql := `
	select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
	from t_meta_user_info
	where del_flag = 0
	and employee_id = ?;
`
	log.Debugf("metadata UserRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, employeeID)

	result, err := ur.Execute(sql, employeeID)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByEmployeeID(): data does not exists, id: %s", employeeID))
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.New(fmt.Sprintf("metadata UserInfo.GetByEmployeeID(): duplicate key exists, id: %s", employeeID))
	}
}

// Create creates a user in the middleware
func (ur *UserRepo) Create(user metadata.User) (metadata.User, error) {
	sql := `insert into t_meta_user_info(user_name, department_name, employee_id, account_name, email , telephone , mobile, role) values(?,?,?,?,?,?,?,?);`
	log.Debugf("metadata UserRepo.Create() insert sql: %s", sql)
	// execute
	userInfo := user.(*UserInfo)
	_, err := ur.Execute(sql, userInfo.UserName, userInfo.DepartmentName, userInfo.EmployeeID, userInfo.AccountName, userInfo.Email, userInfo.Telephone, userInfo.Mobile, userInfo.Role)
	if err != nil {
		return nil, err
	}
	// get id
	id, err := ur.GetID(user.GetAccountName())
	if err != nil {
		return nil, err
	}
	// get user
	return ur.GetByID(id)
}

// Update updates a user in the middleware
func (ur *UserRepo) Update(user metadata.User) error {
	sql := `update t_meta_user_info set user_name = ?, del_flag = ?, department_name = ?, employee_id = ?, account_name = ?, email = ?, telephone = ?, mobile = ?, role = ? where id = ?;`
	log.Debugf("metadata UserRepo.Update() update sql: %s", sql)
	userInfo := user.(*UserInfo)
	_, err := ur.Execute(sql, userInfo.UserName, userInfo.DelFlag, userInfo.DepartmentName, userInfo.EmployeeID, userInfo.AccountName, userInfo.Email, userInfo.Telephone, userInfo.Mobile, userInfo.Role, userInfo.ID)

	return err
}

// Delete deletes the user of given id in the middleware
func (ur *UserRepo) Delete(id int) error {
	sql := `delete from t_meta_user_info where id = ?;`
	log.Debugf("metadata UserRepo.Delete() update sql: %s", sql)

	_, err := ur.Execute(sql, id)

	return err
}

// GetAppsByUserID gets app list that this user owns
func (ur *UserRepo) GetAppsByUserID(userID int) ([]metadata.App, error) {
	sql := `
		select app.id, app.app_name, app.level, app.owner_id, app.del_flag
			, app.create_time, app.last_update_time
		from t_meta_app_info as app
			inner join t_meta_app_user_map as map on app.id = map.app_id
			inner join t_meta_user_info as user on user.id = map.user_id
		where app.del_flag = 0 
			and map.del_flag = 0 
			and user.del_flag = 0
			and user.id = ?;
	`
	log.Debugf("metadata UserRepo.GetAppsByUserID() sql: \n%s\nplaceholders: %d", sql, userID)

	result, err := ur.Execute(sql, userID)
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

// GetDBsByUserID gets db list that this user owns
func (ur *UserRepo) GetDBsByUserID(userID int) ([]metadata.DB, error) {
	sql := `
		select db.id, db.db_name, db.cluster_id, db.cluster_type, db.owner_id, db.env_id, db.del_flag, db.create_time, db.last_update_time
		from t_meta_db_info as db
			inner join t_meta_db_user_map as map on db.id = map.db_id
			inner join t_meta_user_info as user on user.id = map.user_id
		where db.del_flag = 0 
			and map.del_flag = 0 
			and user.del_flag = 0
			and user.id = ?;
	`
	log.Debugf("metadata UserRepo.GetDBsByUserID() sql: \n%s\nplaceholders: %d", sql, userID)

	result, err := ur.Execute(sql, userID)
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

// GetMiddlewareClustersByUserID gets middlewarecluster list that this user owns
func (ur *UserRepo) GetMiddlewareClustersByUserID(userID int) ([]metadata.MiddlewareCluster, error) {
	sql := `
		select middlewarecluster.id, middlewarecluster.cluster_name, middlewarecluster.owner_id, middlewarecluster.env_id, middlewarecluster.del_flag, middlewarecluster.create_time, middlewarecluster.last_update_time
		from t_meta_middleware_cluster_info as middlewarecluster
			inner join t_meta_middleware_cluster_user_map as map on middlewarecluster.id = map.middleware_cluster_id
			inner join t_meta_user_info as user on user.id = map.user_id
		where middlewarecluster.del_flag = 0 
			and map.del_flag = 0 
			and user.del_flag = 0
			and user.id = ?;
	`
	log.Debugf("metadata UserRepo.GetMiddlewareClustersByUserID() sql: \n%s\nplaceholders: %d", sql, userID)

	result, err := ur.Execute(sql, userID)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	middlewareclusterList := make([]metadata.MiddlewareCluster, resultNum)

	for row := 0; row < resultNum; row++ {
		middlewareclusterList[row] = NewEmptyMiddlewareClusterInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(middlewareclusterList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return middlewareclusterList, nil
}

// GetMySQLClustersByUserID gets mysqlcluster list that this user owns
func (ur *UserRepo) GetMySQLClustersByUserID(userID int) ([]metadata.MySQLCluster, error) {
	sql := `
		select mysqlcluster.id, mysqlcluster.cluster_name, mysqlcluster.middleware_cluster_id, mysqlcluster.monitor_system_id, mysqlcluster.owner_id, mysqlcluster.env_id, mysqlcluster.del_flag, mysqlcluster.create_time, mysqlcluster.last_update_time
		from t_meta_mysql_cluster_info as mysqlcluster
			inner join t_meta_mysql_cluster_user_map as map on mysqlcluster.id = map.mysql_cluster_id
			inner join t_meta_user_info as user on user.id = map.user_id
		where mysqlcluster.del_flag = 0 
			and map.del_flag = 0 
			and user.del_flag = 0
			and user.id = ?;
	`
	log.Debugf("metadata UserRepo.GetMySQLClustersByUserID() sql: \n%s\nplaceholders: %d", sql, userID)

	result, err := ur.Execute(sql, userID)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	mysqlclusterList := make([]metadata.MySQLCluster, resultNum)

	for row := 0; row < resultNum; row++ {
		mysqlclusterList[row] = NewEmptyMySQLClusterInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(mysqlclusterList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return mysqlclusterList, nil
}
