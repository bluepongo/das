package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type TableStatistic interface {
	GetTableSchema() string
	GetTableName() string
	GetRows() int
	GetSize() int
	GetSizeMB() int
	GetAvgRowLength() int
	GetAutoIncrement() int
	GetEngine() string
	GetCharSet() string
	GetCollation() string
	GetCreateTime() time.Time
	MarshalJSON() ([]byte, error)
}

type IndexStatistic interface {
	GetTableSchema() string
	GetTableName() string
	GetIndexName() string
	GetSequence() int
	GetColumnName() string
	GetCardinality() int
	IsUnique() bool
	IsNullable() bool
	MarshalJSON() ([]byte, error)
}

type Table interface {
	GetTableSchema() string
	GetTableName() string
	GetTableStatistics() []TableStatistic
	GetIndexStatistics() []IndexStatistic
	GetCreateStatement() string
	MarshalJSON() ([]byte, error)
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type TableRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// Transaction returns a middleware.Transaction that could execute multiple commands as a transaction
	Transaction() (middleware.Transaction, error)
	GetTableStatistics(tableSchema, tableName string) ([]TableStatistic, error)
	GetIndexStatistics(tableSchema, tableName string) ([]IndexStatistic, error)
	GetCreateStatement(tableSchema, tableName string) (string, error)
	AnalyzeTableByDBIDAndTableName(dbID int, tableName, userName string) error
	AnalyzeTableByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, userName string) error
}

type TableService interface {
	GetTables() []Table
	GetTableByDBIDAndTableName(dbID int, tableName string) error
	GetTableByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName string) error
	AnalyzeTableByDBIDAndTableName(dbID int, tableName, userName string) error
	AnalyzeTableByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, userName string) error
	Marshal() ([]byte, error)
	MarshalWithFields(fields ...string) ([]byte, error)
}
