package usecase

import (
	"context"
	"fmt"
	"log"
	"opengine/v2/internal/domain/components"
)

type ServicesServ interface {
	CreatedServices(ctx context.Context, serv *components.Services) error
	ListServices(ctx context.Context) ([]*components.Services, error)
}

type ServicesImpl struct {
	Repo components.InterfaceService
}

func NewServicesServ(repo components.InterfaceService) ServicesServ {
	return &ServicesImpl{
		Repo: repo,
	}
}

func (s *ServicesImpl) CreatedServices(ctx context.Context, serv *components.Services) error {

	err := s.Repo.CreatedServices(ctx, serv)
	if err != nil {
		log.Printf("Error al procesar la peticion: %v", err.Error())
		return fmt.Errorf("Error al procesar la peticion %v", err.Error())
	}

	return nil
}

func (s *ServicesImpl) ListServices(ctx context.Context) ([]*components.Services, error) {
	services, err := s.Repo.ListServices(ctx)

	if err != nil {
		log.Println("Error al procesar la peticion ", err.Error())
		return nil, fmt.Errorf("Error al procesar la peticion %v", err.Error())
	}

	return services, nil
}
