package service

import (
	"context"
	"fmt"
	"github.com/hiroki90/goringorin/schedule-tool/internal/infra"
	"github.com/hiroki90/goringorin/schedule-tool/internal/model"
)

// web層，infra層でできなかったことを実装，何もなければそのまま引き渡すだけでok

type Schedules interface{
	Find(ctx context.Context, accountID string) (*model.Schedule, error)
	Create(ctx context.Context, r *model.Schedule) error
}


func NewSchedulesService(schedulesRepository *infra.SchedulesRepository) *SchedulesService{
	return &SchedulesService{schedulesRepository: schedulesRepository}
}

type SchedulesService struct {
	schedulesRepository *infra.SchedulesRepository
}


func (s SchedulesService)Find(ctx context.Context, accountID string) (*model.Schedule, error){
	schedule, err := s.schedulesRepository.Find(ctx, accountID)
	if err != nil{
		return nil, fmt.Errorf("%w", err)	// fmt=>エラー型とエラー関数は定義するのが慣例
	}
	return schedule, nil
}

func (s SchedulesService)Create(ctx context.Context, schedule *model.Schedule) error{
	err := s.schedulesRepository.Create(ctx, schedule)
	if err != nil{
		return fmt.Errorf("%w", err)
	}
	return nil
}