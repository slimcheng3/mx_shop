package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(ctx *gin.Context) {
	result := make([]interface{}, 0)
	result = append(result, map[string]string{"name": "slim"})
	ctx.JSON(http.StatusOK, result)
}