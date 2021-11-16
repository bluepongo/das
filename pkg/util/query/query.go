package query

import (
	"time"

	"github.com/romberli/das/internal/app/query"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/sql/parser"
)

type Range struct {
	StartTime string `json:"start_time" bind:"required"`
	EndTime   string `json:"end_time" bind:"required"`
	Limit     int    `json:"limit" bind:"required"`
	Offset    int    `json:"offset" bind:"required"`
}

func (r *Range) GetConfig() (*query.Config, error) {
	return getConfig(r.StartTime, r.EndTime, r.Limit, r.Offset)
}

type ServerRange struct {
	MysqlServerID int    `json:"mysql_server_id" bind:"required"`
	StartTime     string `json:"start_time" bind:"required"`
	EndTime       string `json:"end_time" bind:"required"`
	Limit         int    `json:"limit" bind:"required"`
	Offset        int    `json:"offset" bind:"required"`
}

func (sr *ServerRange) GetConfig() (*query.Config, error) {
	return getConfig(sr.StartTime, sr.EndTime, sr.Limit, sr.Offset)
}

func getConfig(startTime, endTime string, limit, offset int) (*query.Config, error) {
	st, err := time.ParseInLocation(constant.TimeLayoutSecond, startTime, time.Local)
	if err != nil {
		return nil, err
	}
	et, err := time.ParseInLocation(constant.TimeLayoutSecond, endTime, time.Local)
	if err != nil {
		return nil, err
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
