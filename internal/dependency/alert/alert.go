package alert

import (
	"github.com/romberli/go-util/middleware"
)

type Config interface {
	// Get returns the value of the given key
	Get(key string) string
	// Set sets the value of the given key
	Set(key string, value string)
	// Delete deletes the given key from the config
	Delete(key string)
	// String returns the string value of the config
	String() string
}

type Repository interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	// Save saves alert message to the middleware
	Save(url, toAddrs, ccAddrs, subject, content, config, message string) error
}

type Sender interface {
	// GetConfig return the config
	GetConfig() Config
	// GetURL returns the url
	GetURL() string
	// Send sends the email
	Send() error
}

type Service interface {
	// GetRepository returns the repository of the service
	GetRepository() Repository
	// GetConfig returns the config of the service
	GetConfig() Config
	// SendEmail sends the email
	SendEmail(toAddrs, ccAddrs, subject, content string) error
	// Save saves the email into the middleware
	Save(toAddrs, ccAddrs, subject, content, message string) error
}
