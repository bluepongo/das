package metadata

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/das/pkg/message"
)

func init() {
	initDebugMonitorSystemMessage()
	initInfoMonitorSystemMessage()
	initErrorMonitorSystemMessage()
}

const (
	// debug
	DebugMetadataGetMonitorSystemAll        = 101501
	DebugMetadataGetMonitorSystemByEnv      = 101502
	DebugMetadataGetMonitorSystemByID       = 101503
	DebugMetadataGetMonitorSystemByHostInfo = 101504
	DebugMetadataAddMonitorSystem           = 101505
	DebugMetadataUpdateMonitorSystem        = 101506
	DebugMetadataDeleteMonitorSystem        = 101507
	// info
	InfoMetadataGetMonitorSystemAll        = 201501
	InfoMetadataGetMonitorSystemByEnv      = 201502
	InfoMetadataGetMonitorSystemByID       = 201503
	InfoMetadataGetMonitorSystemByHostInfo = 201504
	InfoMetadataAddMonitorSystem           = 201505
	InfoMetadataUpdateMonitorSystem        = 201506
	InfoMetadataDeleteMonitorSystem        = 201507
	// error
	ErrMetadataGetMonitorSystemAll        = 401501
	ErrMetadataGetMonitorSystemByEnv      = 401502
	ErrMetadataGetMonitorSystemByID       = 401503
	ErrMetadataGetMonitorSystemByHostInfo = 401504
	ErrMetadataAddMonitorSystem           = 401505
	ErrMetadataUpdateMonitorSystem        = 401506
	ErrMetadataDeleteMonitorSystem        = 401507
)

func initDebugMonitorSystemMessage() {
	message.Messages[DebugMetadataGetMonitorSystemAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMonitorSystemAll, "metadata: get all monitor systems. message: %s")
	message.Messages[DebugMetadataGetMonitorSystemByEnv] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMonitorSystemByEnv, "metadata: get monitor systems by environment. message: %s")
	message.Messages[DebugMetadataGetMonitorSystemByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMonitorSystemByID, "metadata: get monitor system by id. message: %s")
	message.Messages[DebugMetadataGetMonitorSystemByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMonitorSystemByHostInfo, "metadata: get monitor system by host info. message: %s")
	message.Messages[DebugMetadataAddMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddMonitorSystem, "metadata: add new monitor system. message: %s")
	message.Messages[DebugMetadataUpdateMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateMonitorSystem, "metadata: update monitor system. message: %s")
	message.Messages[DebugMetadataDeleteMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteMonitorSystem, "metadata: delete monitor system. message: %s")
}

func initInfoMonitorSystemMessage() {
	message.Messages[InfoMetadataGetMonitorSystemAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMonitorSystemAll, "metadata: get all monitor systems completed")
	message.Messages[InfoMetadataGetMonitorSystemByEnv] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMonitorSystemByEnv, "metadata: get monitor systems by environment completed. env_id: %d")
	message.Messages[InfoMetadataGetMonitorSystemByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMonitorSystemByID, "metadata: get monitor system by id completed. id: %d")
	message.Messages[InfoMetadataGetMonitorSystemByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMonitorSystemByHostInfo, "metadata: get monitor system by host info completed. host_ip: %s, port_num: %d")
	message.Messages[InfoMetadataAddMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddMonitorSystem, "metadata: add new monitor system completed. system_name: %s, system_type: %d, host_ip: %s, port_num: %d, port_num_slow: %d, base_url: %s, env_id: %d")
	message.Messages[InfoMetadataUpdateMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateMonitorSystem, "metadata: update monitor system completed. id: %d")
	message.Messages[InfoMetadataDeleteMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDeleteMonitorSystem, "metadata: delete monitor system completed. id: %d")
}

func initErrorMonitorSystemMessage() {
	message.Messages[ErrMetadataGetMonitorSystemAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMonitorSystemAll, "metadata: get all monitor systems failed")
	message.Messages[ErrMetadataGetMonitorSystemByEnv] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMonitorSystemByEnv, "metadata: get monitor systems by environment failed. env_id: %d")
	message.Messages[ErrMetadataGetMonitorSystemByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMonitorSystemByID, "metadata: get monitor system by id failed. id: %d")
	message.Messages[ErrMetadataGetMonitorSystemByHostInfo] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMonitorSystemByHostInfo, "metadata: get monitor system by host info failed. host_ip: %s, port_num: %d")
	message.Messages[ErrMetadataAddMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddMonitorSystem, "metadata: add new monitor system failed. system_name: %s, system_type: %d, host_ip: %s, port_num: %d, port_num_slow: %d, base_url: %s, env_id: %d")
	message.Messages[ErrMetadataUpdateMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateMonitorSystem, "metadata: update monitor system failed. id: %d")
	message.Messages[ErrMetadataDeleteMonitorSystem] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteMonitorSystem, "metadata: delete monitor system failed. id: %d")
}
