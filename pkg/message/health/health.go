package health

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initHealthDebugMessage()
	initHealthInfoMessage()
	initHealthErrorMessage()
}

const (
	// info
	InfoHealthPing = 209001
)

func initHealthDebugMessage() {

}

func initHealthInfoMessage() {
	message.Messages[InfoHealthPing] = config.NewErrMessage(message.DefaultMessageHeader, InfoHealthPing, "health: ping completed")
}

func initHealthErrorMessage() {

}
