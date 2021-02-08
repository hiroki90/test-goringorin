package model

//import validation "github.com/go-ozzo/ozzo-validation"

type Account struct {
	ID           string
	Name         string
	AllSchedules Schedules //アカウントに紐づくScheduleの集合
	AllEvents Events	// 追加
}
type Accounts []Account
type Events []Event	// 追加
//func (a Account) Validate() error {
//	return validation.ValidateStruct(&a,
//		validation.Field(&a.Name, validation.Required),
//		validation.Field(&a.Age, validation.Required),
//	)
//}
