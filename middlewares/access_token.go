package middlewares

import (
	"github.com/RockRockWhite/minio-client/dtos"
	"github.com/RockRockWhite/minio-client/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

// AccessToken AccessToken验证中间件
func AccessToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" || strings.Fields(token)[0] != "Bearer" {
		// 没有传token参数
		c.JSON(http.StatusUnauthorized, dtos.ErrorDto{
			Message:          "Requires bearer token in filed {Authorization}.",
			DocumentationUrl: viper.GetString("Document.Url"),
		})

		c.Abort()
		return
	}
	token = strings.Fields(token)[1]

	claims, err := utils.ParseJwtToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, dtos.ErrorDto{
			Message:          "Token expired or the other error occurred",
			DocumentationUrl: viper.GetString("Document.Url"),
		})

		c.Abort()
		return
	}
	// Claims写入上下文
	c.Set("claims", claims)
}
