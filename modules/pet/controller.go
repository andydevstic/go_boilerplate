package pet

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andydevstic/boilerplate-backend/core"
	"github.com/andydevstic/boilerplate-backend/shared/constants"
	"github.com/andydevstic/boilerplate-backend/shared/custom"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type IPetController interface {
	Find(*gin.Context)
	Create(*gin.Context)
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
	var dto FindPetsDTO = ctx.MustGet(constants.ParsedDtoKey).(FindPetsDTO)

	appState := core.GetAppState()
	db := appState.Db

	err := dto.AddFilterToStatement(db)
	if err != nil {
		custom.HandleCustomError(ctx, fmt.Errorf("add filter to statement: %w", err))

		return
	}

	pets := make([]Pet, 0, 10)

	tx := db.Find(&pets)

	if tx.Error != nil {
		custom.HandleCustomError(ctx, fmt.Errorf("find pets: %w", err))

		return
	}

	result, err := json.Marshal(pets)
	if err != nil {
		custom.HandleCustomError(ctx, fmt.Errorf("marshal pets: %w", err))

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}
