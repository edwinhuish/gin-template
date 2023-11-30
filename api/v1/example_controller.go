package v1

import (
	"os"
	"time"

	"gin-template/global"
	"gin-template/utils/response"

	"github.com/gin-gonic/gin"
)

type ExampleController struct{}

func NewExampleController() *ExampleController {
	return &ExampleController{}
}

func (controller *ExampleController) Hello(c *gin.Context) {

	response.OkWithMessage(c, "hello world!")
}

func (controller *ExampleController) Upload(c *gin.Context) {
	// 单文件
	file, err := c.FormFile("Filedata")
	if err != nil {
		global.LOG.Error("上传文件出错：" + err.Error())
		return
	}

	filename := c.PostForm("Filename")

	datepath := time.Now().Format("2006-01-02")
	os.MkdirAll("cache/"+datepath, os.ModePerm)

	fullpath := "cache/" + datepath + "/" + filename

	err = c.SaveUploadedFile(file, fullpath)
	if err != nil {
		response.FailWithMessage(c, "保存文件出错")
		global.LOG.Error(err.Error())
		return
	}

	response.OkWithMessage(c, "上传文件成功，文件保存在："+fullpath)
}
