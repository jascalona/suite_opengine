package components

import (
	"context"
	"database/sql"
	"log"
	"opengine/v2/internal/domain/components"
)

type ServicesRepo struct {
	DB *sql.DB
}

func NewServicesRepo(db *sql.DB) components.InterfaceService {
	return &ServicesRepo{DB: db}
}

func (r *ServicesRepo) CreatedServices(ctx context.Context, serv *components.Services) error {

	query := `INSERT INTO services(name)VALUES($1)`

	_, err := r.DB.ExecContext(ctx, query, serv.Name)

	if err != nil {
		log.Println("Error al correr el query: ", err.Error())
		return err
	}

	return nil
}

func (r *ServicesRepo) ListServices(ctx context.Context) ([]*components.Services, error) {
	query := `SELECT id, name, created_at FROM services ORDER BY created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query)

	if err != nil {
		log.Println("Error al correr el query")
		return nil, err
	}

	services := make([]*components.Services, 0)

	for rows.Next() {
		ls := &components.Services{}
		err := rows.Scan(
			&ls.Id,
			&ls.Name,
			&ls.Created_at,
		)

		if err != nil {
			log.Println("Error al aplicar el scanner ", err.Error())
			return nil, err
		}

		services = append(services, ls)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return services, nil

}
