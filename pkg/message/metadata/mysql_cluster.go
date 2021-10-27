package metadata

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/das/pkg/message"
)

func init() {
	initDebugMySQLCLusterMessage()
	initInfoMySQLCLusterMessage()
	initErrorMySQLCLusterMessage()
}

const (
	// debug
	DebugMetadataGetMySQLClusterAll    = 100701
	DebugMetadataGetMySQLClusterByEnv  = 100702
	DebugMetadataGetMySQLClusterByID   = 100703
	DebugMetadataGetMySQLClusterByName = 100704
	DebugMetadataGetMySQLServers       = 100705
	DebugMetadataGetMasterServers      = 100706
	DebugMetadataGetDBs                = 100707
	DebugMetadataGetAppOwners          = 100708
	DebugMetadataGetDBOwners           = 100709
	DebugMetadataGetAllOwners          = 100710
	DebugMetadataAddMySQLCluster       = 100711
	DebugMetadataUpdateMySQLCluster    = 100712
	DebugMetadataDeleteMySQLCluster    = 100713
	// debug
	InfoMetadataGetMySQLClusterAll    = 200701
	InfoMetadataGetMySQLClusterByEnv  = 200702
	InfoMetadataGetMySQLClusterByID   = 200703
	InfoMetadataGetMySQLClusterByName = 200704
	InfoMetadataGetMySQLServers       = 200705
	InfoMetadataGetMasterServers      = 200706
	InfoMetadataGetDBs                = 200707
	InfoMetadataGetAppOwners          = 200708
	InfoMetadataGetDBOwners           = 200709
	InfoMetadataGetAllOwners          = 200710
	InfoMetadataAddMySQLCluster       = 200711
	InfoMetadataUpdateMySQLCluster    = 200712
	InfoMetadataDeleteMySQLCluster    = 200713
	// error
	ErrMetadataGetMySQLClusterAll    = 400701
	ErrMetadataGetMySQLClusterByEnv  = 400702
	ErrMetadataGetMySQLClusterByID   = 400703
	ErrMetadataGetMySQLClusterByName = 400704
	ErrMetadataGetMySQLServers       = 400705
	ErrMetadataGetMasterServers      = 400706
	ErrMetadataGetDBs                = 400707
	ErrMetadataGetAppOwners          = 400708
	ErrMetadataGetDBOwners           = 400709
	ErrMetadataGetAllOwners          = 400710
	ErrMetadataAddMySQLCluster       = 400711
	ErrMetadataUpdateMySQLCluster    = 400712
	ErrMetadataDeleteMySQLCluster    = 400713
)

func initDebugMySQLCLusterMessage() {
	message.Messages[DebugMetadataGetMySQLClusterAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterAll,
		"metadata: get all mysql clusters message: %s")
	message.Messages[DebugMetadataGetMySQLClusterByEnv] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterByEnv,
		"metadata: get mysql cluster by env message: %s")
	message.Messages[DebugMetadataGetMySQLClusterByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterByID,
		"metadata: get mysql cluster by id message: %s")
	message.Messages[DebugMetadataGetMySQLClusterByName] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLClusterByName,
		"metadata: get mysql cluster by name message: %s")
	message.Messages[DebugMetadataGetMySQLServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMySQLServers,
		"metadata: get mysql servers from mysql cluster message: %s")
	message.Messages[DebugMetadataGetMasterServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetMasterServers,
		"metadata: get master servers from mysql cluster message: %s")
	message.Messages[DebugMetadataGetDBs] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetDBs,
		"metadata: get databases from mysql cluster message: %s")
	message.Messages[DebugMetadataGetAppOwners] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetAppOwners,
		"metadata: get app owners from mysql cluster message: %s")
	message.Messages[DebugMetadataGetDBOwners] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetDBOwners,
		"metadata: get database owners from mysql cluster message: %s")
	message.Messages[DebugMetadataGetAllOwners] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataGetAllOwners,
		"metadata: get all owners from mysql cluster message: %s")
	message.Messages[DebugMetadataAddMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataAddMySQLCluster,
		"metadata: add new mysql cluster message: %s")
	message.Messages[DebugMetadataUpdateMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataUpdateMySQLCluster,
		"metadata: update mysql cluster message: %s")
	message.Messages[DebugMetadataDeleteMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		DebugMetadataDeleteMySQLCluster,
		"metadata: delete mysql cluster message: %s")
}

