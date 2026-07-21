package components

import (
	"database/sql"
	"log"
	"opengine/v2/internal/domain/components"

	"golang.org/x/net/context"
)

type AccountRepo struct {
	DB *sql.DB
}

func NewAccountRepo(db *sql.DB) components.InterfaceAccount {
	return &AccountRepo{DB: db}
}

func (r *AccountRepo) CreatedAccount(ctx context.Context, account *components.Account) error {

	query := `
		INSERT INTO account_list(
			environments_id,
			account_origin,
			name,
			document_id,
			agent,
			cnta,
			cele,
			is_active,
			collector,
			contract)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	_, err := r.DB.ExecContext(ctx, query,
		account.EnvironmentId,
		account.AccountOrigin,
		account.Name,
		account.DocumentsId,
		account.Agent,
		account.Cnta,
		account.Cele,
		account.IsActive,
		account.Collector,
		account.Contract,
	)

	if err != nil {
		log.Println("Error al correr el query: ", err.Error())
		return err
	}

	return nil
}

func (r *AccountRepo) ListAccount(ctx context.Context) ([]*components.Account, error) {
	query := `
		SELECT
			id,
			environments_id,
			account_origin,
			name,
			document_id,
			agent,
			cnta,
			cele,
			is_active,
			collector,
			contract,
			created_at
		FROM account_list ORDER BY created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query)

	if err != nil {
		log.Println("Error al correr el query: ", err.Error())
		return nil, err
	}

	defer rows.Close()

	account_list := make([]*components.Account, 0)

	for rows.Next() {
		ls := &components.Account{}
		err := rows.Scan(
			&ls.Id,
			&ls.EnvironmentId,
			&ls.AccountOrigin,
			&ls.Name,
			&ls.DocumentsId,
			&ls.Agent,
			&ls.Cnta,
			&ls.Cele,
			&ls.IsActive,
			&ls.Collector,
			&ls.Contract,
			&ls.CreatedAt,
		)

		if err != nil {
			log.Println("Error al aplicar el scanner", err.Error())
			return nil, err
		}

		account_list = append(account_list, ls)
	}

	return account_list, nil

}
