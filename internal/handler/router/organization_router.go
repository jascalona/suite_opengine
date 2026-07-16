package router

import (
	"opengine/v2/internal/handler"

	"github.com/gin-gonic/gin"
)

type RouterOrganization struct {
	users_h *handler.UserHandler
}

func NewRouterOrganization(
	users *handler.UserHandler,
) *RouterOrganization {
	return &RouterOrganization{
		users_h: users,
	}
}

func (r *RouterOrganization) RegisterOrganization(rg *gin.RouterGroup) {

	users := rg.Group("users")
	{
		users.GET("", r.users_h.GetAllUsers)
		users.POST("", r.users_h.CreateUser)
	}

}
