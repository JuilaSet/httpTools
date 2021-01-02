package staticFileInfo

import (
	"github.com/gin-gonic/gin"
	"httpTools/src/domain/httpServer/vo"
	"log"
)

type Statics []*StaticFileInfo
type StaticFileInfo struct {
	Dir   *vo.VDir   `json:"dir"`
	Route *vo.VRoute `json:"route"`
}

func NewStaticFileInfo(builders ...Builder) *StaticFileInfo {
	c := &StaticFileInfo{
		Dir:   vo.NewVDir("."),
		Route: vo.NewVRoute(""),
	}
	Builders(builders).apply(c)
	return c
}

// 构建静态服务器
func (statics Statics) Apply(engine *gin.Engine) {
	log.Println("Static infos", statics)
	for _, static := range statics {
		static.Apply(engine)
	}
}

func (static *StaticFileInfo) Apply(engine *gin.Engine) {
	log.Println("Static", static.Route.Path, static.Dir.Path)
	engine.Static(static.Route.Path, static.Dir.Path)
}