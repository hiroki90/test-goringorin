package internal

import (
	"context"
	"database/sql"
	"fmt"
)

func NewAccountsRepository(conn *sql.DB) *AccountsRepository {
	return &AccountsRepository{conn: conn}
}

type AccountsRepository struct {
	conn *sql.DB
}

func (a AccountsRepository) Create(ctx context.Context, r *Account) error {
	ins, err := a.conn.Prepare("INSERT INTO accounts(name,age) VALUES(?,?);")
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	_, err = ins.Exec(r.Name, r.Age)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (a AccountsRepository) FindByID(ctx context.Context, id string) (*Account, error) {
	row := a.conn.QueryRow("SELECT id,name,age FROM accounts;")

	var result Account
	err := row.Scan(&result.ID, &result.Name, &result.Age)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &result, nil
}
