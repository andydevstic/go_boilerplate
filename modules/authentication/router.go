package authentication

import (
	model "github.com/andydevstic/boilerplate-backend/models"
	"github.com/andydevstic/boilerplate-backend/shared/middlewares"
	"github.com/gin-gonic/gin"
)

type router struct {
	controller IAuthController
}

func NewRouter(controller IAuthController) router {
	return router{controller: controller}
}

func (r *router) Route(rg *gin.RouterGroup) {
	authRouter := rg.Group("/auth")

	authRouter.POST("/register", middlewares.JsonValidationMiddleware[model.RegisterUserDTO], r.controller.Register)
	authRouter.POST("/login", middlewares.JsonValidationMiddleware[model.LoginDTO], r.controller.Login)
}
