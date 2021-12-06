package healthcheck

import (
	"testing"
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
)

const (
	testResultOperationID                       = testHealthcheckOperationID
	testResultHostIP                            = "192.168.10.219"
	testResultPortNum                           = 3306
	testResultWeightedAverageScore              = 1
	testResultDBConfigScore                     = 1
	testResultDBConfigData                      = "db config data"
	testResultDBConfigAdvice                    = "db config advice"
	testResultAvgBackupFailedRatioScore         = 80
	testResultAvgBackupFailedRatioData          = "cpu usage data"
	testResultAvgBackupFailedRatioHigh          = "cpu usage high"
	testResultStatisticFailedRatioScore         = 80
	testResultStatisticFailedRatioData          = "cpu usage data"
	testResultStatisticFailedRatioHigh          = "cpu usage high"
	testResultCPUUsageScore                     = 80
	testResultCPUUsageData                      = "cpu usage data"
	testResultCPUUsageHigh                      = "cpu usage high"
	testResultIOUtilScore                       = 80
	testResultIOUtilData                        = "io util data"
	testResultIOUtilHigh                        = "io util high"
	testResultDiskCapacityUsageScore            = 80
	testResultDiskCapacityUsageData             = "disk capacity usage data"
	testResultDiskCapacityUsageHigh             = "disk capacity usage high"
	testResultConnectionUsageScore              = 80
	testResultConnectionUsageData               = "connection usage data"
	testResultConnectionUsageHigh               = "connection usage high"
	testResultAverageActiveSessionPercentsScore = 80
	testResultAverageActiveSessionPercentsData  = "average active session num data"
	testResultAverageActiveSessionPercentsHigh  = "average active session num high"
	testResultCacheMissRatioScore               = 80
	testResultCacheMissRatioData                = "cache miss ratio data"
	testResultCacheMissRatioHigh                = "cache miss ratio high"
	testResultTableRowsScore                    = 80
	testResultTableRowsData                     = "table rows data"
	testResultTableRowsHigh                     = "table rows high"
	testResultTableSizeScore                    = 80
	testResultTableSizeData                     = "table size data"
	testResultTableSizeHigh                     = "table size high"
	testResultSlowQueryScore                    = 80
	testResultSlowQueryData                     = "slow query data"
	testResultSlowQueryAdvice                   = "slow query advice"
	testResultAccuracyReview                    = 0
	testResultDelFlag                           = 0

	testResultNewOperationID       = 2
	testResultUpdateOperationID    = 3
	testResultUpdateAccuracyReview = 1
)

var testResult *Result

func init() {
	testResult = testInitResult()
}

func testInitResult() *Result {
	return NewResult(
		testDASRepo,
		testResultOperationID,
		testResultHostIP,
		testResultPortNum,
		testResultWeightedAverageScore,
		testResultDBConfigScore,
		testResultDBConfigData,
		testResultDBConfigAdvice,
		testResultAvgBackupFailedRatioScore,
		testResultAvgBackupFailedRatioData,
		testResultAvgBackupFailedRatioHigh,
		testResultStatisticFailedRatioScore,
		testResultStatisticFailedRatioData,
		testResultStatisticFailedRatioHigh,
		testResultCPUUsageScore,
		testResultCPUUsageData,
		testResultCPUUsageHigh,
		testResultIOUtilScore,
		testResultIOUtilData,
		testResultIOUtilHigh,
		testResultDiskCapacityUsageScore,
		testResultDiskCapacityUsageData,
		testResultDiskCapacityUsageHigh,
		testResultConnectionUsageScore,
		testResultConnectionUsageData,
		testResultConnectionUsageHigh,
		testResultAverageActiveSessionPercentsScore,
		testResultAverageActiveSessionPercentsData,
		testResultAverageActiveSessionPercentsHigh,
		testResultCacheMissRatioScore,
		testResultCacheMissRatioData,
		testResultCacheMissRatioHigh,
		testResultTableRowsScore,
		testResultTableRowsData,
		testResultTableRowsHigh,
		testResultTableSizeScore,
		testResultTableSizeData,
		testResultTableSizeHigh,
		testResultSlowQueryScore,
		testResultSlowQueryData,
		testResultSlowQueryAdvice,
	)
}

