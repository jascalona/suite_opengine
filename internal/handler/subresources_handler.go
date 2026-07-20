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

type SubResourcesHandler struct {
	Services usecase.SubResourcesServ
}

func NewSubResourcesHandler(s usecase.SubResourcesServ) *SubResourcesHandler {
	return &SubResourcesHandler{Services: s}
}

func (h *SubResourcesHandler) CreatedSubResources(c *gin.Context) {
	var reqSubr components.ValidationSubResources

	if err := c.ShouldBindJSON(&reqSubr); err != nil {
		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Println("Error en la validacion del mensaje: ", err)
			c.JSON(http.StatusConflict, gin.H{"Error de formato": errors})
			return
		}

		log.Println("Error: json mal formado", err)
		c.JSON(http.StatusBadRequest, gin.H{"Error json mal formado: ": errors})
		return
	}

	subvalidation := components.SubResources{
		ResourcesId:   reqSubr.ResourcesId,
		Name:          reqSubr.Name,
		CodSubProduct: reqSubr.CodSubProduct,
		Description:   reqSubr.Description,
	}

	if err := h.Services.CreatedSubResources(c.Request.Context(), &subvalidation); err != nil {
		log.Println("Error interno en el srv: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error interno: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Solicitud procesada con exito")

}

func (h *SubResourcesHandler) SubResourcesByResources(c *gin.Context) {
	idStr := c.Query("resourceid")

	resource_id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": "El Id del servicio es requerido"})
		return
	}

	list_subr, err := h.Services.SubResourcesByResources(c.Request.Context(), resource_id)
	if err != nil {
		log.Printf("Error al obtener los registros asociados: %d: %v", resource_id, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error interno: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list_subr)
}
