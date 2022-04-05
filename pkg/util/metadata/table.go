package metadata

// TablesByDBID is the struct bind to JSON param
type TablesByDBID struct {
	DBID      int    `json:"db_id" binding:"required"`
	LoginName string `json:"login_name" binding:"required"`
}

func (t *TablesByDBID) GetDBID() int {
	return t.DBID
}

// GetLoginName is getter of login name
func (t *TablesByDBID) GetLoginName() string {
	return t.LoginName
}

// TableByDBIDAndTableName is the struct bind to JSON param
type TableByDBIDAndTableName struct {
	DBID      int    `json:"db_id" binding:"required"`
	TableName string `json:"table_name" binding:"required"`
	LoginName string `json:"login_name" binding:"required"`
}

// GetDBID is getter of db id
func (t *TableByDBIDAndTableName) GetDBID() int {
	return t.DBID
}

// GetTableName is getter of table name
func (t *TableByDBIDAndTableName) GetTableName() string {
	return t.TableName
}

// GetLoginName is getter of login name
func (t *TableByDBIDAndTableName) GetLoginName() string {
	return t.LoginName
}

// TableByHostInfoAndDBNameAndTableName is the struct bind to JSON param
type TableByHostInfoAndDBNameAndTableName struct {
	HostIP    string `json:"host_ip" binding:"required"`
	PortNum   int    `json:"port_num" binding:"required"`
	DBName    string `json:"db_name" binding:"required"`
	TableName string `json:"table_name" binding:"required"`
	LoginName string `json:"login_name" binding:"required"`
}

// GetHostIP is getter of host ip
func (t *TableByHostInfoAndDBNameAndTableName) GetHostIP() string {
	return t.HostIP
}

// GetPortNum is getter of port num
func (t *TableByHostInfoAndDBNameAndTableName) GetPortNum() int {
	return t.PortNum
}

// GetDBName is getter of db name
func (t *TableByHostInfoAndDBNameAndTableName) GetDBName() string {
	return t.DBName
}

// GetTableName is getter of table name
func (t *TableByHostInfoAndDBNameAndTableName) GetTableName() string {
	return t.TableName
}

// GetLoginName is getter of login name
func (t *TableByHostInfoAndDBNameAndTableName) GetLoginName() string {
	return t.LoginName
}
