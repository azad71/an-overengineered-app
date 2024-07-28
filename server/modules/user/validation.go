package users

import (
	users "an-overengineered-social-media-app/modules/user/models"
	"time"

	"github.com/go-playground/validator/v10"
)

type SignupBody struct {
	Username     string        `json:"username" binding:"required,alphanum,max=50"`
	Password     string        `json:"password" binding:"required,min=6,max=250"`
	Email        string        `json:"email" binding:"required,max=300,email"`
	FirstName    *string       `json:"firstName" binding:"omitempty,max=100,alpha"`
	LastName     *string       `json:"lastName" binding:"omitempty,max=100,alpha"`
	BirthDate    *string       `json:"birthDate" binding:"omitempty,validateBirthDate"`
	Gender       *users.Gender `json:"gender" binding:"omitempty,oneof= MALE FEMALE OTHERS"`
	Address      *string       `json:"address" binding:"omitempty,max=1000"`
	UserTimezone *string       `json:"userTimezone" binding:"required"`
	Description  *string       `json:"description" binding:"omitempty,max=5000"`
}

func ValidateBirthDate(data validator.FieldLevel) bool {
	date := data.Field().Interface().(string)
	value, err := time.Parse("2006-01-02", date)

	if err != nil {
		return false
	}

	minDate := time.Date(time.Now().Year()-12, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	return value.Before(minDate)
}
