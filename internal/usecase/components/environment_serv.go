package usecase

import (
	"context"
	"fmt"
	"log"
	"opengine/v2/internal/domain"
)

type EnvironmentServ interface {
	CreatedEnvironment(ctx context.Context, envi *domain.Environment) error
}

type EnvironmentImpl struct {
	Repo domain.InterfaceEnvironment
}

func NewEnvironmentServ(repo domain.InterfaceEnvironment) EnvironmentServ {
	return &EnvironmentImpl{
		Repo: repo,
	}
}

func (s *EnvironmentImpl) CreatedEnvironment(ctx context.Context, envi *domain.Environment) error {

	err := s.Repo.CreatedEnvironment(ctx, envi)
	if err != nil {
		log.Println("Error al procesar la peticion (created environments): ", err.Error())
		return fmt.Errorf("Error al procesar la peticion (created environments): %v", err.Error())
	}

	return nil
}
