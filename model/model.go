package model

type Employee struct {
	CommonModel
	Name       string `json:"name" validate:"required"`
	Position   string `json:"position" validate:"required"`
	Department string `json:"department" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
}
