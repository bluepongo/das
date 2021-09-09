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
				   and qcm.rows_examined_max >= ?
                 group by qcm.query_class_id
                 order by qcm.rows_examined_max desc
                  limit ? offset ?) m
                 inner join query_classes qc on m.query_class_id = qc.query_class_id
                 left join query_examples qe on m.query_class_id = qe.query_class_id;
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
                 		  inner join query_examples qe on qcm.query_class_id = qe.query_class_id
                 where i.name in (%s)
				   and qe.db = ?
                   and qcm.start_ts >= ?
                   and qcm.start_ts < ?
				   and qcm.rows_examined_max >= ?
                 group by qcm.query_class_id
                 order by qcm.rows_examined_max desc
				  limit ? offset ?) m
                 inner join query_classes qc on m.query_class_id = qc.query_class_id
                 left join query_examples qe on m.query_class_id = qe.query_class_id;
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
						  inner join query_classes qc on qcm.query_class_id = qc.query_class_id
                 where i.name in (%s)
				   and qc.checksum = ?
                   and qcm.start_ts >= ?
                   and qcm.start_ts < ?
                 group by query_class_id) m
                 inner join query_classes qc on m.query_class_id = qc.query_class_id
                 left join query_examples qe on m.query_class_id = qe.query_class_id
        limit 1;
    `
	clickhouseQueryWithServiceNames = `
        select sm.sql_id,
               m.fingerprint,
               m.example,
               m.db_name,
               sm.exec_count,
               sm.total_exec_time,
               sm.avg_exec_time,
               sm.rows_examined_max
        
        from (
                 select queryid                                               as sql_id,
                        sum(num_queries)                                      as exec_count,
                        truncate(sum(m_query_time_sum), 2)                    as total_exec_time,
                        truncate(sum(m_query_time_sum) / sum(num_queries), 2) as avg_exec_time,
                        max(m_rows_examined_max)                              as rows_examined_max
                 from metrics
                 where service_type = 'mysql'
                   and service_name in (%s)
                   and period_start >= ?
                   and period_start < ?
                   and m_rows_examined_max >= ?
                 group by queryid
                 order by rows_examined_max desc
                 limit ? offset ? ) sm
                 left join (select queryid          as sql_id,
                                   max(fingerprint) as fingerprint,
                                   max(example)     as example,
                                   max(database)    as db_name
                            from metrics
                            where service_type = 'mysql'
                              and service_name in (%s)
                              and period_start >= ?
                              and period_start < ?
                              and m_rows_examined_max >= ?
                            group by queryid) m
                           on sm.sql_id = m.sql_id;
    `
	clickhouseQueryWithDBName = `
        select sm.sql_id,
               m.fingerprint,
               m.example,
               m.db_name,
               sm.exec_count,
               sm.total_exec_time,
               sm.avg_exec_time,
               sm.rows_examined_max
        
        from (
                 select queryid                                               as sql_id,
                        sum(num_queries)                                      as exec_count,
                        truncate(sum(m_query_time_sum), 2)                    as total_exec_time,
                        truncate(sum(m_query_time_sum) / sum(num_queries), 2) as avg_exec_time,
                        max(m_rows_examined_max)                              as rows_examined_max
                 from metrics
                 where service_type = 'mysql'
                   and service_name in (%s)
                   and (database = ? or schema = ?)
                   and period_start >= ?
                   and period_start < ?
                   and m_rows_examined_max >= ?
                 group by queryid
                 order by rows_examined_max desc
                 limit ? offset ? ) sm
                 left join (select queryid          as sql_id,
                                   max(fingerprint) as fingerprint,
                                   max(example)     as example,
                                   max(database)    as db_name
                            from metrics
                            where service_type = 'mysql'
                              and service_name in (%s)
                              and (database = ? or schema = ?)
                              and period_start >= ?
                              and period_start < ?
                              and m_rows_examined_max >= ?
                            group by queryid) m
                           on sm.sql_id = m.sql_id;
    `
	clickhouseQueryWithSQLID = `
        select sm.sql_id,
               m.fingerprint,
               m.example,
               m.db_name,
               sm.exec_count,
               sm.total_exec_time,
               sm.avg_exec_time,
               sm.rows_examined_max
        
        from (
                 select queryid                                               as sql_id,
                        sum(num_queries)                                      as exec_count,
                        truncate(sum(m_query_time_sum), 2)                    as total_exec_time,
                        truncate(sum(m_query_time_sum) / sum(num_queries), 2) as avg_exec_time,
                        max(m_rows_examined_max)                              as rows_examined_max
                 from metrics
                 where service_type = 'mysql'
                   and service_name in (%s)
                   and queryid = ?
                   and period_start >= ?
                   and period_start < ?
                   and m_rows_examined_max >= ?
                 group by queryid
                 order by rows_examined_max desc
                 limit ? offset ? ) sm
                 left join (select queryid          as sql_id,
                                   max(fingerprint) as fingerprint,
                                   max(example)     as example,
                                   max(database)    as db_name
                            from metrics
                            where service_type = 'mysql'
                              and service_name in (%s)
                              and queryid = ?
                              and period_start >= ?
                              and period_start < ?
                              and m_rows_examined_max >= ?
                            group by queryid) m
                           on sm.sql_id = m.sql_id;
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
func (dr *DASRepo) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := dr.Database.Get()
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
func (dr *DASRepo) Transaction() (middleware.Transaction, error) {
	return dr.Database.Transaction()
}

