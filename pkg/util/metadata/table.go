package metadata

// GetTablesByDBID is the struct bind to JSON param
type GetTablesByDBID struct {
	LoginName string `json:"login_name" binding:"required"`
}

// GetLoginName is getter of login name
func (g *GetTablesByDBID) GetLoginName() string {
	return g.LoginName
}

// GetTableByDBIDAndTableName is the struct bind to JSON param
type GetTableByDBIDAndTableName struct {
	DBID      int    `json:"db_id" binding:"required"`
	TableName string `json:"table_name" binding:"required"`
	LoginName string `json:"login_name" binding:"required"`
}

// GetDBID is getter of db id
func (g *GetTableByDBIDAndTableName) GetDBID() int {
	return g.DBID
}

// GetTableName is getter of table name
func (g *GetTableByDBIDAndTableName) GetTableName() string {
	return g.TableName
}

// GetLoginName is getter of login name
func (g *GetTableByDBIDAndTableName) GetLoginName() string {
	return g.LoginName
}

// GetTableByHostInfoAndDBNameAndTableName is the struct bind to JSON param
type GetTableByHostInfoAndDBNameAndTableName struct {
	HostIP    string `json:"host_ip" binding:"required"`
	PortNum   int    `json:"port_num" binding:"required"`
	DBName    string `json:"db_name" binding:"required"`
	TableName string `json:"table_name" binding:"required"`
	LoginName string `json:"login_name" binding:"required"`
}

// GetHostIP is getter of host ip
func (g *GetTableByHostInfoAndDBNameAndTableName) GetHostIP() string {
	return g.HostIP
}

// GetPortNum is getter of port num
func (g *GetTableByHostInfoAndDBNameAndTableName) GetPortNum() int {
	return g.PortNum
}

// GetDBName is getter of db name
func (g *GetTableByHostInfoAndDBNameAndTableName) GetDBName() string {
	return g.DBName
}

// GetTableName is getter of table name
func (g *GetTableByHostInfoAndDBNameAndTableName) GetTableName() string {
	return g.TableName
}

// GetLoginName is getter of login name
func (g *GetTableByHostInfoAndDBNameAndTableName) GetLoginName() string {
	return g.LoginName
}
