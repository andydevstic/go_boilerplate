package authentication

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andydevstic/boilerplate-backend/modules/user"
	"github.com/andydevstic/boilerplate-backend/shared"
	"github.com/andydevstic/boilerplate-backend/shared/constants"
	"github.com/andydevstic/boilerplate-backend/shared/custom"
	"github.com/gin-gonic/gin"
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
	authService IAuthService
}

func NewController(userService user.IUserService, authService IAuthService) IAuthController {
	return &AuthController{
		userService: userService,
		authService: authService,
	}
}

func (controller *AuthController) Login(context *gin.Context) {
	var dto LoginDTO = context.MustGet(constants.ParsedDtoKey).(LoginDTO)

	user, err := controller.userService.FindOne(context, map[string]any{"email": dto.Email})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "email or password is incorrect",
			})

			return
		}

		log.Error().Msgf("find user by email: %s", err)
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": constants.InternalServerErrorMsg,
		})

		return
	}

	if err = controller.authService.ValidateUserPassword(dto.Password, user.Password); err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "email or password is incorrect",
		})

		return
	}

	jwtToken, err := controller.authService.GenerateJwtFromUser(&user)
	if err != nil {
		err = fmt.Errorf("generate jwt: %w", err)
		custom.HandleCustomError(context, err)

		return
	}

	context.JSON(http.StatusOK, LoginResponse{
		Token: jwtToken,
		User: shared.UserAuthPayload{
			Id:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			Name:     user.Name,
			Type:     user.Type,
			Status:   user.Status,
		},
	})

}

func (controller *AuthController) Register(context *gin.Context) {
	var dto RegisterUserDTO = context.MustGet(constants.ParsedDtoKey).(RegisterUserDTO)

	var err error

	defer func() {
		if err != nil {
			log.Error().Msgf("register: %s", err)
		}
	}()

	var payload map[string]any

	err = mapstructure.Decode(dto, &payload)
	if err != nil {
		err = fmt.Errorf("decode user payload: %w", err)
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": constants.InternalServerErrorMsg,
		})

		return
	}

	err = controller.userService.Create(context, payload)
	if err != nil {
		err = fmt.Errorf("create user: %w", err)
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": constants.InternalServerErrorMsg,
		})

		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Registered successfully!",
	})
}
