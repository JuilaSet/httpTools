package proxyInfo

// builder
type Builder func(model *ProxyInfo)
type Builders []Builder

func (builders Builders) apply(model *ProxyInfo) {
	for _, f := range builders {
		f(model)
	}
}

// build methods
func WithHttpMethod(method string) Builder {
	return func(model *ProxyInfo) {
		model.HttpMethod.Method = method
	}
}

func WithTargetURL(targetUrl string) Builder {
	return func(model *ProxyInfo) {
		model.TargetURL.URL = targetUrl
	}
}

func WithRoute(path string) Builder {
	return func(model *ProxyInfo) {
		model.Route.Path = path
	}
}
