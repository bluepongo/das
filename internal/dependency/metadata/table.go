package metadata

import (
	"time"

	"github.com/romberli/go-util/middleware"
)

type TableStatistic interface {
	// GetDBName returns the table schema
	GetDBName() string
	// GetTableName returns the table name
	GetTableName() string
	// GetTableRows returns the rows of the table
	GetTableRows() int
	// GetSize returns the size of the table
	GetSize() int
	// GetSizeMB returns the size(MB) of the table
	GetSizeMB() float64
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
	// GetDBName returns the table schema
	GetDBName() string
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
	// GetDBName returns the table schema
	GetDBName() string
	// GetTableName returns the table name
	GetTableName() string
	// GetTableStatistics returns the table statistics
	GetTableStatistics() ([]TableStatistic, error)
	// GetIndexStatistics returns the index statistics
	GetIndexStatistics() ([]IndexStatistic, error)
	// GetCreateStatement returns the create statement
	GetCreateStatement() (string, error)
	// MarshalJSON marshals Table to json string
	MarshalJSON() ([]byte, error)
	// MarshalJSONWithFields marshals only specified field of the Table to json string
	MarshalJSONWithFields(fields ...string) ([]byte, error)
}

type TableRepo interface {
	// Close closes the mysql connection
	Close() error
	// InitMySQLConn initialize the mysql connection of the repository
	InitMySQLConn(hostIP string, portNum int, dbName string) error
	// Execute executes given command and placeholders on the middleware
	Execute(command string, args ...interface{}) (middleware.Result, error)
	// GetByDBName gets the table info by db name from middleware
	GetByDBName(dbName string) ([]Table, error)
	// GetTableStatistics gets table statistics from the middleware
	GetTableStatistics(dbName, tableName string) ([]TableStatistic, error)
	// GetIndexStatistics gets index statistics from the middleware
	GetIndexStatistics(dbName, tableName string) ([]IndexStatistic, error)
	// GetCreateStatement gets the create statement of the table
	GetCreateStatement(dbName, tableName string) (string, error)
	// GetStatisticsByDBNameAndTableName gets the full table info by db name and table name from middleware
	GetStatisticsByDBNameAndTableName(dbName, tableName string) ([]TableStatistic, []IndexStatistic, string, error)
	// AnalyzeTableByDBNameAndTableName analyzes the table by db name and table name
	AnalyzeTableByDBNameAndTableName(dbName, tableName string) error
}

type TableService interface {
	// GetTables returns the tables list
	GetTables() []Table
	// GetByHostInfoAndDBName returns tables info by db name
	GetByDBID(dbID int, loginName string) error
	// GetStatisticsByDBIDAndTableName gets the table statistics by db id and table name
	GetStatisticsByDBIDAndTableName(dbID int, tableName, loginName string) error
	// GetStatisticsByHostInfoAndDBNameAndTableName gets the table statistics by host info and db name and table name
	GetStatisticsByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, loginName string) error
	// AnalyzeTableByDBIDAndTableName analyzes the table by db id and table name
	AnalyzeTableByDBIDAndTableName(dbID int, tableName, loginName string) error
	// AnalyzeTableByHostInfoAndDBNameAndTableName analyzes the table by host info and db name and table name
	AnalyzeTableByHostInfoAndDBNameAndTableName(hostIP string, portNum int, dbName, tableName, loginName string) error
	// Marshal marshals TableService.Tables to json bytes
	Marshal() ([]byte, error)
	// MarshalWithFields marshals only specified fields of the TableService to json bytes
	MarshalWithFields(fields ...string) ([]byte, error)
}
