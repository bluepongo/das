package message

import (
	"github.com/romberli/go-util/config"
)

const (
	DebugMetadataGetEnvAll  = 100301
	DebugMetadataGetEnvByID = 100302
	DebugMetadataAddEnv     = 100303
	DebugMetadataUpdateEnv  = 100304

	DebugMetadataGetMiddlewareServerAll  = 100601
	DebugMetadataGetMiddlewareServerByID = 100602
	DebugMetadataAddMiddlewareServer     = 100603
	DebugMetadataUpdateMiddlewareServer  = 100604

	DebugMetadataGetMySQLClusterAll  = 100701
	DebugMetadataGetMySQLClusterByID = 100702
	DebugMetadataAddMySQLCluster     = 100703
	DebugMetadataUpdateMySQLCluster  = 100704

	DebugMetadataGetMySQLServerAll  = 100801
	DebugMetadataGetMySQLServerByID = 100802
	DebugMetadataAddMySQLServer     = 100803
	DebugMetadataUpdateMySQLServer  = 100804

	DebugMetadataGetUserAll  = 100901
	DebugMetadataGetUserByID = 100902
	DebugMetadataAddUser     = 100903
	DebugMetadataUpdateUser  = 100904

	DebugMetadataGetMonitorSystemAll  = 100601
	DebugMetadataGetMonitorSystemByID = 100602
	DebugMetadataAddMonitorSystem     = 100603
	DebugMetadataUpdateMonitorSystem  = 100604
)

func initDebugMessage() {
	Messages[DebugMetadataGetEnvAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetEnvAll, "metadata: get all environments message: %s")
	Messages[DebugMetadataGetEnvByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetEnvByID, "metadata: get environment by id message: %s")
	Messages[DebugMetadataAddEnv] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddEnv, "metadata: add new environment message: %s")
	Messages[DebugMetadataUpdateEnv] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateEnv, "metadata: update environment message: %s")

	Messages[DebugMetadataGetMySQLClusterAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMySQLClusterAll, "metadata: get all mysql clusters message: %s")
	Messages[DebugMetadataGetMySQLClusterByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMySQLClusterByID, "metadata: get mysql cluster by id message: %s")
	Messages[DebugMetadataAddMySQLCluster] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddMySQLCluster, "metadata: add new mysql cluster message: %s")
	Messages[DebugMetadataUpdateMySQLCluster] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateMySQLCluster, "metadata: update mysql cluster message: %s")

	Messages[DebugMetadataGetMySQLServerAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMySQLServerAll, "metadata: get all mysql servers message: %s")
	Messages[DebugMetadataGetMySQLServerByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMySQLServerByID, "metadata: get mysql server by id message: %s")
	Messages[DebugMetadataAddMySQLServer] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddMySQLServer, "metadata: add new mysql server message: %s")
	Messages[DebugMetadataUpdateMySQLServer] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateMySQLServer, "metadata: update mysql server message: %s")

	Messages[DebugMetadataGetMiddlewareServerAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMiddlewareServerAll, "metadata: get all middleware server message: %s")
	Messages[DebugMetadataGetMiddlewareServerByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMiddlewareServerByID, "metadata: get middleware server by id message: %s")
	Messages[DebugMetadataAddMiddlewareServer] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddMiddlewareServer, "metadata: add new middleware server message: %s")
	Messages[DebugMetadataUpdateMiddlewareServer] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateMiddlewareServer, "metadata: update middleware server message: %s")

	Messages[DebugMetadataGetUserAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetUserAll, "metadata: get all user message: %s")
	Messages[DebugMetadataGetUserByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetUserByID, "metadata: get user by id message: %s")
	Messages[DebugMetadataAddUser] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddUser, "metadata: add new user message: %s")
	Messages[DebugMetadataUpdateUser] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateUser, "metadata: update user message: %s")

	Messages[DebugMetadataGetMonitorSystemAll] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMonitorSystemAll, "metadata: get all monitor systems message: %s")
	Messages[DebugMetadataGetMonitorSystemByID] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataGetMonitorSystemByID, "metadata: get monitor system by id message: %s")
	Messages[DebugMetadataAddMonitorSystem] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataAddMonitorSystem, "metadata: add new monitor system message: %s")
	Messages[DebugMetadataUpdateMonitorSystem] = config.NewErrMessage(DefaultMessageHeader, DebugMetadataUpdateMonitorSystem, "metadata: update monitor system message: %s")
}
