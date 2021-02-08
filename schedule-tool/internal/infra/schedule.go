package infra

import (
	"context"
	"fmt"
	"github.com/hiroki90/goringorin/schedule-tool/internal/model"
	"gorm.io/gorm"
)
type Schedules interface{	// 引数と戻り値の型を定義
	Find(ctx context.Context, accountID string) (*model.Schedules, error)
	Create(ctx context.Context, r *model.Schedule) error
}
//	CreateSchedule(Block, State, AccountID) error
//  FindSchedule(AccountID) (Block, State, error)
// TODO: Schedule->account->eventの順にinfra作っていく(0129)

func NewSchedulesRepository(conn *gorm.DB) *SchedulesRepository {
	return &SchedulesRepository{conn: conn}
}

type SchedulesRepository struct {
	conn *gorm.DB
}


func (s SchedulesRepository) Find(ctx context.Context, accountID string) (*model.Schedule, error){
	var schedule model.Schedule
	return &schedule, nil
}

func (s SchedulesRepository) FindByAccountID(ctx context.Context, accountID string) (*model.Schedules, error) {
	var schedules model.Schedules
	if err := s.conn.Where("account_id = ?", accountID).Find(&schedules).Error; err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &schedules, nil
}

func (s SchedulesRepository) Create(ctx context.Context, r *model.Schedule) error {
	result := s.conn.Create(&r)
	if result.Error != nil {
		return fmt.Errorf("%w", result.Error)
	}
	return nil
}



//func (a SchedulesRepository) FindByID(ctx context.Context, id string) (*model.Schedule, error) {
//	var schedule model.Schedule
//	if err := a.conn.First(&schedule, id).Error; err != nil {
//		return nil, fmt.Errorf("%w", err)
//	}
//	return &schedule, nil
//}

//func (a SchedulesRepository) FindByBlockAndState(ctx context.Context, block int, state string) (*model.Accounts, error) {
//	var schedules model.Schedules
//	var accounts model.Accounts
//	//TODO: blockを範囲指定にする(block time型で範囲指定する) sqlのBETWEEN
//	//// BETWEEN
//	//db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
//	//// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
//	if err := a.conn.Where("block = ? AND state = ?", block, state).Find(&schedules).Error; err != nil {
//		return nil, fmt.Errorf("%w", err)
//	}
//	return &accounts, nil
//}

//TODO: (a SchedulesRepository) FindByAccountNameAndState(ctx context.Context, ) (*model.Schedules, error)

//func (a SchedulesRepository) Update(ctx context.Context, r *model.Schedule) error {
//	if err := a.conn.Updates(&r).Error; err != nil {
//		return fmt.Errorf("%w", err)
//	}
//	return nil
//}
//
//func (a SchedulesRepository) Delete(ctx context.Context, id string) error {
//	if err := a.conn.Delete(&model.Schedule{}, id).Error; err != nil {
//		return fmt.Errorf("%w", err)
//	}
//	return nil
//}
