package privilege

import (
	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/app/metadata"
	depmeta "github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/internal/dependency/privilege"
	"github.com/romberli/das/pkg/message"
	msgpriv "github.com/romberli/das/pkg/message/privilege"
	"github.com/romberli/go-util/constant"
	"github.com/spf13/viper"
)

var _ privilege.Service = (*Service)(nil)

type Service struct {
	User depmeta.User
}

// NewService returns privilege.Service with given user
func NewService(user depmeta.User) privilege.Service {
	return newService(user)
}

// newService returns privilege.Service with given user
func newService(user depmeta.User) *Service {
	return &Service{
		User: user,
	}
}

// GetUser returns the user
func (s *Service) GetUser() depmeta.User {
	return s.User
}

// CheckMySQLServerByID checks if given user has privilege to the mysql server with mysql server id
func (s *Service) CheckMySQLServerByID(mysqlServerID int) error {
	return s.checkMySQLServerByID(mysqlServerID)
}

// CheckMySQLServerByHostInfo checks if given user has privilege to the mysql server with host ip and port number
func (s *Service) CheckMySQLServerByHostInfo(hostIP string, portNum int) error {
	mysqlServerService := metadata.NewMySQLServerServiceWithDefault()
	err := mysqlServerService.GetByHostInfo(hostIP, portNum)
	if err != nil {
		return err
	}

	return s.checkMySQLServerByID(mysqlServerService.GetMySQLServers()[constant.ZeroInt].Identity())
}

// CheckDBByID checks if given user has privilege to the database with db id
func (s *Service) CheckDBByID(dbID int) error {
	return s.checkDBByID(dbID)
}

// CheckDBByNameAndClusterInfo checks if given user has privilege to the database with db name, mysql cluster id and mysql cluster type
func (s *Service) CheckDBByNameAndClusterInfo(dbName string, mysqlClusterID, mysqlClusterType int) error {
	dbService := metadata.NewDBServiceWithDefault()
	err := dbService.GetDBByNameAndClusterInfo(dbName, mysqlClusterID, mysqlClusterType)
	if err != nil {
		return err
	}

	return s.checkDBByID(dbService.GetDBs()[constant.ZeroInt].Identity())
}

// CheckDBByNameAndHostInfo checks if given user has privilege to the database with db name, host ip and port number
func (s *Service) CheckDBByNameAndHostInfo(dbName string, hostIP string, portNum int) error {
	dbService := metadata.NewDBServiceWithDefault()
	err := dbService.GetDBByNameAndHostInfo(dbName, hostIP, portNum)
	if err != nil {
		return err
	}

	return s.checkDBByID(dbService.GetDBs()[constant.ZeroInt].Identity())
}

// checkMySQLServerByID checks if given user has privilege to the mysql server with mysql server id
func (s *Service) checkMySQLServerByID(mysqlServerID int) error {
	if !viper.GetBool(config.PrivilegeEnabledKey) {
		return nil
	}

	if s.GetUser().GetRole() >= config.MetadataUserDBARole {
		// this user is dba or admin
		return nil
	}
	// get all mysql servers
	mysqlServerList, err := s.GetUser().GetAllMySQLServers()
	if err != nil {
		return err
	}
	for _, mysqlServer := range mysqlServerList {
		if mysqlServer.Identity() == mysqlServerID {
			// user has the privilege to the given mysql
			return nil
		}
	}

	return message.NewMessage(msgpriv.ErrPrivilegeNotEnoughPrivilege, s.GetUser().GetUserName(), s.GetUser().GetAccountName(), mysqlServerID)
}

// checkDBByID checks if given user has privilege to the database with db id
func (s *Service) checkDBByID(dbID int) error {
	dbService := metadata.NewDBServiceWithDefault()
	err := dbService.GetByID(dbID)
	if err != nil {
		return err
	}

	mysqlClusterService := metadata.NewMySQLClusterServiceWithDefault()
	err = mysqlClusterService.GetMasterServersByID(dbService.GetDBs()[constant.ZeroInt].GetClusterID())
	if err != nil {
		return err
	}

	return s.checkMySQLServerByID(mysqlClusterService.GetMySQLServers()[constant.ZeroInt].Identity())
}
