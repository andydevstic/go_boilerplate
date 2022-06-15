package custom

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, message error) CustomError {
	return CustomError{Code: code, Message: message.Error()}
}

func (c CustomError) Error() string {
	return c.Message
}

func HandleCustomError(context *gin.Context, err error) {
	customError, isCustomError := err.(CustomError)
	log.Error().Msg(err.Error())

	if !isCustomError {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal Server Error",
		})

		return
	}

	context.AbortWithStatusJSON(customError.Code, gin.H{
		"success": false,
		"message": customError.Message,
	})
}
