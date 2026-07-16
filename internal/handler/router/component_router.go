package router

import (
	"opengine/v2/internal/handler"

	"github.com/gin-gonic/gin"
)

type RouterComponent struct {
	Envi *handler.EnvironmentHandler
	Serv *handler.ServicesHandler
}

func NewRouterComponent(
	envi *handler.EnvironmentHandler,
	serv *handler.ServicesHandler,
) *RouterComponent {
	return &RouterComponent{
		Envi: envi,
		Serv: serv,
	}
}

func (r *RouterComponent) RegisterComponents(rg *gin.RouterGroup) {

	environments := rg.Group("environments")
	{
		environments.POST("", r.Envi.CreatedEnvironment)
		environments.GET("", r.Envi.ListEnvironment)
	}

	services := rg.Group("services")
	{
		services.POST("", r.Serv.CreatedServices)
		services.GET("", r.Serv.ListServices)
	}
}
