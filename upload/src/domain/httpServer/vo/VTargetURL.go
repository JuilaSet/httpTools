package vo

// value object
type VTargetURL struct {
	URL string `json:"url"`
}

func NewVTargetURL(URL string) *VTargetURL {
	return &VTargetURL{URL: URL}
}
