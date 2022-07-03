package authentication

import (
	"errors"
	"net/http"

	"github.com/andydevstic/boilerplate-backend/modules/user"
	"github.com/andydevstic/boilerplate-backend/shared"
	"github.com/andydevstic/boilerplate-backend/shared/constants"
	jwtutils "github.com/andydevstic/boilerplate-backend/shared/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type IAuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
}

type AuthController struct {
	userService user.IUserService
}

func NewController(userService user.IUserService) IAuthController {
	return &AuthController{
		userService: userService,
	}
}

type LoginResponse struct {
	Token string                 `json:"token"`
	User  shared.UserAuthPayload `json:"user"`
}

func (controller *AuthController) Login(context *gin.Context) {
	var dto LoginDTO = context.MustGet(constants.ParsedDtoKey).(LoginDTO)

	user, err := controller.userService.FindOne(context, map[string]any{"email": dto.Email})

	if err == nil {
		jwtClaim := jwt.MapClaims{}

		err = mapstructure.Decode(user, jwtClaim)
		if err != nil {
			log.Error().Msgf("decode user into jwt: %s", err)
			context.AbortWithError(http.StatusInternalServerError, errors.New(constants.InternalServerErrorMsg))

			return
		}

		jwtToken, err := jwtutils.SignPayload(jwtClaim)
		if err != nil {
			log.Error().Msgf("sign jwt token: %s", err)
			context.AbortWithError(http.StatusInternalServerError, errors.New(constants.InternalServerErrorMsg))

			return
		}

		context.JSON(http.StatusOK, LoginResponse{
			Token: jwtToken,
			User: shared.UserAuthPayload{
				Id:     user.ID,
				Email:  user.Email,
				Name:   user.Name,
				Type:   user.Type,
				Status: user.Status,
			},
		})
	} else {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithError(http.StatusConflict, errors.New("Email already taken"))

			return
		}

		log.Error().Msgf("find user by email: %s", err)
		context.AbortWithError(http.StatusInternalServerError, errors.New(constants.InternalServerErrorMsg))
	}
}

func (controller *AuthController) Register(context *gin.Context) {
	var dto RegisterUserDTO = context.MustGet(constants.ParsedDtoKey).(RegisterUserDTO)

	var payload map[string]any

	err := mapstructure.Decode(dto, payload)
	if err != nil {
		log.Error().Msgf("decode user payload: %s", err)
		context.AbortWithError(http.StatusInternalServerError, errors.New(constants.InternalServerErrorMsg))

		return
	}

	err = controller.userService.Create(context, payload)
	if err != nil {
		log.Error().Msgf("create user: %s", err)
		context.AbortWithError(http.StatusInternalServerError, errors.New(constants.InternalServerErrorMsg))

		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Registered successfully!",
	})
}
