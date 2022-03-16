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
	DebugMetadataGetUserAll                    = 100901
	DebugMetadataGetUserByID                   = 100902
	DebugMetadataAddUser                       = 100903
	DebugMetadataUpdateUser                    = 100904
	DebugMetadataGetByUserName                 = 100905
	DebugMetadataGetEmployeeID                 = 100906
	DebugMetadataGetAccountName                = 100907
	DebugMetadataGetEmail                      = 100908
	DebugMetadataGetTelephone                  = 100909
	DebugMetadataGetMobile                     = 100910
	DebugMetadataDeleteUserByID                = 100911
	DebugMetadataGetAppsByUserID               = 100912
	DebugMetadataGetDBsByUserID                = 100913
	DebugMetadataGetMiddlewareClustersByUserID = 100914
	DebugMetadataGetMySQLClustersByUserID      = 100915
	DebugMetadataGetByAccountNameOrEmployeeID  = 100916
	DebugMetadataGetAllMySQLServersByUserID    = 100917
	// info
	InfoMetadataGetUserAll                    = 200901
	InfoMetadataGetUserByID                   = 200902
	InfoMetadataAddUser                       = 200903
	InfoMetadataUpdateUser                    = 200904
	InfoMetadataGetByUserName                 = 200905
	InfoMetadataGetEmployeeID                 = 200906
	InfoMetadataGetAccountName                = 200907
	InfoMetadataGetEmail                      = 200908
	InfoMetadataGetTelephone                  = 200909
	InfoMetadataGetMobile                     = 200910
	InfoMetadataDeleteUserByID                = 200911
	InfoMetadataGetAppsByUserID               = 200912
	InfoMetadataGetDBsByUserID                = 200913
	InfoMetadataGetMiddlewareClustersByUserID = 200914
	InfoMetadataGetMySQLClustersByUserID      = 200915
	InfoMetadataGetByAccountNameOrEmployeeID  = 200916
	InfoMetadataGetAllMySQLServersByUserID    = 200917
	// error
	ErrMetadataGetUserAll                    = 400901
	ErrMetadataGetUserByID                   = 400902
	ErrMetadataAddUser                       = 400903
	ErrMetadataUpdateUser                    = 400904
	ErrMetadataGetByUserName                 = 400905
	ErrMetadataGetEmployeeID                 = 400906
	ErrMetadataGetAccountName                = 400907
	ErrMetadataGetEmail                      = 400908
	ErrMetadataGetTelephone                  = 400909
	ErrMetadataGetMobile                     = 400910
	ErrMetadataDeleteUserByID                = 400911
	ErrMetadataGetAppsByUserID               = 400912
	ErrMetadataGetDBsByUserID                = 400913
	ErrMetadataGetMiddlewareClustersByUserID = 400914
	ErrMetadataGetMySQLClustersByUserID      = 400915
	ErrMetadataGetByAccountNameOrEmployeeID  = 400916
	ErrMetadataGetAllMySQLServersByUserID    = 400917
)

func initDebugUserMessage() {
	message.Messages[DebugMetadataGetUserAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetUserAll, "metadata: get all user message: %s")
	message.Messages[DebugMetadataGetUserByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetUserByID, "metadata: get user by id message: %s")
	message.Messages[DebugMetadataAddUser] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddUser, "metadata: add new user message: %s")
	message.Messages[DebugMetadataUpdateUser] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateUser, "metadata: update user message: %s")
	message.Messages[DebugMetadataGetByUserName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetByUserName, "metadata: get user by user name message: %s")
	message.Messages[DebugMetadataGetEmployeeID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetEmployeeID, "metadata: get user by employee id message: %s")
	message.Messages[DebugMetadataGetAccountName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAccountName, "metadata: get user by account name message: %s")
	message.Messages[DebugMetadataGetByAccountNameOrEmployeeID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetByAccountNameOrEmployeeID, "metadata: get user by account name or employee id message: %s")
	message.Messages[DebugMetadataGetEmail] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetEmail, "metadata: get user by email message: %s")
	message.Messages[DebugMetadataGetTelephone] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetTelephone, "metadata: get user by telephone message: %s")
	message.Messages[DebugMetadataGetMobile] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMobile, "metadata: get user by mobile message: %s")
	message.Messages[DebugMetadataDeleteUserByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteUserByID, "metadata: delete user by ID message: %s")
	message.Messages[DebugMetadataGetAppsByUserID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAppsByUserID, "metadata: get app list completed. message: %s")
	message.Messages[DebugMetadataGetDBsByUserID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDBsByUserID, "metadata: get db list completed. message: %s")
	message.Messages[DebugMetadataGetMiddlewareClustersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareClustersByUserID, "metadata: get middleware cluster list completed. message: %s")
	message.Messages[DebugMetadataGetMySQLClustersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMySQLClustersByUserID, "metadata: get mysql cluster list completed. message: %s")
	message.Messages[DebugMetadataGetAllMySQLServersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAllMySQLServersByUserID, "metadata: get mysql server list completed. message: %s")
}

