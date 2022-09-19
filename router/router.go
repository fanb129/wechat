package router

import (
	"github.com/gin-gonic/gin"
	"wechat/controller"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.POST("/hello", controller.Hello)

	return r
}
