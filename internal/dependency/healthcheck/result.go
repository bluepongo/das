package healthcheck

import (
	"time"
)

type Result interface {
	// Identity returns the identity
	Identity() int
	// GetOperationID returns the operation id
	GetOperationID() int
	// GetWeightedAverageScore returns the weighted average score
	GetWeightedAverageScore() int
	// GetDBConfigScore returns the database configuration score
	GetDBConfigScore() int
	// GetDBConfigData returns the database configuration data
	GetDBConfigData() string
	// GetDBConfigAdvice returns the database configuration advice
	GetDBConfigAdvice() string
	// GetAvgBackupFailedRatioScore returns the average backup failed ratio score
	GetAvgBackupFailedRatioScore() int
	// GetAvgBackupFailedRatioData returns the average backup failed ratio data
	GetAvgBackupFailedRatioData() string
	// GetAvgBackupFailedRatioHigh returns the high average backup failed ratio data
	GetAvgBackupFailedRatioHigh() string
	// GetStatisticFailedRatioScore returns the statistic failed ratio score
	GetStatisticFailedRatioScore() int
	// GetStatisticFailedRatioData returns the statistic failed ratio data
	GetStatisticFailedRatioData() string
	// GetStatisticFailedRatioHigh returns the high statistic failed ratio data
	GetStatisticFailedRatioHigh() string
	// GetCPUUsageScore returns the cpu usage score
	GetCPUUsageScore() int
	// GetCPUUsageData returns the cpu usage data
	GetCPUUsageData() string
	// GetCPUUsageHigh returns the high cpu usage data
	GetCPUUsageHigh() string
	// GetIOUtilScore returns the io util score
	GetIOUtilScore() int
	// GetIOUtilData returns the io util data
	GetIOUtilData() string
	// GetIOUtilHigh returns the high io util data
	GetIOUtilHigh() string
	// GetDiskCapacityUsageScore returns the disk capacity usage score
	GetDiskCapacityUsageScore() int
	// GetDiskCapacityUsageData returns the disk capacity usage data
	GetDiskCapacityUsageData() string
	// GetDiskCapacityUsageHigh returns the high disk capacity usage data
	GetDiskCapacityUsageHigh() string
	// GetConnectionUsageScore returns the connection usage score
	GetConnectionUsageScore() int
	// GetConnectionUsageData returns the connection usage data
	GetConnectionUsageData() string
	// GetConnectionUsageHigh returns the high connection usage data
	GetConnectionUsageHigh() string
	// GetAverageActiveSessionPercentsScore returns the average active session number score
	GetAverageActiveSessionPercentsScore() int
	// GetAverageActiveSessionPercentsData returns the average active session number data
	GetAverageActiveSessionPercentsData() string
	// GetAverageActiveSessionPercentsHigh returns the high average active session number data
	GetAverageActiveSessionPercentsHigh() string
	// GetCacheHitRatioScore returns the cache miss ratio score
	GetCacheMissRatioScore() int
	// GetCacheHitRatioData returns the cache miss ratio data
	GetCacheMissRatioData() string
	// GetCacheMissRatioHigh returns the high cache miss ratio data
	GetCacheMissRatioHigh() string
	// GetTableRowsScore returns the table rows score
	GetTableRowsScore() int
	// GetTableRowsData returns the table rows data
	GetTableRowsData() string
	// GetTableRowsHigh returns the high table rows data
	GetTableRowsHigh() string
	// GetTableSizeScore returns the table size score
	GetTableSizeScore() int
	// GetTableSizeData returns the table size data
	GetTableSizeData() string
	// GetTableSizeHigh returns the high table size data
	GetTableSizeHigh() string
	// GetSlowQueryScore returns the slow query score
	GetSlowQueryScore() int
	// GetSlowQueryData returns the slow query data
	GetSlowQueryData() string
	// GetSlowQueryAdvice returns the slow query advice
	GetSlowQueryAdvice() string
	// GetAccuracyReview returns the accuracy review
	GetAccuracyReview() int
	// GetDelFlag returns the delete flag
	GetDelFlag() int
	// GetCreateTime returns the create time
	GetCreateTime() time.Time
	// GetLastUpdateTime returns the last update time
	GetLastUpdateTime() time.Time
	// String returns the string value of the result
	String() string
	// Set sets health check with given fields, key is the field name and value is the relevant value of the key
	Set(fields map[string]interface{}) error
	// MarshalJSON marshals Result to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSON marshals only specified field of the Result to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}
