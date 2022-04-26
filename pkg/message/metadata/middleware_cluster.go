package metadata

import (
	"github.com/romberli/go-util/config"

	"github.com/romberli/das/pkg/message"
)

func init() {
	initDebugMiddlewareClusterMessage()
	initInfoMiddlewareClusterMessage()
	initErrorMiddlewareClusterMessage()
}

const (
	// debug
	DebugMetadataGetMiddlewareClusterAll       = 101301
	DebugMetadataGetMiddlewareClusterByEnv     = 101302
	DebugMetadataGetMiddlewareClusterByID      = 101303
	DebugMetadataGetMiddlewareClusterByName    = 101304
	DebugMetadataGetMiddlewareServers          = 101305
	DebugMetadataAddMiddlewareCluster          = 101306
	DebugMetadataUpdateMiddlewareCluster       = 101307
	DebugMetadataDeleteMiddlewareCluster       = 101308
	DebugMetadataGetUsersByMiddlewareClusterID = 101309
	DebugMetadataMiddlewareClusterAddUser      = 101310
	DebugMetadataMiddlewareClusterDeleteUser   = 101311
	// info
	InfoMetadataGetMiddlewareClusterAll       = 201301
	InfoMetadataGetMiddlewareClusterByEnv     = 201302
	InfoMetadataGetMiddlewareClusterByID      = 201303
	InfoMetadataGetMiddlewareClusterByName    = 201304
	InfoMetadataGetMiddlewareServers          = 201305
	InfoMetadataAddMiddlewareCluster          = 201306
	InfoMetadataUpdateMiddlewareCluster       = 201307
	InfoMetadataDeleteMiddlewareCluster       = 201308
	InfoMetadataGetUsersByMiddlewareClusterID = 201309
	InfoMetadataMiddlewareClusterAddUser      = 201310
	InfoMetadataMiddlewareClusterDeleteUser   = 201311
	// error
	ErrMetadataGetMiddlewareClusterAll       = 401301
	ErrMetadataGetMiddlewareClusterByEnv     = 401302
	ErrMetadataGetMiddlewareClusterByID      = 401303
	ErrMetadataGetMiddlewareClusterByName    = 401304
	ErrMetadataGetMiddlewareServers          = 401305
	ErrMetadataAddMiddlewareCluster          = 401306
	ErrMetadataUpdateMiddlewareCluster       = 401307
	ErrMetadataDeleteMiddlewareCluster       = 401308
	ErrMetadataGetUsersByMiddlewareClusterID = 401309
	ErrMetadataMiddlewareClusterAddUser      = 401310
	ErrMetadataMiddlewareClusterDeleteUser   = 401311
)

func initDebugMiddlewareClusterMessage() {
	message.Messages[DebugMetadataGetMiddlewareClusterAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareClusterAll, "metadata: get all middleware clusters message: %s")
	message.Messages[DebugMetadataGetMiddlewareClusterByEnv] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareClusterByEnv, "metadata: get middleware cluster by environment completed. message: %s")
	message.Messages[DebugMetadataGetMiddlewareClusterByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareClusterByID, "metadata: get middleware cluster by id message: %s")
	message.Messages[DebugMetadataGetMiddlewareClusterByName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareClusterByName, "metadata: get middleware cluster by name message: %s")
	message.Messages[DebugMetadataGetMiddlewareServers] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMiddlewareServers, "metadata: get middleware servers completed. message: %s")
	message.Messages[DebugMetadataAddMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddMiddlewareCluster, "metadata: add new middleware cluster message: %s")
	message.Messages[DebugMetadataUpdateMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateMiddlewareCluster, "metadata: update middleware cluster message: %s")
	message.Messages[DebugMetadataDeleteMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteMiddlewareCluster, "metadata: delete middleware cluster completed. message: %s")
	message.Messages[DebugMetadataGetUsersByMiddlewareClusterID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetUsersByMiddlewareClusterID, "metadata: get users by middleware cluster id completed. message: %s")
	message.Messages[DebugMetadataMiddlewareClusterAddUser] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataMiddlewareClusterAddUser, "metadata: add map of middleware cluster and user completed. message: %s")
	message.Messages[DebugMetadataMiddlewareClusterDeleteUser] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataMiddlewareClusterDeleteUser, "metadata: delete map of middleware cluster and user completed. message: %s")
}

