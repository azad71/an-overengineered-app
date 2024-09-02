package httpResponse

import (
	"an-overengineered-app/internal/helpers"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type CustomError struct {
	Message    string            `json:"message"`
	StatusCode int               `json:"statusCode"`
	Errors     map[string]string `json:"errors,omitempty"`
	RequestID  string            `json:"requestId,omitempty"`
}

func (e CustomError) Error() string {
	return "CustomError"
}

func BadRequestError(message string) CustomError {
	if message == "" {
		message = "Something went wrong"
	}

	return CustomError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func InternerServerError(message string) CustomError {
	if message == "" {
		message = "Something went wrong"
	}

	return CustomError{
		Message:    message,
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
		StatusCode: http.StatusConflict,
	}
}
