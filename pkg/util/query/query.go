package query

import (
	"time"

	"github.com/pingcap/errors"
	"github.com/romberli/das/internal/app/query"
	depquery "github.com/romberli/das/internal/dependency/query"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/sql/parser"
)

type MySQLClusterRange struct {
	MySQLClusterID int    `json:"mysql_cluster_id" bind:"required"`
	StartTime      string `json:"start_time" bind:"required"`
	EndTime        string `json:"end_time" bind:"required"`
	Limit          int    `json:"limit" bind:"required"`
	Offset         int    `json:"offset" bind:"required"`
}

func (mcr *MySQLClusterRange) GetMySQLClusterID() int {
	return mcr.MySQLClusterID
}

func (mcr *MySQLClusterRange) GetStartTime() string {
	return mcr.StartTime
}

func (mcr *MySQLClusterRange) GetEndTime() string {
	return mcr.EndTime
}

func (mcr *MySQLClusterRange) GetLimit() int {
	return mcr.Limit
}

func (mcr *MySQLClusterRange) GetOffset() int {
	return mcr.Offset
}

func (mcr *MySQLClusterRange) GetConfig() (depquery.Config, error) {
	return getConfig(mcr.GetStartTime(), mcr.GetEndTime(), mcr.GetLimit(), mcr.GetOffset())
}

type MySQLServerRange struct {
	MySQLServerID int    `json:"mysql_server_id" bind:"required"`
	StartTime     string `json:"start_time" bind:"required"`
	EndTime       string `json:"end_time" bind:"required"`
	Limit         int    `json:"limit" bind:"required"`
	Offset        int    `json:"offset" bind:"required"`
}

func (msr *MySQLServerRange) GetMySQLServerID() int {
	return msr.MySQLServerID
}

func (msr *MySQLServerRange) GetStartTime() string {
	return msr.StartTime
}

func (msr *MySQLServerRange) GetEndTime() string {
	return msr.EndTime
}

func (msr *MySQLServerRange) GetLimit() int {
	return msr.Limit
}

func (msr *MySQLServerRange) GetOffset() int {
	return msr.Offset
}

func (msr *MySQLServerRange) GetConfig() (depquery.Config, error) {
	return getConfig(msr.GetStartTime(), msr.GetEndTime(), msr.GetLimit(), msr.GetOffset())
}

type HostInfoRange struct {
	HostIP    string `json:"host_ip" bind:"required"`
	PortNum   int    `json:"port_num" bind:"required"`
	StartTime string `json:"start_time" bind:"required"`
	EndTime   string `json:"end_time" bind:"required"`
	Limit     int    `json:"limit" bind:"required"`
	Offset    int    `json:"offset" bind:"required"`
}

func (hir *HostInfoRange) GetHostIP() string {
	return hir.HostIP
}

func (hir *HostInfoRange) GetPortNum() int {
	return hir.PortNum
}

func (hir *HostInfoRange) GetStartTime() string {
	return hir.StartTime
}

func (hir *HostInfoRange) GetEndTime() string {
	return hir.EndTime
}

func (hir *HostInfoRange) GetLimit() int {
	return hir.Limit
}

func (hir *HostInfoRange) GetOffset() int {
	return hir.Offset
}

func (hir *HostInfoRange) GetConfig() (depquery.Config, error) {
	return getConfig(hir.GetStartTime(), hir.GetEndTime(), hir.GetLimit(), hir.GetOffset())
}

type DBRange struct {
	DBID      int    `json:"db_id" bind:"required"`
	StartTime string `json:"start_time" bind:"required"`
	EndTime   string `json:"end_time" bind:"required"`
	Limit     int    `json:"limit" bind:"required"`
	Offset    int    `json:"offset" bind:"required"`
}

func (dr *DBRange) GetDBID() int {
	return dr.DBID
}

func (dr *DBRange) GetStartTime() string {
	return dr.StartTime
}

func (dr *DBRange) GetEndTime() string {
	return dr.EndTime
}

func (dr *DBRange) GetLimit() int {
	return dr.Limit
}

func (dr *DBRange) GetOffset() int {
	return dr.Offset
}

func (dr *DBRange) GetConfig() (depquery.Config, error) {
	return getConfig(dr.GetStartTime(), dr.GetEndTime(), dr.GetLimit(), dr.GetOffset())
}

type SQLIDRange struct {
	MySQLServerID int    `json:"mysql_server_id" bind:"required"`
	SQLID         string `json:"sql_id" bind:"required"`
	StartTime     string `json:"start_time" bind:"required"`
	EndTime       string `json:"end_time" bind:"required"`
	Limit         int    `json:"limit" bind:"required"`
	Offset        int    `json:"offset" bind:"required"`
}

func (sir *SQLIDRange) GetMySQLServerID() int {
	return sir.MySQLServerID
}

func (sir *SQLIDRange) GetSQLID() string {
	return sir.SQLID
}

func (sir *SQLIDRange) GetStartTime() string {
	return sir.StartTime
}

func (sir *SQLIDRange) GetEndTime() string {
	return sir.EndTime
}

func (sir *SQLIDRange) GetLimit() int {
	return sir.Limit
}

func (sir *SQLIDRange) GetOffset() int {
	return sir.Offset
}

func (sir *SQLIDRange) GetConfig() (depquery.Config, error) {
	return getConfig(sir.GetStartTime(), sir.GetEndTime(), sir.GetLimit(), sir.GetOffset())
}

func getConfig(startTime, endTime string, limit, offset int) (depquery.Config, error) {
	st, err := time.ParseInLocation(constant.TimeLayoutSecond, startTime, time.Local)
	if err != nil {
		return nil, errors.Trace(err)
	}
	et, err := time.ParseInLocation(constant.TimeLayoutSecond, endTime, time.Local)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return query.NewConfig(st, et, limit, offset), nil
}

func GetDBName(sql string) (string, error) {
	p := parser.NewParserWithDefault()
	r, err := p.Parse(sql)
	if err != nil {
		return constant.EmptyString, err
	}

	if len(r.GetDBNames()) > constant.ZeroInt {
		return r.GetDBNames()[constant.ZeroInt], nil
	}

	return constant.EmptyString, nil
}

func GetTableNames(sql string) ([]string, error) {
	p := parser.NewParserWithDefault()
	r, err := p.Parse(sql)
	if err != nil {
		return nil, err
	}

	return r.GetTableNames(), nil
}