func initInfoMySQLCLusterMessage() {
	message.Messages[InfoMetadataGetMySQLClusterAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLClusterAll,
		"metadata: get mysql cluster all completed")
	message.Messages[InfoMetadataGetMySQLClusterByEnv] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLClusterByEnv,
		"metadata: get mysql cluster by env completed. env_id: %s")
	message.Messages[InfoMetadataGetMySQLClusterByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLClusterByID,
		"metadata: get mysql cluster by id completed. id: %s")
	message.Messages[InfoMetadataGetMySQLClusterByName] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLClusterByName,
		"metadata: get mysql cluster by name completed. cluster_name: %s")
	message.Messages[InfoMetadataGetMySQLServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMySQLServers,
		"metadata: get mysql servers from mysql cluster completed. id: %d")
	message.Messages[InfoMetadataGetMasterServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetMasterServers,
		"metadata: get master servers from mysql cluster completed. id: %d")
	message.Messages[InfoMetadataGetDBs] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetDBs,
		"metadata: get databases from mysql cluster completed. id: %d")
	message.Messages[InfoMetadataGetAppOwners] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetAppOwners,
		"metadata: get app owners from mysql cluster completed. id: %d")
	message.Messages[InfoMetadataGetDBOwners] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetDBOwners,
		"metadata: get database owners from mysql cluster completed. id: %d")
	message.Messages[InfoMetadataGetAllOwners] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataGetAllOwners,
		"metadata: get all owners from mysql cluster completed. id: %d")
	message.Messages[InfoMetadataAddMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataAddMySQLCluster,
		"metadata: add new mysql cluster completed. cluster_name: %s, env_id: %s")
	message.Messages[InfoMetadataUpdateMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataUpdateMySQLCluster,
		"metadata: update mysql cluster completed. id: %s")
	message.Messages[InfoMetadataDeleteMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		InfoMetadataDeleteMySQLCluster,
		"metadata: delete mysql cluster completed. id: %s")
}

func initErrorMySQLCLusterMessage() {
	message.Messages[ErrMetadataGetMySQLClusterAll] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLClusterAll,
		"metadata: get all mysql cluster failed.\n%s")
	message.Messages[ErrMetadataGetMySQLClusterByEnv] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLClusterByEnv,
		"metadata: get mysql cluster by env failed. env_id: %d\n%s")
	message.Messages[ErrMetadataGetMySQLClusterByID] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLClusterByID,
		"metadata: get mysql cluster by id failed. id: %d\n%s")
	message.Messages[ErrMetadataGetMySQLClusterByName] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLClusterByName,
		"metadata: get mysql cluster by name failed. cluster_name: %s\n%s")
	message.Messages[ErrMetadataGetMySQLServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMySQLServers,
		"metadata: get mysql servers from mysql cluster failed. id: %d\n%s")
	message.Messages[ErrMetadataGetMasterServers] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetMasterServers,
		"metadata: get master servers from mysql cluster failed. id: %d\n%s")
	message.Messages[ErrMetadataGetDBs] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetDBs,
		"metadata: get databases from mysql cluster failed. id: %d\n%s")
	message.Messages[ErrMetadataGetAppOwners] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetAppOwners,
		"metadata: get app owners from mysql cluster failed. id: %d\n%s")
	message.Messages[ErrMetadataGetDBOwners] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetDBOwners,
		"metadata: get database owners from mysql cluster failed. id: %d\n%s")
	message.Messages[ErrMetadataGetAllOwners] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataGetAllOwners,
		"metadata: get all owners from mysql cluster failed. id: %d\n%s")
	message.Messages[ErrMetadataAddMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataAddMySQLCluster,
		"metadata: add new mysql cluster failed. cluster_name: %s, env_id: %d\n%s")
	message.Messages[ErrMetadataUpdateMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataUpdateMySQLCluster,
		"metadata: update mysql cluster failed. id: %d\n%s")
	message.Messages[ErrMetadataDeleteMySQLCluster] = config.NewErrMessage(
		message.DefaultMessageHeader,
		ErrMetadataDeleteMySQLCluster,
		"metadata: delete mysql cluster failed. id: %d\n%s")
}
