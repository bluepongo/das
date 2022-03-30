package sqladvisor

type SQLInfo struct {
	DBID    int    `json:"db_id" binding:"required"`
	SQLText string `json:"sql_text" binding:"required"`
}

func (si *SQLInfo) GetDBID() int {
	return si.DBID
}

func (si *SQLInfo) GetSQLText() string {
	return si.SQLText
}
