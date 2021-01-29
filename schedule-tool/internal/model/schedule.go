package model

import "time"

type Schedule struct {
	ID        string
	Block     time.Time
	State     State
	AccountID string
}

type State int // enum（列挙型）
const (
	Unknown State = iota
	OK
	NG
)

type Schedules []Schedule

//
