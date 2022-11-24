package controllers

import (
	"fmt"
	"github.com/RockRockWhite/minio-client/config"
	"github.com/RockRockWhite/minio-client/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var _addr, _port string

func init() {
	_addr = config.GetString("HttpServer.ExternalAddr")
	_port = config.GetString("HttpServer.Port")
}

func GetFile(c *gin.Context) {
	escaped := c.Param("objectname")
	objname, err := url.PathUnescape(escaped)
	if err != nil {
		utils.GetLogger().Fatalf("[fatalf] Failed to decode url %s, %s", escaped, err.Error())
		return
	}

	obj, err := utils.GetObject(objname)
	if err != nil {
		utils.GetLogger().Fatalf("[fatalf] Failed to get object %s, %s", objname, err.Error())
		return
	}

	res, _ := io.ReadAll(obj)
	fileContentDisposition := "attachment;objname=\"" + objname + "\""
	c.Header("Content-Disposition", fileContentDisposition)
	c.Data(http.StatusOK, "", res)
}

func UploadObject(c *gin.Context) {
	file, _ := c.FormFile("file")
	obj, err := file.Open()
	if err != nil {
		utils.GetLogger().Fatalf("[fatalf] Failed to open file, %s", err.Error())
	}
	prefix, postfix := utils.GetPrefixAndPosfix(file.Filename)

	id, _ := uuid.NewUUID()
	prefix += strings.Replace(id.String(), "-", "", -1)

	err = utils.PutObject(fmt.Sprintf("%s.%s", prefix, postfix), obj, file.Size)
	if err != nil {
		utils.GetLogger().Fatalf("[fatalf] Failed to put object %s, %s", fmt.Sprintf("%s.%s", prefix, postfix), err.Error())
	}

	c.JSON(http.StatusCreated, struct {
		Url string
	}{
		Url: fmt.Sprintf("%s/buckets/%s.%s", _addr, url.PathEscape(prefix), url.PathEscape(postfix)),
	})
}
