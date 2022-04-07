package privilege

import (
	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/privilege"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/log"
)

var _ privilege.Repository = (*Repository)(nil)

type Repository struct {
	Database middleware.Pool
}

// NewRepository returns privilege.Repository with given middleware.Pool
func NewRepository(db middleware.Pool) privilege.Repository {
	return newRepository(db)
}

// NewRepositoryWithGlobal returns privilege.Repository with global mysql pool
func NewRepositoryWithGlobal() privilege.Repository {
	return newRepository(global.DASMySQLPool)
}

// newRepository returns *Repository with given middleware.Pool
func newRepository(db middleware.Pool) privilege.Repository {
	return &Repository{Database: db}
}

// Execute executes given command and placeholders on the middleware
func (r *Repository) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := r.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("sqladvisor DASRepo.Execute(): close database connection failed.\n%+v", err)
		}
	}()

	return conn.Execute(command, args...)
}

// GetMySQLServerClusterIDByLoginName gets mysql cluster id list by login name
func (r *Repository) GetMySQLServerClusterIDListByLoginName(loginName string) ([]int, error) {
	sql := `
		select mci.id
		from t_meta_user_info ui
			inner join t_meta_mysql_cluster_user_map mcum on ui.id = mcum.user_id
			inner join t_meta_mysql_cluster_info mci on mcum.mysql_cluster_id = mci.id
		where ui.del_flag = 0
			and mcum.del_flag = 0
			and mci.del_flag = 0
			and ui.account_name = ?
			or ui.employee_id = ?
		union
		select mci.id
		from t_meta_user_info ui
			inner join t_meta_db_user_map dum on ui.id = dum.user_id
			inner join t_meta_db_info di on dum.db_id = di.id
			inner join t_meta_mysql_cluster_info mci on di.cluster_id = mci.id and di.cluster_type = 1
		where ui.del_flag = 0
			and dum.del_flag = 0
			and di.del_flag = 0
			and ui.account_name = ?
			or ui.employee_id = ?
		union
		select mci.id
		from t_meta_user_info ui
			inner join t_meta_app_user_map aum on ui.id = aum.user_id
			inner join t_meta_app_db_map adm on aum.app_id = adm.app_id
			inner join t_meta_db_info di on adm.db_id = di.id
			inner join t_meta_mysql_cluster_info mci on di.cluster_id = mci.id and di.cluster_type = 1
		where ui.del_flag = 0
			and aum.del_flag = 0
			and adm.del_flag = 0
			and di.del_flag = 0
			and mci.del_flag = 0
			and ui.account_name = ?
			or ui.employee_id = ?
	`
	log.Debugf("privilege Repository.GetMySQLServerClusterIDListByLoginName() sql: \n%s\nplaceholders: %s", sql, loginName)

	result, err := r.Execute(sql, loginName, loginName, loginName, loginName, loginName, loginName)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	mysqlClusterIDs := make([]int, resultNum)
	for i := range mysqlClusterIDs {
		mysqlClusterIDs[i], err = result.GetInt(i, constant.ZeroInt)
		if err != nil {
			return nil, err
		}
	}

	return mysqlClusterIDs, nil
}

// GetMySQLClusterIDByMySQLServerID gets mysql cluster id by mysql server id
func (r *Repository) GetMySQLClusterIDByMySQLServerID(mysqlServerID int) (int, error) {
	sql := `
		select mci.id
		from t_meta_mysql_server_info msi
			inner join t_meta_mysql_cluster_info mci on msi.cluster_id = mci.id
		where msi.del_flag = 0
		  and mci.del_flag = 0
		  and msi.id = ? ;
	`
	log.Debugf("privilege Repository.GetMySQLClusterIDByMySQLServerID() sql: \n%s\nplaceholders: %d", sql, mysqlServerID)

	result, err := r.Execute(sql, mysqlServerID)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// GetMySQLClusterIDByHostInfo gets mysql cluster id by mysql server host ip and port number
func (r *Repository) GetMySQLClusterIDByHostInfo(hostIP string, portNum int) (int, error) {
	sql := `
		select mci.id
		from t_meta_mysql_server_info msi
			inner join t_meta_mysql_cluster_info mci on msi.cluster_id = mci.id
		where msi.del_flag = 0
		  and mci.del_flag = 0
		  and msi.host_ip = ?
		  and msi.port_num = ? ;
	`
	log.Debugf("privilege Repository.GetMySQLClusterIDByMySQLServerID() sql: \n%s\nplaceholders: %s, %d", sql, hostIP, portNum)

	result, err := r.Execute(sql, hostIP, portNum)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// GetMySQLClusterIDByDBID gets mysql cluster id by db id
func (r *Repository) GetMySQLClusterIDByDBID(dbID int) (int, error) {
	sql := `
		select mci.id
		from t_meta_db_info di
			inner join t_meta_mysql_cluster_info mci on di.cluster_id = mci.id and di.cluster_type = 1
		where di.del_flag = 0
		  and mci.del_flag = 0
		  and di.id = ? ;
	`
	log.Debugf("privilege Repository.GetMySQLClusterIDByMySQLServerID() sql: \n%s\nplaceholders: %d", sql, dbID)

	result, err := r.Execute(sql, dbID)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}
