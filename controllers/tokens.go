package controllers

import (
	"fmt"
	"github.com/RockRockWhite/minio-client/dtos"
	"github.com/RockRockWhite/minio-client/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

// CreateToken 创建token
func CreateToken(c *gin.Context) {
	accessKey := struct {
		AccessKeyID     string
		SecretAccessKey string
	}{}

	if err := c.ShouldBind(&accessKey); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "Bind Model Error",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	// 验证AccessKey
	if accessKey.AccessKeyID != viper.GetString("Minio-Client.AccessKeyID") ||
		accessKey.SecretAccessKey != viper.GetString("Minio-Client.SecretAccessKey") {
		c.JSON(http.StatusBadRequest, dtos.ErrorDto{
			Message:          "invalid access key.",
			DocumentationUrl: viper.GetString("Document.Url"),
		})
		return
	}

	token, err := utils.GenerateJwtToken(&utils.JwtClaims{})
	if err != nil {
		panic(fmt.Sprintf("Failed to generate JwtToken"))
	}

	c.JSON(http.StatusCreated, struct {
		AccessToken string
	}{AccessToken: token})
}
