package metadata

// GetStatisticsByDBIDAndTableName is the struct bind to JSON param
type GetStatisticsByDBIDAndTableName struct {
	DBID      int    `json:"db_id" binding:"required"`
	TableName string `json:"table_name" binding:"required"`
}

// GetDBID is getter of db id
func (g *GetStatisticsByDBIDAndTableName) GetDBID() int {
	return g.DBID
}

// GetTableName is getter of table name
func (g *GetStatisticsByDBIDAndTableName) GetTableName() string {
	return g.TableName
}

// AnalyzeTableByDBIDAndTableName is the struct bind to JSON param
type AnalyzeTableByDBIDAndTableName struct {
	DBID      int    `json:"db_id" binding:"required"`
	TableName string `json:"table_name" binding:"required"`
	LoginName string `json:"account_name" binding:"required"`
}

// GetDBID is getter of db id
func (a *AnalyzeTableByDBIDAndTableName) GetDBID() int {
	return a.DBID
}

// GetTableName is getter of table name
func (a *AnalyzeTableByDBIDAndTableName) GetTableName() string {
	return a.TableName
}

// GetLoginName is getter of login name
func (a *AnalyzeTableByDBIDAndTableName) GetLoginName() string {
	return a.LoginName
}

// GetStatisticsByHostInfoAndDBNameAndTableName is the struct bind to JSON param
type GetStatisticsByHostInfoAndDBNameAndTableName struct {
	HostIP    string `json:"host_ip" binding:"required"`
	PortNum   int    `json:"port_num" binding:"required"`
	DBName    string `json:"db_name" binding:"required"`
	TableName string `json:"table_name" binding:"required"`
}

// GetHostIP is getter of host ip
func (g *GetStatisticsByHostInfoAndDBNameAndTableName) GetHostIP() string {
	return g.HostIP
}

// GetPortNum is getter of port num
func (g *GetStatisticsByHostInfoAndDBNameAndTableName) GetPortNum() int {
	return g.PortNum
}

// GetDBName is getter of db name
func (g *GetStatisticsByHostInfoAndDBNameAndTableName) GetDBName() string {
	return g.DBName
}

// GetTableName is getter of table name
func (g *GetStatisticsByHostInfoAndDBNameAndTableName) GetTableName() string {
	return g.TableName
}

// AnalyzeTableByHostInfoAndDBNameAndTableName is the struct bind to JSON param
type AnalyzeTableByHostInfoAndDBNameAndTableName struct {
	HostIP    string `json:"host_ip" binding:"required"`
	PortNum   int    `json:"port_num" binding:"required"`
	DBName    string `json:"db_name" binding:"required"`
	TableName string `json:"table_name" binding:"required"`
	LoginName string `json:"account_name" binding:"required"`
}

// GetHostIP is getter of host ip
func (a *AnalyzeTableByHostInfoAndDBNameAndTableName) GetHostIP() string {
	return a.HostIP
}

// GetPortNum is getter of port num
func (a *AnalyzeTableByHostInfoAndDBNameAndTableName) GetPortNum() int {
	return a.PortNum
}

// GetDBName is getter of db name
func (a *AnalyzeTableByHostInfoAndDBNameAndTableName) GetDBName() string {
	return a.DBName
}

// GetTableName is getter of table name
func (a *AnalyzeTableByHostInfoAndDBNameAndTableName) GetTableName() string {
	return a.TableName
}

// GetLoginName is getter of login name
func (a *AnalyzeTableByHostInfoAndDBNameAndTableName) GetLoginName() string {
	return a.LoginName
}
