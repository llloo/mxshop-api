package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-api/router"
)

func InitRouters() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/v1")
	router.InitUserRouter(apiGroup)

	return r
}
