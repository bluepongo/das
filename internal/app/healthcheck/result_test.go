package healthcheck

import (
	"testing"
	"time"

	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
)

const (
	defaultResultDBAddress = "192.168.10.210:3306"
	defaultResultHostIP    = "127.0.0.1"
	defaultResultPortNum   = 3306
	defaultResultDBName    = "das"
	defaultResultDBUser    = "root"
	defaultResultDBPass    = "root"

	defaultResultID                                = 1
	defaultResultOperationID                       = 1
	defaultResultWeightedAverageScore              = 1
	defaultResultDBConfigScore                     = 1
	defaultResultDBConfigData                      = "db config data"
	defaultResultDBConfigAdvice                    = "db config advice"
	defaultResultBackupScore                       = 80
	defaultResultBackupData                        = "cpu usage data"
	defaultResultBackupHigh                        = "cpu usage high"
	defaultResultStatisticScore                    = 80
	defaultResultStatisticData                     = "cpu usage data"
	defaultResultStatisticHigh                     = "cpu usage high"
	defaultResultCPUUsageScore                     = 80
	defaultResultCPUUsageData                      = "cpu usage data"
	defaultResultCPUUsageHigh                      = "cpu usage high"
	defaultResultIOUtilScore                       = 80
	defaultResultIOUtilData                        = "io util data"
	defaultResultIOUtilHigh                        = "io util high"
	defaultResultDiskCapacityUsageScore            = 80
	defaultResultDiskCapacityUsageData             = "disk capacity usage data"
	defaultResultDiskCapacityUsageHigh             = "disk capacity usage high"
	defaultResultConnectionUsageScore              = 80
	defaultResultConnectionUsageData               = "connection usage data"
	defaultResultConnectionUsageHigh               = "connection usage high"
	defaultResultAverageActiveSessionPercentsScore = 80
	defaultResultAverageActiveSessionPercentsData  = "average active session num data"
	defaultResultAverageActiveSessionPercentsHigh  = "average active session num high"
	defaultResultCacheMissRatioScore               = 80
	defaultResultCacheMissRatioData                = "cache miss ratio data"
	defaultResultCacheMissRatioHigh                = "cache miss ratio high"
	defaultResultTableRowsScore                    = 80
	defaultResultTableRowsData                     = "table rows data"
	defaultResultTableRowsHigh                     = "table rows high"
	defaultResultTableSizeScore                    = 80
	defaultResultTableSizeData                     = "table size data"
	defaultResultTableSizeHigh                     = "table size high"
	defaultResultSlowQueryScore                    = 80
	defaultResultSlowQueryData                     = "slow query data"
	defaultResultSlowQueryAdvice                   = "slow query advice"
	defaultResultAccuracyReview                    = 0
	defaultResultDelFlag                           = 0
)

func rInitRepository() *DASRepo {
	pool, err := mysql.NewPoolWithDefault(defaultResultDBAddress, defaultResultDBName, defaultResultDBUser, defaultResultDBPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initRepository() failed", err))
		return nil
	}
	return NewDASRepo(pool)
}

var rRepo = rInitRepository()

func rCreateService() (*Service, error) {
	var result = NewResult(rRepo,
		defaultResultOperationID, defaultResultWeightedAverageScore,
		defaultResultDBConfigScore, defaultResultDBConfigData, defaultResultDBConfigAdvice,
		defaultResultBackupScore, defaultResultBackupData, defaultResultBackupHigh,
		defaultResultStatisticScore, defaultResultStatisticData, defaultResultStatisticHigh,
		defaultResultCPUUsageScore, defaultResultCPUUsageData, defaultResultCPUUsageHigh,
		defaultResultIOUtilScore, defaultResultIOUtilData, defaultResultIOUtilHigh,
		defaultResultDiskCapacityUsageScore, defaultResultDiskCapacityUsageData, defaultResultDiskCapacityUsageHigh,
		defaultResultConnectionUsageScore, defaultResultConnectionUsageData, defaultResultConnectionUsageHigh,
		defaultResultAverageActiveSessionPercentsScore, defaultResultAverageActiveSessionPercentsData, defaultResultAverageActiveSessionPercentsHigh,
		defaultResultCacheMissRatioScore, defaultResultCacheMissRatioData, defaultResultCacheMissRatioHigh,
		defaultResultTableRowsScore, defaultResultTableRowsData, defaultResultTableRowsHigh,
		defaultResultTableSizeScore, defaultResultTableSizeData, defaultResultTableSizeHigh,
		defaultResultSlowQueryScore, defaultResultSlowQueryData, defaultResultSlowQueryAdvice)
	err := rRepo.SaveResult(result)
	if err != nil {
		return nil, err
	}
	return &Service{
		DASRepo: rRepo,
		Result:  result,
	}, nil
}

