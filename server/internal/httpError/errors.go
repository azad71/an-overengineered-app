package httpError

import (
	"an-overengineered-social-media-app/internal/helpers"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type CustomError struct {
	Message    string            `json:"message"`
	StatusCode int               `json:"statusCode"`
	Errors     map[string]string `json:"errors,omitempty"`
	Data       map[string]any    `json:"data,omitempty"`
}

func (e CustomError) Error() string {
	return "CustomError"
}

func BadRequestError(message string, err error) CustomError {
	if message == "" {
		message = "Something went wrong"
	}

	return CustomError{
		Message:    message,
		Errors:     nil,
		Data:       nil,
		StatusCode: http.StatusBadRequest,
	}
}

func InternerServerError(message string, err error) CustomError {
	if message == "" {
		message = "Something went wrong"
	}

	return CustomError{
		Message:    message,
		Errors:     nil,
		Data:       nil,
		StatusCode: http.StatusInternalServerError,
	}
}

func ValidationError(message string, errs validator.ValidationErrors) CustomError {
	if message == "" {
		message = "Request validation failed"
	}
	return CustomError{
		Message:    message,
		Errors:     helpers.FormatValidationError(errs),
		Data:       nil,
		StatusCode: http.StatusUnprocessableEntity,
	}

}

func ConflictError(message string, err map[string]string) CustomError {
	if message == "" {
		message = "Request failed due to resource conflict"
	}

	return CustomError{
		Message:    message,
		Errors:     err,
		Data:       nil,
		StatusCode: http.StatusConflict,
	}
}
