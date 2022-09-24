package controllers

import (
	"fmt"
	"github.com/RockRockWhite/minio-client/config"
	"github.com/RockRockWhite/minio-client/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strings"
)

var _addr, _port string

func init() {
	_addr = config.GetString("HttpServer.Addr")
	_port = config.GetString("HttpServer.Port")
}

func GetFile(c *gin.Context) {
	filename := c.Param("filename")
	obj, err := utils.GetObject(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, _ := io.ReadAll(obj)
	fileContentDisposition := "attachment;filename=\"" + filename + "\""
	c.Header("Content-Disposition", fileContentDisposition)
	c.Data(http.StatusOK, "", res)
}

func UploadObject(c *gin.Context) {

	file, _ := c.FormFile("file")
	r, err := file.Open()
	if err != nil {
		fmt.Println(err)
	}
	filename, postfix := utils.GetFileNameAndPosfix(file.Filename)
	id, _ := uuid.NewUUID()
	filename += strings.Replace(id.String(), "-", "", -1)

	err = utils.PutObject(fmt.Sprintf("%s.%s", filename, postfix), r, file.Size)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, struct {
		Url string
	}{
		Url: fmt.Sprintf("%s%s/buckets/%s.%s", _addr, _port, filename, postfix),
	})
}
