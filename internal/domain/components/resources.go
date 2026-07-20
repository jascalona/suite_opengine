package components

import "context"

type Resources struct {
	Id            int    `json:"id" db:"id"`
	ServiceId     int    `json:"service_id" db:"service_id"`
	EnvironmentId int    `json:"environment_id" db:"environment_id"`
	Name          string `json:"name" db:"name"`
	CreatedAt     string `json:"created_at" db:"created_at"`
}

type ValidationResources struct {
	ServiceId     int    `json:"service_id" binding:"required"`
	EnvironmentId int    `json:"environment_id" binding:"required"`
	Name          string `json:"name" binding:"required"`
}

type InterfaceResources interface {
	CreatedResources(ctx context.Context, resource *Resources) error
	ResourcesByService(ctx context.Context, id int) ([]*Resources, error)
}
