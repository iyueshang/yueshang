package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type IndexController struct{}

func (i *IndexController) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":      200,
		"msg":       "Welcome.",
		"data":      nil,
		"timestamp": time.Now().Unix(),
	})

	return
}
