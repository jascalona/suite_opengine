package router

import (
	"opengine/v2/internal/handler"

	"github.com/gin-gonic/gin"
)

type RouterComponent struct {
	Envi *handler.EnvironmentHandler
}

func NewRouterComponent(
	envi *handler.EnvironmentHandler,
) *RouterComponent {
	return &RouterComponent{
		Envi: envi,
	}
}

func (r *RouterComponent) RegisterComponents(rg *gin.RouterGroup) {

	environments := rg.Group("environments")
	{
		environments.POST("", r.Envi.CreatedEnvironment)
	}
}
