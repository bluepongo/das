package healthcheck

import (
	"testing"
	"time"

	"github.com/romberli/das/internal/dependency/healthcheck"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testSleepTime = 5 * time.Second
	testLoginName = "zhangs"
)

var testService healthcheck.Service

func init() {
	testInitDASMySQLPool()
	testInitViper()

	testService = NewServiceWithDefault()
}

func deleteByOperationID(operationID int) error {
	tx, err := testDASRepo.Transaction()
	if err != nil {
		return err
	}
	err = tx.Begin()
	if err != nil {
		return err
	}

	sql := `delete from t_hc_result where operation_id = ?`
	_, err = testDASRepo.Execute(sql, operationID)
	if err != nil {
		return err
	}
	sql = `delete from t_hc_operation_history where id = ?`
	_, err = testDASRepo.Execute(sql, operationID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func TestService_All(t *testing.T) {
	TestService_GetResult(t)
	TestService_GetResultByOperationID(t)
	TestService_Check(t)
	TestService_ReviewAccuracy(t)
	TestService_Marshal(t)
	TestService_MarshalWithFields(t)
}

func TestService_GetResult(t *testing.T) {
	asst := assert.New(t)

	operationID, err := testService.Check(testHealthcheckMySQLServerID, time.Now().Add(-constant.Week), time.Now(), testHealthcheckStep, testLoginName)
	asst.Nil(err, common.CombineMessageWithError("test GetResult() failed", err))
	time.Sleep(testSleepTime)
	err = testService.GetResultByOperationID(operationID)
	asst.Nil(err, common.CombineMessageWithError("test GetResult() failed", err))
	t.Log(testService.GetResult().String())
	// delete
	err = deleteByOperationID(operationID)
	asst.Nil(err, common.CombineMessageWithError("test GetResult() failed", err))
}

func TestService_GetResultByOperationID(t *testing.T) {
	asst := assert.New(t)

	operationID, err := testService.Check(testHealthcheckMySQLServerID, time.Now().Add(-constant.Week), time.Now(), testHealthcheckStep, testLoginName)
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
	time.Sleep(testSleepTime)
	err = testService.GetResultByOperationID(operationID)
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
	t.Log(testService.GetResult().String())
	// delete
	err = deleteByOperationID(operationID)
	asst.Nil(err, common.CombineMessageWithError("test GetResultByOperationID() failed", err))
}

// bug
func TestService_Check(t *testing.T) {
	asst := assert.New(t)

	operationID, err := testService.Check(testHealthcheckMySQLServerID, time.Now().Add(-constant.Week), time.Now(), testHealthcheckStep, testLoginName)
	asst.Nil(err, common.CombineMessageWithError("test Check() failed", err))
	time.Sleep(testSleepTime)
	err = testService.GetResultByOperationID(operationID)
	asst.Nil(err, common.CombineMessageWithError("test Check() failed", err))
	t.Log(testService.GetResult().String())
	// delete
	err = deleteByOperationID(operationID)
	asst.Nil(err, common.CombineMessageWithError("test Check() failed", err))
}

func TestService_CheckByHostInfo(t *testing.T) {
	asst := assert.New(t)

	operationID, err := testService.CheckByHostInfo(
		testOperationInfo.GetMySQLServer().GetHostIP(),
		testOperationInfo.GetMySQLServer().GetPortNum(),
		time.Now().Add(-constant.Week),
		time.Now(),
		testHealthcheckStep,
		testLoginName,
	)
	asst.Nil(err, common.CombineMessageWithError("test CheckByHostInfo() failed", err))
	time.Sleep(testSleepTime)
	err = testService.GetResultByOperationID(operationID)
	asst.Nil(err, common.CombineMessageWithError("test CheckByHostInfo() failed", err))
	t.Log(testService.GetResult().String())
	// delete
	err = deleteByOperationID(operationID)
	asst.Nil(err, common.CombineMessageWithError("test CheckByHostInfo() failed", err))
}

func TestService_ReviewAccuracy(t *testing.T) {
	asst := assert.New(t)

	operationID, err := testService.Check(testHealthcheckMySQLServerID, time.Now().Add(-constant.Week), time.Now(), testHealthcheckStep, testLoginName)
	asst.Nil(err, common.CombineMessageWithError("test ReviewAccuracy() failed", err))
	time.Sleep(testSleepTime)
	err = testService.ReviewAccuracy(operationID, testResultUpdateAccuracyReview)
	asst.Nil(err, common.CombineMessageWithError("test ReviewAccuracy() failed", err))
	err = testService.GetResultByOperationID(operationID)
	asst.Nil(err, common.CombineMessageWithError("test ReviewAccuracy() failed", err))
	asst.Equal(testResultUpdateAccuracyReview, testService.GetResult().GetAccuracyReview(), "test ReviewAccuracy() failed")
	// delete
	err = deleteByOperationID(operationID)
	asst.Nil(err, common.CombineMessageWithError("test ReviewAccuracy() failed", err))
}

func TestService_Marshal(t *testing.T) {
	asst := assert.New(t)

	operationID, err := testService.Check(testHealthcheckMySQLServerID, time.Now().Add(-constant.Week), time.Now(), testHealthcheckStep, testLoginName)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	time.Sleep(testSleepTime)
	err = testService.GetResultByOperationID(operationID)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	_, err = testService.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	// t.Log(jsonBytes)
	// delete
	err = deleteByOperationID(operationID)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
}

func TestService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)
	operationID, err := testService.Check(testHealthcheckMySQLServerID, time.Now().Add(-constant.Week), time.Now(), testHealthcheckStep, testLoginName)
	asst.Nil(err, common.CombineMessageWithError("test healthcheckResultStruct() failed", err))
	time.Sleep(testSleepTime)
	err = testService.GetResultByOperationID(operationID)
	asst.Nil(err, common.CombineMessageWithError("test healthcheckResultStruct() failed", err))
	_, err = testService.MarshalWithFields(healthcheckResultStruct)
	asst.Nil(err, common.CombineMessageWithError("test healthcheckResultStruct() failed", err))
	// t.Log(jsonBytes)
	// delete
	err = deleteByOperationID(operationID)
	asst.Nil(err, common.CombineMessageWithError("test healthcheckResultStruct() failed", err))
}
