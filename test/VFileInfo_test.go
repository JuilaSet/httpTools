package test

import (
	"httpTools/src/domain/httpServer/model/fileUploader"
	"testing"
)

func TestNewVFileInfo(t *testing.T) {
	fileUploader.NewVUploadDir("./upload")
}