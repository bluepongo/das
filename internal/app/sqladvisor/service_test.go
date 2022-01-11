package sqladvisor

import (
	"testing"

	"github.com/romberli/go-util/common"
	"github.com/stretchr/testify/assert"
)

const (
	testFingerprint = "select * from t_meta_db_info where create_time<?"
	testSQLID       = "B95017DB61875675"
)

var service *Service

func init() {
	testInitDASMySQLPool()
	testInitViper()

	service = NewServiceWithDefault()
}

func TestService_All(t *testing.T) {
	TestService_GetFingerprint(t)
	TestService_GetFingerprint(t)
	TestService_GetSQLID(t)
	TestService_Advise(t)
}

func TestService_GetFingerprint(t *testing.T) {
	asst := assert.New(t)

	fingerprint := service.GetFingerprint(testSQLText)
	asst.Equal(testFingerprint, fingerprint, "test GetFingerprint() failed")
}

func TestService_GetSQLID(t *testing.T) {
	asst := assert.New(t)

	sqlID := service.GetSQLID(testSQLText)
	asst.Equal(testSQLID, sqlID, "test GetSQLID() failed")
}

func TestService_Advise(t *testing.T) {
	asst := assert.New(t)

	advice, err := service.Advise(testDBID, testSQLText)
	asst.Nil(err, common.CombineMessageWithError("test Advise() failed", err))
	asst.NotEmpty(advice, "test Advise() failed")
	t.Log(advice)
}
