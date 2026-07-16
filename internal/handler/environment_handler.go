package handler

import (
	"log"
	"net/http"
	"opengine/v2/internal/domain"
	usecase "opengine/v2/internal/usecase/components"
	"opengine/v2/pkg/utils"

	"github.com/gin-gonic/gin"
)

type EnvironmentHandler struct {
	Services usecase.EnvironmentServ
}

func NewEnvironmentHandler(s usecase.EnvironmentServ) *EnvironmentHandler {
	return &EnvironmentHandler{Services: s}
}

func (h *EnvironmentHandler) CreatedEnvironment(c *gin.Context) {
	var reqEnvironment domain.ValidationEnvironment

	if err := c.ShouldBindJSON(&reqEnvironment); err != nil {
		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Println("Error en la validacion del mensaje: ", err)
			c.JSON(http.StatusConflict, gin.H{"Error de formato: ": errors})
			return
		}

		log.Println("Error: json mal formado: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"Error json mal formado ": errors})
		return
	}

	// validacion para los campos de tipo jsonb
	global_h := reqEnvironment.GlobalHeaders
	global_a := reqEnvironment.GlobalAuth

	envi := domain.Environment{
		Name:          reqEnvironment.Name,
		GlobalDomain:  reqEnvironment.GlobalDomain,
		GlobalHeaders: global_h,
		GlobalAuth:    global_a,
	}

	if err := h.Services.CreatedEnvironment(c.Request.Context(), &envi); err != nil {
		log.Println("Error interno en el srv: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error interno: ": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Solicitud procesada con exito")

}
