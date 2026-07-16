package domain

import "context"

type Environment struct {
	Id            int                    `json:"id" db:"id"`
	Name          string                 `json:"name" db:"name"`
	GlobalDomain  string                 `json:"global_domain" db:"global_domain"`
	GlobalHeaders map[string]string      `json:"global_headers" db:"global_headers"`
	GlobalAuth    map[string]interface{} `json:"global_auth" db:"global_auth"`
	Created_at    string                 `json:"created_at" db:"created_at"`
}

type ValidationEnvironment struct {
	Name          string                 `json:"name" binding:"required,min=3,max=225"`
	GlobalDomain  string                 `json:"global_domain" binding:"required,min=3,max=225"`
	GlobalHeaders map[string]string      `json:"global_headers" binding:"required,dive,required"`
	GlobalAuth    map[string]interface{} `json:"global_auth" binding:"required,dive,required"`
}

type InterfaceEnvironment interface {
	CreatedEnvironment(ctx context.Context, envi *Environment) error
}
