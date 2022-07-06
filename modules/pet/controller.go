package pet

import (
	"fmt"
	"net/http"

	"github.com/andydevstic/boilerplate-backend/shared"
	"github.com/andydevstic/boilerplate-backend/shared/constants"
	"github.com/andydevstic/boilerplate-backend/shared/custom"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type IPetController interface {
}

type PetController struct {
	service IPetService
}

func NewController(service IPetService) IPetController {
	return &PetController{
		service: service,
	}
}

func (controller *PetController) Create(ctx *gin.Context) {
	var dto CreatePetDTO = ctx.MustGet(constants.ParsedDtoKey).(CreatePetDTO)

	createPayload := make(map[string]any)
	err := mapstructure.Decode(dto, &createPayload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": constants.UnprocessableEntityErrorMsg,
			"detail":  err.Error(),
		})

		return
	}

	err = controller.service.Create(ctx, createPayload)
	if err != nil {
		custom.HandleCustomError(ctx, fmt.Errorf("create new pet: %w", err))

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}

func (controller *PetController) Find(ctx *gin.Context) {
	var dto shared.FindDTO = ctx.MustGet(constants.ParsedDtoKey).(shared.FindDTO)

	fmt.Printf("%v \n", dto)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    nil,
	})
}