// GetMonitorSystemByDBID gets a metadata.MonitorSystem by database identity
func (dr *DASRepo) GetMonitorSystemByDBID(dbID int) (demetadata.MonitorSystem, error) {
	dbService := metadata.NewDBServiceWithDefault()
	err := dbService.GetByID(dbID)
	if err != nil {
		return nil, err
	}

	return dr.GetMonitorSystemByClusterID(dbService.GetDBs()[constant.ZeroInt].GetClusterID())
}

// GetMonitorSystemByMySQLServerID gets a metadata.MonitorSystem by mysql server identity
func (dr *DASRepo) GetMonitorSystemByMySQLServerID(mysqlServerID int) (demetadata.MonitorSystem, error) {
	mysqlServerService := metadata.NewMySQLServerServiceWithDefault()
	err := mysqlServerService.GetByID(mysqlServerID)
	if err != nil {
		return nil, err
	}

	return dr.GetMonitorSystemByClusterID(mysqlServerService.GetMySQLServers()[constant.ZeroInt].GetClusterID())
}

// GetMonitorSystemByClusterID gets a metadata.MonitorSystem by mysql cluster identify
func (dr *DASRepo) GetMonitorSystemByClusterID(clusterID int) (demetadata.MonitorSystem, error) {
	mysqlClusterService := metadata.NewMySQLClusterServiceWithDefault()
	err := mysqlClusterService.GetByID(clusterID)
	if err != nil {
		return nil, err
	}

	monitorSystemService := metadata.NewMonitorSystemServiceWithDefault()
	err = monitorSystemService.GetByID(mysqlClusterService.GetMySQLClusters()[constant.ZeroInt].GetMonitorSystemID())
	if err != nil {
		return nil, err
	}

	return monitorSystemService.GetMonitorSystems()[constant.ZeroInt], nil
}

// Save saves dasInfo into table
func (dr *DASRepo) Save(mysqlClusterID, mysqlServerID, dbID int, sqlID string, startTime, endTime time.Time, limit, offset int) error {
	sql := `
		insert into t_query_operation_info(mysql_cluster_id, mysql_server_id, db_id, sql_id, start_time, end_time, limit, offset)
		values(?, ?, ?, ?, ?, ?, ?, ?);
	`

	_, err := dr.Execute(sql, mysqlClusterID, mysqlServerID, dbID, sqlID, startTime.Format(constant.DefaultTimeLayout),
		endTime.Format(constant.DefaultTimeLayout), limit, offset)

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

// Close closes the connection
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

	return mr.execute(sql,
		mr.getConfig().GetStartTime().Format(constant.DefaultTimeLayout),
		mr.getConfig().GetEndTime().Format(constant.DefaultTimeLayout),
		minRowsExamined,
		mr.getConfig().GetLimit(),
		mr.getConfig().GetOffset(),
	)
}

// GetByDBName returns query.query list by dbName
func (mr *MySQLRepo) GetByDBName(serviceName, dbName string) ([]query.Query, error) {
	interfaces, err := common.ConvertInterfaceToSliceInterface([]string{serviceName})
	if err != nil {
		return nil, err
	}

	services, err := middleware.ConvertSliceToString(interfaces...)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf(mysqlQueryWithDBName, services)

	return mr.execute(sql,
		dbName,
		mr.getConfig().GetStartTime().Format(constant.DefaultTimeLayout),
		mr.getConfig().GetEndTime().Format(constant.DefaultTimeLayout),
		minRowsExamined,
		mr.getConfig().GetLimit(),
		mr.getConfig().GetOffset())
}

