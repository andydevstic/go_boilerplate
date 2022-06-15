package user

import (
	"github.com/gin-gonic/gin"
)

type router struct {
	controller IUserController
}

func NewRouter(controller IUserController) router {
	return router{controller: controller}
}

func (r *router) Route(rg *gin.RouterGroup) {
	_ = rg.Group("/users")
}
