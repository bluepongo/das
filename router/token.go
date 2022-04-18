package router

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
	"github.com/romberli/das/global"
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/das/pkg/message/router"
	"github.com/romberli/das/pkg/resp"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware"
	"github.com/romberli/log"
)

const (
	tokenTokenJSON = "token"
)

type TokenAuth struct {
	Database middleware.Pool
}

func NewTokenAuth(database middleware.Pool) *TokenAuth {
	return newTokenAuth(database)
}

func NewTokenAuthWithGlobal() *TokenAuth {
	return newTokenAuth(global.DASMySQLPool)
}

func newTokenAuth(database middleware.Pool) *TokenAuth {
	return &TokenAuth{database}
}

func (ta *TokenAuth) Execute(command string, args ...interface{}) (middleware.Result, error) {
	conn, err := ta.Database.Get()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Errorf("router TokenAuth.Execute(): close database connection failed.\n%+v", err)
		}
	}()

	return conn.Execute(command, args...)
}

func (ta *TokenAuth) GetTokens() ([]string, error) {
	var tokens []string

	sql := `select token from t_sys_token_info where del_flag = 0;`
	log.Debugf("router TokenAuth.GetTokens() sql: \n%s", sql)

	result, err := ta.Execute(sql)
	if err != nil {
		return nil, err
	}

	for i := constant.ZeroInt; i < result.RowNumber(); i++ {
		token, err := result.GetString(i, constant.ZeroInt)
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}

func (ta *TokenAuth) GetHandlerFunc(tokens []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var fields map[string]string
		// get data
		data, err := c.GetRawData()
		if err != nil {
			resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
			c.Abort()
			return
		}
		// set body back so that body can be read in the router
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		// unmarshal data
		err = json.Unmarshal(data, &fields)
		if err != nil {
			resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
			c.Abort()
			return
		}
		// check if http dody has token field
		token, ok := fields[tokenTokenJSON]
		if !ok || token == constant.EmptyString {
			resp.ResponseNOK(c, message.ErrFieldNotExists, tokenTokenJSON)
			c.Abort()
			return
		}
		if !common.StringInSlice(tokens, token) {
			// not a valid token
			resp.ResponseNOK(c, router.ErrRouterValidateToken, token, c.ClientIP())
			c.Abort()
			return
		}
	}
}
