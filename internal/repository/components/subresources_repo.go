package components

import (
	"context"
	"database/sql"
	"log"
	"opengine/v2/internal/domain/components"
)

type SubResourcesRepo struct {
	DB *sql.DB
}

func NewSubResourcesRepo(db *sql.DB) components.InterfaceSubResources {
	return &SubResourcesRepo{DB: db}
}

func (r *SubResourcesRepo) CreatedSubResources(ctx context.Context, subr *components.SubResources) error {
	query := `
		INSERT INTO subresources(
			resource_id,
			name, 
			cod_sub_product,
			description)
		VALUES($1,$2,$3,$4)`

	_, err := r.DB.ExecContext(ctx, query,
		subr.ResourcesId,
		subr.Name,
		subr.CodSubProduct,
		subr.Description,
	)

	if err != nil {
		log.Println("Error al correr el query: ", err.Error())
		return err
	}

	return nil

}

func (r *SubResourcesRepo) SubResourcesByResources(ctx context.Context, id int) ([]*components.SubResources, error) {

	query := `
		SELECT 
			id,
			resource_id,
			name,
			cod_sub_product,
			description,
			created_at
		FROM subresources WHERE resource_id = $1`

	rows, err := r.DB.QueryContext(ctx, query, id)

	if err != nil {
		log.Println("Error al correr el query")
		return nil, err
	}

	lis_subresources := make([]*components.SubResources, 0)

	for rows.Next() {
		ls := &components.SubResources{}
		err := rows.Scan(
			&ls.Id,
			&ls.ResourcesId,
			&ls.Name,
			&ls.CodSubProduct,
			&ls.Description,
			&ls.CreatedAt,
		)

		if err != nil {
			log.Println("Error al correr el query")
			return nil, err
		}

		lis_subresources = append(lis_subresources, ls)
	}
	return lis_subresources, nil

}
