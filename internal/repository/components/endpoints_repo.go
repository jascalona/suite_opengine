package components

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"opengine/v2/internal/domain/components"
)

type EndpointsRepo struct {
	DB *sql.DB
}

func NewEndpointsRepo(db *sql.DB) components.InterfaceEndpoints {
	return &EndpointsRepo{DB: db}
}

func (r *EndpointsRepo) CreatedEndpoints(ctx context.Context, endp *components.Endpoints) error {
	query := `
		INSERT INTO endpoints_manager(
			subresource_id,
			path,
			method,
			default_headers)VALUES($1,$2,$3,$4)`

	default_headers, errJson := json.Marshal(endp.DefaultHeaders)
	if errJson != nil {
		return fmt.Errorf("error al convertir el Default Headers a json: %w, ", errJson)
	}

	_, err := r.DB.ExecContext(ctx, query,
		endp.SubResourcesId,
		endp.Path,
		endp.Method,
		default_headers,
	)
	if err != nil {
		log.Println("Error al correr el query: ", err.Error())
		return err
	}

	return nil
}
