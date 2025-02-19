package sharedutil

import (
	"github.com/go-playground/validator/v10"
)

type validate = validator.Validate

type Validator struct {
	*validate
}

func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

func (cv *Validator) Validate(t any) error {
	return cv.Struct(t)
}
