package metadata

import (
	"encoding/json"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"

	"github.com/romberli/das/pkg/message"

	"github.com/romberli/das/internal/dependency"
)

const (
	clusterIDStruct      = "ClusterID"
	serverNameStruct     = "ServerName"
	hostIPStruct         = "HostIP"
	portNumStruct        = "PortNum"
	deploymentTypeStruct = "DeploymentType"
	versionStruct        = "Version"
)

var _ dependency.Service = (*MySQLServerService)(nil)

// MySQLServerService implements Service interface
type MySQLServerService struct {
	dependency.Repository
	Entities []dependency.Entity
}

// NewMySQLServerService returns a new *MySQLServerService
func NewMySQLServerService(repo dependency.Repository) *MySQLServerService {
	return &MySQLServerService{repo, []dependency.Entity{}}
}

// NewMySQLServerServiceWithDefault returns a new *MySQLServerService with default repository
func NewMySQLServerServiceWithDefault() *MySQLServerService {
	return NewMySQLServerService(NewMySQLServerRepoWithGlobal())
}

// GetEntities returns entities of the service
func (mss *MySQLServerService) GetEntities() []dependency.Entity {
	entityList := make([]dependency.Entity, len(mss.Entities))
	for i := range entityList {
		entityList[i] = mss.Entities[i]
	}

	return entityList
}

// GetAll gets all mysql server entities from the middleware
func (mss *MySQLServerService) GetAll() error {
	var err error
	mss.Entities, err = mss.Repository.GetAll()

	return err
}

// GetByID gets an mysql server entity that contains the given id from the middleware
func (mss *MySQLServerService) GetByID(id string) error {
	entity, err := mss.Repository.GetByID(id)
	if err != nil {
		return err
	}

	mss.Entities = append(mss.Entities, entity)

	return err
}

// Create creates a new mysql server entity and insert it into the middleware
func (mss *MySQLServerService) Create(fields map[string]interface{}) error {
	// generate new map
	if _, ok := fields[clusterIDStruct]; !ok {
		return message.NewMessage(message.ErrFieldNotExists, clusterIDStruct)
	}
	if _, ok := fields[serverNameStruct]; !ok {
		return message.NewMessage(message.ErrFieldNotExists, serverNameStruct)
	}
	if _, ok := fields[hostIPStruct]; !ok {
		return message.NewMessage(message.ErrFieldNotExists, hostIPStruct)
	}
	if _, ok := fields[portNumStruct]; !ok {
		return message.NewMessage(message.ErrFieldNotExists, portNumStruct)
	}
	if _, ok := fields[deploymentTypeStruct]; !ok {
		return message.NewMessage(message.ErrFieldNotExists, deploymentTypeStruct)
	}
	if _, ok := fields[versionStruct]; !ok {
		return message.NewMessage(message.ErrFieldNotExists, versionStruct)
	}
	// create a new entity
	mysqlServerInfo, err := NewMySQLServerInfoWithMapAndRandom(fields)
	if err != nil {
		return err
	}
	// insert into middleware
	entity, err := mss.Repository.Create(mysqlServerInfo)
	if err != nil {
		return err
	}

	mss.Entities = append(mss.Entities, entity)
	return nil
}

// Update gets an mysql server entity that contains the given id from the middleware,
// and then update its fields that was specified in fields argument,
// key is the filed name and value is the new field value,
// it saves the changes to the middleware
func (mss *MySQLServerService) Update(id string, fields map[string]interface{}) error {
	err := mss.GetByID(id)
	if err != nil {
		return err
	}
	err = mss.Entities[constant.ZeroInt].Set(fields)
	if err != nil {
		return err
	}

	return mss.Repository.Update(mss.Entities[constant.ZeroInt])
}

// Delete deletes the mysql server entity that contains the given id in the middleware
func (mss *MySQLServerService) Delete(id string) error {
	return mss.Repository.Delete(id)
}

// Marshal marshals service.Envs
func (mss *MySQLServerService) Marshal() ([]byte, error) {
	return json.Marshal(mss.Entities)
}

// MarshalWithFields marshals service.Envs with given fields
func (mss *MySQLServerService) MarshalWithFields(fields ...string) ([]byte, error) {
	interfaceList := make([]interface{}, len(mss.Entities))
	for i := range interfaceList {
		entity, err := common.CopyStructWithFields(mss.Entities[i], fields...)
		if err != nil {
			return nil, err
		}
		interfaceList[i] = entity
	}

	return json.Marshal(interfaceList)
}
