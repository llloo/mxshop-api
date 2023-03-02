package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-api/api"
)

func InitUserRouter(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/list", api.GetUserList)
	}
}
