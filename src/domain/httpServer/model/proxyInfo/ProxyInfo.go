package proxyInfo

import (
	"github.com/gin-gonic/gin"
	"httpProxyDDD/src/domain/httpServer/vo"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

type Proxies []*ProxyInfo
type ProxyInfo struct {
	HttpMethod *vo.VHttpMethod `json:"http_method"`
	TargetURL  *vo.VTargetURL  `json:"target_url"`
	Route      *vo.VRoute      `json:"route"`
}

func NewProxy(builders ...Builder) *ProxyInfo {
	c := &ProxyInfo{
		HttpMethod: vo.NewVHttpMethod("GET"),
		TargetURL:  vo.NewVTargetURL("127.0.0.1"),
		Route:      vo.NewVRoute(""),
	}
	Builders(builders).apply(c)
	return c
}

// 构建代理工具
func (infos Proxies) Apply(engine *gin.Engine) {
	log.Println("Proxy infos", infos)
	for _, info := range infos {
		info.Apply(engine)
	}
}

func (info *ProxyInfo) Apply(engine *gin.Engine) {
	log.Println("Proxy", info.HttpMethod.Method, info.Route.Path+"/*next")
	engine.Handle(info.HttpMethod.Method, info.Route.Path+"/*next", reverseProxy(info.TargetURL.URL, info))
}

// private
func reverseProxy(target string, proxyInfo *ProxyInfo) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, err := url.Parse(target)
		if err != nil {
			panic(err)
		}

		// fmt.Printf("%s\n", )// [bc ca olang]

		director := func(req *http.Request) {
			// /kube/xxx -> /test/xxx
			// /kube/ -> /test/
			// /kube/api -> /test/api

			proxyRoute := proxyInfo.Route.Path
			reqPath := req.URL.Path

			//fmt.Println("ProxyRoute: ", proxyRoute)
			//fmt.Println("ReqPath: ", reqPath)
			reg := regexp.MustCompile(proxyRoute)
			cutPrefix := reg.ReplaceAllString(reqPath, "")
			if cutPrefix == "/" {
				cutPrefix = ""
			}

			//fmt.Println("CutPrefix: ", cutPrefix)
			tarUrl := u.Path + cutPrefix

			//fmt.Println("Target Url: ", tarUrl)
			req.URL.Scheme = "http"
			req.URL.Host = u.Host
			req.URL.Path = tarUrl
		}
		reverseProxy := &httputil.ReverseProxy{Director: director}
		reverseProxy.ServeHTTP(c.Writer, c.Request)
	}
}