func rDeleteHCResultByOperationID(operationID int) error {
	sql := `delete from t_hc_result where operation_id = ?`
	_, err := rRepo.Execute(sql, operationID)
	return err
}

func TestResultAll(t *testing.T) {
	TestResult_Identity(t)
	TestResult_GetOperationID(t)
	TestResult_GetWeightedAverageScore(t)
	TestResult_GetDBConfigScore(t)
	TestResult_GetDBConfigData(t)
	TestResult_GetDBConfigAdvice(t)
	TestResult_GetBackupScore(t)
	TestResult_GetBackupData(t)
	TestResult_GetBackupHigh(t)
	TestResult_GetStatisticScore(t)
	TestResult_GetStatisticData(t)
	TestResult_GetStatisticHigh(t)
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
	TestResult_Set(t)
	TestResult_MarshalJSON(t)
	TestResult_MarshalJSONWithFields(t)
}

func TestResult_Identity(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test Identity() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test Identity() failed", err))
	result := service.GetResult()
	id := result.Identity()
	asst.IsType(defaultResultID, id, "test Identity() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test Identity() failed", err))
}

func TestResult_GetOperationID(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetOperationID() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetOperationID() failed", err))
	result := service.GetResult()
	operationID := result.GetOperationID()
	asst.Equal(defaultResultOperationID, operationID, "test GetOperationID() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetOperationID() failed", err))
}

func TestResult_GetWeightedAverageScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetWeightedAverageScore() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetWeightedAverageScore() failed", err))
	result := service.GetResult()
	weightedAverageScore := result.GetWeightedAverageScore()
	asst.Equal(defaultResultWeightedAverageScore, weightedAverageScore, "test GetWeightedAverageScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetWeightedAverageScore() failed", err))
}

func TestResult_GetDBConfigScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigScore() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigScore() failed", err))
	result := service.GetResult()
	dbConfigScore := result.GetDBConfigScore()
	asst.Equal(defaultResultDBConfigScore, dbConfigScore, "test GetDBConfigScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigScore() failed", err))
}

func TestResult_GetDBConfigData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigData() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigData() failed", err))
	result := service.GetResult()
	dbConfigData := result.GetDBConfigData()
	asst.Equal(defaultResultDBConfigData, dbConfigData, "test GetDBConfigData() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigData() failed", err))
}

func TestResult_GetDBConfigAdvice(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigAdvice() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigAdvice() failed", err))
	result := service.GetResult()
	dbConfigAdvice := result.GetDBConfigAdvice()
	asst.Equal(defaultResultDBConfigAdvice, dbConfigAdvice, "test GetDBConfigAdvice() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDBConfigAdvice() failed", err))
}

func TestResult_GetBackupScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetBackupScore() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetBackupScore() failed", err))
	result := service.GetResult()
	BackupScore := result.GetBackupScore()
	asst.Equal(defaultResultBackupScore, BackupScore, "test GetBackupScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetBackupScore() failed", err))
}

func TestResult_GetBackupData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetBackupData() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetBackupData() failed", err))
	result := service.GetResult()
	BackupData := result.GetBackupData()
	asst.Equal(defaultResultBackupData, BackupData, "test GetBackupData() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetBackupData() failed", err))
}

func TestResult_GetBackupHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetBackupHigh() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetBackupHigh() failed", err))
	result := service.GetResult()
	BackupHigh := result.GetBackupHigh()
	asst.Equal(defaultResultBackupHigh, BackupHigh, "test GetBackupHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetBackupHigh() failed", err))
}

func TestResult_GetStatisticScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetStatisticScore() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetStatisticScore() failed", err))
	result := service.GetResult()
	StatisticScore := result.GetStatisticScore()
	asst.Equal(defaultResultStatisticScore, StatisticScore, "test GetStatisticScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetStatisticScore() failed", err))
}

