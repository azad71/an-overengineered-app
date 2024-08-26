package helpers

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func IsDateBefore(data validator.FieldLevel) bool {
	date := data.Field().Interface().(string)
	value, err := time.Parse("2006-01-02", date)

	if err != nil {
		return false
	}

	minDate := time.Date(time.Now().Year()-12, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	return value.Before(minDate)
}
