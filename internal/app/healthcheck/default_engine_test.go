package healthcheck

import (
	"testing"
	"time"

	"github.com/romberli/das/config"
	_ "github.com/romberli/das/internal/app/alert"
	"github.com/romberli/das/internal/dependency/healthcheck"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	testHealthcheckOperationID          = 1
	testHealthcheckMySQLServerID        = 1
	testHealthcheckResultUpdateStatus   = 1
	testHealthcheckResultAccuracyReview = 1
	testHealthcheckStep                 = time.Minute

	testSoarBin    = "../../../bin/soar"
	testSoarConfig = "../../../config/soar.yaml"

	testSMTPURL  = "smtp.163.com:465"
	testSMTPUser = "allinemailtest@163.com"
	testSMTPPass = "EHEBEAVXSVLXEMFM"
	testSMTPFrom = "allinemailtest@163.com"
)

func init() {
	testInitViper()
}

func testInitViper() {
	viper.Set(config.HealthcheckAlertOwnerTypeKey, config.HealthcheckAlertOwnerTypeAll)
	viper.Set(config.DBApplicationMySQLUserKey, config.DefaultDBApplicationMySQLUser)
	viper.Set(config.DBApplicationMySQLPassKey, config.DefaultDBApplicationMySQLPass)
	viper.Set(config.DBMonitorPrometheusUserKey, config.DefaultDBMonitorPrometheusUser)
	viper.Set(config.DBMonitorPrometheusPassKey, config.DefaultDBMonitorPrometheusPass)
	viper.Set(config.DBMonitorMySQLUserKey, config.DefaultDBMonitorMySQLUser)
	viper.Set(config.DBMonitorMySQLPassKey, config.DefaultDBMonitorMySQLPass)
	viper.Set(config.DBMonitorClickhouseUserKey, config.DefaultDBMonitorClickhouseUser)
	viper.Set(config.DBMonitorClickhousePassKey, config.DefaultDBMonitorClickhousePass)
	// alert
	viper.Set(config.AlertSMTPEnabledKey, true)
	viper.Set(config.AlertSMTPFormatKey, config.AlertSMTPTextFormat)
	viper.Set(config.AlertSMTPURLKey, testSMTPURL)
	viper.Set(config.AlertSMTPUserKey, testSMTPUser)
	viper.Set(config.AlertSMTPPassKey, testSMTPPass)
	viper.Set(config.AlertSMTPFromKey, testSMTPFrom)
	// sqladvisor
	viper.Set(config.SQLAdvisorSoarBinKey, testSoarBin)
	viper.Set(config.SQLAdvisorSoarConfigKey, testSoarConfig)
}

func TestDefaultEngine_All(t *testing.T) {
	TestDefaultEngineConfig_Validate(t)
	TestDefaultEngine_Run(t)
}

func TestDefaultEngineConfig_Validate(t *testing.T) {
	asst := assert.New(t)
	// load config
	sql := `
		select id, item_name, item_weight, low_watermark, high_watermark, unit, score_deduction_per_unit_high, max_score_deduction_high,
		score_deduction_per_unit_medium, max_score_deduction_medium, del_flag, create_time, last_update_time
		from t_hc_default_engine_config
		where del_flag = 0;
	`
	result, err := testDASRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Validate() failed", err))
	itemConfigList := make([]healthcheck.ItemConfig, result.RowNumber())
	for i := range itemConfigList {
		itemConfigList[i] = NewEmptyDefaultItemConfig()
	}
	err = result.MapToStructSlice(itemConfigList, constant.DefaultMiddlewareTag)
	asst.Nil(err, common.CombineMessageWithError("test Validate() failed", err))
	engineConfig := NewEmptyDefaultEngineConfig()
	for _, itemConfig := range itemConfigList {
		engineConfig.SetItemConfig(itemConfig.GetItemName(), itemConfig)
	}
	// validate config
	err = engineConfig.Validate()
	asst.Nil(err, common.CombineMessageWithError("test Validate() failed", err))
}

func TestDefaultEngine_Run(t *testing.T) {
	asst := assert.New(t)

	id, err := testDASRepo.InitOperation(
		testOperationInfo.GetUser().Identity(),
		testHealthcheckMySQLServerID,
		time.Now().Add(-constant.Week),
		time.Now(),
		testHealthcheckStep,
	)
	asst.Nil(err, common.CombineMessageWithError("test Run() failed", err))
	operationInfo := NewOperationInfo(
		id,
		testOperationInfo.GetUser(),
		testOperationInfo.GetApps(),
		testOperationInfo.GetMySQLServer(),
		testOperationInfo.GetMonitorSystem(),
		testOperationInfo.GetStartTime(),
		testOperationInfo.GetEndTime(),
		testOperationInfo.GetStep(),
	)

	de := newDefaultEngine(operationInfo, testDASRepo, testApplicationMySQLRepo, testPrometheusRepo, testQueryRepo)
	err = de.run()
	asst.Nil(err, common.CombineMessageWithError("test Run() failed", err))
	r, err := testDASRepo.GetResultByOperationID(de.GetOperationInfo().GetOperationID())
	asst.Nil(err, common.CombineMessageWithError("test Run() failed", err))
	t.Log(r.String())
}
