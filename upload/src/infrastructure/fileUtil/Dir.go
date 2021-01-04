package fileUtil

import (
	"fmt"
	"io/ioutil"
)

// relativePath: "/aaa", pathname: "/name" -> "/aaa/name"
func GetAllFile(pathname string, relativePath string) (list []string, err error) {
	rd, err := ioutil.ReadDir(relativePath + pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname + "/" + fi.Name())
			li, err := GetAllFile(pathname + "/" + fi.Name(), relativePath)
			if err != nil {
				return nil, err
			}
			list = append(list, li...)
		} else {
			fmt.Println(pathname + "/" + fi.Name())
			list = append(list, pathname + "/" + fi.Name())
		}
	}
	return
}

