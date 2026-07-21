package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"opengine/v2/internal/domain/components"
	"slices"
)

type AccountServ interface {
	CreatedAccount(ctx context.Context, account *components.Account) error
	ListAccount(ctx context.Context) ([]*components.Account, error)
}

type AccountImpl struct {
	Repo components.InterfaceAccount
}

func NewAccountServ(repo components.InterfaceAccount) AccountServ {
	return &AccountImpl{
		Repo: repo,
	}
}

var ErrInvalidBankPrefix = errors.New("El número de cuenta no coincide con el código bancario del agente")
var ErrInvalidDocType = errors.New("El tipo de documento es invalido")

func (s *AccountImpl) CreatedAccount(ctx context.Context, account *components.Account) error {

	// APLICACION REGLAS DE NEGOCIO
	account_agent := account.Agent
	account_cnta := account.Cnta[:4]

	if account_cnta != account_agent {
		return ErrInvalidBankPrefix
	}

	tp_document := string(account.DocumentsId[0])
	validTypes := []string{"J", "V", "E", "R", "G", "C", "P"}

	if !slices.Contains(validTypes, tp_document) {
		return fmt.Errorf("%w: '%s'. Tipos admitidos: %v", ErrInvalidDocType, tp_document, validTypes)
	}

	err := s.Repo.CreatedAccount(ctx, account)
	if err != nil {
		log.Println("Error al procesar la peticion: ", err.Error())
		return fmt.Errorf("Error al procesar la peticion: %v", "Revice la trazabilidad de la misma")
	}
	return nil
}

func (s *AccountImpl) ListAccount(ctx context.Context) ([]*components.Account, error) {

	account_lis, err := s.Repo.ListAccount(ctx)
	if err != nil {
		log.Println("Error al procesar la peticion: ", err.Error())
		return nil, fmt.Errorf("Error al procesar la peticion %v", "Revice la trazabilidad de la misma")
	}

	return account_lis, nil
}
