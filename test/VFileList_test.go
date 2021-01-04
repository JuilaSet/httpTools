package test

import (
	"httpTools/src/domain/httpServer/model/fileUploader"
	"testing"
)

func TestNewVFileList(t *testing.T) {
	for _, v := range fileUploader.NewVFileList(fileUploader.NewVUploadDir("./upload")).List {
		t.Log("Test: ", v)
	}
}

