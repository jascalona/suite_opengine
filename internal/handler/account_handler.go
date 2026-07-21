package handler

import (
	"log"
	"net/http"
	"opengine/v2/internal/domain/components"
	usecase "opengine/v2/internal/usecase/components"
	"opengine/v2/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	Service usecase.AccountServ
}

func NewAccountHandler(s usecase.AccountServ) *AccountHandler {
	return &AccountHandler{Service: s}
}

func (h *AccountHandler) CreatedAccount(c *gin.Context) {
	var reqAccount components.ValidationAccount

	if err := c.ShouldBindJSON(&reqAccount); err != nil {

		errors := utils.GetValidationError(err)

		if errors != nil {
			log.Println("Error service: en la validacion del mensaje: ", err.Error())
			c.JSON(http.StatusConflict, errors)
			return
		}

		log.Println("Error: json mal formado", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Error: json mal formado": errors})
	}

	toNull := func(v string) *string {
		if v == "" {
			return nil
		}
		return &v
	}

	vali_account := components.Account{
		EnvironmentId: reqAccount.EnvironmentId,
		AccountOrigin: reqAccount.AccountOrigin,
		Name:          reqAccount.Name,
		DocumentsId:   reqAccount.DocumentsId,
		Agent:         reqAccount.Agent,
		Cnta:          reqAccount.Cnta,
		Cele:          reqAccount.Cele,
		IsActive:      reqAccount.IsActive,
		Collector:     reqAccount.Collector,
		Contract:      toNull(reqAccount.Contract),
	}

	if err := h.Service.CreatedAccount(c.Request.Context(), &vali_account); err != nil {
		log.Println("Error interno en el srv: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"Error interno": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Solicitud procesada con exito")

}

func (h *AccountHandler) ListAccount(c *gin.Context) {
	account_list, err := h.Service.ListAccount(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error interno:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account_list)
}
