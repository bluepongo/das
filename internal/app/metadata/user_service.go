package metadata

import (
	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
)

var _ metadata.UserService = (*UserService)(nil)

const userUsersStruct = "Users"

// UserService struct
type UserService struct {
	metadata.UserRepo
	Users              []metadata.User              `json:"users"`
	Apps               []metadata.App               `json:"apps"`
	DBs                []metadata.DB                `json:"dbs"`
	MiddlewareClusters []metadata.MiddlewareCluster `json:"middleware_clusters"`
	MySQLClusters      []metadata.MySQLCluster      `json:"mysql_clusters"`
	MySQLServers       []metadata.MySQLServer       `json:"mysql_servers"`
}

// NewUserService returns a new *UserService
func NewUserService(repo metadata.UserRepo) *UserService {
	return &UserService{UserRepo: repo}
}

// NewUserServiceWithDefault returns a new *UserService with default repository
func NewUserServiceWithDefault() *UserService {
	return NewUserService(NewUserRepoWithGlobal())
}

// GetUsers returns users of the service
func (us *UserService) GetUsers() []metadata.User {
	return us.Users
}

// GetApps returns the apps of the service
func (us *UserService) GetApps() []metadata.App {
	return us.Apps
}

// GetDBs returns the dbs of the service
func (us *UserService) GetDBs() []metadata.DB {
	return us.DBs
}

// GetMiddlewareClusters returns the middleware clusters of the service
func (us *UserService) GetMiddlewareClusters() []metadata.MiddlewareCluster {
	return us.MiddlewareClusters
}

// GetMySQLClusters returns the mysql clusters of the service
func (us *UserService) GetMySQLClusters() []metadata.MySQLCluster {
	return us.MySQLClusters
}

// GetMySQLServers returns the MySQLServers of the service
func (us *UserService) GetMySQLServers() []metadata.MySQLServer {
	return us.MySQLServers
}

// GetAll gets all users
func (us *UserService) GetAll() error {
	var err error

	us.Users, err = us.UserRepo.GetAll()

	return err
}

// GetByID gets a user by the identity
func (us *UserService) GetByID(id int) error {
	user, err := us.UserRepo.GetByID(id)
	if err != nil {
		return err
	}

	us.Users = nil
	us.Users = append(us.Users, user)

	return err
}

// GetByUserName gets users of given user name
func (us *UserService) GetByUserName(userName string) error {
	var err error

	us.Users, err = us.UserRepo.GetByUserName(userName)

	return err
}

// GetByEmployeeID gets a user of given employee id
func (us *UserService) GetByEmployeeID(employeeID string) error {
	user, err := us.UserRepo.GetByEmployeeID(employeeID)
	if err != nil {
		return err
	}

	us.Users = nil
	us.Users = append(us.Users, user)

	return err
}

// GetByAccountName gets a user of given account name
func (us *UserService) GetByAccountName(accountName string) error {
	user, err := us.UserRepo.GetByAccountName(accountName)
	if err != nil {
		return err
	}

	us.Users = nil
	us.Users = append(us.Users, user)

	return err
}

// GetByAccountNameOrEmployeeID gets a user of given login name
func (us *UserService) GetByAccountNameOrEmployeeID(loginName string) error {
	user, err := us.UserRepo.GetByAccountNameOrEmployeeID(loginName)
	if err != nil {
		return err
	}

	us.Users = nil
	us.Users = append(us.Users, user)

	return err
}

// GetByEmail gets a user of given email
func (us *UserService) GetByEmail(email string) error {
	user, err := us.UserRepo.GetByEmail(email)
	if err != nil {
		return err
	}

	us.Users = nil
	us.Users = append(us.Users, user)

	return err
}

// GetByTelephone gets a user of given telephone
func (us *UserService) GetByTelephone(telephone string) error {
	user, err := us.UserRepo.GetByTelephone(telephone)
	if err != nil {
		return err
	}

	us.Users = nil
	us.Users = append(us.Users, user)

	return err
}

// GetByTelephone gets a user of given mobile
func (us *UserService) GetByMobile(mobile string) error {
	user, err := us.UserRepo.GetByMobile(mobile)
	if err != nil {
		return err
	}

	us.Users = nil
	us.Users = append(us.Users, user)

	return err
}

// GetAppsByUserID gets apps that this user owns
func (us *UserService) GetAppsByUserID(userID int) error {
	var err error

	us.Apps, err = us.UserRepo.GetAppsByUserID(userID)

	return err
}

// GetDBsByUserID gets dbs that this user owns
func (us *UserService) GetDBsByUserID(userID int) error {
	var err error

	us.DBs, err = us.UserRepo.GetDBsByUserID(userID)

	return err
}

// GetMiddlewareClustersByUserID gets middleware clusters that this user owns
func (us *UserService) GetMiddlewareClustersByUserID(userID int) error {
	var err error

	us.MiddlewareClusters, err = us.UserRepo.GetMiddlewareClustersByUserID(userID)

	return err
}

// GetMySQLClustersByUserID gets mysql clusters that this user owns
func (us *UserService) GetMySQLClustersByUserID(userID int) error {
	var err error

	us.MySQLClusters, err = us.UserRepo.GetMySQLClustersByUserID(userID)

	return err
}

// GetAllMySQLServersByUserID gets mysql servers that this user owns
func (us *UserService) GetAllMySQLServersByUserID(id int) error {
	var err error

	us.MySQLServers, err = us.UserRepo.GetAllMySQLServersByUserID(id)
	return err
}

// Create creates a user in the middleware
func (us *UserService) Create(fields map[string]interface{}) error {
	// generate new map
	_, ok := fields[userUserNameStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, userUserNameStruct)
	}
	_, ok = fields[userAccountNameStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, userAccountNameStruct)
	}
	_, ok = fields[userEmailStruct]
	if !ok {
		return message.NewMessage(message.ErrFieldNotExists, userEmailStruct)
	}

	// create a new user
	userInfo, err := NewUserInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}

	// insert into middleware
	user, err := us.UserRepo.Create(userInfo)
	if err != nil {
		return err
	}

	us.Users = nil
	us.Users = append(us.Users, user)

	return nil
}

// Update gets a user of the given id from the middleware,
// and then updates its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (us *UserService) Update(id int, fields map[string]interface{}) error {
	err := us.GetByID(id)
	if err != nil {
		return err
	}
	err = us.Users[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return us.UserRepo.Update(us.Users[constant.ZeroInt])
}

// Delete deletes the user of given id in the middleware
func (us *UserService) Delete(id int) error {
	err := us.GetByID(id)
	if err != nil {
		return err
	}

	return us.UserRepo.Delete(id)
}

// Marshal marshals UserService.Users to json bytes
func (us *UserService) Marshal() ([]byte, error) {
	return us.MarshalWithFields(userUsersStruct)
}

// MarshalWithFields marshals only specified fields of the UserService to json bytes
func (us *UserService) MarshalWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(us, fields...)
}
