package query

import (
	"fmt"

	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/app/metadata"
	depmeta "github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/internal/dependency/query"
	"github.com/romberli/das/pkg/message"
	msgquery "github.com/romberli/das/pkg/message/query"
	"github.com/romberli/go-util/middleware/clickhouse"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/spf13/viper"
)

const (
	pmmMySQLDBName      = "pmm"
	pmmClickhouseDBName = "pmm"
)

var _ query.Query = (*Query)(nil)

type Query struct {
	SQLID           string  `middleware:"sql_id" json:"sql_id"`
	Fingerprint     string  `middleware:"fingerprint" json:"fingerprint"`
	Example         string  `middleware:"example" json:"example"`
	DBName          string  `middleware:"db_name" json:"db_name"`
	ExecCount       int     `middleware:"exec_count" json:"exec_count"`
	TotalExecTime   float64 `middleware:"total_exec_time" json:"total_exec_time"`
	AvgExecTime     float64 `middleware:"avg_exec_time" json:"avg_exec_time"`
	RowsExaminedMax int     `middleware:"rows_examined_max" json:"rows_examined_max"`
}

func NewEmptyQuery() *Query {
	return &Query{}
}

func (q *Query) GetSQLID() string {
	return q.SQLID
}

func (q *Query) GetFingerprint() string {
	return q.Fingerprint
}

func (q *Query) GetExample() string {
	return q.Example
}

func (q *Query) GetDBName() string {
	return q.DBName
}

func (q *Query) GetExecCount() int {
	return q.ExecCount
}

func (q *Query) GetTotalExecTime() float64 {
	return q.TotalExecTime
}

func (q *Query) GetAvgExecTime() float64 {
	return q.AvgExecTime
}

func (q *Query) GetRowsExaminedMax() int {
	return q.RowsExaminedMax
}

type Querier struct {
	config  *Config
	dasRepo *DASRepo
}

func NewQuerier(config *Config, dasRepo *DASRepo) *Querier {
	return newQuerier(config, dasRepo)
}

func NewQuerierWithGlobal(config *Config) *Querier {
	return newQuerier(config, NewDASRepoWithGlobal())
}

func newQuerier(config *Config, dasRepo *DASRepo) *Querier {
	return &Querier{
		config:  config,
		dasRepo: dasRepo,
	}
}

func (q *Querier) getConfig() *Config {
	return q.config
}

func (q *Querier) GetByMySQLClusterID(mysqlClusterID int) ([]query.Query, error) {
	return nil, nil
}

func (q *Querier) GetByMySQLServerID(mysqlServerID int) ([]query.Query, error) {
	// init monitor repos
	monitorRepos, err := q.getMonitorRepo(mysqlServerID)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = monitorRepos.Close()
		if err != nil {
			log.Error(message.NewMessage(msgquery.ErrQueryCloseMonitorRepo, err.Error()).Error())
		}
	}()
	// get mysql server
	mysqlServers, err := q.getMySQLServerByID(mysqlServerID)
	if err != nil {
		return nil, err
	}
	// init service names
	serviceNames := make([]string, len(mysqlServers))
	for i := range mysqlServers {
		serviceNames[i] = mysqlServers[i].GetServiceName()
	}

	return monitorRepos.GetByServiceNames(serviceNames)
}

func (q *Querier) GetByDBID(dbID int) ([]query.Query, error) {
	return nil, nil
}

func (q *Querier) GetBySQLID(mysqlServerID int, sqlID string) ([]query.Query, error) {
	return nil, nil
}

func (q *Querier) getMySQLServerByID(mysqlServerID int) ([]depmeta.MySQLServer, error) {
	service := metadata.NewMySQLServerServiceWithDefault()
	err := service.GetByID(mysqlServerID)
	if err != nil {
		return nil, err
	}

	return service.GetMySQLServers(), nil
}

func (q *Querier) getMySQLServersByClusterID(mysqlClusterID int) ([]depmeta.MySQLServer, error) {
	return nil, nil
}

func (q *Querier) getMonitorSystemByDBID(dbID int) (depmeta.MonitorSystem, error) {
	return nil, nil
}

func (q *Querier) getMonitorSystemByMySQLClusterID(mysqlClusterID int) (depmeta.MonitorSystem, error) {
	return nil, nil
}

func (q *Querier) getMonitorSystemByMySQLServerID(mysqlServerID int) (depmeta.MonitorSystem, error) {
	return nil, nil
}

func (q *Querier) getMonitorRepo(mysqlServerID int) (query.MonitorRepo, error) {
	monitorSystem, err := q.getMonitorSystemByMySQLServerID(mysqlServerID)
	if err != nil {
		return nil, err
	}

	var (
		addr        string
		monitorRepo query.MonitorRepo
	)

	switch monitorSystem.GetSystemType() {
	case 1:
		// pmm 1.x
		addr = fmt.Sprintf("%s:%d", monitorSystem.GetHostIP(), monitorSystem.GetPortNumSlow())
		mysqlConn, err := mysql.NewConn(addr, pmmMySQLDBName, q.getMonitorMySQLUser(), q.getMonitorMySQLPass())
		if err != nil {
			return nil, err
		}
		monitorRepo = NewMySQLRepo(q.getConfig(), mysqlConn)
	case 2:
		// pmm 2.x
		addr = fmt.Sprintf("%s:%d", monitorSystem.GetHostIP(), monitorSystem.GetPortNumSlow())
		clickhouseConn, err := clickhouse.NewConnWithDefault(addr, pmmClickhouseDBName, q.getMonitorClickhouseUser(), q.getMonitorClickhousePass())
		if err != nil {
			return nil, err
		}
		monitorRepo = NewClickHouseRepo(q.getConfig(), clickhouseConn)
	default:
		return nil, message.NewMessage(msgquery.ErrQueryMonitorSystemType, monitorSystem.GetSystemType())
	}

	return monitorRepo, nil
}

// getMonitorClickhouseUser returns clickhouse username of monitor system
func (q *Querier) getMonitorClickhouseUser() string {
	return viper.GetString(config.DBMonitorClickhouseUserKey)
}

// getMonitorClickhousePass returns clickhouse password of monitor system
func (q *Querier) getMonitorClickhousePass() string {
	return viper.GetString(config.DBMonitorClickhousePassKey)
}

// getMonitorMySQLUser returns mysql username of monitor system
func (q *Querier) getMonitorMySQLUser() string {
	return viper.GetString(config.DBMonitorMySQLUserKey)
}

// getMonitorMySQLPass returns mysql password of monitor system
func (q *Querier) getMonitorMySQLPass() string {
	return viper.GetString(config.DBMonitorMySQLPassKey)
}
