package pet

import (
	"github.com/andydevstic/boilerplate-backend/shared/middlewares"
	"github.com/gin-gonic/gin"
)

type router struct {
	controller IPetController
}

func NewRouter(controller IPetController) router {
	return router{controller: controller}
}

func (r *router) Route(rg *gin.RouterGroup) {
	petRoutes := rg.Group("/pets")

	petRoutes.GET("", middlewares.QueryValidationMiddleware[FindPetsDTO], r.controller.Find)
	petRoutes.POST("", middlewares.JsonValidationMiddleware[CreatePetDTO], r.controller.Create)
}
