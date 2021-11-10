package sqladvisor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testDBSoarMySQLUser = "root"
	testDBSoarMySQLPass = "root"
)

var advisor *DefaultAdvisor

func init() {
	advisor = NewDefaultAdvisor(testSoarBin, testConfigFile)
}

func TestDefaultAdvisor_All(t *testing.T) {
	TestDefaultAdvisor_GetFingerprint(t)
	TestDefaultAdvisor_GetSQLID(t)
	TestDefaultAdvisor_Advise(t)
}

func TestDefaultAdvisor_GetFingerprint(t *testing.T) {
	asst := assert.New(t)

	fingerprint := advisor.GetFingerprint(testSQLText)
	asst.Equal(testFingerprint, fingerprint, "test GetFingerprint() failed")
}

func TestDefaultAdvisor_GetSQLID(t *testing.T) {
	asst := assert.New(t)

	sqlID := advisor.GetSQLID(testSQLText)
	asst.Equal(testSQLID, sqlID, "test GetSQLID() failed")
}

func TestDefaultAdvisor_Advise(t *testing.T) {
	asst := assert.New(t)

	advice, message, err := advisor.advise(testDBID, testSQLText, testDBSoarMySQLUser, testDBSoarMySQLPass)
	asst.Nil(err, "test Advise() failed")
	asst.NotEmpty(advice, "test Advise() failed")
	t.Log(message)
	t.Log(advice)
}
