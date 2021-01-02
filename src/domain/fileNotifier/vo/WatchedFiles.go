package vo

import (
	"httpProxyDDD/src/infrastructure/fileSys"
	"log"
)

type WatchedFiles []*WatchedFile
type WatchedFile struct {
	Filename string
}

func NewWatchedFile(fileName string) *WatchedFile {
	if !fileSys.IsExist(fileName) {
		log.Println("file not exist")
	}
	return &WatchedFile{
		Filename: fileName,
	}
}