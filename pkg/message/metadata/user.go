package metadata

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initDebugUserMessage()
	initInfoUserMessage()
	initErrorUserMessage()
}

const (
	// debug
	DebugMetadataGetUserAll                    = 101001
	DebugMetadataGetUserByID                   = 101002
	DebugMetadataAddUser                       = 101003
	DebugMetadataUpdateUser                    = 101004
	DebugMetadataGetByUserName                 = 101005
	DebugMetadataGetEmployeeID                 = 101006
	DebugMetadataGetAccountName                = 101007
	DebugMetadataGetEmail                      = 101008
	DebugMetadataGetTelephone                  = 101009
	DebugMetadataGetMobile                     = 101010
	DebugMetadataDeleteUserByID                = 101011
	DebugMetadataGetAppsByUserID               = 101012
	DebugMetadataGetDBsByUserID                = 101013
	DebugMetadataGetMiddlewareClustersByUserID = 101014
	DebugMetadataGetMySQLClustersByUserID      = 101015
	DebugMetadataGetByAccountNameOrEmployeeID  = 101016
	DebugMetadataGetAllMySQLServersByUserID    = 101017
	// info
	InfoMetadataGetUserAll                    = 201001
	InfoMetadataGetUserByID                   = 201002
	InfoMetadataAddUser                       = 201003
	InfoMetadataUpdateUser                    = 201004
	InfoMetadataGetByUserName                 = 201005
	InfoMetadataGetEmployeeID                 = 201006
	InfoMetadataGetAccountName                = 201007
	InfoMetadataGetEmail                      = 201008
	InfoMetadataGetTelephone                  = 201009
	InfoMetadataGetMobile                     = 201010
	InfoMetadataDeleteUserByID                = 201011
	InfoMetadataGetAppsByUserID               = 201012
	InfoMetadataGetDBsByUserID                = 201013
	InfoMetadataGetMiddlewareClustersByUserID = 201014
	InfoMetadataGetMySQLClustersByUserID      = 201015
	InfoMetadataGetByAccountNameOrEmployeeID  = 201016
	InfoMetadataGetAllMySQLServersByUserID    = 201017
	// error
	ErrMetadataGetUserAll                    = 401001
	ErrMetadataGetUserByID                   = 401002
	ErrMetadataAddUser                       = 401003
	ErrMetadataUpdateUser                    = 401004
	ErrMetadataGetByUserName                 = 401005
	ErrMetadataGetEmployeeID                 = 401006
	ErrMetadataGetAccountName                = 401007
	ErrMetadataGetEmail                      = 401008
	ErrMetadataGetTelephone                  = 401009
	ErrMetadataGetMobile                     = 401010
	ErrMetadataDeleteUserByID                = 401011
	ErrMetadataGetAppsByUserID               = 401012
	ErrMetadataGetDBsByUserID                = 401013
	ErrMetadataGetMiddlewareClustersByUserID = 401014
	ErrMetadataGetMySQLClustersByUserID      = 401015
	ErrMetadataGetByAccountNameOrEmployeeID  = 401016
	ErrMetadataGetAllMySQLServersByUserID    = 401017
)

func initDebugUserMessage() {
	message.Messages[DebugMetadataGetUserAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetUserAll, "metadata: get all user. message: %s")
	message.Messages[DebugMetadataGetUserByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetUserByID, "metadata: get user by id. message: %s")
	message.Messages[DebugMetadataAddUser] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddUser, "metadata: add new user. message: %s")
	message.Messages[DebugMetadataUpdateUser] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateUser, "metadata: update user. message: %s")
	message.Messages[DebugMetadataGetByUserName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetByUserName, "metadata: get user by user name. message: %s")
	message.Messages[DebugMetadataGetEmployeeID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetEmployeeID, "metadata: get user by employee id. message: %s")
	message.Messages[DebugMetadataGetAccountName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAccountName, "metadata: get user by account name. message: %s")
	message.Messages[DebugMetadataGetByAccountNameOrEmployeeID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetByAccountNameOrEmployeeID, "metadata: get user by account name or employee id. message: %s")
	message.Messages[DebugMetadataGetEmail] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetEmail, "metadata: get user by email. message: %s")
	message.Messages[DebugMetadataGetTelephone] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetTelephone, "metadata: get user by telephone. message: %s")
	message.Messages[DebugMetadataGetMobile] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMobile, "metadata: get user by mobile. message: %s")
	message.Messages[DebugMetadataDeleteUserByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteUserByID, "metadata: delete user by id. message: %s")
	message.Messages[DebugMetadataGetAppsByUserID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAppsByUserID, "metadata: get app list completed. message: %s")
	message.Messages[DebugMetadataGetDBsByUserID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDBsByUserID, "metadata: get db list completed. message: %s")
	message.Messages[DebugMetadataGetMiddlewareClustersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareClustersByUserID, "metadata: get middleware cluster list completed. message: %s")
	message.Messages[DebugMetadataGetMySQLClustersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMySQLClustersByUserID, "metadata: get mysql cluster list completed. message: %s")
	message.Messages[DebugMetadataGetAllMySQLServersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAllMySQLServersByUserID, "metadata: get all mysql server list completed. message: %s")
}

