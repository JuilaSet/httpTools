package vo

import (
	"errors"
	"httpTools/src/infrastructure/fileSys"
)

type VDir struct {
	Path string `json:"path"`
}

func NewVDir(dirPath string) *VDir {
	if !fileSys.IsExist(dirPath) {
		panic(errors.New("file not exist"))
	}
	return &VDir{Path: dirPath}
}
