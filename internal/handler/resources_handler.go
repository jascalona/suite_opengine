package handler

import (
	"log"
	"net/http"
	"opengine/v2/internal/domain/components"
	usecase "opengine/v2/internal/usecase/components"
	"opengine/v2/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResourcesHandler struct {
	Services usecase.ResourcesServ
}

func NewRosourcesHandler(s usecase.ResourcesServ) *ResourcesHandler {
	return &ResourcesHandler{Services: s}
}

func (h *ResourcesHandler) CreatedResources(c *gin.Context) {
	var reqResource components.ValidationResources

	if err := c.ShouldBindJSON(&reqResource); err != nil {
		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Println("Error en la validacion del mensaje: ", err)
			c.JSON(http.StatusConflict, errors)
			return
		}

		log.Println("Error: json mal formado: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"Error: json mal formado": errors})
		return
	}

	resources := components.Resources{
		ServiceId:     reqResource.ServiceId,
		EnvironmentId: reqResource.EnvironmentId,
		Name:          reqResource.Name,
	}

	if err := h.Services.CreatedResources(c.Request.Context(), &resources); err != nil {
		log.Println("Error interno en el srv: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error interno: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Solicitud procesada con exito")

}

func (h *ResourcesHandler) ResourcesByService(c *gin.Context) {
	idStr := c.Query("servid")

	service_id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": "El Id del servicio es requerido"})
		return
	}

	list_resources, err := h.Services.ResourcesByService(c.Request.Context(), service_id)
	if err != nil {
		log.Printf("Error al obtener los registros asociados: %d: %v", service_id, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error interno: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list_resources)

}