func TestResult_GetStatisticData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetStatisticData() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetStatisticData() failed", err))
	result := service.GetResult()
	StatisticData := result.GetStatisticData()
	asst.Equal(defaultResultStatisticData, StatisticData, "test GetStatisticData() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetStatisticData() failed", err))
}

func TestResult_GetStatisticHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetStatisticHigh() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetStatisticHigh() failed", err))
	result := service.GetResult()
	StatisticHigh := result.GetStatisticHigh()
	asst.Equal(defaultResultStatisticHigh, StatisticHigh, "test GetStatisticHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetStatisticHigh() failed", err))
}

func TestResult_GetCPUUsageScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageScore() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageScore() failed", err))
	result := service.GetResult()
	cpuUsageScore := result.GetCPUUsageScore()
	asst.Equal(defaultResultCPUUsageScore, cpuUsageScore, "test GetCPUUsageScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageScore() failed", err))
}

func TestResult_GetCPUUsageData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageData() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageData() failed", err))
	result := service.GetResult()
	cpuUsageData := result.GetCPUUsageData()
	asst.Equal(defaultResultCPUUsageData, cpuUsageData, "test GetCPUUsageData() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageData() failed", err))
}

func TestResult_GetCPUUsageHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageHigh() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageHigh() failed", err))
	result := service.GetResult()
	cpuUsageHigh := result.GetCPUUsageHigh()
	asst.Equal(defaultResultCPUUsageHigh, cpuUsageHigh, "test GetCPUUsageHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCPUUsageHigh() failed", err))
}

func TestResult_GetIOUtilScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilScore() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilScore() failed", err))
	result := service.GetResult()
	ioUtilScore := result.GetIOUtilScore()
	asst.Equal(defaultResultIOUtilScore, ioUtilScore, "test GetIOUtilScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilScore() failed", err))
}

func TestResult_GetIOUtilData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilScore() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilScore() failed", err))
	result := service.GetResult()
	ioUtilData := result.GetIOUtilData()
	asst.Equal(defaultResultIOUtilData, ioUtilData, "test GetIOUtilData() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilScore() failed", err))
}

func TestResult_GetIOUtilHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilData() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilData() failed", err))
	result := service.GetResult()
	ioUtilHigh := result.GetIOUtilHigh()
	asst.Equal(defaultResultIOUtilHigh, ioUtilHigh, "test GetIOUtilHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetIOUtilData() failed", err))
}

func TestResult_GetDiskCapacityUsageScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageScore() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageScore() failed", err))
	result := service.GetResult()
	diskCapacityUsageScore := result.GetDiskCapacityUsageScore()
	asst.Equal(defaultResultDiskCapacityUsageScore, diskCapacityUsageScore, "test GetDiskCapacityUsageScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageScore() failed", err))
}

func TestResult_GetDiskCapacityUsageData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageScore() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageScore() failed", err))
	result := service.GetResult()
	diskCapacityUsageData := result.GetDiskCapacityUsageData()
	asst.Equal(defaultResultDiskCapacityUsageData, diskCapacityUsageData, "test GetDiskCapacityUsageData() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageScore() failed", err))
}

func TestResult_GetDiskCapacityUsageHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageHigh() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageHigh() failed", err))
	result := service.GetResult()
	diskCapacityUsageHigh := result.GetDiskCapacityUsageHigh()
	asst.Equal(defaultResultDiskCapacityUsageHigh, diskCapacityUsageHigh, "test GetDiskCapacityUsageHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDiskCapacityUsageHigh() failed", err))
}

func TestResult_GetConnectionUsageScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageScore() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageScore() failed", err))
	result := service.GetResult()
	connectionUsageScore := result.GetConnectionUsageScore()
	asst.Equal(defaultResultConnectionUsageScore, connectionUsageScore, "test GetConnectionUsageScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageScore() failed", err))
}

func TestResult_GetConnectionUsageData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageData() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageData() failed", err))
	result := service.GetResult()
	connectionUsageData := result.GetConnectionUsageData()
	asst.Equal(defaultResultConnectionUsageData, connectionUsageData, "test GetConnectionUsageData() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageData() failed", err))
}

