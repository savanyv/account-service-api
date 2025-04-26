package utils

import "github.com/go-playground/validator/v10"

type CustomerValidator struct {
	validator *validator.Validate
}

func NewValidator() *CustomerValidator {
	return &CustomerValidator{
		validator: validator.New(),
	}
}

func (cv *CustomerValidator) Validate(i interface{}) error {
	if cv == nil || cv.validator == nil {
		return nil
	}

	return cv.validator.Struct(i)
}
