package query

import (
	"fmt"
	"time"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/app/metadata"
	demetadata "github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/internal/dependency/query"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/go-util/middleware/clickhouse"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
)

const (
	mysqlQueryWithServiceNames = `
		select qc.checksum as sql_id,
			   qc.fingerprint,
			   qe.query    as example,
			   qe.db       as db_name,
			   m.exec_count,
			   m.total_exec_time,
			   m.avg_exec_time,
			   m.rows_examined_max
		from (
				 select qcm.query_class_id,
						sum(qcm.query_count)                                        as exec_count,
						truncate(sum(qcm.query_time_sum), 2)                        as total_exec_time,
						truncate(sum(qcm.query_time_sum) / sum(qcm.query_count), 2) as avg_exec_time,
						qcm.rows_examined_max
				 from query_class_metrics qcm
						  inner join instances i on qcm.instance_id = i.instance_id
				 where i.name in (%s)
				   and qcm.start_ts >= ?
				   and qcm.start_ts < ?
				 group by query_class_id
				 order by rows_examined_max desc
 				 limit ? offset ?) m
				 inner join query_examples qe on m.query_class_id = qe.query_class_id
				 inner join query_classes qc on m.query_class_id = qc.query_class_id
	`
	mysqlQueryWithDBName = `
		select qc.checksum as sql_id,
			   qc.fingerprint,
			   qe.query    as example,
			   qe.db       as db_name,
			   m.exec_count,
			   m.total_exec_time,
			   m.avg_exec_time,
			   m.rows_examined_max
		from (
				 select qcm.query_class_id,
						sum(qcm.query_count)                                        as exec_count,
						truncate(sum(qcm.query_time_sum), 2)                        as total_exec_time,
						truncate(sum(qcm.query_time_sum) / sum(qcm.query_count), 2) as avg_exec_time,
						qcm.rows_examined_max
				 from query_class_metrics qcm
						  inner join instances i on qcm.instance_id = i.instance_id
				 where i.name in (%s)
				   and qcm.start_ts >= ?
				   and qcm.start_ts < ?
				 group by query_class_id
				 order by rows_examined_max desc) m
				 inner join query_examples qe on m.query_class_id = qe.query_class_id
				 inner join query_classes qc on m.query_class_id = qc.query_class_id
		where qe.db = ?
	    limit ? offset ?
	`
	mysqlQueryWithSQLID = `
		select qc.checksum as sql_id,
			   qc.fingerprint,
			   qe.query    as example,
			   qe.db       as db_name,
			   m.exec_count,
			   m.total_exec_time,
			   m.avg_exec_time,
			   m.rows_examined_max
		from (
				 select qcm.query_class_id,
						sum(qcm.query_count)                                        as exec_count,
						truncate(sum(qcm.query_time_sum), 2)                        as total_exec_time,
						truncate(sum(qcm.query_time_sum) / sum(qcm.query_count), 2) as avg_exec_time,
						qcm.rows_examined_max
				 from query_class_metrics qcm
						  inner join instances i on qcm.instance_id = i.instance_id
				 where i.name in (%s)
				   and qcm.query_class_id in (select query_class_id from query_classes where checksum = ?)
				   and qcm.start_ts >= ?
				   and qcm.start_ts < ?
				 group by query_class_id) m
				 inner join query_examples qe on m.query_class_id = qe.query_class_id
				 inner join query_classes qc on m.query_class_id = qc.query_class_id
		where qc.checksum = ?
        order by qe.query_time desc
        limit 1
	`
	clickhouseQueryWithServiceNames = `
		select queryid                                                       as sql_id,
			   fingerprint,
			   (select example from metrics where queryid = queryid limit 1) as example,
			   database                                                      as db_name,
			   sum(num_queries)                                              as exec_count,
			   truncate(sum(m_query_time_sum), 2)                            as total_exec_time,
			   truncate(sum(m_query_time_sum) / sum(num_queries), 2)         as avg_exec_time,
			   max(m_rows_examined_max)                                      as rows_examined_max
		from metrics
		where service_type = 'mysql'
		  and service_name in (%s)
		  and period_start >= ?
		  and period_start < ?
		group by queryid, fingerprint
		order by rows_examined_max desc
		limit ? offset ?
	`
	clickhouseQueryWithDBName = `
		select queryid                                                       as sql_id,
			   fingerprint,
			   (select example from metrics where queryid = queryid limit 1) as example,
			   database                                                      as db_name,
			   sum(num_queries)                                              as exec_count,
			   truncate(sum(m_query_time_sum), 2)                            as total_exec_time,
			   truncate(sum(m_query_time_sum) / sum(num_queries), 2)         as avg_exec_time,
			   max(m_rows_examined_max)                                      as rows_examined_max
		from metrics
		where service_type = 'mysql'
		  and service_name in (%s)
		  and database = ?
		  and period_start >= ?
		  and period_start < ?
		group by queryid, fingerprint
		order by rows_examined_max desc
		limit ? offset ?
	`
	clickhouseQueryWithSQLID = `
		select queryid                                                       as sql_id,
			   fingerprint,
			   (select example from metrics where queryid = queryid limit 1) as example,
			   database                                                      as db_name,
			   sum(num_queries)                                              as exec_count,
			   truncate(sum(m_query_time_sum), 2)                            as total_exec_time,
			   truncate(sum(m_query_time_sum) / sum(num_queries), 2)         as avg_exec_time,
			   max(m_rows_examined_max)                                      as rows_examined_max
		from metrics
		where service_type = 'mysql'
		  and service_name in (%s)
		  and queryid = ?
		  and period_start >= ?
		  and period_start < ?
		group by queryid, fingerprint
		order by rows_examined_max desc
		limit ?, offset ?
	`
)

