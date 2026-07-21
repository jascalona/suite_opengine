package router

import (
	"opengine/v2/internal/handler"

	"github.com/gin-gonic/gin"
)

type RouterComponent struct {
	Envi   *handler.EnvironmentHandler
	Serv   *handler.ServicesHandler
	Resour *handler.ResourcesHandler
	SubR   *handler.SubResourcesHandler
	Endp   *handler.EndpointsHandler
	Acco   *handler.AccountHandler
}

func NewRouterComponent(
	envi *handler.EnvironmentHandler,
	serv *handler.ServicesHandler,
	resour *handler.ResourcesHandler,
	subR *handler.SubResourcesHandler,
	endp *handler.EndpointsHandler,
	acco *handler.AccountHandler,
) *RouterComponent {
	return &RouterComponent{
		Envi:   envi,
		Serv:   serv,
		Resour: resour,
		SubR:   subR,
		Endp:   endp,
		Acco:   acco,
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

	subresources := rg.Group("subresources")
	{
		subresources.POST("", r.SubR.CreatedSubResources)
		subresources.GET("", r.SubR.SubResourcesByResources)

	}

	endpoints := rg.Group("endpoints")
	{
		endpoints.POST("", r.Endp.CreatedEndpoints)
	}

	accounts := rg.Group("accounts")
	{
		accounts.POST("", r.Acco.CreatedAccount)
		accounts.GET("", r.Acco.ListAccount)
	}
}
