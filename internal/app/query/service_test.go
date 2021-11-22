package query

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/stretchr/testify/assert"
)

var testService *Service

func init() {
	testPMMVersion = 2
	testInitMySQLInfo()
	testInitDASMySQLPool()
	testInitViper()

	testService = newService(NewConfigWithDefault(), testDASRepo)
}

func TestService_All(t *testing.T) {
	testPMMVersion = 1
	TestService_GetByMySQLClusterID(t)
	TestService_GetByMySQLServerID(t)
	TestService_GetByDBID(t)
	TestService_GetBySQLID(t)
	TestService_Marshal(t)

	testPMMVersion = 2
	TestService_GetByMySQLClusterID(t)
	TestService_GetByMySQLServerID(t)
	TestService_GetByDBID(t)
	TestService_GetBySQLID(t)
	TestService_Marshal(t)
}

func TestService_GetByMySQLClusterID(t *testing.T) {
	asst := assert.New(t)

	err := testService.GetByMySQLClusterID(testMySQLClusterID)
	asst.Nil(err, common.CombineMessageWithError("test GetByMySQLClusterID() failed", err))
}

func TestService_GetByMySQLServerID(t *testing.T) {
	asst := assert.New(t)

	err := testService.GetByMySQLServerID(testMySQLServerID)
	asst.Nil(err, common.CombineMessageWithError("test GetByMySQLServerID() failed", err))
}

func TestService_GetByDBID(t *testing.T) {
	asst := assert.New(t)

	err := testService.GetByDBID(testMySQLServerID, testDBID)
	asst.Nil(err, common.CombineMessageWithError("test GetByDBID() failed", err))
}

func TestService_GetBySQLID(t *testing.T) {
	asst := assert.New(t)

	err := testService.GetBySQLID(testMySQLServerID, testSQLID)
	asst.Nil(err, common.CombineMessageWithError("test GetBySQLID() failed", err))
}

func TestService_Marshal(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testService.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	t.Log(string(jsonBytes))
}
