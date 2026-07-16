package router

import (
	"opengine/v2/internal/domain"
	"opengine/v2/internal/handler"

	"github.com/gin-gonic/gin"
)

type MainRouter struct {
	RouterComponent    *RouterComponent
	RouterOrganization *RouterOrganization
}

// funcion de carga manual para los usuarios iniciales (depuracion)
func RegisterUser(r *gin.Engine, ru *handler.UserHandler) {
	register := r.Group("register")
	{
		register.POST("", ru.CreateUser)
	}
}

func SetupRouter(r *gin.Engine, authService domain.AuthServices, routers MainRouter) {
	// GESTION DE ACCESO PUBLICO
	auth := &handler.AuthHandler{Service: authService}

	r.POST("/login", auth.Login)
	r.GET("/health", func(c *gin.Context) { c.Status(200) })

	// GESTION DE SERVICIOS PROTEGIDOS

	api_v1 := r.Group("/api/v1")
	api_v1.Use(handler.Auth(authService))

	{
		routers.RouterComponent.RegisterComponents(api_v1.Group("/components"))
		routers.RouterOrganization.RegisterOrganization(api_v1.Group("/organization"))

	}

}
