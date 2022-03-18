package metadata

import (
	"time"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
)

var _ metadata.TableStatistic = (*TableStatistic)(nil)

// TableStatistic include statistics of table
type TableStatistic struct {
	TableSchema   string    `middleware:"table_schema" json:"table_schema"`
	TableName     string    `middleware:"table_name" json:"table_name"`
	Rows          int       `middleware:"table_rows" json:"table_rows"`
	Size          int       `middleware:"size" json:"size"`
	SizeMB        float64   `middleware:"size_mb" json:"size_mb"`
	AvgRowLength  int       `middleware:"avg_row_length" json:"avg_row_length"`
	AutoIncrement int       `middleware:"auto_increment" json:"auto_increment"`
	Engine        string    `middleware:"engine" json:"engine"`
	CharSet       string    `middleware:"char_set" json:"char_set"`
	Collation     string    `middleware:"collation" json:"collation"`
	CreateTime    time.Time `middleware:"create_time" json:"create_time"`
}

// NewEmptyTableStatistic returns an empty *TableStatistic
func NewEmptyTableStatistic() *TableStatistic {
	return &TableStatistic{}
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
func (ts TableStatistic) GetSizeMB() float64 {
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

// IndexStatistic include statistics of index
type IndexStatistic struct {
	TableSchema string `middleware:"table_schema" json:"table_schema"`
	TableName   string `middleware:"table_name" json:"table_name"`
	IndexName   string `middleware:"index_name" json:"index_name"`
	Sequence    int    `middleware:"sequence" json:"sequence"`
	ColumnName  string `middleware:"column_name" json:"column_name"`
	Cardinality int    `middleware:"cardinality" json:"cardinality"`
	Unique      bool   `middleware:"non_unique" json:"non_unique"`
	Nullable    bool   `middleware:"nullable" json:"nullable"`
}

// NewEmptyIndexStatistic returns an empty *IndexStatistic
func NewEmptyIndexStatistic() *IndexStatistic {
	return &IndexStatistic{}
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

// TableInfo include information of logic table
type TableInfo struct {
	TableRepo   metadata.TableRepo
	TableSchema string `middleware:"table_schema" json:"table_schema"`
	TableName   string `middleware:"table_name" json:"table_name"`
	// TableStatistics []metadata.TableStatistic `middleware:"table_statistics" json:"table_statistics"`
	// IndexStatistics []metadata.IndexStatistic `middleware:"index_statistics" json:"index_statistics"`
	CreateStatement string `middleware:"create_statement" json:"create_statement"`
}

// NewTableInfo returns a new TableInfo
func NewTableInfo(
	repo metadata.TableRepo,
	tableSchema string,
	tableName string,
	// tableStatistics []metadata.TableStatistic,
	// indexStatistics []metadata.IndexStatistic,
	createStatement string,
) *TableInfo {
	return &TableInfo{
		repo,
		tableSchema,
		tableName,
		// tableStatistics,
		// indexStatistics,
		createStatement,
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

// GetTableStatistics returns the table statistics
func (ti *TableInfo) GetTableStatistics() ([]metadata.TableStatistic, error) {
	return ti.TableRepo.GetTableStatistics(ti.TableSchema, ti.TableName)
}

// GetIndexStatistics returns the index statistics
func (ti *TableInfo) GetIndexStatistics() ([]metadata.IndexStatistic, error) {
	return ti.TableRepo.GetIndexStatistics(ti.TableSchema, ti.TableName)
}

// GetCreateStatement returns the create statement
func (ti *TableInfo) GetCreateStatement() string {
	return ti.CreateStatement
}

// MarshalJSON marshals Table to json string
func (ti *TableInfo) MarshalJSON() ([]byte, error) {
	return common.MarshalStructWithTag(ti, constant.DefaultMarshalTag)
}

// MarshalJSONWithFields marshals only specified field of the Table to json string
func (ti *TableInfo) MarshalJSONWithFields(fields ...string) ([]byte, error) {
	return common.MarshalStructWithFields(ti, fields...)
}
