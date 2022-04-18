package query

import (
	"time"

	"github.com/pingcap/errors"
	"github.com/romberli/das/internal/app/query"
	depquery "github.com/romberli/das/internal/dependency/query"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/sql/parser"
)

type Range struct {
	StartTime string `json:"start_time" bind:"required"`
	EndTime   string `json:"end_time" bind:"required"`
	Limit     int    `json:"limit" bind:"required"`
	Offset    int    `json:"offset" bind:"required"`
}

func (r *Range) GetStartTime() string {
	return r.StartTime
}

func (r *Range) GetEndTime() string {
	return r.EndTime
}

func (r *Range) GetLimit() int {
	return r.Limit
}

func (r *Range) GetOffset() int {
	return r.Offset
}

func (r *Range) GetConfig() (depquery.Config, error) {
	return getConfig(r.GetStartTime(), r.GetEndTime(), r.GetLimit(), r.GetOffset())
}

type ServerRange struct {
	MySQLServerID int    `json:"mysql_server_id" bind:"required"`
	StartTime     string `json:"start_time" bind:"required"`
	EndTime       string `json:"end_time" bind:"required"`
	Limit         int    `json:"limit" bind:"required"`
	Offset        int    `json:"offset" bind:"required"`
}

func (sr *ServerRange) GetMySQLServerID() int {
	return sr.MySQLServerID
}

func (sr *ServerRange) GetStartTime() string {
	return sr.StartTime
}

func (sr *ServerRange) GetEndTime() string {
	return sr.EndTime
}

func (sr *ServerRange) GetLimit() int {
	return sr.Limit
}

func (sr *ServerRange) GetOffset() int {
	return sr.Offset
}

func (sr *ServerRange) GetConfig() (depquery.Config, error) {
	return getConfig(sr.GetStartTime(), sr.GetEndTime(), sr.GetLimit(), sr.GetOffset())
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
