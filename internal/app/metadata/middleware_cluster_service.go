package metadata

import (
	"fmt"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/pkg/message"
)

const middlewareClusterMiddlewareClustersStruct = "MiddlewareClusters"

var _ metadata.MiddlewareClusterService = (*MiddlewareClusterService)(nil)

type MiddlewareClusterService struct {
	metadata.MiddlewareClusterRepo
	MiddlewareClusters []metadata.MiddlewareCluster `json:"middleware_clusters"`
	MiddlewareServers  []metadata.MiddlewareServer  `json:"middleware_servers"`
	Users              []metadata.User              `json:"users"`
}

// NewMiddlewareClusterService returns a new *MiddlewareClusterService
func NewMiddlewareClusterService(repo metadata.MiddlewareClusterRepo) *MiddlewareClusterService {
	return &MiddlewareClusterService{
		MiddlewareClusterRepo: repo,
		MiddlewareClusters:    []metadata.MiddlewareCluster{},
		MiddlewareServers:     []metadata.MiddlewareServer{},
	}
}

// NewMiddlewareClusterServiceWithDefault returns a new *MiddlewareClusterService with default repository
func NewMiddlewareClusterServiceWithDefault() *MiddlewareClusterService {
	return NewMiddlewareClusterService(NewMiddlewareClusterRepoWithGlobal())
}

// GetMiddlewareClusters returns middleware clusters of the service
func (mcs *MiddlewareClusterService) GetMiddlewareClusters() []metadata.MiddlewareCluster {
	return mcs.MiddlewareClusters
}

// GetMiddlewareServers returns middleware servers of the service
func (mcs *MiddlewareClusterService) GetMiddlewareServers() []metadata.MiddlewareServer {
	return mcs.MiddlewareServers
}

// GetAll gets all middleware cluster entities from the middleware
func (mcs *MiddlewareClusterService) GetAll() error {
	var err error

	mcs.MiddlewareClusters, err = mcs.MiddlewareClusterRepo.GetAll()

	return err
}

// GetByEnv gets middleware clusters of given env id
func (mcs *MiddlewareClusterService) GetByEnv(envID int) error {
	var err error

	mcs.MiddlewareClusters, err = mcs.MiddlewareClusterRepo.GetByEnv(envID)

	return err
}

// GetByID gets an middleware cluster entity that contains the given id from the middleware
func (mcs *MiddlewareClusterService) GetByID(id int) error {
	middlewareCluster, err := mcs.MiddlewareClusterRepo.GetByID(id)
	if err != nil {
		return err
	}

	mcs.MiddlewareClusters = nil
	mcs.MiddlewareClusters = append(mcs.MiddlewareClusters, middlewareCluster)

	return err
}

// GetByName gets a middleware cluster of given cluster name
func (mcs *MiddlewareClusterService) GetByName(clusterName string) error {
	middlewareCluster, err := mcs.MiddlewareClusterRepo.GetByName(clusterName)
	if err != nil {
		return err
	}

	mcs.MiddlewareClusters = nil
	mcs.MiddlewareClusters = append(mcs.MiddlewareClusters, middlewareCluster)
	return nil
}

// GetUsers returns users of the service
func (mcs *MiddlewareClusterService) GetUsers() []metadata.User {
	return mcs.Users
}

// GetMiddlewareServersByID gets the middleware server id list of given cluster id
func (mcs *MiddlewareClusterService) GetMiddlewareServersByID(clusterID int) error {
	middlewareCluster, err := mcs.MiddlewareClusterRepo.GetByID(clusterID)
	if err != nil {
		return err
	}

	mcs.MiddlewareServers, err = middlewareCluster.GetMiddlewareServers()

	return err
}

// GetUsersByMiddlewareClusterID gets user list that own the middleware cluster
func (mcs *MiddlewareClusterService) GetUsersByMiddlewareClusterID(clusterID int) error {
	var err error
	mcs.Users, err = mcs.MiddlewareClusterRepo.GetUsersByMiddlewareClusterID(clusterID)
	return err
}

// Create creates a new middleware cluster entity and insert it into the middleware
func (mcs *MiddlewareClusterService) Create(fields map[string]interface{}) error {
	// generate new map
	_, clusterNameExists := fields[middlewareClusterClusterNameStruct]
	_, envIDExists := fields[middlewareClusterEnvIDStruct]
	if !clusterNameExists || !envIDExists {
		return message.NewMessage(message.ErrFieldNotExists, fmt.Sprintf("%s and %s", middlewareClusterClusterNameStruct, middlewareClusterClusterNameStruct))
	}
	// create a new entity
	middlewareClusterInfo, err := NewMiddlewareClusterInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	middlewareCluster, err := mcs.MiddlewareClusterRepo.Create(middlewareClusterInfo)
	if err != nil {
		return err
	}

	mcs.MiddlewareClusters = nil
	mcs.MiddlewareClusters = append(mcs.MiddlewareClusters, middlewareCluster)
	return nil
}

// Update gets an middleware cluster entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (mcs *MiddlewareClusterService) Update(id int, fields map[string]interface{}) error {
	err := mcs.GetByID(id)
	if err != nil {
		return err
	}
	err = mcs.MiddlewareClusters[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return mcs.MiddlewareClusterRepo.Update(mcs.MiddlewareClusters[constant.ZeroInt])
}

// Delete deletes the middleware cluster entity that contains the given id in the middleware
func (mcs *MiddlewareClusterService) Delete(id int) error {
	err := mcs.GetByID(id)
	if err != nil {
		return err
	}

	return mcs.MiddlewareClusterRepo.Delete(id)
}

// AddUser adds a new map of middleware cluster and user in the middleware
func (mcs *MiddlewareClusterService) AddUser(middlewareClusterID, userID int) error {
	err := mcs.MiddlewareClusterRepo.MiddlewareClusterAddUser(middlewareClusterID, userID)
	if err != nil {
		return err
	}
	return mcs.GetUsersByMiddlewareClusterID(middlewareClusterID)
}

// DeleteUser deletes the map of middleware cluster and user in the middleware
func (mcs *MiddlewareClusterService) DeleteUser(middlewareClusterID, userID int) error {
	err := mcs.MiddlewareClusterRepo.MiddlewareClusterDeleteUser(middlewareClusterID, userID)
	if err != nil {
		return err
	}
	return mcs.GetUsersByMiddlewareClusterID(middlewareClusterID)
}

// Marshal marshals service.Envs
func (mcs *MiddlewareClusterService) Marshal() ([]byte, error) {
	return mcs.MarshalWithFields(middlewareClusterMiddlewareClustersStruct)
}

// MarshalWithFields marshals service.Envs with given fields
func (mcs *MiddlewareClusterService) MarshalWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(mcs, fields...)
}
