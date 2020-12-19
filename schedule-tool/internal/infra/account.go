package infra

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

func NewAccountRepository(conn *gorm.DB) *AccountsRepository{
	return &AccountsRepository{conn: conn}
}

type AccountsRepository struct{
	conn *gorm.DB
}

func (a AccountsRepository) Create(ctx context.Context, r *Account) error {
	result := a.conn.Create(&r)

	if result.Error != nil   {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// go get -u gorm.io/gorm github.com/gin-gonic/gin