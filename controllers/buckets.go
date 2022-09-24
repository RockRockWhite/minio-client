package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFile(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func PutFile(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
