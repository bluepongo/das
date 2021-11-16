package healthcheck

type Check struct {
	ServerID  int    `json:"server_id" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
	Step      string `json:"step" binding:"required"`
}

type CheckByHostInfo struct {
	HostIP    string `json:"host_ip" binding:"required"`
	PortNum   int    `json:"port_num" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
	Step      string `json:"step" binding:"required"`
}
