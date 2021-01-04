package fileUploader

import (
	"errors"
	"httpTools/src/infrastructure/fileUtil"
	"io/ioutil"
)

type VUploadDir struct {
	Path string `json:"path"`
}

func NewVUploadDir(filepath string) *VUploadDir {
	if !fileUtil.IsExist(filepath) {
		panic(errors.New("file not exist"))
	}
	return &VUploadDir{Path: filepath}
}

// dir目录下的所有文件
func (d *VUploadDir) GetFileList() ([]string, error) {
	var list []string
	rd, err := ioutil.ReadDir(d.Path)
	for _, fi := range rd {
		if fi.IsDir() {
			var innerList []string
			innerList, err = fileUtil.GetAllFile("/" + fi.Name(), d.Path)
			list = append(list, innerList...)
		} else {
			list = append(list, "/" + fi.Name())
		}
	}
	return list, err
}

