package query

import (
	"time"

	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/query"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/go-util/middleware/clickhouse"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
)

var _ query.Repository = (*Repository)(nil)

type Repository struct {
	Database              middleware.Pool
	monitorClickhouseConn *clickhouse.Conn
	monitorMySQLConn      *mysql.Conn
}

// NewRepository returns *Repository with given middleware.Pool
func NewRepository(db middleware.Pool) *Repository {
	return &Repository{Database: db}
}

// NewRepository returns *Repository with global mysql pool
func NewRepositoryWithGlobal() *Repository {
	return NewRepository(global.DASMySQLPool)
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
			log.Errorf("query Repository.Execute(): close database connection failed.\n%s", err.Error())
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
func (r *Repository) Transaction() (middleware.Transaction, error) {
	return r.Database.Transaction()
}

func (r *Repository) GetAll(startTime, endTime time.Time) ([]query.Query, error) {
	return nil, nil
}

func (r *Repository) GetByDBID(dbID int, startTime, endTime time.Time, limit, offset int) ([]query.Query, error) {
	return nil, nil
}

func (r *Repository) GetByID(id, dbID int, startTime, endTime time.Time) (query.Query, error) {
	return nil, nil
}
