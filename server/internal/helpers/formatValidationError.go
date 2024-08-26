package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	if _, ok := err.(*validator.InvalidValidationError); ok {
		errors["validation_error"] = err.Error()
		return errors
	}

	for _, err := range err.(validator.ValidationErrors) {
		var field string
		var msg string

		switch err.Tag() {
		case "required":
			msg = fmt.Sprintf("%s is required", err.Field())
		case "alphanum":
			msg = fmt.Sprintf("%s must contain only alphanumeric characters", err.Field())
		case "max":
			msg = fmt.Sprintf("%s cannot be longer than %s characters", err.Field(), err.Param())
		case "min":
			msg = fmt.Sprintf("%s must be at least %s characters", err.Field(), err.Param())
		case "email":
			msg = fmt.Sprintf("%s must be a valid email address", err.Field())
		case "alpha":
			msg = fmt.Sprintf("%s must contain only alphabetic characters", err.Field())
		case "validateBirthDate":
			msg = "User must be at least 12 years old"
		default:
			msg = fmt.Sprintf("%s is invalid", err.Field())
		}

		field = strings.ToLower(err.Field())
		errors[field] = msg
	}

	return errors
}
