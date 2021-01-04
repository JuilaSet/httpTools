package staticFileInfo

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wxnacy/wgo/file"
	"httpTools/src/domain/httpServer/vo"
	"httpTools/src/infrastructure/fileUtil"
	"httpTools/src/infrastructure/httpUtil"
	"io"
	"io/ioutil"
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
	filename := c.Param("Name")

	// 当前文件夹下所有文件
	if filename == "/" {
		var list = make([]string, 0)
		rd, _ := ioutil.ReadDir(static.Dir.Path)
		for _, fi := range rd {
			if err = os.RemoveAll(static.Dir.Path + "/" + fi.Name()); err != nil {
				// 删除失败
				return err
			} else {
				list = append(list, fi.Name())
			}
		}
		// 回显文件名url
		c.JSON(http.StatusOK, list)
	} else {
		fileInfo := vo.NewFileNameInfo(filename)
		// 删除文件
		filepath := fileInfo.FilePath(static.Dir.Path)
		if file.IsDir(filepath) {
			if err = os.RemoveAll(filepath); err != nil {
				// 删除失败
				return err
			}
		} else {
			if err = os.Remove(filepath); err != nil {
				// 删除失败
				return err
			}
		}

		// 回显文件名url
		c.String(http.StatusOK, fileInfo.FilePath(static.Route.Path))
	}
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

	// 如果源文件存在，就删除
	if fileUtil.IsExist(fileInfo.FilePath(static.Dir.Path)) {
		os.Remove(fileInfo.FilePath(static.Dir.Path))
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

