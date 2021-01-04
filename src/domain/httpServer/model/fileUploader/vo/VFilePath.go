package vo

type VExcludes struct {
	FilePathList []string `json:"file_path"`
}

func NewVExclude(filePaths []string) *VExcludes {
	return &VExcludes{FilePathList: filePaths}
}

