package query

import (
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
	"testing"
)

var service = createService()

func createService() *Service {
	return NewServiceWithDefault(NewConfigWithDefault())
}

func TestService_All(t *testing.T) {
	TestService_GetConfig(t)
	TestService_GetQueries(t)
	TestService_GetByMySQLClusterID(t)
	TestService_GetByMySQLServerID(t)
	TestService_GetByDBID(t)
	TestService_GetBySQLID(t)
	TestService_Marshal(t)
	TestService_MarshalWithFields(t)
}

func TestService_GetConfig(t *testing.T) {
	asst := assert.New(t)

	limit := service.GetConfig().GetLimit()
	asst.Equal(DefaultLimit, limit, "test GetConfig() failed")
}

func TestService_GetQueries(t *testing.T) {
	asst := assert.New(t)

	entities := service.GetQueries()
	asst.Equal(service.queries, entities, "test GetQueries() failed")
}

func TestService_GetByMySQLClusterID(t *testing.T) {
	asst := assert.New(t)

	err := service.GetByMySQLClusterID(constant.DefaultRandomInt)
	asst.Nil(err, common.CombineMessageWithError("test GetByMySQLClusterID() failed", err))
}

func TestService_GetByMySQLServerID(t *testing.T) {
	asst := assert.New(t)

	err := service.GetByMySQLServerID(constant.DefaultRandomInt)
	asst.Nil(err, common.CombineMessageWithError("test GetByMySQLServerID() failed", err))
}

func TestService_GetByDBID(t *testing.T) {
	asst := assert.New(t)

	err := service.GetByDBID(constant.DefaultRandomInt, constant.DefaultRandomInt)
	asst.Nil(err, common.CombineMessageWithError("test GetByDBID() failed", err))
}

func TestService_GetBySQLID(t *testing.T) {
	asst := assert.New(t)

	err := service.GetBySQLID(constant.DefaultRandomInt, constant.DefaultRandomString)
	asst.Nil(err, common.CombineMessageWithError("test GetBySQLID() failed", err))
}

func TestService_Marshal(t *testing.T) {
	asst := assert.New(t)

	_, err := service.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
}

func TestService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	_, err := service.MarshalWithFields(queryQueriesStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields(fields ...string) failed", err))
}
