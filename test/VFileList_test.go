package test

import (
	"httpTools/src/domain/httpServer/model/fileUploader/vo"
	"testing"
)

func TestNewVFileList(t *testing.T) {
	for _, v := range vo.NewVFileList(vo.NewVUploadDir("./upload")).List {
		t.Log("Test: ", v)
	}
}

