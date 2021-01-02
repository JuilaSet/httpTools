package vo

import (
	"httpTools/src/infrastructure/fileSys"
	"log"
)

type WatchedFiles []*VWatchedFile
type VWatchedFile struct {
	Filename string
}

func NewWatchedFile(fileName string) *VWatchedFile {
	if !fileSys.IsExist(fileName) {
		log.Println("file not exist")
	}
	return &VWatchedFile{
		Filename: fileName,
	}
}