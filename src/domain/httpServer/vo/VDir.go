package vo

import (
	"errors"
	"httpTools/src/infrastructure/fileUtil"
	"strings"
)

type VDir struct {
	Path string `json:"path"`
}

func NewVDir(dirPath string) *VDir {
	if strings.Contains(dirPath, "..") {
		panic(errors.New("cannot use relative path for safety"))
	}
	if !fileUtil.IsExist(dirPath) {
		panic(errors.New("file not exist"))
	}
	return &VDir{Path: dirPath}
}
