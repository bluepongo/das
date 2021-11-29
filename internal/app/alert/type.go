package alert

import (
	"fmt"
	"strings"

	"github.com/romberli/das/config"
	"github.com/romberli/go-util/constant"
	"github.com/spf13/viper"
	"github.com/tidwall/pretty"
)

const configContentJSON = "content"

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

// String returns the string value of the config
func (c Config) String() string {
	s := constant.LeftBraceString

	for key, value := range c {
		if key != configContentJSON {
			value = fmt.Sprintf(`"%s"`, value)
		}
		s += fmt.Sprintf(`"%s":%s,`, key, value)
	}

	return string(pretty.Pretty([]byte(strings.Trim(s, constant.CommaString) + constant.RightBraceString)))
}
