package vo

import (
	"errors"
	"regexp"
	"strings"
)

type VFileNameInfo struct {
	Paths []string
	Path  string
	Dir   string
	Name  string
}

// "/www/aaa/bbb" -> Dir: "/www/aaa" , Name: "bbb"
func NewFileNameInfo(path string) *VFileNameInfo {
	e, _ := regexp.Compile(`^(/[^/]+)+$`)
	if !e.MatchString(path) {
		panic(errors.New("filename incorrect"))
	}

	paths := strings.Split(path, "/")[1:]
	dir := strings.Join(paths[0:len(paths)-1], "/")
	filename := paths[len(paths)-1]
	return &VFileNameInfo{
		Paths: paths,
		Dir:   dir,
		Name:  filename,
		Path:  path,
	}
}

func (info *VFileNameInfo) FileName() string {
	return info.Name
}

func (info *VFileNameInfo) FileDir(base string) string {
	return base + "/" + info.Dir
}

func (info *VFileNameInfo) FilePath(base string) string {
	return info.FileDir(base) + "/" + info.FileName()
}
