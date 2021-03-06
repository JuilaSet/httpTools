package fileUploader

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"httpTools/src/domain/httpServer/model/fileUploader/vo"
	"httpTools/src/infrastructure/exception"
	"httpTools/src/infrastructure/httpUtil"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type DirUploads []*DirUploader

func (u DirUploads) Apply(engine *gin.Engine) {
	for _, up := range u {
		up.Apply(engine)
	}
}

type DirUploader struct {
	List   *vo.VFileList  `json:"list"`
	Route  *vo.VRoute     `json:"route"`
	Target *vo.VTarget    `json:"target"`
	Exclude *vo.VExcludes `json:"exclude"`
}

func NewDirUploader(VDir *vo.VUploadDir, route *vo.VRoute, target *vo.VTarget, exclude *vo.VExcludes) *DirUploader {
	return &DirUploader{
		vo.NewVFileList(VDir),
		route,
		target,
		exclude,
	}
}

func (d *DirUploader) Apply(engine *gin.Engine) {
	engine.Use(d.refresh())
	engine.POST(d.Route.Path, httpUtil.ErrorWrapper(
		d.uploadApi,
		func(err error) string {
			return err.Error()
		},
	))
	engine.PUT(d.Route.Path, httpUtil.ErrorWrapper(
		d.syncDirApi,
		func(err error) string {
			return err.Error()
		},
	))
}

func (d *DirUploader) refresh() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 刷新列表
		d.List.Refresh()
	}
}

func (d *DirUploader) syncDirApi(c *gin.Context) error {
	payload := strings.NewReader("")
	req, _ := http.NewRequest(http.MethodDelete, d.Target.URL, payload)
	_, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	return d.uploadApi(c)
}

func (d *DirUploader) uploadApi(c *gin.Context) error {
	// 刷新列表
	d.List.Refresh()

	var list = make([]string, 0)
	d.UploadDir(d.Target, func(resp *http.Response, err error) {
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		list = append(list, string(respBody))
		fmt.Println("RESULT: ", resp.Status, string(respBody))
	})
	c.JSON(http.StatusOK, list)
	return nil
}

func (d *DirUploader) UploadDir(target *vo.VTarget, callBack func(resp *http.Response, err error)) {
	for _, filepath := range d.List.GetFileList() {
		// ./uploadApi/aaa/txt -> (./uploadApi/aaa/txt, http://xxxx/www/aaa/txt)
		b := true
		for _, ex := range d.Exclude.FilePathList {
			if strings.HasPrefix(filepath, ex) {
				b = false
				break
			}
		}
		if b {
			log.Println("UPLOAD: ", d.List.GetDirPath()+filepath, target.URL+filepath)
			resp, err := uploadFile(d.List.GetDirPath()+filepath, target.URL+filepath)
			callBack(resp, err)
		}
	}
}

// filepath: /xxx
func uploadFile(filepath, targetUrl string) (*http.Response, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// 关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("file", filepath)
	if err != nil {
		fmt.Println("error writing to buffer")
	}

	// 打开文件句柄操作
	fh, err := os.Open(filepath)
	if err != nil {
		panic(errors.New("error opening file"))
	}
	defer fh.Close()

	// io copy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		panic(err)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		panic(exception.ConnectTimeoutException)
	}

	//// 返回结果
	//defer resp.Body.Close()
	//respBody, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(resp.Status)
	return resp, err
}
