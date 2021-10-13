package alert

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initHTTTPSenderDebugMessage()
	initHTTTPSenderInfoMessage()
	initHTTTPSenderErrorMessage()
}

const (
	// debug

	// info

	// error
	ErrHTTTPSenderCallAlertAPI = 405001
)

func initHTTTPSenderDebugMessage() {

}

func initHTTTPSenderInfoMessage() {

}

func initHTTTPSenderErrorMessage() {
	message.Messages[ErrHTTTPSenderCallAlertAPI] = config.NewErrMessage(
		message.DefaultMessageHeader, ErrHTTTPSenderCallAlertAPI,
		"got error when calling alert api.\n%s")
}
