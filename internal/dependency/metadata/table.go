package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type TableStatistic interface {
	// GetTableSchema returns the table schema
	GetTableSchema() string
	// GetTableName returns the table name
	GetTableName() string
	// GetTableRows returns the rows of the table
	GetTableRows() int
	// GetSize returns the size of the table
	GetSize() int
	// GetSizeMB returns the size(MB) of the table
	GetSizeMB() int
	// GetAvgRowLength returns the average row length of the table
	GetAvgRowLength() int
	// GetAutoIncrement returns the state of auto increment
	GetAutoIncrement() int
	// GetEngine returns the engine type of the table
	GetEngine() string
	// GetCharSet returns the charset of the table
	GetCharSet() string
	// GetCollation returns the collation of the table
	GetCollation() string
	// GetCreateTime returns the create time of the table
	GetCreateTime() time.Time
	// MarshalJSON marshals Table to json string
	MarshalJSON() ([]byte, error)
}

type IndexStatistic interface {
	// GetTableSchema returns the table schema
	GetTableSchema() string
	// GetTableName returns the table name
	GetTableName() string
	// GetIndexName returns the index name
	GetIndexName() string
	// GetSequence returns the sequence
	GetSequence() int
	// GetColumnName returns the column name
	GetColumnName() string
	// GetCardinality returns the cardinality
	GetCardinality() int
	// IsUnique returns unique state of index
	IsUnique() bool
	// IsNullable returns the index is nullable or not
	IsNullable() bool
	// MarshalJSON marshals Index to json string
	MarshalJSON() ([]byte, error)
}

type Table interface {
	// GetTableSchema returns the table schema
	GetTableSchema() string
	// GetTableName returns the table name
	GetTableName() string
	// GetTableStatistics returns the table statistics
	GetTableStatistics() ([]TableStatistic, error)
	// GetIndexStatistics returns the index statistics
	GetIndexStatistics() ([]IndexStatistic, error)
	// GetCreateStatement returns the create statement
	GetCreateStatement() string
	// MarshalJSON marshals Table to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified field of the Table to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type TableRepo interface {
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// GetTableStatistics gets table statistics from the middleware
	GetTableStatistics(tableSchema, tableName string) ([]TableStatistic, error)
	// GetIndexStatistics gets index statistics from the middleware
	GetIndexStatistics(tableSchema, tableName string) ([]IndexStatistic, error)
	// GetCreateStatement returns the create statement of the table
	GetCreateStatement(tableSchema, tableName string) (string, error)
	// AnalyzeTableByDBIDAndTableName analyzes the table by DBID and TableName
	AnalyzeTableByDBIDAndTableName(dbID int, tableName, userName string) error
	// AnalyzeTableByHostInfoAndDBNameAndTableName analyzes the table by host info、DB name and table name
	AnalyzeTableByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, userName string) error
}

type TableService interface {
	// GetTables returns the tables list
	GetTables() []Table
	// GetTableByDBIDAndTableName returns the table info by DBID and table name
	GetTableByDBIDAndTableName(dbID int, tableName string) error
	// GetTableByHostInfoAndDBNameAndTableName returns the table info by host info、 DB name and table name
	GetTableByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName string) error
	// AnalyzeTableByDBIDAndTableName analyzes the table by DBID and TableName
	AnalyzeTableByDBIDAndTableName(dbID int, tableName, accountName string) error
	// AnalyzeTableByHostInfoAndDBNameAndTableName analyzes the table by host info、DB name and table name
	AnalyzeTableByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, accountName string) error
	// Marshal marshals TableService.Tables to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the TableService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
