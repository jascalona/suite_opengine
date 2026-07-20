package components

import (
	"database/sql"
	"log"
	"opengine/v2/internal/domain/components"

	"golang.org/x/net/context"
)

type ResourcesRepo struct {
	DB *sql.DB
}

func NewRosourcesRepo(db *sql.DB) components.InterfaceResources {
	return &ResourcesRepo{DB: db}
}

func (r *ResourcesRepo) CreatedResources(ctx context.Context, resource *components.Resources) error {

	query := `
		INSERT INTO resources(
			service_id,
			environment_id,
			name
		)VALUES($1,$2,$3)`

	_, err := r.DB.ExecContext(ctx, query,
		resource.ServiceId,
		resource.EnvironmentId,
		resource.Name,
	)

	if err != nil {
		log.Println("Error al correr el query ", err)
		return err
	}
	return nil

}

func (r *ResourcesRepo) ResourcesByService(ctx context.Context, id int) ([]*components.Resources, error) {

	query := `
		SELECT 
			id,
			service_id,
			environment_id,
			name,
			created_at
		FROM resources WHERE service_id = $1`

	rows, err := r.DB.QueryContext(ctx, query, id)

	if err != nil {
		log.Println("Error al correr el query ", err)
		return nil, err
	}

	defer rows.Close()

	var resources []*components.Resources

	for rows.Next() {
		ls := &components.Resources{}
		err := rows.Scan(
			&ls.Id,
			&ls.ServiceId,
			&ls.EnvironmentId,
			&ls.Name,
			&ls.CreatedAt,
		)
		if err != nil {
			log.Println("Error al aplicar el scanner ", err.Error())
			return nil, err
		}
		resources = append(resources, ls)
	}

	return resources, nil

}
