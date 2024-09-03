package users

import (
	"an-overengineered-app/internal/mailer"
	users "an-overengineered-app/modules/user/models"
	"context"
	"time"
)

func BuildNewUserObj(userData SignupBody, hashedPassword []byte) users.User {
	birthDate, _ := time.Parse("2006-01-02", *userData.BirthDate)

	return users.User{
		Username:     userData.Username,
		Password:     string(hashedPassword),
		Email:        userData.Email,
		FirstName:    userData.FirstName,
		LastName:     userData.LastName,
		BirthDate:    &birthDate,
		Gender:       userData.Gender,
		Address:      userData.Address,
		UserTimezone: userData.UserTimezone,
		Description:  userData.Description,
	}
}

func SendSignupMail(ctx context.Context, email string, otp string) error {
	mailContent := mailer.GetSignupContent(ctx, otp)
	return mailer.SendMail(ctx, email, []byte(mailContent), "auth")
}

func BuildOTPObj(otp string, userData users.User, otpType string) users.OtpCodes {
	return users.OtpCodes{
		Otp:        otp,
		Username:   userData.Username,
		Email:      userData.Email,
		OtpType:    "SIGNUP",
		RetryCount: 0,
	}
}
