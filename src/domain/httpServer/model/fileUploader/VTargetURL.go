package fileUploader

// value object
type VTarget struct {
	URL string `json:"url"`
}

func NewVTarget(URL string) *VTarget {
	return &VTarget{URL: URL}
}
