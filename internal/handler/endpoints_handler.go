package handler

import (
	"log"
	"net/http"
	"opengine/v2/internal/domain/components"
	usecase "opengine/v2/internal/usecase/components"
	"opengine/v2/pkg/utils"

	"github.com/gin-gonic/gin"
)

type EndpointsHandler struct {
	Service usecase.EndpontsServ
}

func NewEndpointsHandler(s usecase.EndpontsServ) *EndpointsHandler {
	return &EndpointsHandler{Service: s}
}

func (h *EndpointsHandler) CreatedEndpoints(c *gin.Context) {

	var reqEndp components.ValidationEndpoints

	if err := c.ShouldBindJSON(&reqEndp); err != nil {
		errors := utils.GetValidationError(err)
		if errors != nil {
			log.Println("Error en la validacion del mensaje: ", err)
			c.JSON(http.StatusConflict, errors)
			return
		}

		log.Println("Error: json mal formado: ", err)
		c.JSON(http.StatusBadRequest, errors)
		return
	}

	endpoints := components.Endpoints{
		SubResourcesId: reqEndp.SubResourcesId,
		Path:           reqEndp.Path,
		Method:         reqEndp.Method,
		DefaultHeaders: reqEndp.DefaultHeaders,
	}

	if err := h.Service.CreatedEndpoints(c.Request.Context(), &endpoints); err != nil {
		log.Println("Error interno en el srv; ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error interno": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Solicitud procesada con exito")

}
