package privilege

import (
	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/dependency/privilege"
	"github.com/romberli/das/pkg/message"
	msgpriv "github.com/romberli/das/pkg/message/privilege"
)

var _ privilege.Service = (*Service)(nil)

type Service struct {
	privilege.Repository
	loginName string
}

// NewService returns privilege.Service with given user
func NewService(repo privilege.Repository, loginName string) privilege.Service {
	return newService(repo, loginName)
}

// NewServiceWithDefault returns privilege.Service with default value
func NewServiceWithDefault(loginName string) privilege.Service {
	return newService(NewRepositoryWithGlobal(), loginName)
}

// newService returns privilege.Service with given user
func newService(repo privilege.Repository, loginName string) *Service {
	return &Service{
		Repository: repo,
		loginName:  loginName,
	}
}

// GetLoginName returns the login name
func (s *Service) GetLoginName() string {
	return s.loginName
}

// CheckMySQLServerByID checks if given user has privilege to the mysql server with mysql server id
func (s *Service) CheckMySQLServerByID(mysqlServerID int) error {
	mysqlClusterID, err := s.Repository.GetMySQLClusterIDByMySQLServerID(mysqlServerID)
	if err != nil {
		return err
	}

	ok, err := s.checkPrivilege(mysqlClusterID)
	if err != nil {
		return err
	}
	if !ok {
		return message.NewMessage(msgpriv.ErrPrivilegeNotEnoughPrivilegeByMySQLServerID, mysqlServerID, s.GetLoginName())
	}

	return nil
}

// CheckMySQLServerByHostInfo checks if given user has privilege to the mysql server with host ip and port number
func (s *Service) CheckMySQLServerByHostInfo(hostIP string, portNum int) error {
	mysqlClusterID, err := s.Repository.GetMySQLClusterIDByHostInfo(hostIP, portNum)
	if err != nil {
		return err
	}

	ok, err := s.checkPrivilege(mysqlClusterID)
	if err != nil {
		return err
	}
	if !ok {
		return message.NewMessage(msgpriv.ErrPrivilegeNotEnoughPrivilegeByHostInfo, hostIP, portNum, s.GetLoginName())
	}

	return nil
}

// CheckDBByID checks if given user has privilege to the database with db id
func (s *Service) CheckDBByID(dbID int) error {
	mysqlClusterID, err := s.Repository.GetMySQLClusterIDByDBID(dbID)
	if err != nil {
		return err
	}

	ok, err := s.checkPrivilege(mysqlClusterID)
	if err != nil {
		return err
	}
	if !ok {
		return message.NewMessage(msgpriv.ErrPrivilegeNotEnoughPrivilegeByDBID, dbID, s.GetLoginName())
	}

	return nil
}

func (s *Service) checkPrivilege(mysqlClusterID int) (bool, error) {
	userRole, err := s.Repository.GetUserRoleByLoginName(s.GetLoginName())
	if userRole == config.MetadataUserDBARole || userRole == config.MetadataUserAdminRole {
		return true, nil
	}

	mysqlClusterIDList, err := s.Repository.GetMySQLClusterIDListByLoginName(s.GetLoginName())
	if err != nil {
		return false, err
	}

	for _, clusterID := range mysqlClusterIDList {
		if clusterID == mysqlClusterID {
			return true, nil
		}
	}

	return false, nil
}
