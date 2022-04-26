package metadata

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initDebugMiddlewareServerMessage()
	initInfoMiddlewareServerMessage()
	initErrorMiddlewareServerMessage()
}

const (
	// debug
	DebugMetadataGetMiddlewareServerAll        = 101401
	DebugMetadataGetMiddlewareSeverByClusterID = 101402
	DebugMetadataGetMiddlewareServerByID       = 101403
	DebugMetadataGetMiddlewareServerByHostInfo = 101404
	DebugMetadataAddMiddlewareServer           = 101405
	DebugMetadataUpdateMiddlewareServer        = 101406
	DebugMetadataDeleteMiddlewareServer        = 101407

	// info
	InfoMetadataGetMiddlewareServerAll        = 201401
	InfoMetadataGetMiddlewareSeverByClusterID = 201402
	InfoMetadataGetMiddlewareServerByID       = 201403
	InfoMetadataGetMiddlewareServerByHostInfo = 201404
	InfoMetadataAddMiddlewareServer           = 201405
	InfoMetadataUpdateMiddlewareServer        = 201406
	InfoMetadataDeleteMiddlewareServer        = 201407
	// error
	ErrMetadataGetMiddlewareServerAll        = 401401
	ErrMetadataGetMiddlewareSeverByClusterID = 401402
	ErrMetadataGetMiddlewareServerByID       = 401403
	ErrMetadataGetMiddlewareServerByHostInfo = 401404
	ErrMetadataAddMiddlewareServer           = 401405
	ErrMetadataUpdateMiddlewareServer        = 401406
	ErrMetadataDeleteMiddlewareServer        = 401407
)

func initDebugMiddlewareServerMessage() {
	message.Messages[DebugMetadataGetMiddlewareServerAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareServerAll, "metadata: get all middleware server message: %s")
	message.Messages[DebugMetadataGetMiddlewareSeverByClusterID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareSeverByClusterID, "metadata: get middleware cluster by cluster completed. message: %s")
	message.Messages[DebugMetadataGetMiddlewareServerByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareServerByID, "metadata: get middleware server by id message: %s")
	message.Messages[DebugMetadataGetMiddlewareServerByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareServerByHostInfo, "metadata: get middleware cluster by host info message: %s")
	message.Messages[DebugMetadataAddMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddMiddlewareServer, "metadata: add new middleware server message: %s")
	message.Messages[DebugMetadataUpdateMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateMiddlewareServer, "metadata: update middleware server message: %s")
	message.Messages[DebugMetadataDeleteMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteMiddlewareServer, "metadata: delete middleware server completed. message: %s")
}

func initInfoMiddlewareServerMessage() {
	message.Messages[InfoMetadataGetMiddlewareServerAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareServerAll, "metadata: get middleware server all completed. id: %d")
	message.Messages[InfoMetadataGetMiddlewareSeverByClusterID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareSeverByClusterID, "metadata: get middleware clusters by cluster completed. cluster_id: %d")
	message.Messages[InfoMetadataGetMiddlewareServerByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareServerByID, "metadata: get middleware server by id completed. id: %d")
	message.Messages[InfoMetadataGetMiddlewareServerByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareServerByHostInfo, "metadata: get middleware cluster by host info completed. host ip: %s, port num: %d")
	message.Messages[InfoMetadataAddMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddMiddlewareServer, "metadata: add new middleware server completed. server_name: %s")
	message.Messages[InfoMetadataUpdateMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateMiddlewareServer, "metadata: update middleware server completed. id: %d")
	message.Messages[InfoMetadataDeleteMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDeleteMiddlewareServer, "metadata: delete middleware server completed. server_name: %s")
}

func initErrorMiddlewareServerMessage() {
	message.Messages[ErrMetadataGetMiddlewareServerAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareServerAll, "metadata: get all middleware server failed")
	message.Messages[ErrMetadataGetMiddlewareSeverByClusterID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareSeverByClusterID, "metadata: get middleware cluster by cluster failed")
	message.Messages[ErrMetadataGetMiddlewareServerByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareServerByID, "metadata: get middleware server by id failed. id: %d")
	message.Messages[ErrMetadataGetMiddlewareServerByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareServerByHostInfo, "metadata: get middleware cluster by host info failed. host ip: %s, port num: %d")
	message.Messages[ErrMetadataAddMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddMiddlewareServer, "metadata: add new middleware server failed. server_name: %s")
	message.Messages[ErrMetadataUpdateMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateMiddlewareServer, "metadata: update middleware server failed. id: %d")
	message.Messages[ErrMetadataDeleteMiddlewareServer] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteMiddlewareServer, "metadata: delete middleware server failed. server_name: %s")
}