// GetBySQLID return query.query by SQL ID
func (mr *MySQLRepo) GetBySQLID(serviceName, sqlID string) (query.Query, error) {
	interfaces, err := common.ConvertInterfaceToSliceInterface([]string{serviceName})
	if err != nil {
		return nil, err
	}

	services, err := middleware.ConvertSliceToString(interfaces...)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf(mysqlQueryWithSQLID, services)

	queries, err := mr.execute(sql,
		sqlID,
		mr.getConfig().GetStartTime().Format(constant.DefaultTimeLayout),
		mr.getConfig().GetEndTime().Format(constant.DefaultTimeLayout),
	)
	if len(queries) == 0 {
		return nil, fmt.Errorf("sql(id=%s) in service(name=%s) is not found", sqlID, serviceName)
	}
	return queries[constant.ZeroInt], err
}

// execute executes the SQL with args
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

// NewClickHouseRepo returns a new ClickHouseRepo
func NewClickHouseRepo(config *Config, conn *clickhouse.Conn) *ClickhouseRepo {
	return &ClickhouseRepo{
		config: config,
		conn:   conn,
	}
}

// getConfig returns the configuration
func (cr *ClickhouseRepo) getConfig() *Config {
	return cr.config
}

// Close closes the connection
func (cr *ClickhouseRepo) Close() error {
	return cr.conn.Close()
}

// GetByServiceNames returns query.Query list by serviceNames
func (cr *ClickhouseRepo) GetByServiceNames(serviceNames []string) ([]query.Query, error) {
	interfaces, err := common.ConvertInterfaceToSliceInterface(serviceNames)
	if err != nil {
		return nil, err
	}

	services, err := middleware.ConvertSliceToString(interfaces...)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf(clickhouseQueryWithServiceNames, services, services)

	return cr.execute(
		sql,
		cr.getConfig().GetStartTime(),
		cr.getConfig().GetEndTime(),
		minRowsExamined,
		cr.getConfig().GetLimit(),
		cr.getConfig().GetOffset(),
		cr.getConfig().GetStartTime(),
		cr.getConfig().GetEndTime(),
		minRowsExamined,
	)
}

// GetByDBName returns query.Query list by dbNameS
func (cr *ClickhouseRepo) GetByDBName(serviceName, dbName string) ([]query.Query, error) {
	interfaces, err := common.ConvertInterfaceToSliceInterface([]string{serviceName})
	if err != nil {
		return nil, err
	}

	services, err := middleware.ConvertSliceToString(interfaces...)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf(clickhouseQueryWithDBName, services, services)

	return cr.execute(sql,
		dbName,
		dbName,
		cr.getConfig().GetStartTime(),
		cr.getConfig().GetEndTime(),
		minRowsExamined,
		cr.getConfig().GetLimit(),
		cr.getConfig().GetOffset(),
		dbName,
		dbName,
		cr.getConfig().GetStartTime(),
		cr.getConfig().GetEndTime(),
		minRowsExamined,
	)
}

// GetBySQLID returns query.Query by SQL ID
func (cr *ClickhouseRepo) GetBySQLID(serviceName, sqlID string) (query.Query, error) {
	interfaces, err := common.ConvertInterfaceToSliceInterface([]string{serviceName})
	if err != nil {
		return nil, err
	}

	services, err := middleware.ConvertSliceToString(interfaces...)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf(clickhouseQueryWithSQLID, services, services)

	queries, err := cr.execute(sql,
		sqlID,
		cr.getConfig().GetStartTime(),
		cr.getConfig().GetEndTime(),
		minRowsExamined,
		cr.getConfig().GetLimit(),
		cr.getConfig().GetOffset(),
		sqlID,
		cr.getConfig().GetStartTime(),
		cr.getConfig().GetEndTime(),
		minRowsExamined,
	)
	if len(queries) == 0 {
		return nil, fmt.Errorf("sql(id=%s) in service(name=%s) is not found", sqlID, serviceName)
	}

	return queries[constant.ZeroInt], err
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
