package authentication

import (
	"net/http"

	"github.com/andydevstic/boilerplate-backend/core"
	"github.com/andydevstic/boilerplate-backend/shared/constants"
	"github.com/andydevstic/boilerplate-backend/shared/custom"
	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
}

type AuthController struct {
	service  IAuthService
	appState *core.AppState
}

func NewController() IAuthController {
	return &AuthController{
		service:  NewService(),
		appState: core.GetAppState(),
	}
}

type LoginResponse struct {
	Token string          `json:"token"`
	User  UserAuthPayload `json:"user"`
}

func (controller *AuthController) Login(context *gin.Context) {
	var dto LoginDTO = context.MustGet(constants.ParsedDtoKey).(LoginDTO)

	user, jwtToken, err := controller.service.Login(context, controller.appState.Store, &dto)

	defer func() {
		if err != nil {
			custom.HandleCustomError(context, err)
		}
	}()

	if err != nil {
		return
	}

	context.JSON(http.StatusOK, LoginResponse{
		Token: jwtToken,
		User: UserAuthPayload{
			Id:     user.Id,
			Email:  user.Email,
			Name:   user.Name,
			Type:   user.Type,
			Status: user.Status,
		},
	})
}

func (controller *AuthController) Register(context *gin.Context) {
	var dto RegisterUserDTO = context.MustGet(constants.ParsedDtoKey).(RegisterUserDTO)

	if err := controller.service.Register(context, controller.appState.Store, &dto); err != nil {
		custom.HandleCustomError(context, err)

		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}
