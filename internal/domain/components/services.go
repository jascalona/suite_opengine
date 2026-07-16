package components

import "context"

type Services struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Created_at string `json:"created_at" db:"created_at"`
}

type ValidateServices struct {
	Name string `json:"name" binding:"required,min=3"`
}

type InterfaceService interface {
	CreatedServices(ctx context.Context, serv *Services) error
	ListServices(ctx context.Context) ([]*Services, error)
}
