package metadata

type HostInfo struct {
	HostIP  string `json:"host_ip" bind:"required"`
	PortNum int    `json:"port_num" bind:"required"`
}

func (hi *HostInfo) GetHostIP() string {
	return hi.HostIP
}

func (hi *HostInfo) GetPortNum() int {
	return hi.PortNum
}
