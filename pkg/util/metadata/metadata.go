package metadata

type HostInfo struct {
	HostIP  string `json:"host_ip" bind:"required"`
	PortNum int    `json:"port_num" bind:"required"`
}

type NameAndClusterInfo struct {
	DBName      string `json:"db_name" bind:"required"`
	ClusterID   int    `json:"cluster_id" bind:"required"`
	ClusterType int    `json:"cluster_type" bind:"required"`
}

func (hi *HostInfo) GetHostIP() string {
	return hi.HostIP
}

func (hi *HostInfo) GetPortNum() int {
	return hi.PortNum
}

func (nc *NameAndClusterInfo) GetDBName() string {
	return nc.DBName
}

func (nc *NameAndClusterInfo) GetClusterID() int {
	return nc.ClusterID
}

func (nc *NameAndClusterInfo) GetClusterType() int {
	return nc.ClusterType
}
