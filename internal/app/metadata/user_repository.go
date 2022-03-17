package metadata

import (
	"github.com/pingcap/errors"
	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/log"
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
			log.Errorf("metadata UserRepo.Execute(): close database connection failed.\n%+v", err)
		}
	}()

	return conn.Execute(command, args...)
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
		return nil, errors.Errorf("metadata UserInfo.GetByID(): data does not exists, id: %d", id)
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.Errorf("metadata UserInfo.GetByID(): duplicate key exists, id: %d", id)
	}
}

// GetByUserName gets users of given username from the middleware
func (ur *UserRepo) GetByUserName(userName string) ([]metadata.User, error) {
	sql := `
	select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
	from t_meta_user_info
	where del_flag = 0
	and user_name = ?;
`
	log.Debugf("metadata UserRepo.GetByUserName() sql: \n%s\nplaceholders: %s", sql, userName)

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

// GetByEmployeeID gets a user of given employee id from the middleware
func (ur *UserRepo) GetByEmployeeID(employeeID string) (metadata.User, error) {
	sql := `
	select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
	from t_meta_user_info
	where del_flag = 0
	and employee_id = ?;
`
	log.Debugf("metadata UserRepo.GetByEmployeeID() sql: \n%s\nplaceholders: %s", sql, employeeID)

	result, err := ur.Execute(sql, employeeID)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.Errorf("metadata UserInfo.GetByEmployeeID(): data does not exists, id: %s", employeeID)
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.Errorf("metadata UserInfo.GetByEmployeeID(): duplicate key exists, id: %s", employeeID)
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
	log.Debugf("metadata UserRepo.GetByAccountName() sql: \n%s\nplaceholders: %s", sql, accountName)

	result, err := ur.Execute(sql, accountName)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.Errorf("metadata UserInfo.GetByAccountName(): data does not exists, id: %s", accountName)
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.Errorf("metadata UserInfo.GetByAccountName(): duplicate key exists, id: %s", accountName)
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
	log.Debugf("metadata UserRepo.GetByEmail() sql: \n%s\nplaceholders: %s", sql, email)

	result, err := ur.Execute(sql, email)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.Errorf("metadata UserInfo.GetByEmail(): data does not exists, id: %s", email)
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.Errorf("metadata UserInfo.GetByEmail(): duplicate key exists, id: %s", email)
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
	log.Debugf("metadata UserRepo.GetByTelephone() sql: \n%s\nplaceholders: %s", sql, telephone)

	result, err := ur.Execute(sql, telephone)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.Errorf("metadata UserInfo.GetByTelephone(): data does not exists, id: %s", telephone)
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.Errorf("metadata UserInfo.GetByTelephone(): duplicate key exists, id: %s", telephone)
	}
}

// GetByTelephone gets a user of given mobile from the middleware
func (ur *UserRepo) GetByMobile(mobile string) (metadata.User, error) {
	sql := `
		select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
		from t_meta_user_info
		where del_flag = 0
		and mobile = ?;
`
	log.Debugf("metadata UserRepo.GetByMobile() sql: \n%s\nplaceholders: %s", sql, mobile)

	result, err := ur.Execute(sql, mobile)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.Errorf("metadata UserInfo.GetByMobile(): data does not exists, id: %s", mobile)
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.Errorf("metadata UserInfo.GetByMobile(): duplicate key exists, id: %s", mobile)
	}
}

// GetByAccountNameOrEmployeeID gets a user of given loginName from the middleware
func (ur *UserRepo) GetByAccountNameOrEmployeeID(loginName string) (metadata.User, error) {
	sql := `
	select id, user_name, department_name, employee_id, account_name, email, telephone, mobile, role, del_flag, create_time, last_update_time
	from t_meta_user_info
	where del_flag = 0
	and account_name = ? or employee_id = ?;
`
	log.Debugf("metadata UserRepo.GetByAccountNameOrEmployeeID() sql: \n%s\nplaceholders: %s", sql, loginName)

	result, err := ur.Execute(sql, loginName, loginName)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.Errorf("metadata UserInfo.GetByAccountNameOrEmployeeID(): data does not exists, login name: %s", loginName)
	case 1:
		userInfo := NewEmptyUserInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(userInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return userInfo, nil
	default:
		return nil, errors.Errorf("metadata UserInfo.GetByAccountNameOrEmployeeID(): duplicate key exists, login name: %s", loginName)
	}
}

