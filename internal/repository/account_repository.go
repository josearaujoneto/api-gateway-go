package repository

import (
	"database/sql"

	"github.com/josearaujoneto/api-gateway-go/internal/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (repository *AccountRepository) Create(account *domain.Account) error {
	statement, err := repository.db.Prepare(`
	INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(
		account.ID,
		account.Name,
		account.Email,
		account.APIKey,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
