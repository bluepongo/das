package metadata

import "github.com/romberli/das/internal/dependency/metadata"

var _ metadata.TableService = (*TableService)(nil)

type TableService struct {
	metadata.TableRepo
	Tables []metadata.Table `middleware:"tables" json:"tables"`
	TableStatistics []metadata.TableStatistic `middleware:"table_statistics" json:"table_statistics"`
	IndexStatistics []metadata.IndexStatistic `middleware:"index_statistics" json:"index_statistics"`
	CreateStatement string `middleware:"create_statement" json:"create_statement"`
}

// NewTableService returns a new TableService
func NewTableService(repo metadata.TableRepo) *TableService {
	return &TableService{repo,
		[]metadata.Table{},
		[]metadata.TableStatistic{},
		[]metadata.IndexStatistic{},
		"",
	}
}

// GetTables returns the tables list
func (ts *TableService) GetTables() []metadata.Table {
	 return ts.Tables
}

// GetTableByDBIDAndTableName returns the table info by DBID and table name
func (ts *TableService) GetTableByDBIDAndTableName(dbID int, tableName string) error {

}