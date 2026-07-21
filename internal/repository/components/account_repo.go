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
