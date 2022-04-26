package metadata

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initDebugEnvMessage()
	initInfoEnvMessage()
	initErrorEnvMessage()
}

const (
	// debug
	DebugMetadataGetEnvAll     = 101201
	DebugMetadataGetEnvByID    = 101202
	DebugMetadataAddEnv        = 101203
	DebugMetadataUpdateEnv     = 101204
	DebugMetadataGetEnvByName  = 101205
	DebugMetadataDeleteEnvByID = 101206
	// info
	InfoMetadataGetEnvAll     = 201201
	InfoMetadataGetEnvByID    = 201202
	InfoMetadataAddEnv        = 201203
	InfoMetadataUpdateEnv     = 201204
	InfoMetadataGetEnvByName  = 201205
	InfoMetadataDeleteEnvByID = 201206
	// error
	ErrMetadataGetEnvAll     = 401201
	ErrMetadataGetEnvByID    = 401202
	ErrMetadataAddEnv        = 401203
	ErrMetadataUpdateEnv     = 401204
	ErrMetadataGetEnvByName  = 401205
	ErrMetadataDeleteEnvByID = 401206
)

func initDebugEnvMessage() {
	message.Messages[DebugMetadataGetEnvAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetEnvAll, "metadata: get all environments message: %s")
	message.Messages[DebugMetadataGetEnvByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetEnvByID, "metadata: get environment by id message: %s")
	message.Messages[DebugMetadataAddEnv] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddEnv, "metadata: add new environment message: %s")
	message.Messages[DebugMetadataUpdateEnv] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateEnv, "metadata: update environment message: %s")
	message.Messages[DebugMetadataGetEnvByName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetEnvByName, "metadata: get environment by name message: %s")
	message.Messages[DebugMetadataDeleteEnvByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteEnvByID, "metadata: delete environment by ID message: %s")
}

func initInfoEnvMessage() {
	message.Messages[InfoMetadataGetEnvAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetEnvAll, "metadata: get environment all completed")
	message.Messages[InfoMetadataGetEnvByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetEnvByID, "metadata: get environment by id completed. id: %d")
	message.Messages[InfoMetadataAddEnv] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddEnv, "metadata: add new environment completed. env_name: %s")
	message.Messages[InfoMetadataUpdateEnv] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateEnv, "metadata: update environment completed. id: %d")
	message.Messages[InfoMetadataGetEnvByName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetEnvByName, "metadata: get environment by name completed. id: %d")
	message.Messages[InfoMetadataDeleteEnvByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDeleteEnvByID, "metadata: delete environment by ID completed. id: %d")
}

func initErrorEnvMessage() {
	message.Messages[ErrMetadataGetEnvAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetEnvAll, "metadata: get all environment failed")
	message.Messages[ErrMetadataGetEnvByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetEnvByID, "metadata: get environment by id failed. id: %d")
	message.Messages[ErrMetadataAddEnv] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddEnv, "metadata: add new environment failed. env_name: %s")
	message.Messages[ErrMetadataUpdateEnv] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateEnv, "metadata: update environment failed. id: %d")
	message.Messages[ErrMetadataGetEnvByName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetEnvByName, "metadata: get environment by name failed. env_name: %s")
	message.Messages[ErrMetadataDeleteEnvByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteEnvByID, "metadata: delete environment by ID failed. id: %d")
}
