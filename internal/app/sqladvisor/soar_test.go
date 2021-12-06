package sqladvisor

import (
	"testing"

	"github.com/romberli/das/config"
	"github.com/romberli/go-util/common"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	testSoarBin    = "/Users/romber/work/source_code/go/src/github.com/romberli/das/bin/soar"
	testConfigFile = "/Users/romber/work/source_code/go/src/github.com/romberli/das/config/soar.yaml"

	testDBSoarMySQLUser = "root"
	testDBSoarMySQLPass = "root"
)

var advisor *DefaultAdvisor

func init() {
	testInitViper()

	advisor = NewDefaultAdvisor(testSoarBin, testConfigFile)
}

func testInitViper() {
	viper.Set(config.SQLAdvisorSoarBinKey, testSoarBin)
	viper.Set(config.SQLAdvisorSoarConfigKey, testConfigFile)
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
	asst.Nil(err, common.CombineMessageWithError("test Advise() failed", err))
	asst.NotEmpty(advice, "test Advise() failed")
	t.Log(message)
	t.Log(advice)
}
