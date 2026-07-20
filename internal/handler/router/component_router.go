package router

import (
	"opengine/v2/internal/handler"

	"github.com/gin-gonic/gin"
)

type RouterComponent struct {
	Envi   *handler.EnvironmentHandler
	Serv   *handler.ServicesHandler
	Resour *handler.ResourcesHandler
}

func NewRouterComponent(
	envi *handler.EnvironmentHandler,
	serv *handler.ServicesHandler,
	resour *handler.ResourcesHandler,
) *RouterComponent {
	return &RouterComponent{
		Envi:   envi,
		Serv:   serv,
		Resour: resour,
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

	resources := rg.Group("resources")
	{
		resources.POST("", r.Resour.CreatedResources)
		resources.GET("", r.Resour.ResourcesByService)
	}
}
