package model

type Event struct {
	ID       string
	Name     string
	Accounts Accounts
	Owner    Account	//スケジュールの幹事のアカウント
}
