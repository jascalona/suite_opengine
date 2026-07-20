package components

import "context"

type SubResources struct {
	Id            int    `json:"id" db:"id"`
	ResourcesId   int    `json:"resource_id" db:"resource_id"`
	Name          string `json:"name" db:"name"`
	CodSubProduct string `json:"cod_sub_product" db:"cod_sub_product"`
	Description   string `json:"description" db:"description"`
	CreatedAt     string `json:"created_at" db:"created_at"`
}

type ValidationSubResources struct {
	ResourcesId   int    `json:"resource_id" binding:"required"`
	Name          string `json:"name" binding:"required,min=3"`
	CodSubProduct string `json:"cod_sub_product" binding:"required"`
	Description   string `json:"description" binding:"required"`
}

type InterfaceSubResources interface {
	CreatedSubResources(ctx context.Context, subr *SubResources) error
	SubResourcesByResources(ctx context.Context, id int) ([]*SubResources, error)
}
