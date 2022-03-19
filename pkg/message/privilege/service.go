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
	ErrPrivilegeNotEnoughPrivilege = 406001
)

func initServiceDebugMessage() {

}

func initServiceInfoMessage() {

}

func initServiceErrorMessage() {
	message.Messages[ErrPrivilegeNotEnoughPrivilege] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrPrivilegeNotEnoughPrivilege,
		"user %s(%s) does not have privilege of this mysql server. mysql server id: %d")
}