func TestResult_GetConnectionUsageHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageHigh() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageHigh() failed", err))
	result := service.GetResult()
	connectionUsageHigh := result.GetConnectionUsageHigh()
	asst.Equal(defaultResultConnectionUsageHigh, connectionUsageHigh, "test GetConnectionUsageHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetConnectionUsageHigh() failed", err))
}

func TestResult_GetAverageActiveSessionPercentsScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionPercentsScore() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionPercentsScore() failed", err))
	result := service.GetResult()
	averageActiveSessionPercentsScore := result.GetAverageActiveSessionPercentsScore()
	asst.Equal(defaultResultAverageActiveSessionPercentsScore, averageActiveSessionPercentsScore, "test GetAverageActiveSessionPercentsScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionPercentsScore() failed", err))
}

func TestResult_GetAverageActiveSessionPercentsData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionPercentsData() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionPercentsData() failed", err))
	result := service.GetResult()
	averageActiveSessionPercentsData := result.GetAverageActiveSessionPercentsData()
	asst.Equal(defaultResultAverageActiveSessionPercentsData, averageActiveSessionPercentsData, "test GetAverageActiveSessionPercentsData() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionPercentsData() failed", err))
}

func TestResult_GetAverageActiveSessionPercentsHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionPercentsHigh() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionPercentsHigh() failed", err))
	result := service.GetResult()
	averageActiveSessionPercentsHigh := result.GetAverageActiveSessionPercentsHigh()
	asst.Equal(defaultResultAverageActiveSessionPercentsHigh, averageActiveSessionPercentsHigh, "test GetAverageActiveSessionPercentsHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAverageActiveSessionPercentsHigh() failed", err))
}

func TestResult_GetCacheMissRatioScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioScore() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioScore() failed", err))
	result := service.GetResult()
	cacheMissRatioScore := result.GetCacheMissRatioScore()
	asst.Equal(defaultResultCacheMissRatioScore, cacheMissRatioScore, "test GetCacheMissRatioScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioScore() failed", err))
}

func TestResult_GetCacheMissRatioData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioData() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioData() failed", err))
	result := service.GetResult()
	cacheMissRatioData := result.GetCacheMissRatioData()
	asst.Equal(defaultResultCacheMissRatioData, cacheMissRatioData, "test GetCacheMissRatioData() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioData() failed", err))
}

func TestResult_GetCacheMissRatioHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioHigh() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioHigh() failed", err))
	result := service.GetResult()
	cacheMissRatioHigh := result.GetCacheMissRatioHigh()
	asst.Equal(defaultResultCacheMissRatioHigh, cacheMissRatioHigh, "test GetCacheMissRatioHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCacheMissRatioHigh() failed", err))
}

func TestResult_GetTableRowsScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetTableRowsScore() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableRowsScore() failed", err))
	result := service.GetResult()
	tableRowsScore := result.GetTableRowsScore()
	asst.Equal(defaultResultTableRowsScore, tableRowsScore, "test GetTableRowsScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableRowsScore() failed", err))
}

func TestResult_GetTableRowsData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetTableRowsData() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableRowsData() failed", err))
	result := service.GetResult()
	tableRowsData := result.GetTableRowsData()
	asst.Equal(defaultResultTableRowsData, tableRowsData, "test GetTableRowsData() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableRowsData() failed", err))
}

func TestResult_GetTableRowsHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetTableRowsHigh() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableRowsHigh() failed", err))
	result := service.GetResult()
	tableRowsHigh := result.GetTableRowsHigh()
	asst.Equal(defaultResultTableRowsHigh, tableRowsHigh, "test GetTableRowsHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableRowsHigh() failed", err))
}

func TestResult_GetTableSizeScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeScore() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeScore() failed", err))
	result := service.GetResult()
	tableSizeScore := result.GetTableSizeScore()
	asst.Equal(defaultResultTableSizeScore, tableSizeScore, "test GetTableSizeScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeScore() failed", err))
}

func TestResult_GetTableSizeData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeData() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeData() failed", err))
	result := service.GetResult()
	tableSizeData := result.GetTableSizeData()
	asst.Equal(defaultResultTableSizeData, tableSizeData, "test GetTableSizeData() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeData() failed", err))
}

