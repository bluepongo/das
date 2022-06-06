package privilege

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initServiceDebugMessage()
	initServiceInfoMessage()
	initServiceErrorMessage()
}

// Message code
const (
	// debug

	// info

	// error
	ErrPrivilegeNotEnoughPrivilegeByMySQLServerID = 407001
	ErrPrivilegeNotEnoughPrivilegeByHostInfo      = 407002
	ErrPrivilegeNotEnoughPrivilegeByDBID          = 407003
)

func initServiceDebugMessage() {

}

func initServiceInfoMessage() {

}

func initServiceErrorMessage() {
	message.Messages[ErrPrivilegeNotEnoughPrivilegeByMySQLServerID] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrPrivilegeNotEnoughPrivilegeByMySQLServerID,
		"user does not have privilege of this mysql server. mysql server id: %d, login name: %s")
	message.Messages[ErrPrivilegeNotEnoughPrivilegeByHostInfo] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrPrivilegeNotEnoughPrivilegeByHostInfo,
		"user does not have privilege of this mysql server. host ip: %s, port number: %d, login name: %s")
	message.Messages[ErrPrivilegeNotEnoughPrivilegeByDBID] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrPrivilegeNotEnoughPrivilegeByDBID,
		"user does not have privilege of this db. db id: %d, login name: %s")
}