var _ query.DASRepo = (*DASRepo)(nil)
var _ query.MonitorRepo = (*MySQLRepo)(nil)
var _ query.MonitorRepo = (*ClickhouseRepo)(nil)

type DASRepo struct {
	Database middleware.Pool
}

// NewDASRepo returns *DASRepo
func NewDASRepo(db middleware.Pool) *DASRepo {
	return newDASRepo(db)
}

// NewDASRepoWithGlobal returns *DASRepo with global mysql pool
func NewDASRepoWithGlobal() *DASRepo {
	return NewDASRepo(global.DASMySQLPool)
}

// NewDASRepo returns *DASRepo
func newDASRepo(db middleware.Pool) *DASRepo {
	return &DASRepo{Database: db}
}

// Execute executes given command and placeholders on the middleware
func (r *DASRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := r.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("query DASRepo.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
func (r *DASRepo) Transaction() (middleware.Transaction, error) {
	return r.Database.Transaction()
}

// GetMonitorSystemByDBID returns a metadata.MonitorSystem by Database ID
func (r *DASRepo) GetMonitorSystemByDBID(dbID int) (demetadata.MonitorSystem, error) {
	dbInfo := metadata.NewDBServiceWithDefault()
	err := dbInfo.GetByID(dbID)
	if err != nil {
		return nil, err
	}
	dbs := dbInfo.DBs[constant.ZeroInt]
	clusterID := dbs.GetClusterID()
	return r.GetMonitorSystemByClusterID(clusterID)
}

// GetMonitorSystemByMySQLServerID returns a metadata.MonitorSystem by mysqlServerID
func (r *DASRepo) GetMonitorSystemByMySQLServerID(mysqlServerID int) (demetadata.MonitorSystem, error) {
	serverInfo := metadata.NewMySQLServerServiceWithDefault()
	err := serverInfo.GetByID(mysqlServerID)
	if err != nil {
		return nil, err
	}
	ss := serverInfo.MySQLServers[constant.ZeroInt]
	clusterID := ss.GetClusterID()
	return r.GetMonitorSystemByClusterID(clusterID)
}

// GetMonitorSystemByClusterID returns a metadata.MonitorSystem by clusterID
func (r *DASRepo) GetMonitorSystemByClusterID(clusterID int) (demetadata.MonitorSystem, error) {
	clusterInfo := metadata.NewMySQLClusterServiceWithDefault()
	err := clusterInfo.GetByID(clusterID)
	if err != nil {
		return nil, err
	}
	mcs := clusterInfo.MySQLClusters[constant.ZeroInt]
	monitorSystemID := mcs.GetMonitorSystemID()

	monitorSystemInfo := metadata.NewMonitorSystemServiceWithDefault()
	err = clusterInfo.GetByID(monitorSystemID)
	if err != nil {
		return nil, err
	}
	msi := monitorSystemInfo.MonitorSystems[constant.ZeroInt]
	return msi, nil
}

// Save save dasInfo into table
func (r *DASRepo) Save(mysqlClusterID, mysqlServerID, dbID int, sqlID string, startTime, endTime time.Time, limit, offset int) error {
	sql := `
		insert into t_query_operation_info(mysql_cluster_id, mysql_server_id, db_id, sql_id, start_time, end_time, limit, offset
		values(?, ?, ?, ?, ?, ?, ?, ?);
    `
	_, err := r.Execute(sql, mysqlClusterID, mysqlServerID, dbID, sqlID, startTime.Format(constant.DefaultTimeLayout), endTime.Format(constant.DefaultTimeLayout), limit, offset)

	return err
}

type MySQLRepo struct {
	config *Config
	conn   *mysql.Conn
}

// NewMySQLRepo returns a new mysqlRepo
func NewMySQLRepo(config *Config, conn *mysql.Conn) *MySQLRepo {
	return &MySQLRepo{
		config: config,
		conn:   conn,
	}
}

// getConfig gets Config
func (mr *MySQLRepo) getConfig() *Config {
	return mr.config
}

// Close close the connection
func (mr *MySQLRepo) Close() error {
	return mr.conn.Close()
}

// GetByServiceNames return query.query list by serviceName
func (mr *MySQLRepo) GetByServiceNames(serviceName []string) ([]query.Query, error) {
	interfaces, err := common.ConvertInterfaceToSliceInterface(serviceName)
	if err != nil {
		return nil, err
	}

	services, err := middleware.ConvertSliceToString(interfaces...)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf(mysqlQueryWithServiceNames, services)

	return mr.execute(sql, mr.getConfig().GetStartTime(), mr.getConfig().GetEndTime(), mr.getConfig().GetLimit(), mr.getConfig().GetOffset())
}

// GetByDBName returns query.query list by dbName
func (mr *MySQLRepo) GetByDBName(serviceName, dbName string) ([]query.Query, error) {
	interfaces, err := common.ConvertInterfaceToSliceInterface(serviceName)
	if err != nil {
		return nil, err
	}

	services, err := middleware.ConvertSliceToString(interfaces...)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf(mysqlQueryWithDBName, services)

	return mr.execute(sql, dbName, mr.getConfig().GetStartTime(), mr.getConfig().GetEndTime(), mr.getConfig().GetLimit(), mr.getConfig().GetOffset())
}

// GetBySQLID return query.query by SQL ID
func (mr *MySQLRepo) GetBySQLID(serviceName, sqlID string) (query.Query, error) {
	interfaces, err := common.ConvertInterfaceToSliceInterface(serviceName)
	if err != nil {
		return nil, err
	}

	services, err := middleware.ConvertSliceToString(interfaces...)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf(mysqlQueryWithDBName, services)

	querys, err := mr.execute(sql, sqlID, mr.getConfig().GetStartTime(), mr.getConfig().GetEndTime(), mr.getConfig().GetLimit(), mr.getConfig().GetOffset())
	return querys[constant.ZeroInt], err
}

// execute execute the SQL with args
func (mr *MySQLRepo) execute(command string, args ...interface{}) ([]query.Query, error) {
	log.Debugf("query MySQLRepo.execute() sql: %s, args: %v", command, args)

	// get slow queries from the monitor database
	result, err := mr.conn.Execute(command, args...)
	if err != nil {
		return nil, err
	}
	// init queries
	queries := make([]query.Query, result.RowNumber())
	for i := constant.ZeroInt; i < result.RowNumber(); i++ {
		queries[i] = NewEmptyQuery()
	}
	// map result to queries
	err = result.MapToStructSlice(queries, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return queries, nil
}

type ClickhouseRepo struct {
	config *Config
	conn   *clickhouse.Conn
}

// NewClickHouseRepo return a new ClickHouseRepo
func NewClickHouseRepo(config *Config, conn *clickhouse.Conn) *ClickhouseRepo {
	return &ClickhouseRepo{
		config: config,
		conn:   conn,
	}
}

func (cr *ClickhouseRepo) getConfig() *Config {
	return cr.config
}

func (cr *ClickhouseRepo) Close() error {
	return cr.conn.Close()
}

// GetByServiceNames returns query.query list by serviceNames
func (cr *ClickhouseRepo) GetByServiceNames(serviceNames []string) ([]query.Query, error) {
	interfaces, err := common.ConvertInterfaceToSliceInterface(serviceNames)
	if err != nil {
		return nil, err
	}

	services, err := middleware.ConvertSliceToString(interfaces...)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf(clickhouseQueryWithServiceNames, services)

	return cr.execute(sql, cr.getConfig().GetStartTime(), cr.getConfig().GetEndTime(), cr.getConfig().GetLimit(), cr.getConfig().GetOffset())
}

// GetByDBName returns query.query list by dbNameS
func (cr *ClickhouseRepo) GetByDBName(serviceName, dbName string) ([]query.Query, error) {
	interfaces, err := common.ConvertInterfaceToSliceInterface(serviceName)
	if err != nil {
		return nil, err
	}

	services, err := middleware.ConvertSliceToString(interfaces...)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf(clickhouseQueryWithDBName, services)

	return cr.execute(sql, dbName, cr.getConfig().GetStartTime(), cr.getConfig().GetEndTime(), cr.getConfig().GetLimit(), cr.getConfig().GetOffset())
}

// GetBySQLID return query.query by SQL ID
func (cr *ClickhouseRepo) GetBySQLID(serviceName, sqlID string) (query.Query, error) {
	interfaces, err := common.ConvertInterfaceToSliceInterface(serviceName)
	if err != nil {
		return nil, err
	}

	services, err := middleware.ConvertSliceToString(interfaces...)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf(clickhouseQueryWithDBName, services)

	querys, err := cr.execute(sql, sqlID, cr.getConfig().GetStartTime(), cr.getConfig().GetEndTime(), cr.getConfig().GetLimit(), cr.getConfig().GetOffset())
	return querys[constant.ZeroInt], err
}

func (cr *ClickhouseRepo) execute(command string, args ...interface{}) ([]query.Query, error) {
	log.Debugf("query ClickhouseRepo.execute() sql: %s, args: %v", command, args)

	// get slow queries from the monitor database
	result, err := cr.conn.Execute(command, args...)
	if err != nil {
		return nil, err
	}
	// init queries
	queries := make([]query.Query, result.RowNumber())
	for i := constant.ZeroInt; i < result.RowNumber(); i++ {
		queries[i] = NewEmptyQuery()
	}
	// map result to queries
	err = result.MapToStructSlice(queries, constant.DefaultMiddlewareTag)
	if err != nil {
		return nil, err
	}

	return queries, nil
}
