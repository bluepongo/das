package global

import (
	"errors"
	"fmt"

	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/spf13/viper"
)

var DASMySQLPool *mysql.Pool

func InitDASMySQLPool() (err error) {
	dbAddr := viper.GetString("db.das.mysql.addr")
	dbName := viper.GetString("db.das.mysql.name")
	dbUser := viper.GetString("db.das.mysql.user")
	dbPass := viper.GetString("db.das.mysql.pass")
	maxConnections := viper.GetInt("db.pool.maxConnections")
	initConnections := viper.GetInt("db.pool.initConnections")
	maxIdleConnections := viper.GetInt("db.pool.maxIdleConnections")
	maxIdleTime := viper.GetInt("db.pool.maxIdleTime")
	keepAliveInterval := viper.GetInt("db.pool.keepAliveInterval")

	config := mysql.NewConfig(dbAddr, dbName, dbUser, dbPass)
	poolConfig := mysql.NewPoolConfigWithConfig(config, maxConnections, initConnections, maxIdleConnections, maxIdleTime, keepAliveInterval)
	log.Debugf("pool config: %v", poolConfig)
	DASMySQLPool, err = mysql.NewPoolWithPoolConfig(poolConfig)
	if err != nil {
		return errors.New(
			fmt.Sprintf("create das mysql pool failed. addr: %s, db: %s, user: %s. error:\n%s",
				dbAddr, dbName, dbUser, err.Error()))
	}

	return nil
}
