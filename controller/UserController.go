package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

// Add 新增用户
func (controller *UserController) Add(context *gin.Context) {
	name, exist := context.GetPostForm("name")
	if !exist || name == "" {
		context.JSON(http.StatusOK, gin.H{
			"msg": "请输入用户名:name",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

// Get 查询用户
func (controller *UserController) Get(context *gin.Context) {
	id := context.Query("id")
	println("id:", id)
	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