func initInfoUserMessage() {
	message.Messages[InfoMetadataGetUserAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetUserAll, "metadata: get user all completed")
	message.Messages[InfoMetadataGetUserByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetUserByID, "metadata: get user by id completed. id: %d")
	message.Messages[InfoMetadataAddUser] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddUser, "metadata: add new user completed. user_name: %s")
	message.Messages[InfoMetadataUpdateUser] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateUser, "metadata: update user completed. id: %d")
	message.Messages[InfoMetadataGetByUserName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetByUserName, "metadata: get user by user name completed. Name: %s")
	message.Messages[InfoMetadataGetEmployeeID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetEmployeeID, "metadata: get user by employee id completed. employID: %d")
	message.Messages[InfoMetadataGetAccountName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAccountName, "metadata: get user by account name completed. accountName: %s")
	message.Messages[InfoMetadataGetByAccountNameOrEmployeeID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetByAccountNameOrEmployeeID, "metadata: get user by account name or employee id completed. loginName: %s")
	message.Messages[InfoMetadataGetEmail] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetEmail, "metadata: get user by email completed. email: %s")
	message.Messages[InfoMetadataGetTelephone] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetTelephone, "metadata: get user by telephone completed. telephone: %s")
	message.Messages[InfoMetadataGetMobile] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMobile, "metadata: get user by mobile completed.mobile: %s")
	message.Messages[InfoMetadataDeleteUserByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDeleteUserByID, "metadata: delete user by ID completed. id: %d")
	message.Messages[InfoMetadataGetAppsByUserID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAppsByUserID, "metadata: get app list completed. id: %d")
	message.Messages[InfoMetadataGetDBsByUserID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDBsByUserID, "metadata: get db list completed. id: %d")
	message.Messages[InfoMetadataGetMiddlewareClustersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareClustersByUserID, "metadata: get middleware cluster list completed. id: %d")
	message.Messages[InfoMetadataGetMySQLClustersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMySQLClustersByUserID, "metadata: get mysql cluster list completed. id: %d")
	message.Messages[InfoMetadataGetAllMySQLServersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAllMySQLServersByUserID, "metadata: get mysql server list completed. id: %d")
}

func initErrorUserMessage() {
	message.Messages[ErrMetadataGetUserAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetUserAll, "metadata: get all user failed.")
	message.Messages[ErrMetadataGetUserByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetUserByID, "metadata: get user by id failed. id: %d")
	message.Messages[ErrMetadataAddUser] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddUser, "metadata: add new user failed. user_name: %s")
	message.Messages[ErrMetadataUpdateUser] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateUser, "metadata: update user failed. id: %d")
	message.Messages[ErrMetadataGetByUserName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetByUserName, "metadata: get user by user name failed.Name: %s")
	message.Messages[ErrMetadataGetEmployeeID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetEmployeeID, "metadata: get user by employee id failed. employID: %d")
	message.Messages[ErrMetadataGetAccountName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAccountName, "metadata: get user by account name failed. accountName: %s")
	message.Messages[ErrMetadataGetByAccountNameOrEmployeeID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetByAccountNameOrEmployeeID, "metadata: get user by account name or employee id failed. loginName: %s")
	message.Messages[ErrMetadataGetEmail] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetEmail, "metadata: get user by email failed. email: %s")
	message.Messages[ErrMetadataGetTelephone] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetTelephone, "metadata: get user by telephone failed. telephone: %s")
	message.Messages[ErrMetadataGetMobile] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMobile, "metadata: get user by mobile failed. mobile: %s")
	message.Messages[ErrMetadataDeleteUserByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteUserByID, "metadata: delete user by ID failed. id: %d")
	message.Messages[ErrMetadataGetAppsByUserID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAppsByUserID, "metadata: get app list failed. id: %d")
	message.Messages[ErrMetadataGetDBsByUserID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetDBsByUserID, "metadata: get db list failed. id: %d")
	message.Messages[ErrMetadataGetMiddlewareClustersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareClustersByUserID, "metadata: get middleware cluster list failed. id: %d")
	message.Messages[ErrMetadataGetMySQLClustersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMySQLClustersByUserID, "metadata: get mysql cluster list failed. id: %d")
	message.Messages[ErrMetadataGetAllMySQLServersByUserID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAllMySQLServersByUserID, "metadata: get mysql server list failed. id: %d")
}
