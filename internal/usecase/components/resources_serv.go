package usecase

import (
	"context"
	"fmt"
	"log"
	"opengine/v2/internal/domain/components"
)

type ResourcesServ interface {
	CreatedResources(ctx context.Context, resource *components.Resources) error
	ResourcesByService(ctx context.Context, id int) ([]*components.Resources, error)
}

type ResourcesImpl struct {
	Repo components.InterfaceResources
}

func NewRosourcesServ(repo components.InterfaceResources) ResourcesServ {
	return &ResourcesImpl{
		Repo: repo,
	}
}

func (s *ResourcesImpl) CreatedResources(ctx context.Context, resource *components.Resources) error {
	err := s.Repo.CreatedResources(ctx, resource)
	if err != nil {
		log.Printf("Error al procesar la peticion %v", err.Error())
		return fmt.Errorf("Error al procesar la peticion %v ", err.Error())
	}

	return nil
}

func (s *ResourcesImpl) ResourcesByService(ctx context.Context, id int) ([]*components.Resources, error) {
	if id <= 0 {
		return nil, fmt.Errorf("El servicio espera un id de tipo entero y positivo")
	}

	list_resources, err := s.Repo.ResourcesByService(ctx, id)

	if err != nil {
		log.Println("Id enviado: ", id)
		log.Println("Error service: no se pudo obtener los registros asociados a este servicio: ", id)
		return nil, fmt.Errorf("No se pudieron obtener los registros asociados al servicio: (", id, ")", err.Error())
	}

	return list_resources, nil
}
