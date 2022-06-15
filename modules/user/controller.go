package user

import (
	"net/http"
	"strconv"

	"github.com/andydevstic/boilerplate-backend/core"
	"github.com/andydevstic/boilerplate-backend/shared/constants"
	"github.com/andydevstic/boilerplate-backend/shared/custom"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	FindUsersAdmin(context *gin.Context)
	FindUserByEmail(context *gin.Context)
	FindUserById(context *gin.Context)
}

type UserController struct {
	service  *UserService
	appState *core.AppState
}

func NewUser() *UserController {
	return &UserController{
		service:  NewService(),
		appState: core.GetAppState(),
	}
}

func (controller *UserController) FindUsersAdmin(context *gin.Context) {
	var dto FindUsersAdminDTO = context.MustGet(constants.ParsedDtoKey).(FindUsersAdminDTO)

	users, err := controller.service.FindUsersAdmin(context, controller.appState.Store, &dto)

	if err != nil {
		custom.HandleCustomError(context, err)

		return
	}

	context.JSON(http.StatusOK, users)
}

func (controller *UserController) FindUserByEmail(context *gin.Context) {
	var dto FindUserByEmail = context.MustGet(constants.ParsedDtoKey).(FindUserByEmail)

	user, err := controller.service.FindUserByEmail(context, controller.appState.Store, dto.Email)
	if err != nil {
		custom.HandleCustomError(context, err)

		return
	}

	context.JSON(http.StatusOK, user)
}

func (controller *UserController) FindUserById(context *gin.Context) {
	userId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		custom.HandleCustomError(context, err)

		return
	}

	user, err := controller.service.FindUserById(context, controller.appState.Store, userId)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, user)
}

func (controller *UserController) UpdateUserProfile(context *gin.Context) {
	// var dto model.UpdateUserDTO = context.MustGet(constants.ParsedDtoKey).(model.UpdateUserDTO)
	// var user model.UserAuthPayload = context.MustGet(constants.UserAuthPayload).(model.UserAuthPayload)
}
