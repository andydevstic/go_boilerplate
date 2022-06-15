package middlewares

import (
	"errors"
	"io"
	"net/http"

	"github.com/andydevstic/boilerplate-backend/shared/constants"
	"github.com/gin-gonic/gin"
)

func QueryValidationMiddleware[DTO any](context *gin.Context) {
	var dto DTO

	err := context.BindQuery(&dto)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	context.Set(constants.ParsedDtoKey, dto)

	context.Next()
}

func JsonValidationMiddleware[DTO any](context *gin.Context) {
	var dto DTO

	err := context.BindJSON(&dto)

	if err != nil {
		msg := err.Error()

		if errors.Is(err, io.EOF) {
			msg = "Body must not be null"
		}

		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": msg,
		})

		return
	}

	context.Set(constants.ParsedDtoKey, dto)

	context.Next()
}
