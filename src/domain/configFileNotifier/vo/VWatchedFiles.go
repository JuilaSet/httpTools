package vo

import (
	"httpTools/src/infrastructure/fileUtil"
	"log"
)

type WatchedFiles []*VWatchedFile
type VWatchedFile struct {
	Filename string
}

func NewWatchedFile(fileName string) *VWatchedFile {
	if !fileUtil.IsExist(fileName) {
		log.Println("file not exist")
	}
	return &VWatchedFile{
		Filename: fileName,
	}
}