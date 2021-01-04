package fileUploader

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"httpTools/src/infrastructure/httpUtil"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type DirUploads []*DirUploader

func (u DirUploads) Apply(engine *gin.Engine) {
	for _, up := range u {
		up.Apply(engine)
	}
}

type DirUploader struct {
	List   *VFileList `json:"list"`
	Route  *VRoute    `json:"route"`
	Target *VTarget   `json:"target"`
}

func NewDirUploader(VDir *VUploadDir, route *VRoute, target *VTarget) *DirUploader {
	return &DirUploader{NewVFileList(VDir), route, target}
}

func (d *DirUploader) Apply(engine *gin.Engine) {
	engine.POST(d.Route.Path, httpUtil.ErrorWrapper(
		d.uploadApi,
		func(err error) string {
			return err.Error()
		},
	))
}

func (d *DirUploader) uploadApi(c *gin.Context) error {
	var list []string
	d.UploadDir(d.Target, func(resp *http.Response, err error) {
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, string(respBody))
		fmt.Println("RESULT: ", resp.Status, string(respBody))
	})
	c.JSON(http.StatusOK, list)
	return nil
}

func (d *DirUploader) UploadDir(target *VTarget, callBack func(resp *http.Response, err error)) {
	for _, filepath := range d.List.GetFileList() {
		// ./uploadApi/aaa/txt -> (./uploadApi/aaa/txt, http://xxxx/www/aaa/txt)
		log.Println("UPLOAD: ", d.List.GetDirPath()+filepath, target.URL+filepath)
		resp, err := uploadFile(d.List.GetDirPath()+filepath, target.URL+filepath)
		callBack(resp, err)
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
		log.Fatal("error opening file")
	}
	defer fh.Close()

	// io copy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		log.Fatal(err)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		log.Fatal(err)
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
