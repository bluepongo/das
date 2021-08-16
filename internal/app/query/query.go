package query

import (
	"github.com/romberli/das/internal/dependency/query"
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
