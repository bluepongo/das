package metadata

import (
	"fmt"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/das/pkg/message"
)

const mysqlServerMySQLServersStruct = "MySQLServers"

var _ metadata.MySQLServerService = (*MySQLServerService)(nil)

// MySQLServerService implements Service interface
type MySQLServerService struct {
	MySQLServerRepo metadata.MySQLServerRepo
	MySQLServers    []metadata.MySQLServer `json:"mysql_servers"`
	MySQLCluster    metadata.MySQLCluster  `json:"mysql_cluster"`
}

// NewMySQLServerService returns a new *MySQLServerService
func NewMySQLServerService(repo metadata.MySQLServerRepo) *MySQLServerService {
	return &MySQLServerService{
		repo,
		[]metadata.MySQLServer{},
		nil,
	}
}

// NewMySQLServerServiceWithDefault returns a new *MySQLServerService with default repository
func NewMySQLServerServiceWithDefault() *MySQLServerService {
	return NewMySQLServerService(NewMySQLServerRepoWithGlobal())
}

// GetMySQLServers returns the mysql servers of the service
func (mss *MySQLServerService) GetMySQLServers() []metadata.MySQLServer {
	return mss.MySQLServers
}

// GetMySQLCluster returns the mysql cluster of the service
func (mss *MySQLServerService) GetMySQLCluster() metadata.MySQLCluster {
	return mss.MySQLCluster
}

// GetAll gets all mysql server entities from the middleware
func (mss *MySQLServerService) GetAll() error {
	var err error

	mss.MySQLServers, err = mss.MySQLServerRepo.GetAll()

	return err
}

// GetByClusterID gets mysql servers with given cluster id
func (mss *MySQLServerService) GetByClusterID(clusterID int) error {
	var err error

	mss.MySQLServers, err = mss.MySQLServerRepo.GetByClusterID(clusterID)
	if err != nil {
		return err
	}

	return nil
}

// GetByID gets an mysql server entity that contains the given id from the middleware
func (mss *MySQLServerService) GetByID(id int) error {
	entity, err := mss.MySQLServerRepo.GetByID(id)
	if err != nil {
		return err
	}

	mss.MySQLServers = nil
	mss.MySQLServers = append(mss.MySQLServers, entity)

	return nil
}

// GetByHostInfo gets a mysql server with given host ip and port number
func (mss *MySQLServerService) GetByHostInfo(hostIP string, portNum int) error {
	mysqlServer, err := mss.MySQLServerRepo.GetByHostInfo(hostIP, portNum)
	if err != nil {
		return err
	}

	mss.MySQLServers = nil
	mss.MySQLServers = append(mss.MySQLServers, mysqlServer)

	return nil
}

// IsMaster returns if mysql server with given host ip and port number is a master node
func (mss *MySQLServerService) IsMaster(hostIP string, portNum int) (bool, error) {
	return mss.MySQLServerRepo.IsMaster(hostIP, portNum)
}

// GetMySQLClusterByID gets the mysql cluster of the given id
func (mss *MySQLServerService) GetMySQLClusterByID(id int) error {
	err := mss.GetByID(id)
	if err != nil {
		return err
	}

	mss.MySQLCluster, err = mss.GetMySQLServers()[constant.ZeroInt].GetMySQLCluster()

	return err
}

// Create creates a new mysql server entity and insert it into the middleware
func (mss *MySQLServerService) Create(fields map[string]interface{}) error {
	// generate new map
	_, clusterIDExists := fields[mysqlServerClusterIDStruct]
	_, serverNameExists := fields[mysqlServerServerNameStruct]
	_, ServiceNameExists := fields[mysqlServerServiceNameStruct]
	_, hostIPExists := fields[mysqlServerHostIPStruct]
	_, portNumExists := fields[mysqlServerPortNumStruct]
	_, deploymentTypeExists := fields[mysqlServerDeploymentTypeStruct]

	if !clusterIDExists || !serverNameExists || !ServiceNameExists || !hostIPExists || !portNumExists ||
		!deploymentTypeExists {
		return message.NewMessage(
			message.ErrFieldNotExists,
			fmt.Sprintf(
				"%s and %s and %s and %s and %s and %s",
				mysqlServerClusterIDStruct,
				mysqlServerServerNameStruct,
				mysqlServerServiceNameStruct,
				mysqlServerHostIPStruct,
				mysqlServerPortNumStruct,
				mysqlServerDeploymentTypeStruct,
			),
		)
	}

	// create a new entity
	mysqlServerInfo, err := NewMySQLServerInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	entity, err := mss.MySQLServerRepo.Create(mysqlServerInfo)
	if err != nil {
		return err
	}

	mss.MySQLServers = nil
	mss.MySQLServers = append(mss.MySQLServers, entity)

	return nil
}

// Update gets an mysql server entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (mss *MySQLServerService) Update(id int, fields map[string]interface{}) error {
	err := mss.GetByID(id)
	if err != nil {
		return err
	}
	err = mss.MySQLServers[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return mss.MySQLServerRepo.Update(mss.MySQLServers[constant.ZeroInt])
}

// Delete deletes the mysql server entity that contains the given id in the middleware
func (mss *MySQLServerService) Delete(id int) error {
	return mss.MySQLServerRepo.Delete(id)
}

// Marshal marshals service.Envs
func (mss *MySQLServerService) Marshal() ([]byte, error) {
	return mss.MarshalWithFields(mysqlServerMySQLServersStruct)
}

// MarshalWithFields marshals service.Envs with given fields
func (mss *MySQLServerService) MarshalWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(mss, fields...)
}
