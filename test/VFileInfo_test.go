package test

import (
	"httpTools/src/domain/httpServer/model/fileUploader/vo"
	"testing"
)

func TestNewVFileInfo(t *testing.T) {
	vo.NewVUploadDir("./upload")
}