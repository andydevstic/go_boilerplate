package middlewares

import (
	"github.com/gin-gonic/gin"
)

func NoRouteHandler(context *gin.Context) {
	context.JSON(404, gin.H{
		"code":    "PAGE_NOT_FOUND",
		"message": "Route not found!",
	})
}
