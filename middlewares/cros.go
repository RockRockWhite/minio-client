package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
	context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	context.Header("Access-Control-Allow-Credentials", "true")
	if context.Request.Method == "OPTIONS" {
		context.AbortWithStatus(http.StatusNoContent)
		return
	}
	context.Next()
}
