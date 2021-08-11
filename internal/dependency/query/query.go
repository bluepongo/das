package query

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type Query interface {
	GetSQLID() string
	GetFingerprint() string
	GetExample() string
	GetDBName() string
	GetExecCount() int
	GetTotalExecTime() float64
	GetAvgExecTime() float64
	GetRowsExaminedMax() int
}

type Repository interface {
	Execute(command string, args ...interface{}) (middleware.Result, error)
	Transaction() (middleware.Transaction, error)
	GetAll(startTime, endTime time.Time) ([]Query, error)
	GetByDBID(dbID int, startTime, endTime time.Time, limit, offset int) ([]Query, error)
	GetByID(id, dbID int, startTime, endTime time.Time) (Query, error)
}

type Service interface {
	GetQueries() []Query
	GetAll() error
	GetByDBID() error
	GetByID() error
	Marshal() ([]byte, error)
	MarshalWithFields(fields ...string) ([]byte, error)
}
