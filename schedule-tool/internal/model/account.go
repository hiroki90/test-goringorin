package model

//import validation "github.com/go-ozzo/ozzo-validation"

type Account struct {
	ID           string
	Name         string
	AllSchedules Schedules
}
type Accounts []Account

//func (a Account) Validate() error {
//	return validation.ValidateStruct(&a,
//		validation.Field(&a.Name, validation.Required),
//		validation.Field(&a.Age, validation.Required),
//	)
//}
