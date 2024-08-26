package helpers

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestFormatValidationError_InvalidValidationError(t *testing.T) {
	invalidErr := &validator.InvalidValidationError{}

	formattedErrors := FormatValidationError(invalidErr)
	_, exists := formattedErrors["validation_error"]
	assert.True(t, exists, "expected 'validation_error' key in the formatted errors")
}

func TestFormatValidationError_Cases(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("validateBirthDate", IsDateBefore)

	tests := []struct {
		name           string
		testData       interface{}
		expectedErrors map[string]string
	}{
		{
			name: "RequiredField",
			testData: struct {
				RequiredField string `validate:"required"`
			}{},
			expectedErrors: map[string]string{
				"requiredfield": "RequiredField is required",
			},
		},
		{
			name: "AlphaNumField",
			testData: struct {
				AlphaNumField string `validate:"alphanum"`
			}{
				AlphaNumField: "invalid!@#",
			},
			expectedErrors: map[string]string{
				"alphanumfield": "AlphaNumField must contain only alphanumeric characters",
			},
		},
		{
			name: "MaxField",
			testData: struct {
				MaxField string `validate:"max=5"`
			}{
				MaxField: "tooLongValue",
			},
			expectedErrors: map[string]string{
				"maxfield": "MaxField cannot be longer than 5 characters",
			},
		},
		{
			name: "MinField",
			testData: struct {
				MinField string `validate:"min=5"`
			}{
				MinField: "123",
			},
			expectedErrors: map[string]string{
				"minfield": "MinField must be at least 5 characters",
			},
		},
		{
			name: "EmailField",
			testData: struct {
				EmailField string `validate:"email"`
			}{
				EmailField: "invalid-email",
			},
			expectedErrors: map[string]string{
				"emailfield": "EmailField must be a valid email address",
			},
		},
		{
			name: "AlphaField",
			testData: struct {
				AlphaField string `validate:"alpha"`
			}{
				AlphaField: "alpha123",
			},
			expectedErrors: map[string]string{
				"alphafield": "AlphaField must contain only alphabetic characters",
			},
		},
		{
			name: "InvalidDateFormat",
			testData: struct {
				BirthDate string `validate:"validateBirthDate"`
			}{
				BirthDate: "invalid-date-format",
			},
			expectedErrors: map[string]string{
				"birthdate": "User must be at least 12 years old",
			},
		},
		{
			name: "UnsupportedFieldType",
			testData: struct {
				UnsupportedField int `validate:"boolean"`
			}{
				UnsupportedField: 0,
			},
			expectedErrors: map[string]string{
				"unsupportedfield": "UnsupportedField is invalid",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.testData)

			formattedErrors := FormatValidationError(err)

			assert.Equal(t, tt.expectedErrors, formattedErrors)
		})
	}
}
