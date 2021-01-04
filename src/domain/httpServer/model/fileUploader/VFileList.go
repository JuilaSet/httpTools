package fileUploader

type VFileList struct {
	List []string
	dir *VUploadDir
}

// 负责解析文件夹中的所有文件并返回一个文件列表
func NewVFileList(dir *VUploadDir) *VFileList {
	l, err := dir.GetFileList()
	if err != nil {
		panic(err)
	}
	return &VFileList{
		List: l,
		dir: dir,
	}
}

func (f *VFileList) GetFileList() []string {
	return f.List
}

// 返回文件夹的路径
func (f *VFileList) GetDirPath() string {
	return f.dir.Path
}

// 重新加载
func (f *VFileList) Refresh() {
	l, err := f.dir.GetFileList()
	if err != nil {
		panic(err)
	}
	f.List = l
}


