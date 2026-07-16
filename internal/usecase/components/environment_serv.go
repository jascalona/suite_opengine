package usecase

import (
	"context"
	"fmt"
	"log"
	"opengine/v2/internal/domain/components"
)

type EnvironmentServ interface {
	CreatedEnvironment(ctx context.Context, envi *components.Environment) error
	ListEnvironment(ctx context.Context) ([]*components.Environment, error)
}

type EnvironmentImpl struct {
	Repo components.InterfaceEnvironment
}

func NewEnvironmentServ(repo components.InterfaceEnvironment) EnvironmentServ {
	return &EnvironmentImpl{
		Repo: repo,
	}
}

func (s *EnvironmentImpl) CreatedEnvironment(ctx context.Context, envi *components.Environment) error {

	err := s.Repo.CreatedEnvironment(ctx, envi)
	if err != nil {
		log.Println("Error al procesar la peticion (created environments): ", err.Error())
		return fmt.Errorf("Error al procesar la peticion (created environments): %v", err.Error())
	}

	return nil
}

func (s *EnvironmentImpl) ListEnvironment(ctx context.Context) ([]*components.Environment, error) {
	envi_list, err := s.Repo.ListEnvironment(ctx)
	if err != nil {
		log.Println("No se pudo retornar los registros: ", err.Error())
		return nil, fmt.Errorf("Error al procesar la solicitud")
	}

	return envi_list, nil
}
