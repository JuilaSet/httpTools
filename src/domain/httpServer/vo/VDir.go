package vo

import (
	"errors"
	"httpTools/src/infrastructure/fileUtil"
)

type VDir struct {
	Path string `json:"path"`
}

func NewVDir(dirPath string) *VDir {
	if !fileUtil.IsExist(dirPath) {
		panic(errors.New("file not exist"))
	}
	return &VDir{Path: dirPath}
}
