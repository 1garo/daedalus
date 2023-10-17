package types

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type XValidator struct {
	Validator *validator.Validate
}
type ErrorResponse struct {
	Tag         string `json:"description"     validate:"required"`
	FailedField string `json:"field,omitempty"`
}

// TODO: put this in another pkg
func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := v.Validator.Struct(data)
	if errs != nil {
		if _, ok := errs.(*validator.InvalidValidationError); ok {
			fmt.Println(errs)
			return nil
		}

		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = strings.ToLower(err.Field()) // Export struct field name
			elem.Tag = err.Tag()                            // Export struct tag

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
