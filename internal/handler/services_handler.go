package handler

import (
	"log"
	"net/http"
	"opengine/v2/internal/domain/components"
	usecase "opengine/v2/internal/usecase/components"
	"opengine/v2/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ServicesHandler struct {
	Services usecase.ServicesServ
}

func NewServicesHandler(s usecase.ServicesServ) *ServicesHandler {
	return &ServicesHandler{Services: s}
}

func (h *ServicesHandler) CreatedServices(c *gin.Context) {
	var reqServ components.ValidateServices

	if err := c.ShouldBindJSON(&reqServ); err != nil {
		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Println("Error en la validacion del mensaje ", err.Error())
			c.JSON(http.StatusConflict, gin.H{"Error de formato": errors})
			return
		}

		log.Println("Error: json mal formado: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"Error json mal formado": errors})
		return
	}

	services := components.Services{
		Name: reqServ.Name,
	}

	if err := h.Services.CreatedServices(c.Request.Context(), &services); err != nil {
		log.Println("Error interno en el srv: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error interno:": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Solicitud procesada con exito")
}

func (h *ServicesHandler) ListServices(c *gin.Context) {
	services, err := h.Services.ListServices(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error interno": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)

}
