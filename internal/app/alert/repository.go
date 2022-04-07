package alert

import (
	"github.com/romberli/das/global"
	"github.com/romberli/das/internal/dependency/alert"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/log"
)

var _ alert.Repository = (*Repository)(nil)

type Repository struct {
	Database middleware.Pool
}

// NewRepository returns *Repository with given middleware.Pool
func NewRepository(db middleware.Pool) alert.Repository {
	return newRepository(db)
}

// NewRepository returns *Repository with global mysql pool
func NewRepositoryWithGlobal() alert.Repository {
	return newRepository(global.DASMySQLPool)
}

func newRepository(db middleware.Pool) *Repository {
	return &Repository{Database: db}
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
			log.Errorf("alert DASRepo.Execute(): close database connection failed.\n%s", err)
		}
	}()

	return conn.Execute(command, args...)
}

// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
func (r *Repository) Transaction() (middleware.Transaction, error) {
	return r.Database.Transaction()
}

// Save saves sending result into the middleware
func (r *Repository) Save(url, toAddrs, ccAddrs, subject, content, config, message string) error {
	sql := `
		insert into t_alert_operation_info(url, to_addrs, cc_addrs, subject, content, config, message)
		values(?, ?, ?, ?, ?, ?, ?);
    `
	_, err := r.Execute(sql, url, toAddrs, ccAddrs, subject, content, config, message)

	return err
}
