package users

import (
	users "an-overengineered-app/modules/user/models"
)

type SignupBody struct {
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

type VerifyOTPBody struct {
	Otp   string `json:"otp" binding:"required,size=6"`
	Email string `json:"email" binding:"required,max=300,email"`
}
