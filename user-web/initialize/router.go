package initialize

import (
	"mxshop_api/user-web/middlewares"
	"mxshop_api/user-web/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	Router := gin.Default()

	Router.Use(middlewares.GlobalMiddleware)

	Router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"success": true,
		})
	})

	ApiGroup := Router.Group("/u/v1")
	router.InitUserRouter(ApiGroup)

	return Router

}