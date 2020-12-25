package infra

import (
	"context"
	"fmt"
	"github.com/hiroki90/goringorin/schedule-tool/internal/model"
	"gorm.io/gorm"
)

func NewScheduleRepository(conn *gorm.DB) *SchedulesRepository {
	return &SchedulesRepository{conn: conn}
}

type SchedulesRepository struct {
	conn *gorm.DB
}

func (a SchedulesRepository) Create(ctx context.Context, r *model.Schedule) error {
	result := a.conn.Create(&r)

	if result.Error != nil {
		return fmt.Errorf("%w", result.Error)
	}

	return nil
}

func (a SchedulesRepository) FindByID(ctx context.Context, id string) (*model.Schedule, error) {
	var schedule model.Schedule
	if err := a.conn.First(&schedule, id).Error; err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &schedule, nil
}

//TODO: FindByAccountID(ctx context.Context, accountID string) (*model.Schedules, error)

func (a SchedulesRepository) Update(ctx context.Context, r *model.Schedule) error {
	if err := a.conn.Updates(&r).Error; err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (a SchedulesRepository) Delete(ctx context.Context, id string) error {
	if err := a.conn.Delete(&model.Schedule{}, id).Error; err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
