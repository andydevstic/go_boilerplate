package pet

import (
	"github.com/gin-gonic/gin"
)

type router struct {
	controller IPetController
}

func NewRouter(controller IPetController) router {
	return router{controller: controller}
}

func (r *router) Route(rg *gin.RouterGroup) {
	_ = rg.Group("/users")
}
