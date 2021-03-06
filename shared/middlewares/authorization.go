package middlewares

import (
	"net/http"
	"strings"

	"github.com/andydevstic/boilerplate-backend/shared"
	"github.com/andydevstic/boilerplate-backend/shared/constants"
	jwtutils "github.com/andydevstic/boilerplate-backend/shared/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func AuthGuard(context *gin.Context) {
	authHeader := context.GetHeader("authorization")
	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})

		return
	}

	splitted := strings.Split(authHeader, " ")
	if splitted[1] == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})

		return
	}

	mapClaim, err := jwtutils.VerifyToken(splitted[1])
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"error":   err.Error(),
		})

		return
	}

	var userPayload shared.UserAuthPayload

	err = mapstructure.Decode(mapClaim, &userPayload)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Corrupted auth payload",
		})

		return
	}

	context.Set(constants.UserAuthPayload, userPayload)

	context.Next()
}
