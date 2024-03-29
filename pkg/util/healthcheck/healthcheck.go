package healthcheck

type Check struct {
	ServerID  int    `json:"server_id" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
	Step      string `json:"step" binding:"required"`
	LoginName string `json:"login_name" binding:"required"`
}

func (c *Check) GetServerID() int {
	return c.ServerID
}

func (c *Check) GetStartTime() string {
	return c.StartTime
}

func (c *Check) GetEndTime() string {
	return c.EndTime
}

func (c *Check) GetStep() string {
	return c.Step
}

func (c *Check) GetLoginName() string {
	return c.LoginName
}

type CheckByHostInfo struct {
	HostIP    string `json:"host_ip" binding:"required"`
	PortNum   int    `json:"port_num" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
	Step      string `json:"step" binding:"required"`
	LoginName string `json:"login_name" binding:"required"`
}

func (cbhi *CheckByHostInfo) GetHostIP() string {
	return cbhi.HostIP
}

func (cbhi *CheckByHostInfo) GetPortNum() int {
	return cbhi.PortNum
}

func (cbhi *CheckByHostInfo) GetStartTime() string {
	return cbhi.StartTime
}

func (cbhi *CheckByHostInfo) GetEndTime() string {
	return cbhi.EndTime
}

func (cbhi *CheckByHostInfo) GetStep() string {
	return cbhi.Step
}

func (cbhi *CheckByHostInfo) GetLoginName() string {
	return cbhi.LoginName
}

type ReviewAccuracy struct {
	OperationID int `json:"operation_id" binding:"required"`
	Review      int `json:"review" binding:"required"`
}

func (ra *ReviewAccuracy) GetOperationID() int {
	return ra.OperationID
}

func (ra *ReviewAccuracy) GetReview() int {
	return ra.Review
}
