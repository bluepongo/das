package privilege

import (
	"github.com/romberli/das/internal/dependency/metadata"
)

type Service interface {
	// GetUser returns the user
	GetUser() metadata.User
	// CheckMySQLServerByID checks if given user has privilege to the mysql server with mysql server id
	CheckMySQLServerByID(mysqlServerID int) error
	// CheckMySQLServerByHostInfo checks if given user has privilege to the mysql server with host ip and port number
	CheckMySQLServerByHostInfo(hostIP string, portNum int) error
	// CheckDBByID checks if given user has privilege to the database with db id
	CheckDBByID(dbID int) error
	// CheckDBByNameAndClusterInfo checks if given user has privilege to the database with db name, mysql cluster id and mysql cluster type
	CheckDBByNameAndClusterInfo(dbName string, mysqlClusterID, mysqlClusterType int) error
	// CheckDBByNameAndHostInfo checks if given user has privilege to the database with db name, host ip and port number
	CheckDBByNameAndHostInfo(dbName string, hostIP string, portNum int) error
}