func TestResultAll(t *testing.T) {
	TestResult_Identity(t)
	TestResult_GetOperationID(t)
	TestResult_GetWeightedAverageScore(t)
	TestResult_GetDBConfigScore(t)
	TestResult_GetDBConfigData(t)
	TestResult_GetDBConfigAdvice(t)
	TestResult_GetAvgBackupFailedRatioScore(t)
	TestResult_GetAvgBackupFailedRatioData(t)
	TestResult_GetAvgBackupFailedRatioHigh(t)
	TestResult_GetStatisticFailedRatioScore(t)
	TestResult_GetStatisticFailedRatioData(t)
	TestResult_GetStatisticFailedRatioHigh(t)
	TestResult_GetCPUUsageScore(t)
	TestResult_GetCPUUsageData(t)
	TestResult_GetCPUUsageHigh(t)
	TestResult_GetIOUtilScore(t)
	TestResult_GetIOUtilData(t)
	TestResult_GetIOUtilHigh(t)
	TestResult_GetDiskCapacityUsageScore(t)
	TestResult_GetDiskCapacityUsageData(t)
	TestResult_GetDiskCapacityUsageHigh(t)
	TestResult_GetConnectionUsageScore(t)
	TestResult_GetConnectionUsageData(t)
	TestResult_GetConnectionUsageHigh(t)
	TestResult_GetAverageActiveSessionPercentsScore(t)
	TestResult_GetAverageActiveSessionPercentsData(t)
	TestResult_GetAverageActiveSessionPercentsHigh(t)
	TestResult_GetCacheMissRatioScore(t)
	TestResult_GetCacheMissRatioData(t)
	TestResult_GetCacheMissRatioHigh(t)
	TestResult_GetTableRowsScore(t)
	TestResult_GetTableRowsData(t)
	TestResult_GetTableRowsHigh(t)
	TestResult_GetTableSizeScore(t)
	TestResult_GetTableSizeData(t)
	TestResult_GetTableSizeHigh(t)
	TestResult_GetSlowQueryScore(t)
	TestResult_GetSlowQueryData(t)
	TestResult_GetSlowQueryAdvice(t)
	TestResult_GetAccuracyReview(t)
	TestResult_GetDelFlag(t)
	TestResult_GetCreateTime(t)
	TestResult_GetLastUpdateTime(t)
	TestResult_String(t)
	TestResult_Set(t)
	TestResult_MarshalJSON(t)
	TestResult_MarshalJSONWithFields(t)
}

func TestResult_Identity(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(constant.ZeroInt, testResult.Identity(), "test Identity() failed")
}

func TestResult_GetOperationID(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultOperationID, testResult.GetOperationID(), "test GetOperationID() failed")
}

func TestResult_GetHostIP(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultHostIP, testResult.GetHostIP(), "test GetHostIP() failed")
}

func TestResult_GetPortNum(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultPortNum, testResult.GetPortNum(), "test GetPortNum() failed")
}

func TestResult_GetWeightedAverageScore(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultWeightedAverageScore, testResult.GetWeightedAverageScore(), "test GetWeightedAverageScore() failed")
}

func TestResult_GetDBConfigScore(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultDBConfigScore, testResult.GetDBConfigScore(), "test GetDBConfigScore() failed")
}

func TestResult_GetDBConfigData(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultDBConfigData, testResult.GetDBConfigData(), "test GetDBConfigData() failed")
}

func TestResult_GetDBConfigAdvice(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultDBConfigAdvice, testResult.GetDBConfigAdvice(), "test GetDBConfigAdvice() failed")
}

func TestResult_GetAvgBackupFailedRatioScore(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultAvgBackupFailedRatioScore, testResult.GetAvgBackupFailedRatioScore(), "test GetAvgBackupFailedRatioScore() failed")
}

func TestResult_GetAvgBackupFailedRatioData(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultAvgBackupFailedRatioData, testResult.GetAvgBackupFailedRatioData(), "test GetAvgBackupFailedRatioData() failed")
}

func TestResult_GetAvgBackupFailedRatioHigh(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultAvgBackupFailedRatioHigh, testResult.GetAvgBackupFailedRatioHigh(), "test GetAvgBackupFailedRatioHigh() failed")
}

func TestResult_GetStatisticFailedRatioScore(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultStatisticFailedRatioScore, testResult.GetStatisticFailedRatioScore(), "test GetStatisticFailedRatioScore() failed")
}

func TestResult_GetStatisticFailedRatioData(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultStatisticFailedRatioData, testResult.GetStatisticFailedRatioData(), "test GetStatisticFailedRatioData() failed")
}

func TestResult_GetStatisticFailedRatioHigh(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultStatisticFailedRatioHigh, testResult.GetStatisticFailedRatioHigh(), "test GetStatisticFailedRatioHigh() failed")
}

func TestResult_GetCPUUsageScore(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultCPUUsageScore, testResult.GetCPUUsageScore(), "test GetCPUUsageScore() failed")
}

func TestResult_GetCPUUsageData(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultCPUUsageData, testResult.GetCPUUsageData(), "test GetCPUUsageData() failed")
}

func TestResult_GetCPUUsageHigh(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultCPUUsageHigh, testResult.GetCPUUsageHigh(), "test GetCPUUsageHigh() failed")
}

func TestResult_GetIOUtilScore(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultIOUtilScore, testResult.GetIOUtilScore(), "test GetIOUtilScore() failed")
}

func TestResult_GetIOUtilData(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultIOUtilData, testResult.GetIOUtilData(), "test GetIOUtilData() failed")
}

func TestResult_GetIOUtilHigh(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultIOUtilHigh, testResult.GetIOUtilHigh(), "test GetIOUtilHigh() failed")
}

func TestResult_GetDiskCapacityUsageScore(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultDiskCapacityUsageScore, testResult.GetDiskCapacityUsageScore(), "test GetDiskCapacityUsageScore() failed")
}

func TestResult_GetDiskCapacityUsageData(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultDiskCapacityUsageData, testResult.GetDiskCapacityUsageData(), "test GetDiskCapacityUsageData() failed")
}

