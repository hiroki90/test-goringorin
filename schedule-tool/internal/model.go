package internal

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Account struct {
	ID   string
	Name string
	Age  int32
}

func (a Account) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Age, validation.Required),
	)
}
