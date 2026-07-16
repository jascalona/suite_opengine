package components

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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
            global_auth) VALUES($1,$2,$3,$4)`

	global_headers, errJson := json.Marshal(envi.GlobalHeaders)
	if errJson != nil {
		return fmt.Errorf("error al convertir GlobalHeaders a json: %w", errJson)
	}

	global_auth, errJsonii := json.Marshal(envi.GlobalAuth)
	if errJsonii != nil {
		return fmt.Errorf("error al convertir GlobalAuth a json: %w", errJsonii)
	}

	_, err := r.DB.ExecContext(ctx, query,
		envi.Name,
		envi.GlobalDomain,
		global_headers,
		global_auth,
	)

	if err != nil {
		log.Println("Error al correr el query: ", err)
		return err
	}

	return nil
}

func (r *EnvironmentRepo) ListEnvironment(ctx context.Context) ([]*domain.Environment, error) {
	query := `
        SELECT 
            id,
            name,
            global_domain,
            global_headers, 
            global_auth,
            created_at
        FROM environments`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Error al correr el query: ", err.Error())
		return nil, err
	}
	defer rows.Close()

	list_envi := make([]*domain.Environment, 0)

	for rows.Next() {
		lst := &domain.Environment{}

		var g_headers []byte
		var g_auth []byte

		err := rows.Scan(
			&lst.Id,
			&lst.Name,
			&lst.GlobalDomain,
			&g_headers,
			&g_auth,
			&lst.Created_at,
		)
		if err != nil {
			log.Println("Error al aplicar el scanner: ", err.Error())
			return nil, err
		}

		if len(g_headers) > 0 {
			if err := json.Unmarshal(g_headers, &lst.GlobalHeaders); err != nil {
				log.Printf("Error al deserializar headers: %v", err.Error())
				return nil, err
			}
		}

		if len(g_auth) > 0 {
			if err := json.Unmarshal(g_auth, &lst.GlobalAuth); err != nil {
				log.Printf("Error al deserializar auth: %v", err.Error())
				return nil, err
			}
		}

		list_envi = append(list_envi, lst)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list_envi, nil
}
