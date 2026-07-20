package usecase

import (
	"context"
	"fmt"
	"log"
	"opengine/v2/internal/domain/components"
)

type SubResourcesServ interface {
	CreatedSubResources(ctx context.Context, subr *components.SubResources) error
	SubResourcesByResources(ctx context.Context, id int) ([]*components.SubResources, error)
}

type SubResourcesImpl struct {
	Repo components.InterfaceSubResources
}

func NewSubResourcesServ(repo components.InterfaceSubResources) SubResourcesServ {
	return &SubResourcesImpl{
		Repo: repo,
	}
}

func (s *SubResourcesImpl) CreatedSubResources(ctx context.Context, subr *components.SubResources) error {
	err := s.Repo.CreatedSubResources(ctx, subr)
	if err != nil {
		log.Println("Error al procesar la peticion: ", err.Error())
		return fmt.Errorf("Error al procesar la peticion: %v", err.Error())
	}
	return nil
}

func (s *SubResourcesImpl) SubResourcesByResources(ctx context.Context, id int) ([]*components.SubResources, error) {
	if id <= 0 {
		return nil, fmt.Errorf("El Servicio espera un Id de tipo entero positivo ")
	}

	list_subr, err := s.Repo.SubResourcesByResources(ctx, id)

	if err != nil {
		log.Println("Error service: no se pudo obtener los registros asociados a este servicio: ", id)
		return nil, fmt.Errorf("No se pudieron obtener los registros asociados al servicio: (", id, ")", err.Error())
	}
	return list_subr, nil
}
