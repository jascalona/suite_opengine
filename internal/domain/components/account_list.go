package components

import (
	"context"

	"github.com/google/uuid"
)

type Account struct {
	Id            uuid.UUID `json:"id" db:"id"`
	EnvironmentId int       `json:"environments_id" db:"environments_id"`
	AccountOrigin string    `json:"account_origin" db:"account_origin"`
	Name          string    `json:"name" db:"name"`
	DocumentsId   string    `json:"document_id" db:"document_id"`
	Agent         string    `json:"agent" db:"agent"`
	Cnta          string    `json:"cnta" db:"cnta"`
	Cele          string    `json:"cele" db:"cele"`
	IsActive      bool      `json:"is_active" db:"is_active"`
	Collector     bool      `json:"collector" db:"collector"`
	Contract      *string   `json:"contract" db:"contract"`
	CreatedAt     string    `json:"created_at" db:"created_at"`
}

type ValidationAccount struct {
	EnvironmentId int    `json:"environments_id" binding:"required"`
	AccountOrigin string `json:"account_origin" binding:"required,min=3"`
	Name          string `json:"name" binding:"required,min=3"`
	DocumentsId   string `json:"document_id" binding:"required,min=6,max=15"`
	Agent         string `json:"agent" binding:"required,min=4,max=4"`
	Cnta          string `json:"cnta" binding:"required,numeric,len=20"`
	Cele          string `json:"cele" binding:"required,numeric,len=11"`
	IsActive      bool   `json:"is_active" binding:"required"`
	Collector     bool   `json:"collector" binding:"required"`
	Contract      string `json:"contract" binding:"omitempty"`
}

type InterfaceAccount interface {
	CreatedAccount(ctx context.Context, account *Account) error
	ListAccount(ctx context.Context) ([]*Account, error)
}
