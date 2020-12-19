package internal

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/hiroki90/goringorin/schedule-tool/internal/model"
)

func NewAccountsRepository(conn *sql.DB) *AccountsRepository {
	return &AccountsRepository{conn: conn}
}

type AccountsRepository struct {
	conn *sql.DB
}

func (a AccountsRepository) Create(ctx context.Context, r *model.Account) error {
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

func (a AccountsRepository) FindByID(ctx context.Context, id string) (*model.Account, error) {
	row := a.conn.QueryRow(fmt.Sprintf("SELECT id,name,age FROM accounts where id = '%s';", id))
	var result model.Account
	err := row.Scan(&result.ID, &result.Name, &result.Age)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &result, nil
}

func (a AccountsRepository) Update(ctx context.Context, r *model.Account) error {
	upd, preErr := a.conn.Prepare("UPDATE accounts SET name=?, age=? WHERE id=?")
	if preErr != nil {
		return preErr
	}

	if _, exeErr := upd.Exec(r.Name, r.Age, r.ID); exeErr != nil {
		return exeErr
	}

	return nil
}

func (a AccountsRepository) Delete(ctx context.Context, id string) error {
	dlt, preErr := a.conn.Prepare("DELETE FROM accounts WHERE id=?")
	if preErr != nil {
		return preErr
	}

	if _, exeErr := dlt.Exec(id); exeErr != nil {
		return exeErr
	}

	return nil
}
