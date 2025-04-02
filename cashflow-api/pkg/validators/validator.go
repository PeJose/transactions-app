package validators

import (
	"cashflow/pkg/interfaces"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() interfaces.Validator {
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) Validate(data interface{}) error {
	return v.validate.Struct(data)
}

func (v *Validator) Format(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var formattedErrors string
		for _, fieldError := range validationErrors {
			formattedErrors += fieldError.Field() + " is " + fieldError.Tag() + "; "
		}
		return formattedErrors
	}
	return err.Error()
}
