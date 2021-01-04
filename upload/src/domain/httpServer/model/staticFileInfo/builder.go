package staticFileInfo

// builder
type Builder func(model *StaticFileInfo)
type Builders []Builder

func (builders Builders) apply(model *StaticFileInfo) {
	for _, f := range builders {
		f(model)
	}
}

func WithRoute(path string) Builder {
	return func(model *StaticFileInfo) {
		model.Route.Path = path
	}
}

func WithDir(path string) Builder {
	return func(model *StaticFileInfo) {
		model.Dir.Path = path
	}
}
