package metadata

import (
	"fmt"

	"github.com/pingcap/errors"
	"github.com/romberli/das/config"
	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/spf13/viper"
)

var _ metadata.MySQLServerRepo = (*MySQLServerRepo)(nil)

// MySQLServerRepo implements dependency.MySQLServerRepo interface
type MySQLServerRepo struct {
	Database middleware.Pool
}

// NewMySQLServerRepo returns *MySQLServerRepo with given middleware.Pool
func NewMySQLServerRepo(db middleware.Pool) *MySQLServerRepo {
	return &MySQLServerRepo{db}
}

// NewMySQLServerRepoWithGlobal returns *MySQLServerRepo with global mysql pool
func NewMySQLServerRepoWithGlobal() *MySQLServerRepo {
	return NewMySQLServerRepo(global.DASMySQLPool)
}

// Execute implements dependency.MySQLServerRepo interface,
// it executes command with arguments on database
func (msr *MySQLServerRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := msr.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("metadata MySQLServerRepo.Execute(): close database connection failed.\n%+v", err)
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns middleware.PoolConn
func (msr *MySQLServerRepo) Transaction() (middleware.Transaction, error) {
	return msr.Database.Transaction()
}

// GetAll returns all available entities
func (msr *MySQLServerRepo) GetAll() ([]metadata.MySQLServer, error) {
	sql := `
		select id, cluster_id, server_name, service_name, host_ip, port_num, deployment_type, version, del_flag, 
			create_time, last_update_time
		from t_meta_mysql_server_info
		where del_flag = 0
		order by id;
	`
	log.Debugf("metadata MySQLServerRepo.GetAll() sql: \n%s", sql)

	result, err := msr.Execute(sql)
	if err != nil {
		return nil, err
	}
	// init []*MySQLServerInfo
	mysqlServerInfoList := make([]*MySQLServerInfo, result.RowNumber())
	for i := range mysqlServerInfoList {
		mysqlServerInfoList[i] = NewEmptyMySQLServerInfoWithGlobal()
	}
	// map to struct
	err = result.MapToStructSlice(mysqlServerInfoList, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}
	// init []metadata.mysqlServer
	mysqlServerList := make([]metadata.MySQLServer, result.RowNumber())
	for i := range mysqlServerList {
		mysqlServerList[i] = mysqlServerInfoList[i]
	}

	return mysqlServerList, nil
}

// GetByClusterID Select returns an available mysqlServer of the given cluster id
func (msr *MySQLServerRepo) GetByClusterID(clusterID int) ([]metadata.MySQLServer, error) {
	sql := `
		select id, cluster_id, server_name, service_name, host_ip, port_num, deployment_type, version, del_flag, 
			create_time, last_update_time
		from t_meta_mysql_server_info 
		where del_flag = 0
		and cluster_id = ?;
	`
	log.Debugf("metadata MySQLServerRepo.GetByClusterID() sql: \n%s\nplaceholders: %s", sql, clusterID)

	result, err := msr.Execute(sql, clusterID)
	if err != nil {
		return nil, err
	}

	resultNum := result.RowNumber()
	mysqlServerList := make([]metadata.MySQLServer, resultNum)

	for row := 0; row < resultNum; row++ {
		mysqlServerID, err := result.GetInt(row, constant.ZeroInt)
		if err != nil {
			return nil, err
		}

		mysqlServer, err := msr.GetByID(mysqlServerID)
		if err != nil {
			return nil, err
		}

		mysqlServerList[row] = mysqlServer
	}

	return mysqlServerList, nil
}

// GetByID Select returns an available mysqlServer of the given id
func (msr *MySQLServerRepo) GetByID(id int) (metadata.MySQLServer, error) {
	sql := `
		select id, cluster_id, server_name, service_name, host_ip, port_num, deployment_type, version, del_flag, 
			create_time, last_update_time
		from t_meta_mysql_server_info
		where del_flag = 0
		and id = ?;
	`
	log.Debugf("metadata MySQLServerRepo.GetByID() sql: \n%s\nplaceholders: %s", sql, id)

	result, err := msr.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.Trace(fmt.Errorf("metadata MySQLServerRepo.GetByID(): data does not exists, id: %d", id))
	case 1:
		mysqlServerInfo := NewEmptyMySQLServerInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(mysqlServerInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return mysqlServerInfo, nil
	default:
		return nil, errors.Trace(fmt.Errorf("metadata MySQLServerRepo.GetByID(): duplicate key exists, id: %d", id))
	}
}

// GetByHostInfo gets a mysql server with given host ip and port number
func (msr *MySQLServerRepo) GetByHostInfo(hostIP string, portNum int) (metadata.MySQLServer, error) {
	sql := `
		select id, cluster_id, server_name, service_name, host_ip, port_num, deployment_type, version, del_flag, 
			create_time, last_update_time
		from t_meta_mysql_server_info
		where del_flag = 0
		and host_ip = ? and port_num = ?;
	`
	log.Debugf("metadata MySQLServerRepo.GetByHostInfo() sql: \n%s\nplaceholders: %s, %d", sql, hostIP, portNum)

	result, err := msr.Execute(sql, hostIP, portNum)
	if err != nil {
		return nil, err
	}
	switch result.RowNumber() {
	case 0:
		return nil, errors.Trace(fmt.Errorf("metadata MySQLServerInfo.GetByHostInfo(): data does not exists, hostIP: %s, portNum: %d", hostIP, portNum))
	case 1:
		mysqlServerInfo := NewEmptyMySQLServerInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(mysqlServerInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return mysqlServerInfo, nil
	default:
		return nil, errors.Trace(fmt.Errorf("metadata MySQLServerInfo.GetByHostInfo(): duplicate key exists, hostIP: %s, portNum: %d", hostIP, portNum))
	}
}

// GetID gets the identity with given host ip and port number from the mysql
func (msr *MySQLServerRepo) GetID(hostIP string, portNum int) (int, error) {
	sql := `select id from t_meta_mysql_server_info where del_flag = 0 and host_ip = ? and port_num = ?;`
	log.Debugf("metadata MySQLServerRepo.GetID() select sql: %s", sql)
	result, err := msr.Execute(sql, hostIP, portNum)
	if err != nil {
		return constant.ZeroInt, err
	}

	return result.GetInt(constant.ZeroInt, constant.ZeroInt)
}

// IsMaster returns if mysql server of given host ip and port number is a master node
func (msr *MySQLServerRepo) IsMaster(hostIP string, portNum int) (bool, error) {
	// create connection
	conn, err := mysql.NewConn(
		fmt.Sprintf("%s:%d", hostIP, portNum),
		constant.EmptyString,
		viper.GetString(config.DBApplicationMySQLUserKey),
		viper.GetString(config.DBApplicationMySQLPassKey),
	)
	if err != nil {
		return false, err
	}

	// check mysql replication
	sql := `show slave status;`
	log.Debugf("metadata MySQLServerRepo.IsMaster() show slave status sql: %s", sql)

	result, err := conn.Execute(sql)
	if err != nil {
		return false, err
	}
	if result.RowNumber() > constant.ZeroInt {
		return false, nil
	}

	// check if mgr primary
	sql = `select member_role from performance_schema.replication_group_members where member_host = ? and member_port = ? and member_role = 'SECONDARY';`
	log.Debugf("metadata MySQLServerRepo.IsMaster() group replication sql: %s, placeholders: %s, %d", sql, hostIP, portNum)

	result, err = conn.Execute(sql, hostIP, portNum)
	if err != nil {
		return false, err
	}

	return result.RowNumber() == constant.ZeroInt, nil
}

// GetMySQLClusterByID gets the mysql cluster of the given id
func (msr *MySQLServerRepo) GetMySQLClusterByID(id int) (metadata.MySQLCluster, error) {
	sql := `
		select c.id, c.cluster_name, c.middleware_cluster_id, c.monitor_system_id,
			c.env_id, c.del_flag, c.create_time, c.last_update_time
		from t_meta_mysql_cluster_info as c inner join t_meta_mysql_server_info as s
		on c.id = s.cluster_id
		where c.del_flag = 0
			and s.del_flag = 0
			and s.id = ?;
	`
	log.Debugf("metadata MySQLServerRepo.GetMySQLClusterByID() sql: \n%s\nplaceholders: %d", sql, id)

	result, err := msr.Execute(sql, id)
	if err != nil {
		return nil, err
	}

	switch result.RowNumber() {
	case 0:
		return nil, errors.Trace(fmt.Errorf("metadata MySQLServerRepo.GetMySQLClusterByID(): data does not exists, id: %d", id))
	case 1:
		mysqlClusterInfo := NewEmptyMySQLClusterInfoWithGlobal()
		// map to struct
		err = result.MapToStructByRowIndex(mysqlClusterInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
		if err != nil {
			return nil, err
		}

		return mysqlClusterInfo, nil
	default:
		return nil, errors.Trace(fmt.Errorf("metadata MySQLServerRepo.GetMySQLClusterByID(): duplicate key exists, id: %d", id))
	}
}

// GetMonitorSystem gets monitor system with given mysql server id from the mysql
func (msr *MySQLServerRepo) GetMonitorSystem(id int) (metadata.MonitorSystem, error) {
	sql := `
		select monsi.id,
			   monsi.system_name,
			   monsi.system_type,
			   monsi.host_ip,
			   monsi.port_num,
			   monsi.port_num_slow,
			   monsi.base_url,
			   monsi.env_id,
			   monsi.del_flag,
			   monsi.create_time,
			   monsi.last_update_time
		from t_meta_mysql_server_info mysi
				 inner join t_meta_mysql_cluster_info mci on mysi.cluster_id = mci.id
				 inner join t_meta_monitor_system_info monsi on mci.monitor_system_id = monsi.id
		where mysi.del_flag = 0
		  and mci.del_flag = 0
		  and monsi.del_flag = 0
		  and mysi.id = ?;
	`
	log.Debugf("metadata MySQLServerRepo.GetMonitorSystem() select sql: %s", sql)
	// execute
	result, err := msr.Execute(sql, id)
	if err != nil {
		return nil, err
	}
	monitorSystemInfo := NewEmptyMonitorSystemInfoWithGlobal()
	err = result.MapToStructByRowIndex(monitorSystemInfo, constant.ZeroInt, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return monitorSystemInfo, nil
}

// Create creates data with given mysqlServer in the middleware
func (msr *MySQLServerRepo) Create(mysqlServer metadata.MySQLServer) (metadata.MySQLServer, error) {
	sql := `
		insert into t_meta_mysql_server_info(
			cluster_id, server_name, service_name, host_ip, port_num, deployment_type, version) 
		values(?, ?, ?, ?, ?, ?, ?);`
	log.Debugf("metadata MySQLServerRepo.Create() insert sql: %s", sql)
	// execute
	_, err := msr.Execute(sql,
		mysqlServer.GetClusterID(),
		mysqlServer.GetServerName(),
		mysqlServer.GetServiceName(),
		mysqlServer.GetHostIP(),
		mysqlServer.GetPortNum(),
		mysqlServer.GetDeploymentType(),
		mysqlServer.GetVersion(),
	)
	if err != nil {
		return nil, err
	}
	// get id
	id, err := msr.GetID(mysqlServer.GetHostIP(), mysqlServer.GetPortNum())
	if err != nil {
		return nil, err
	}
	// get mysqlServer
	return msr.GetByID(id)
}

// Update updates data with given mysqlServer in the middleware
func (msr *MySQLServerRepo) Update(mysqlServer metadata.MySQLServer) error {
	sql := `
		update t_meta_mysql_server_info set 
			cluster_id = ?, server_name = ?, service_name = ?, host_ip = ?, port_num = ?, deployment_type = ?, 
			version = ?, del_flag = ? 
		where id = ?;`
	log.Debugf("metadata MySQLServerRepo.Update() update sql: %s", sql)
	mysqlServerInfo := mysqlServer.(*MySQLServerInfo)
	_, err := msr.Execute(sql,
		mysqlServerInfo.ClusterID,
		mysqlServerInfo.ServerName,
		mysqlServerInfo.ServiceName,
		mysqlServerInfo.HostIP,
		mysqlServerInfo.PortNum,
		mysqlServerInfo.DeploymentType,
		mysqlServerInfo.Version,
		mysqlServerInfo.DelFlag,
		mysqlServerInfo.ID)

	return err
}

// Delete deletes data in the middleware, it is recommended to use soft deletion,
// therefore use update instead of delete
func (msr *MySQLServerRepo) Delete(id int) error {
	sql := `delete from t_meta_mysql_server_info where id = ?;`
	log.Debugf("metadata MySQLServerRepo.Delete() delete sql(t_meta_mysql_server_info): %s", sql)
	_, err := msr.Execute(sql, id)
	return err
}
