package metadata

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/das/pkg/message"
)

func init() {
	initDebugAppMessage()
	initInfoAppMessage()
	initErrorAppMessage()
}

const (
	// debug
	DebugMetadataGetAppAll       = 101001
	DebugMetadataGetAppByID      = 101002
	DebugMetadataGetAppByName    = 101003
	DebugMetadataGetDBIDList     = 101004
	DebugMetadataAddApp          = 101005
	DebugMetadataUpdateApp       = 101006
	DebugMetadataDeleteApp       = 101007
	DebugMetadataAppAddDB        = 101008
	DebugMetadataAppDeleteDB     = 101009
	DebugMetadataGetDBsByAppID   = 101010
	DebugMetadataGetUsersByAppID = 101011
	DebugMetadataAppAddUser      = 101012
	DebugMetadataAppDeleteUser   = 101013
	// info
	InfoMetadataGetAppAll       = 201001
	InfoMetadataGetAppByID      = 201002
	InfoMetadataGetAppByName    = 201003
	InfoMetadataGetDBIDList     = 201004
	InfoMetadataAddApp          = 201005
	InfoMetadataUpdateApp       = 201006
	InfoMetadataDeleteApp       = 201007
	InfoMetadataAppAddDB        = 201008
	InfoMetadataAppDeleteDB     = 201009
	InfoMetadataGetDBsByAppID   = 201010
	InfoMetadataGetUsersByAppID = 201011
	InfoMetadataAppAddUser      = 201012
	InfoMetadataAppDeleteUser   = 201013

	// error
	ErrMetadataGetAppAll       = 401001
	ErrMetadataGetAppByID      = 401002
	ErrMetadataGetAppByName    = 401003
	ErrMetadataGetDBIDList     = 401004
	ErrMetadataAddApp          = 401005
	ErrMetadataUpdateApp       = 401006
	ErrMetadataDeleteApp       = 401007
	ErrMetadataAppAddDB        = 401008
	ErrMetadataAppDeleteDB     = 401009
	ErrMetadataGetDBsByAppID   = 401010
	ErrMetadataGetUsersByAppID = 401011
	ErrMetadataAppAddUser      = 401012
	ErrMetadataAppDeleteUser   = 401013
)

func initDebugAppMessage() {
	message.Messages[DebugMetadataGetAppAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAppAll, "metadata: get all app completed. message: %s")
	message.Messages[DebugMetadataGetAppByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAppByID, "metadata: get app by id completed. message: %s")
	message.Messages[DebugMetadataGetAppByName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAppByName, "metadata: get app by name completed. message: %s")
	message.Messages[DebugMetadataGetDBIDList] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDBIDList, "metadata: get db id list completed. message: %s")
	message.Messages[DebugMetadataAddApp] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddApp, "metadata: add new app completed. message: %s")
	message.Messages[DebugMetadataUpdateApp] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateApp, "metadata: update app completed. message: %s")
	message.Messages[DebugMetadataDeleteApp] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteApp, "metadata: delete app completed. message: %s")
	message.Messages[DebugMetadataAppAddDB] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAppAddDB, "metadata: add map of app and database completed. message: %s")
	message.Messages[DebugMetadataAppAddUser] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAppAddUser, "metadata: add map of app and user completed. message: %s")
	message.Messages[DebugMetadataAppDeleteDB] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAppDeleteDB, "metadata: delete map of app and database completed. message: %s")
	message.Messages[DebugMetadataAppDeleteUser] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAppDeleteUser, "metadata: delete map of app and user completed. message: %s")
	message.Messages[DebugMetadataGetDBsByAppID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetDBsByAppID, "metadata: get dbs by id completed. message: %s")
	message.Messages[DebugMetadataGetUsersByAppID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetUsersByAppID, "metadata: get dbs by id completed. message: %s")
}

func initInfoAppMessage() {
	message.Messages[InfoMetadataGetAppAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAppAll, "metadata: get app all completed.")
	message.Messages[InfoMetadataGetAppByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAppByID, "metadata: get app by id completed. id: %d")
	message.Messages[InfoMetadataGetAppByName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAppByName, "metadata: get app by name completed. app_name: %s")
	message.Messages[InfoMetadataGetDBIDList] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDBIDList, "metadata: get db id list completed. id: %d")
	message.Messages[InfoMetadataAddApp] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddApp, "metadata: add new app completed. app_name: %s")
	message.Messages[InfoMetadataUpdateApp] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateApp, "metadata: update app completed. id: %d")
	message.Messages[InfoMetadataDeleteApp] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDeleteApp, "metadata: delete app completed. id: %d")
	message.Messages[InfoMetadataAppAddDB] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAppAddDB, "metadata: add map of app and database completed. app_id: %d, db_id: %d")
	message.Messages[InfoMetadataAppAddUser] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAppAddUser, "metadata: add map of app and user completed. app_id: %d, db_id: %d")
	message.Messages[InfoMetadataAppDeleteDB] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAppDeleteDB, "metadata: delete map of app and database completed. app_id: %d, db_id: %d")
	message.Messages[InfoMetadataAppDeleteUser] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAppDeleteUser, "metadata: delete map of app and user completed. app_id: %d, db_id: %d")
	message.Messages[InfoMetadataGetDBsByAppID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetDBsByAppID, "metadata: get dbs by id completed. app_id: %d")
	message.Messages[InfoMetadataGetUsersByAppID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetUsersByAppID, "metadata: get dbs by id completed. app_id: %d")

}

func initErrorAppMessage() {
	message.Messages[ErrMetadataGetAppAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAppAll, "metadata: get all app failed.")
	message.Messages[ErrMetadataGetAppByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAppByID, "metadata: get app by id failed. id: %d")
	message.Messages[ErrMetadataGetAppByName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAppByName, "metadata: get app by name failed. app_name: %s")
	message.Messages[ErrMetadataGetDBIDList] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetDBIDList, "metadata: get db list failed. id: %d")
	message.Messages[ErrMetadataAddApp] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddApp, "metadata: add new app failed. app_name: %s")
	message.Messages[ErrMetadataUpdateApp] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateApp, "metadata: update app failed. id: %d")
	message.Messages[ErrMetadataDeleteApp] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteApp, "metadata: delete app failed. id: %d")
	message.Messages[ErrMetadataAppAddDB] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAppAddDB, "metadata: add map of app and database failed. id: %d")
	message.Messages[ErrMetadataAppAddUser] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAppAddUser, "metadata: add map of app and user failed. id: %d")
	message.Messages[ErrMetadataAppDeleteDB] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAppDeleteDB, "metadata: delete map of app and database failed. id: %d")
	message.Messages[ErrMetadataAppDeleteUser] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAppDeleteUser, "metadata: delete map of app and user failed. id: %d")
	message.Messages[ErrMetadataGetDBsByAppID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetDBsByAppID, "metadata: get dbs by id failed. app_id: %d")
	message.Messages[ErrMetadataGetUsersByAppID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetUsersByAppID, "metadata: get dbs by id failed. app_id: %d")
}
