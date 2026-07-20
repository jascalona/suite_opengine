package usecase

import (
	"context"
	"fmt"
	"log"
	"opengine/v2/internal/domain/components"
)

type EndpontsServ interface {
	CreatedEndpoints(ctx context.Context, endp *components.Endpoints) error
}

type EndpointsImpl struct {
	Repo components.InterfaceEndpoints
}

func NewEndpointsServ(repo components.InterfaceEndpoints) EndpontsServ {
	return &EndpointsImpl{
		Repo: repo,
	}
}

func (s *EndpointsImpl) CreatedEndpoints(ctx context.Context, endp *components.Endpoints) error {

	err := s.Repo.CreatedEndpoints(ctx, endp)
	if err != nil {
		log.Printf("Error al procesar la peticion: %v", err.Error())
		return fmt.Errorf("Error al procesar la peticion: %v", err.Error())
	}

	return nil

}