func TestResult_GetDiskCapacityUsageHigh(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultDiskCapacityUsageHigh, testResult.GetDiskCapacityUsageHigh(), "test GetDiskCapacityUsageHigh() failed")
}

func TestResult_GetConnectionUsageScore(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultConnectionUsageScore, testResult.GetConnectionUsageScore(), "test GetConnectionUsageScore() failed")
}

func TestResult_GetConnectionUsageData(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultConnectionUsageData, testResult.GetConnectionUsageData(), "test GetConnectionUsageData() failed")
}

func TestResult_GetConnectionUsageHigh(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultConnectionUsageHigh, testResult.GetConnectionUsageHigh(), "test GetConnectionUsageHigh() failed")
}

func TestResult_GetAverageActiveSessionPercentsScore(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultAverageActiveSessionPercentsScore, testResult.GetAverageActiveSessionPercentsScore(), "test GetAverageActiveSessionPercentsScore() failed")
}

func TestResult_GetAverageActiveSessionPercentsData(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultAverageActiveSessionPercentsData, testResult.GetAverageActiveSessionPercentsData(), "test GetAverageActiveSessionPercentsData() failed")
}

func TestResult_GetAverageActiveSessionPercentsHigh(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultAverageActiveSessionPercentsHigh, testResult.GetAverageActiveSessionPercentsHigh(), "test GetAverageActiveSessionPercentsHigh() failed")
}

func TestResult_GetCacheMissRatioScore(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultCacheMissRatioScore, testResult.GetCacheMissRatioScore(), "test GetCacheMissRatioScore() failed")
}

func TestResult_GetCacheMissRatioData(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultCacheMissRatioData, testResult.GetCacheMissRatioData(), "test GetCacheMissRatioData() failed")
}

func TestResult_GetCacheMissRatioHigh(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultCacheMissRatioHigh, testResult.GetCacheMissRatioHigh(), "test GetCacheMissRatioHigh() failed")
}

func TestResult_GetTableRowsScore(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultTableRowsScore, testResult.GetTableRowsScore(), "test GetTableRowsScore() failed")
}

func TestResult_GetTableRowsData(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultTableRowsData, testResult.GetTableRowsData(), "test GetTableRowsData() failed")
}

func TestResult_GetTableRowsHigh(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultTableRowsHigh, testResult.GetTableRowsHigh(), "test GetTableRowsHigh() failed")
}

func TestResult_GetTableSizeScore(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultTableSizeScore, testResult.GetTableSizeScore(), "test GetTableSizeScore() failed")
}

func TestResult_GetTableSizeData(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultTableSizeData, testResult.GetTableSizeData(), "test GetTableSizeData() failed")
}

func TestResult_GetTableSizeHigh(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultTableSizeHigh, testResult.GetTableSizeHigh(), "test GetTableSizeHigh() failed")
}

func TestResult_GetSlowQueryScore(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultSlowQueryScore, testResult.GetSlowQueryScore(), "test GetSlowQueryScore() failed")
}

func TestResult_GetSlowQueryData(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultSlowQueryData, testResult.GetSlowQueryData(), "test GetSlowQueryData() failed")
}

func TestResult_GetSlowQueryAdvice(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultSlowQueryAdvice, testResult.GetSlowQueryAdvice(), "test GetSlowQueryAdvice() failed")
}

func TestResult_GetAccuracyReview(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultAccuracyReview, testResult.GetAccuracyReview(), "test GetAccuracyReview() failed")
}

func TestResult_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	asst.Equal(testResultDelFlag, testResult.GetDelFlag(), "test GetDelFlag() failed")
}

func TestResult_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	createTime := time.Now()
	err := testResult.Set(map[string]interface{}{resultCreateTimeStruct: createTime})
	asst.Nil(err, common.CombineMessageWithError("test GetCreateTime() failed", err))
	asst.Equal(createTime, testResult.GetCreateTime(), "test GetCreateTime() failed")
}

func TestResult_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	lastUpdateTime := time.Now()
	err := testResult.Set(map[string]interface{}{resultLastUpdateTimeStruct: lastUpdateTime})
	asst.Nil(err, common.CombineMessageWithError("test GetLastUpdateTime() failed", err))
	asst.Equal(lastUpdateTime, testResult.GetLastUpdateTime(), "test GetLastUpdateTime() failed")
}

func TestResult_String(t *testing.T) {
	t.Log(testResult.String())
}

func TestResult_Set(t *testing.T) {
	asst := assert.New(t)

	err := testResult.Set(map[string]interface{}{resultOperationIDStruct: testResultUpdateOperationID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testResultUpdateOperationID, testResult.GetOperationID(), "test Set() failed")
	err = testResult.Set(map[string]interface{}{resultOperationIDStruct: testResultOperationID})
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	asst.Equal(testResultOperationID, testResult.GetOperationID(), "test Set() failed")
}

func TestResult_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testResult.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	t.Log(string(jsonBytes))
}

func TestResult_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	jsonBytes, err := testResult.MarshalJSONWithFields(resultOperationIDStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	t.Log(string(jsonBytes))
}
