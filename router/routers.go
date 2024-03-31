package router

import (
	"gin-http-server/controller"
	"github.com/gin-gonic/gin"
)

// RegisterRouter 路由设置
func RegisterRouter(router *gin.Engine) {
	routerUser(router)
}

// 用户路由
func routerUser(engine *gin.Engine) {
	var group = engine.Group("/api/user")
	{
		c := &controller.UserController{}
		group.POST("/add", c.Add)
		group.GET("/get", c.Get)
	}
}
