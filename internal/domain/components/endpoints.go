package components

import "context"

type Endpoints struct {
	Id             int                `json:"id" db:"id"`
	SubResourcesId int                `json:"subresource_id" db:"subresource_id"`
	Path           string             `json:"path" db:"path"`
	Method         string             `json:"method" db:"method"`
	DefaultHeaders *map[string]string `json:"default_headers" db:"default_headers"`
	CreatedAt      string             `json:"created_at" db:"created_at"`
}

type ValidationEndpoints struct {
	SubResourcesId int                `json:"subresource_id" binding:"required"`
	Path           string             `json:"path" binding:"required"`
	Method         string             `json:"method" binding:"required"`
	DefaultHeaders *map[string]string `json:"default_headers" binding:"omitempty"`
}

type InterfaceEndpoints interface {
	CreatedEndpoints(ctx context.Context, endp *Endpoints) error
}
