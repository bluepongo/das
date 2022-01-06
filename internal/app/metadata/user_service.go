package metadata

import (
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/pkg/message"

	"github.com/romberli/das/internal/dependency/metadata"
)

var _ metadata.UserService = (*UserService)(nil)

const userUsersStruct = "Users"

// UserService struct
type UserService struct {
	metadata.UserRepo
	Users []metadata.User `json:"users"`
	Apps  []metadata.App  `jaon:"apps"`
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

// GetAll gets all users
func (us *UserService) GetAll() error {
	var err error

	us.Users, err = us.UserRepo.GetAll()

	return err
}

// GetApps returns the apps of the service
func (us *UserService) GetApps() []metadata.App {
	return us.Apps
}

// GetByName gets users of given user name
func (us *UserService) GetByName(userName string) error {
	var err error

	us.Users, err = us.UserRepo.GetByName(userName)

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

// GetAppsByID gets apps that this user owns
func (us *UserService) GetAppsByID(userID int) error {
	var err error

	us.Apps, err = us.UserRepo.GetAppsByID(userID)

	return err
}

// AddApp adus a new map of app and user in the middleware
func (us *UserService) AddApp(userID, appID int) error {
	err := us.UserRepo.AddApp(userID, appID)
	if err != nil {
		return err
	}

	return us.GetAppsByID(userID)
}

// DeleteApp deletes the map of app and user in the middleware
func (us *UserService) DeleteApp(userID, appID int) error {
	err := us.UserRepo.DeleteApp(userID, appID)
	if err != nil {
		return err
	}

	return us.GetAppsByID(userID)
}
