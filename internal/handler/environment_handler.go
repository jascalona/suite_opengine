package handler

import "opengine/v2/internal/usecase"

type EnvironmentHandler struct {
	Services usecase.EnvironmentServ
}

func NewEnvironmentHandler(s usecase.EnvironmentServ) *EnvironmentHandler {

}
