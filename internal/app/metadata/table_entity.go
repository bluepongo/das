package metadata

import (
	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"time"
)

var _ metadata.TableStatistic = (*TableStatistic)(nil)

type TableStatistic struct {
	TableSchema   string    `middleware:"table_schema" json:"table_schema"`
	TableName     string    `middleware:"table_name" json:"table_name"`
	Rows          int       `middleware:"rows" json:"table_rows"`
	Size          int       `middleware:"size" json:"size"`
	SizeMB        int       `middleware:"size_mb" json:"size_mb"`
	AvgRowLength  int       `middleware:"avg_row_length" json:"avg_row_length"`
	AutoIncrement int       `middleware:"auto_increment" json:"auto_increment"`
	Engine        string    `middleware:"engine" json:"engine"`
	CharSet       string    `middleware:"char_set" json:"char_set"`
	Collation     string    `middleware:"collation" json:"collation"`
	CreateTime    time.Time `middleware:"create_time" json:"create_time"`
}

// GetTableSchema returns the table schema
func (ts TableStatistic) GetTableSchema() string {
	return ts.TableSchema
}

// GetTableName returns the table name
func (ts TableStatistic) GetTableName() string {
	return ts.TableName
}

// GetTableRows returns the rows of the table
func (ts TableStatistic) GetTableRows() int {
	return ts.Rows
}

// GetSize returns the size of the table
func (ts TableStatistic) GetSize() int {
	return ts.Size
}

// GetSizeMB returns the size(MB) of the table
func (ts TableStatistic) GetSizeMB() int {
	return ts.SizeMB
}

// GetAvgRowLength returns the average row length of the table
func (ts TableStatistic) GetAvgRowLength() int {
	return ts.AvgRowLength
}

// GetAutoIncrement returns the state of auto increment
func (ts TableStatistic) GetAutoIncrement() int {
	return ts.AutoIncrement
}

// GetEngine returns the engine type of the table
func (ts TableStatistic) GetEngine() string {
	return ts.Engine
}

// GetCharSet returns the charset of the table
func (ts TableStatistic) GetCharSet() string {
	return ts.CharSet
}

// GetCollation returns the collation of the table
func (ts TableStatistic) GetCollation() string {
	return ts.Collation
}

// GetCreateTime returns the create time of the table
func (ts TableStatistic) GetCreateTime() time.Time {
	return ts.CreateTime
}

// MarshalJSON marshals Table to json string
func (ts TableStatistic) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(ts, constant.DefaultMarshalTag)
}

var _ metadata.IndexStatistic = (*IndexStatistic)(nil)

type IndexStatistic struct {
	TableSchema string `middleware:"table_schema" json:"table_schema"`
	TableName   string `middleware:"table_name" json:"table_name"`
	IndexName   string `middleware:"index_name" json:"index_name"`
	Sequence    int    `middleware:"sequence" json:"sequence"`
	ColumnName  string `middleware:"column_name" json:"column_name"`
	Cardinality int    `middleware:"cardinality" json:"cardinality"`
	Unique      bool   `middleware:"unique" json:"unique"`
	Nullable    bool   `middleware:"nullable" json:"nullable"`
}

// GetTableSchema returns the table schema
func (is IndexStatistic) GetTableSchema() string {
	return is.TableSchema
}

// GetTableName returns the table name
func (is IndexStatistic) GetTableName() string {
	return is.TableName
}

// GetIndexName returns the index name
func (is IndexStatistic) GetIndexName() string {
	return is.IndexName
}

// GetSequence returns the sequence
func (is IndexStatistic) GetSequence() int {
	return is.Sequence
}

// GetColumnName returns the column name
func (is IndexStatistic) GetColumnName() string {
	return is.ColumnName
}

// GetCardinality returns the cardinality
func (is IndexStatistic) GetCardinality() int {
	return is.Cardinality
}

// IsUnique returns unique state of index
func (is IndexStatistic) IsUnique() bool {
	return is.Unique
}

// IsNullable returns the index is nullable or not
func (is IndexStatistic) IsNullable() bool {
	return is.Nullable
}

// MarshalJSON marshals Index to json string
func (is IndexStatistic) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(is, constant.DefaultMarshalTag)
}

var _ metadata.Table = (*TableInfo)(nil)

type TableInfo struct {
	tableRepo   metadata.TableRepo
	TableSchema string `middleware:"table_schema" json:"table_schema"`
	TableName   string `middleware:"table_name" json:"table_name"`
}

// NewTableInfo returns a new TableInfo
func NewTableInfo(repo metadata.TableRepo, tableSchema string, tableName string) *TableInfo {
	return &TableInfo{
		repo,
		tableSchema,
		tableName,
	}
}

// NewTableInfoWithMapAndRandom returns a new *TableInfo with given map
func NewTableInfoWithMapAndRandom(fields map[string]interface{}) (*TableInfo, error) {
	ti := &TableInfo{}
	err := common.SetValuesWithMapAndRandom(ti, fields)
	if err != nil {
		return nil, err
	}

	return ti, nil
}

// GetTableSchema returns the table schema
func (ti *TableInfo) GetTableSchema() string {
	return ti.TableSchema
}

// GetTableName returns the table name
func (ti *TableInfo) GetTableName() string {
	return ti.TableName
}

// GetTableStatistics returns the table statistic list
func (ti *TableInfo) GetTableStatistics() ([]metadata.TableStatistic, error) {
	return ti.tableRepo.GetTableStatistics(ti.TableSchema, ti.TableName)
}

// GetIndexStatistics returns the index statistic list
func (ti *TableInfo) GetIndexStatistics() ([]metadata.IndexStatistic, error) {
	return ti.tableRepo.GetIndexStatistics(ti.TableSchema, ti.TableName)
}

// MarshalJSON marshals Table to json string
func (ti *TableInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(ti, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only specified field of the Table to json string
func (ti *TableInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(ti, fields...)
}
