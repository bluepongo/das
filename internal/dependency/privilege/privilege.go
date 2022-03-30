package privilege

import (
	"github.com/romberli/go-util/middleware"
)

type Repository interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// GetMySQLServerClusterIDByLoginName gets mysql cluster id list by login name
	GetMySQLServerClusterIDListByLoginName(loginName string) ([]int, error)
	// GetMySQLClusterIDByMySQLServerID gets mysql cluster id by mysql server id
	GetMySQLClusterIDByMySQLServerID(mysqlServerID int) (int, error)
	// GetMySQLClusterIDByHostInfo gets mysql cluster id by mysql server host ip and port number
	GetMySQLClusterIDByHostInfo(hostIP string, portNum int) (int, error)
	// GetMySQLClusterIDByDBID gets mysql cluster id by db id
	GetMySQLClusterIDByDBID(dbID int) (int, error)
}

type Service interface {
	// GetLoginName returns the login name
	GetLoginName() string
	// CheckMySQLServerByID checks if given user has privilege to the mysql server with mysql server id
	CheckMySQLServerByID(mysqlServerID int) error
	// CheckMySQLServerByHostInfo checks if given user has privilege to the mysql server with host ip and port number
	CheckMySQLServerByHostInfo(hostIP string, portNum int) error
	// CheckDBByID checks if given user has privilege to the database with db id
	CheckDBByID(dbID int) error
}
