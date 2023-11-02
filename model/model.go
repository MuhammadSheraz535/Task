package model

import "github.com/go-playground/validator/v10"

type Employee struct {
	CommonModel
	Name       string `json:"name" validate:"required,min=3"`
	Position   string `json:"position" validate:"required,min=2"`
	Department string `json:"department" validate:"required,min=2"`
	Email      string `json:"email" validate:"required,email"`
}

func (a *Employee) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