func TestResult_GetTableSizeHigh(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeHigh() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeHigh() failed", err))
	result := service.GetResult()
	tableSizeHigh := result.GetTableSizeHigh()
	asst.Equal(defaultResultTableSizeHigh, tableSizeHigh, "test GetTableSizeHigh() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetTableSizeHigh() failed", err))
}

func TestResult_GetSlowQueryScore(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryScore() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryScore() failed", err))
	result := service.GetResult()
	slowQueryScore := result.GetSlowQueryScore()
	asst.Equal(defaultResultSlowQueryScore, slowQueryScore, "test GetSlowQueryScore() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryScore() failed", err))
}

func TestResult_GetSlowQueryData(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryData() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryData() failed", err))
	result := service.GetResult()
	slowQueryData := result.GetSlowQueryData()
	asst.Equal(defaultResultSlowQueryData, slowQueryData, "test GetSlowQueryData() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryData() failed", err))
}

func TestResult_GetSlowQueryAdvice(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryAdvice() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryAdvice() failed", err))
	result := service.GetResult()
	slowQueryAdvice := result.GetSlowQueryAdvice()
	asst.Equal(defaultResultSlowQueryAdvice, slowQueryAdvice, "test GetSlowQueryAdvice() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetSlowQueryAdvice() failed", err))
}

func TestResult_GetAccuracyReview(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetAccuracyReview() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAccuracyReview() failed", err))
	result := service.GetResult()
	accuracyReview := result.GetAccuracyReview()
	asst.Equal(defaultResultAccuracyReview, accuracyReview, "test GetAccuracyReview() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetAccuracyReview() failed", err))
}

func TestResult_GetDelFlag(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetDelFlag() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDelFlag() failed", err))
	result := service.GetResult()
	delFlag := result.GetDelFlag()
	asst.Equal(defaultResultDelFlag, delFlag, "test GetDelFlag() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetDelFlag() failed", err))
}

func TestResult_GetCreateTime(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetCreateTime() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCreateTime() failed", err))
	result := service.GetResult()
	createTime := result.GetCreateTime()
	asst.IsType(time.Now(), createTime, "test GetCreateTime() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetCreateTime() failed", err))
}

func TestResult_GetLastUpdateTime(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test GetLastUpdateTime() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetLastUpdateTime() failed", err))
	result := service.GetResult()
	lastUpdateTime := result.GetLastUpdateTime()
	asst.IsType(time.Now(), lastUpdateTime, "test GetLastUpdateTime() failed")
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test GetLastUpdateTime() failed", err))
}

func TestResult_Set(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
	result := service.GetResult()

	fields := make(map[string]interface{})
	fields["ID"] = defaultResultID
	fields["operationID"] = defaultResultOperationID

	err = result.Set(fields)
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))

	// field XX does not exist
	fields["XX"] = 100
	err = result.Set(fields)
	asst.NotNil(err, common.CombineMessageWithError("test Set() failed", err))

	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test Set() failed", err))
}

func TestResult_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	result := service.GetResult()
	_, err = result.MarshalJSON()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSON() failed", err))
}

func TestResult_MarshalJSONWithFields(t *testing.T) {
	asst := assert.New(t)

	service, err := rCreateService()
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	err = service.GetResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	result := service.GetResult()
	_, err = result.MarshalJSONWithFields("ID", "operationID", "WeightedAverageScore", "DBConfigScore", "DBConfigData", "DBConfigAdvice", "CPUUsageScore", "CPUUsageData", "CPUUsageHigh", "IOUtilScore", "IOUtilData", "IOUtilHigh", "DiskCapacityUsageScore", "DiskCapacityUsageData", "DiskCapacityUsageHigh", "ConnectionUsageScore", "ConnectionUsageData", "ConnectionUsageHigh", "AverageActiveSessionPercentsScore", "AverageActiveSessionPercentsData", "AverageActiveSessionPercentsHigh", "CacheMissRatioScore", "CacheMissRatioData", "CacheMissRatioHigh", "TableSizeScore", "TableSizeData", "TableSizeHigh", "SlowQueryScore", "SlowQueryData", "SlowQueryAdvice")
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
	// delete
	err = rDeleteHCResultByOperationID(defaultResultOperationID)
	asst.Nil(err, common.CombineMessageWithError("test MarshalJSONWithFields() failed", err))
}

// go test ./result_test.go ./result.go ./query.go ./service.go ./default_engine.go
