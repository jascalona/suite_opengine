package components

import (
	"context"
	"database/sql"
	"log"
	"opengine/v2/internal/domain"
)

type EnvironmentRepo struct {
	DB *sql.DB
}

func NewEnvironmentRepo(db *sql.DB) domain.InterfaceEnvironment {
	return &EnvironmentRepo{DB: db}
}

func (r *EnvironmentRepo) CreatedEnvironment(ctx context.Context, envi *domain.Environment) error {
	query := `
		INSERT INTO environments(
			name,
			global_domain,
			global_headers,
			global_auth)VALUES($1,$2,$3,$4)`

	_, err := r.DB.ExecContext(ctx, query,
		envi.Name,
		envi.GlobalDomain,
		envi.GlobalHeaders,
		envi.GlobalAuth,
	)

	if err != nil {
		log.Println("Error al correr el query: ", err)
		return err
	}

	return nil
}