func initInfoMiddlewareClusterMessage() {
	message.Messages[InfoMetadataGetMiddlewareClusterAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareClusterAll, "metadata: get middleware clusters all completed. id: %d")
	message.Messages[InfoMetadataGetMiddlewareClusterByEnv] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareClusterByEnv, "metadata: get middleware clusters by environment completed. env_id: %d")
	message.Messages[InfoMetadataGetMiddlewareClusterByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareClusterByID, "metadata: get middleware cluster by id completed. id: %d")
	message.Messages[InfoMetadataGetMiddlewareClusterByName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareClusterByName, "metadata: get middleware cluster by name completed. id: %d")
	message.Messages[InfoMetadataGetMiddlewareServers] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMiddlewareServers, "metadata: get middleware servers completed. id: %d")
	message.Messages[InfoMetadataAddMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddMiddlewareCluster, "metadata: add new middleware cluster completed. cluster_name: %s")
	message.Messages[InfoMetadataUpdateMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateMiddlewareCluster, "metadata: update middleware cluster completed. id: %d")
	message.Messages[InfoMetadataDeleteMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDeleteMiddlewareCluster, "metadata: delete middleware cluster completed. cluster_name: %s")
	message.Messages[InfoMetadataGetUsersByMiddlewareClusterID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetUsersByMiddlewareClusterID, "metadata: get users by middleware cluster id completed. id: %d")
	message.Messages[InfoMetadataMiddlewareClusterAddUser] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataMiddlewareClusterAddUser, "metadata: add map of middleware cluster and user completed. middleware_cluster_id: %d, user_id: %d")
	message.Messages[InfoMetadataMiddlewareClusterDeleteUser] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataMiddlewareClusterDeleteUser, "metadata: delete map of middleware cluster and user completed. middleware_cluster_id: %d, user_id: %d")
}

func initErrorMiddlewareClusterMessage() {
	message.Messages[ErrMetadataGetMiddlewareClusterAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareClusterAll, "metadata: get all middleware clusters failed")
	message.Messages[ErrMetadataGetMiddlewareClusterByEnv] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareClusterByEnv, "metadata: get middleware cluster by environment failed")
	message.Messages[ErrMetadataGetMiddlewareClusterByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareClusterByID, "metadata: get middleware cluster by id failed. id: %d")
	message.Messages[ErrMetadataGetMiddlewareClusterByName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareClusterByName, "metadata: get middleware cluster by name failed. id: %d")
	message.Messages[ErrMetadataGetMiddlewareServers] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMiddlewareServers, "metadata: get middleware servers failed")
	message.Messages[ErrMetadataAddMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddMiddlewareCluster, "metadata: add new middleware cluster failed. env_name: %s")
	message.Messages[ErrMetadataUpdateMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateMiddlewareCluster, "metadata: update middleware cluster failed. id: %d")
	message.Messages[ErrMetadataDeleteMiddlewareCluster] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteMiddlewareCluster, "metadata: delete middleware cluster failed. cluster_name: %s")
	message.Messages[ErrMetadataGetUsersByMiddlewareClusterID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetUsersByMiddlewareClusterID, "metadata: get users by middleware cluster id failed")
	message.Messages[ErrMetadataMiddlewareClusterAddUser] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataMiddlewareClusterAddUser, "metadata: add map of middleware cluster and user failed. id: %d")
	message.Messages[ErrMetadataMiddlewareClusterDeleteUser] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataMiddlewareClusterDeleteUser, "metadata: delete map of middleware cluster and user failed. id: %d")
}
