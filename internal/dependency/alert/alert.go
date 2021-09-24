package alert

import (
	"github.com/romberli/go-util/middleware"
)

type Repository interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// Save saves alert message to the middleware
	Save(url string, toAddr, ccAddr []string, content string, status int) error
}

type EmailAlerter interface {
	GetURL() string
	GetToAddr() []string
	GetCCAddr() []string
	GetContent() string
	Send() error
}

type Service interface {
	SendEmail(url string, toAddr, ccAddr []string, content string) error
}
