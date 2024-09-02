package httpResponse

import (
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestBadRequestError(t *testing.T) {

	t.Run("Should return provided error message", func(t *testing.T) {
		err := BadRequestError("Invalid request body")
		assert.Equal(t, "Invalid request body", err.Message)
	})

	t.Run("Should return http 404", func(t *testing.T) {
		err := BadRequestError("")
		assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	})

	t.Run("Should return default error message if empty string provided", func(t *testing.T) {
		err := BadRequestError("")
		assert.Equal(t, "Something went wrong", err.Message)
	})
}

func TestInternalServerError(t *testing.T) {

	t.Run("Should return provided error message", func(t *testing.T) {
		err := InternerServerError("Failed to save user")
		assert.Equal(t, "Failed to save user", err.Message)
	})

	t.Run("Should return http 500", func(t *testing.T) {
		err := InternerServerError("")
		assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	})

	t.Run("Should return default error message if empty string provided", func(t *testing.T) {
		err := InternerServerError("")
		assert.Equal(t, "Something went wrong", err.Message)
	})
}

func prepareValidation(value string) validator.ValidationErrors {
	validate := validator.New()

	type testStruct struct {
		Name string `validate:"required"`
	}

	testObj := testStruct{Name: value}
	return validate.Struct(testObj).(validator.ValidationErrors)

}

func TestValidationError(t *testing.T) {

	t.Run("Should return formatted validation errors", func(t *testing.T) {

		unformattedErrors := prepareValidation("")
		err := ValidationError("Please provide name", unformattedErrors)

		expectedErrors := map[string]string{
			"name": "Name is required",
		}

		assert.Equal(t, "Please provide name", err.Message)
		assert.Equal(t, http.StatusUnprocessableEntity, err.StatusCode)
		assert.NotNil(t, err.Errors)
		assert.Equal(t, expectedErrors, err.Errors)
	})

	t.Run("Should return defautl validation error message if empty msg string given", func(t *testing.T) {

		unformattedErrors := prepareValidation("")
		err := ValidationError("", unformattedErrors)

		assert.Equal(t, "Request validation failed", err.Message)

	})
}

func TestConflictError(t *testing.T) {
	errorObj := map[string]string{
		"email": "User with this email already exists",
	}
	t.Run("Should return http 409", func(t *testing.T) {
		err := ConflictError("Email already exists", errorObj)

		assert.Equal(t, http.StatusConflict, err.StatusCode)
	})

	t.Run("Should return provided error object", func(t *testing.T) {
		err := ConflictError("Email already exists", errorObj)
		expectedErrorObj := errorObj

		assert.Equal(t, expectedErrorObj, err.Errors)
	})

	t.Run("Should return provided error message", func(t *testing.T) {
		err := ConflictError("Email already exists", errorObj)

		assert.Equal(t, "Email already exists", err.Message)
	})

	t.Run("Should return default error message if empty error msg string given", func(t *testing.T) {
		err := ConflictError("", errorObj)

		assert.Equal(t, "Request failed due to resource conflict", err.Message)
	})
}

func TestCustomError_Error(t *testing.T) {
	err := CustomError{}
	assert.Equal(t, "CustomError", err.Error())
}
