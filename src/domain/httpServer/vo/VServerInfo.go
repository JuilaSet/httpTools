package vo

type VServerInfo struct {
	Port string `json:"port"`
}

func NewVServiceInfo(port string) *VServerInfo {
	return &VServerInfo{Port: port}
}