// GetAppsByUserID gets app list that this user owns
func (ur *UserRepo) GetAppsByUserID(userID int) ([]metadata.App, error) {
	sql := `
		select ai.id, ai.app_name, ai.level, ai.del_flag, ai.create_time, ai.last_update_time
		from t_meta_app_info ai
			inner join t_meta_app_user_map aum on ai.id = aum.app_id
			inner join t_meta_user_info ui on ui.id = aum.user_id
		where ai.del_flag = 0
			and aum.del_flag = 0
			and ui.del_flag = 0
			and ui.id = ?;
	`
	log.Debugf("metadata UserRepo.GetAppsByUserID() sql: \n%s\nplaceholders: %d", sql, userID)

	result, err := ur.Execute(sql, userID)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	appList := make([]metadata.App, resultNum)

	for row := constant.ZeroInt; row < resultNum; row++ {
		appList[row] = NewEmptyAppInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(appList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return appList, nil
}

// GetDBsByUserID gets dbs that this user owns
func (ur *UserRepo) GetDBsByUserID(userID int) ([]metadata.DB, error) {
	sql := `
		select di.id, di.db_name, di.cluster_id, di.cluster_type, di.env_id, db.del_flag, di.create_time, di.last_update_time
		from t_meta_db_info di
			inner join t_meta_db_user_map dum on di.id = dum.db_id
			inner join t_meta_user_info ui on ui.id = dum.user_id
		where di.del_flag = 0
			and dum.del_flag = 0
			and ui.del_flag = 0
			and ui.id = ?;
	`
	log.Debugf("metadata UserRepo.GetDBsByUserID() sql: \n%s\nplaceholders: %d", sql, userID)

	result, err := ur.Execute(sql, userID)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	dbList := make([]metadata.DB, resultNum)

	for row := constant.ZeroInt; row < resultNum; row++ {
		dbList[row] = NewEmptyDBInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(dbList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return dbList, nil
}

// GetMiddlewareClustersByUserID gets middleware clusters that this user owns
func (ur *UserRepo) GetMiddlewareClustersByUserID(userID int) ([]metadata.MiddlewareCluster, error) {
	sql := `
		select mci.id, mci.cluster_name, mci.env_id, mci.del_flag, mci.create_time, mci.last_update_time
		from t_meta_middleware_cluster_info mci
			inner join t_meta_middleware_cluster_user_map mcum on mci.id = mcum.middleware_cluster_id
			inner join t_meta_user_info ui on ui.id = mcum.user_id
		where mci.del_flag = 0
			and mcum.del_flag = 0
			and ui.del_flag = 0
			and ui.id = ?;
	`
	log.Debugf("metadata UserRepo.GetMiddlewareClustersByUserID() sql: \n%s\nplaceholders: %d", sql, userID)

	result, err := ur.Execute(sql, userID)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	middlewareClusterList := make([]metadata.MiddlewareCluster, resultNum)

	for i := constant.ZeroInt; i < resultNum; i++ {
		middlewareClusterList[i] = NewEmptyMiddlewareClusterInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(middlewareClusterList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return middlewareClusterList, nil
}

// GetMySQLClustersByUserID gets mysql clusters that this user owns
func (ur *UserRepo) GetMySQLClustersByUserID(userID int) ([]metadata.MySQLCluster, error) {
	sql := `
		select mci.id, mci.cluster_name, mci.middleware_cluster_id, mci.monitor_system_id, mci.env_id, mci.del_flag, mci.create_time, mci.last_update_time
		from t_meta_mysql_cluster_info mci
			inner join t_meta_mysql_cluster_user_map mcum on mci.id = mcum.mysql_cluster_id
			inner join t_meta_user_info ui on ui.id = mcum.user_id
		where mci.del_flag = 0
			and mcum.del_flag = 0
			and ui.del_flag = 0
			and ui.id = ?;
	`
	log.Debugf("metadata UserRepo.GetMySQLClustersByUserID() sql: \n%s\nplaceholders: %d", sql, userID)

	result, err := ur.Execute(sql, userID)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	mysqlClusterList := make([]metadata.MySQLCluster, resultNum)

	for i := constant.ZeroInt; i < resultNum; i++ {
		mysqlClusterList[i] = NewEmptyMySQLClusterInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(mysqlClusterList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return mysqlClusterList, nil
}

// GetAllMySQLServersByUserID gets mysqlserver list that this user owns
func (ur *UserRepo) GetAllMySQLServersByUserID(id int) ([]metadata.MySQLServer, error) {
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
		from t_meta_user_info ui
			inner join t_meta_mysql_cluster_user_map mcum on ui.id = mcum.user_id
			inner join t_meta_mysql_server_info msi on mcum.mysql_cluster_id = msi.cluster_id
		where ui.del_flag = 0
		and mcum.del_flag = 0
		and msi.del_flag = 0
		and ui.id = ?
		union
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
		from t_meta_user_info ui
			inner join t_meta_db_user_map dum on ui.id = dum.user_id
			inner join t_meta_db_info di on dum.db_id = di.id
			inner join t_meta_mysql_server_info msi on di.cluster_id = msi.cluster_id and di.cluster_type = 1
		where ui.del_flag = 0
		and dum.del_flag = 0
		and di.del_flag = 0
		and msi.del_flag = 0
		and ui.id = ?
		union
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
		from t_meta_user_info ui
			inner join t_meta_app_user_map aum on ui.id = aum.user_id
			inner join t_meta_app_db_map adm on aum.app_id = adm.app_id
			inner join t_meta_db_info di on adm.db_id = di.id
			inner join t_meta_mysql_server_info msi on di.cluster_id = msi.cluster_id and di.cluster_type = 1
		where ui.del_flag = 0
		and aum.del_flag = 0
		and adm.del_flag = 0
		and di.del_flag = 0
		and msi.del_flag = 0
		and ui.id = ?;
	`
	log.Debugf("metadata UserRepo.GetAllMySQLServersByUserID() sql: \n%s\nplaceholders: %d", sql, id)

	result, err := ur.Execute(sql, id, id, id)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	mysqlServerList := make([]metadata.MySQLServer, resultNum)

	for i := constant.ZeroInt; i < resultNum; i++ {
		mysqlServerList[i] = NewEmptyMySQLServerInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(mysqlServerList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return mysqlServerList, nil
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
