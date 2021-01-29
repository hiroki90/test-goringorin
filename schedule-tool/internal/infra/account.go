package infra

import (
	"context"
	"fmt"
	"github.com/hiroki90/goringorin/schedule-tool/internal/model"
	"gorm.io/gorm"
)

//  CreateAccount(Name) error
//  UpdateAccount(ID, Schedules) error
//	FindByID(ID) (model.Account, error)

func NewAccountRepository(conn *gorm.DB) *AccountsRepository {
	return &AccountsRepository{conn: conn}
}

type AccountsRepository struct {
	conn *gorm.DB
}

func (a AccountsRepository) Create(ctx context.Context, r *model.Account) error {
	result := a.conn.Create(&r)

	if result.Error != nil {
		return fmt.Errorf("%w", result.Error)
	}

	return nil
}

func (a AccountsRepository) FindByID(ctx context.Context, id string) (*model.Account, error) {
	var account model.Account
	if err := a.conn.First(&account, id).Error; err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &account, nil
}

func (a AccountsRepository) Update(ctx context.Context, r *model.Account) error {
	if err := a.conn.Updates(&r).Error; err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (a AccountsRepository) Delete(ctx context.Context, id string) error {
	if err := a.conn.Delete(&model.Account{}, id).Error; err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
