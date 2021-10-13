package alert

import (
	"github.com/romberli/das/config"
	"github.com/spf13/viper"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewEmptyResponse returns a new empty *Response
func NewEmptyResponse() *Response {
	return &Response{}
}

// GetCode returns the code
func (r *Response) GetCode() int {
	return r.Code
}

// GetMessage returns the message
func (r *Response) GetMessage() string {
	return r.Message
}

type Config map[string]string

// NewEmptyConfig returns a new empty Config
func NewEmptyConfig() Config {
	return make(map[string]string)
}

// NewConfigFromFile returns a new Config which reads data from config file
func NewConfigFromFile() Config {
	return viper.GetStringMapString(config.AlertHTTPConfigKey)
}

// Get returns the value of the given key
func (c Config) Get(key string) string {
	return c[key]
}

// Set sets the value of the given key
func (c Config) Set(key string, value string) {
	c[key] = value
}

// Delete deletes the given key from the config
func (c Config) Delete(key string) {
	delete(c, key)
}
