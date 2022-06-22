package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GlobalMiddleware(ctx *gin.Context) {
	zap.S().Info("#############hello####################")
	ctx.Next()
}