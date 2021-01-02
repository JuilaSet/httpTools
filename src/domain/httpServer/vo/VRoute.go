package vo

// value object
type VRoute struct {
	Path string `json:"path"`
}

func NewVRoute(path string) *VRoute {
	return &VRoute{Path: path}
}
