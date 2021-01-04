package fileUploader

import (
	"errors"
	"regexp"
)

// value object
type VRoute struct {
	Path string `json:"path"`
}

func NewVRoute(path string) *VRoute {
	// route: "/www"
	if reg , err := regexp.Compile(`(/[^/]+)+`); err != nil {
		panic(err)
	} else if !reg.MatchString(path) {
		panic(errors.New("invalid path: " + path))
	}
	return &VRoute{Path: path}
}
