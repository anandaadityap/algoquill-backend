package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		var message string
		switch e.ActualTag() {
		case "required":
			message = fmt.Sprintf("field %v must be %v", e.Field(), e.Tag())
		case "email":
			message = fmt.Sprintf("field %v must be %v", e.Field(), e.Tag())
		case "min":
			message = fmt.Sprintf("field %v minimum %v character", e.Field(), e.Param())
		case "max":
			message = fmt.Sprintf("field %v maximum %v character", e.Field(), e.Param())
		default:
			message = fmt.Sprintf("Field '%s' is invalid", e.Field())
		}
		errors = append(errors, message)
	}
	return errors
}
