package routes

import (
	"github.com/gin-gonic/gin"
	"yueshang/app/controllers"
	"yueshang/websockets"
)

var wsCtl = new(websockets.WsController)
var indexCtl = new(controllers.IndexController)

func SetupRouter() *gin.Engine{
	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("/", indexCtl.Welcome)
	v1 := router.Group("/v1")
	{
		v1.GET("/ws", func(c *gin.Context) {
			wsCtl.WsHandler(c.Writer, c.Request)
		})

	}
	return router
}