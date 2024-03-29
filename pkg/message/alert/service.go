package alert

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initServiceDebugMessage()
	initServiceInfoMessage()
	initServiceErrorMessage()
}

const (
	// debug
	DebugServiceSendEmail = 106101

	// info
	InfoServiceSendEmail = 206101

	// error
	ErrServiceSendEmail = 406101
)

func initServiceDebugMessage() {
	message.Messages[DebugServiceSendEmail] = config.NewErrMessage(
		message.DefaultMessageHeader, DebugServiceSendEmail,
		"sending email completed. config: %s, to addresses: %s, cc addresses: %s, subject: %s, content: %s.")
}

func initServiceInfoMessage() {
	message.Messages[InfoServiceSendEmail] = config.NewErrMessage(
		message.DefaultMessageHeader, InfoServiceSendEmail,
		"sending email completed. config: %s, to addresses: %s, cc addresses: %s, subject: %s.")
}

func initServiceErrorMessage() {
	message.Messages[ErrServiceSendEmail] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrServiceSendEmail,
		"got error when sending email. config: %s, to addresses: %s, cc addresses: %s, subject: %s.\n%s")
}
