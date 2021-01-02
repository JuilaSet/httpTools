package staticFileInfo

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"httpTools/src/domain/httpServer/vo"
	"httpTools/src/infrastructure/httpUtil"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type Statics []*StaticFileInfo
type StaticFileInfo struct {
	Dir   *vo.VDir   `json:"Dir"`
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

	// 文件浏览
	engine.Static(static.Route.Path, static.Dir.Path)

	// 文件删除
	engine.DELETE(static.Route.Path+"/*Name", httpUtil.ErrorWrapper(static.fileDeleteHandler, func(err error) string {
		return err.Error()
	}))

	// 文件存储
	engine.MaxMultipartMemory = 8 << 20
	engine.POST(static.Route.Path+"/*Name", httpUtil.ErrorWrapper(static.filePostHandler, func(err error) string {
		return err.Error()
	}))
}

func (static *StaticFileInfo) fileDeleteHandler(c *gin.Context) (err error) {
	fileInfo := vo.NewFileNameInfo(c.Param("Name"))

	// 删除文件
	if err = os.Remove(fileInfo.FilePath(static.Dir.Path)); err != nil {
		// 删除失败
		return err
	}

	// 回显文件名
	c.String(http.StatusOK, fileInfo.FilePath(static.Route.Path))
	return
}

func (static *StaticFileInfo) filePostHandler(c *gin.Context) (err error) {
	fileInfo := vo.NewFileNameInfo(c.Param("Name"))
	var (
		file multipart.File
		dst *os.File
	)
	defer func() {
		if dst != nil {
			dst.Close()
		}
		if file != nil {
			file.Close()
		}
	}()

	fmt.Printf("fullFilename %s\n", fileInfo.FilePath(static.Dir.Path))
	fmt.Printf("Dir %s\n", fileInfo.FileDir(static.Dir.Path))
	fmt.Printf("Name %s\n", fileInfo.FileName())

	// 接受文件
	file, _, err = c.Request.FormFile("file")
	if err != nil {
		c.AbortWithError(500, errors.New("upload error"))
	}

	err = os.MkdirAll(fileInfo.FileDir(static.Dir.Path), os.ModePerm)
	if err != nil {
		c.String(500, "save Dir error")
		log.Fatal(err.Error())
	}

	dst, err = os.Create(fileInfo.FilePath(static.Dir.Path))
	if err != nil {
		c.String(500, "save file error")
		log.Fatal(err.Error())
	}

	// 将文件拷贝到指定路径下，或者其他文件操作
	_, err = io.Copy(dst, file)
	if err != nil {
		c.String(500, "save copy error")
		log.Fatal(err.Error())
	}

	// 回显文件名
	c.String(http.StatusOK, fileInfo.FilePath(static.Route.Path))
	return
}