func initInfoUserMessage() {
	message.Messages[InfoMetadataGetUserAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetUserAll, "metadata: get user all completed")
	message.Messages[InfoMetadataGetUserByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetUserByID, "metadata: get user by id completed. id: %d")
	message.Messages[InfoMetadataAddUser] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddUser, "metadata: add new user completed. user_name: %s")
	message.Messages[InfoMetadataUpdateUser] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateUser, "metadata: update user completed. id: %d")
	message.Messages[InfoMetadataGetByUserName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetByUserName, "metadata: get user by user name completed. user_name: %s")
	message.Messages[InfoMetadataGetEmployeeID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetEmployeeID, "metadata: get user by employee id completed. employee_id: %d")
	message.Messages[InfoMetadataGetAccountName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAccountName, "metadata: get user by account name completed. account_name: %s")
	message.Messages[InfoMetadataGetByAccountNameOrEmployeeID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetByAccountNameOrEmployeeID, "metadata: get user by account name or employee id completed. login_name: %s")
	message.Messages[InfoMetadataGetEmail] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetEmail, "metadata: get user by email completed. email: %s")
	message.Messages[InfoMetadataGetTelephone] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetTelephone, "metadata: get user by telephone completed. telephone: %s")
	message.Messages[InfoMetadataGetMobile] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMobile, "metadata: get user by mobile completed.mobile: %s")
	message.Messages[InfoMetadataDeleteUserByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDeleteUserByID, "metadata: delete user by id completed. id: %d")
	message.Messages[InfoMetadataGetAppsByUserID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAppsByUserID, "metadata: get app list completed. id: %d")
	message.Messages[InfoMetadataGetDBsByUserID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDBsByUserID, "metadata: get db list completed. id: %d")
	message.Messages[InfoMetadataGetMiddlewareClustersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareClustersByUserID, "metadata: get middleware cluster list completed. id: %d")
	message.Messages[InfoMetadataGetMySQLClustersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMySQLClustersByUserID, "metadata: get mysql cluster list completed. id: %d")
	message.Messages[InfoMetadataGetAllMySQLServersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAllMySQLServersByUserID, "metadata: get all mysql server list completed. id: %d")
}

func initErrorUserMessage() {
	message.Messages[ErrMetadataGetUserAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetUserAll, "metadata: get all user failed.")
	message.Messages[ErrMetadataGetUserByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetUserByID, "metadata: get user by id failed. id: %d")
	message.Messages[ErrMetadataAddUser] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddUser, "metadata: add new user failed. user_name: %s")
	message.Messages[ErrMetadataUpdateUser] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateUser, "metadata: update user failed. id: %d")
	message.Messages[ErrMetadataGetByUserName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetByUserName, "metadata: get user by user name failed. user_name: %s")
	message.Messages[ErrMetadataGetEmployeeID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetEmployeeID, "metadata: get user by employee id failed. employee_id: %d")
	message.Messages[ErrMetadataGetAccountName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAccountName, "metadata: get user by account name failed. account_name: %s")
	message.Messages[ErrMetadataGetByAccountNameOrEmployeeID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetByAccountNameOrEmployeeID, "metadata: get user by account name or employee id failed. login_name: %s")
	message.Messages[ErrMetadataGetEmail] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetEmail, "metadata: get user by email failed. email: %s")
	message.Messages[ErrMetadataGetTelephone] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetTelephone, "metadata: get user by telephone failed. telephone: %s")
	message.Messages[ErrMetadataGetMobile] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMobile, "metadata: get user by mobile failed. mobile: %s")
	message.Messages[ErrMetadataDeleteUserByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteUserByID, "metadata: delete user by id failed. id: %d")
	message.Messages[ErrMetadataGetAppsByUserID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAppsByUserID, "metadata: get app list failed. id: %d")
	message.Messages[ErrMetadataGetDBsByUserID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetDBsByUserID, "metadata: get db list failed. id: %d")
	message.Messages[ErrMetadataGetMiddlewareClustersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareClustersByUserID, "metadata: get middleware cluster list failed. id: %d")
	message.Messages[ErrMetadataGetMySQLClustersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMySQLClustersByUserID, "metadata: get mysql cluster list failed. id: %d")
	message.Messages[ErrMetadataGetAllMySQLServersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAllMySQLServersByUserID, "metadata: get all mysql server list failed. id: %d")
}
