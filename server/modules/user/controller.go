package users

import (
	"an-overengineered-social-media-app/pkg/config"
	"an-overengineered-social-media-app/pkg/helpers"
	"an-overengineered-social-media-app/pkg/httpError"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func SignupUser(ctx *gin.Context) {

	var userData SignupBody

	if err := ctx.ShouldBindJSON(&userData); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			ctx.Error(httpError.ValidationError("", errs))
			return
		}

		ctx.Error(httpError.BadRequestError("Failed to parse request data", nil))
		return
	}

	isNewUser, err := IsEmailAndUsernameUnique(userData.Email, userData.Username)

	if err != nil {
		ctx.Error(httpError.InternerServerError("", nil))
		return
	}

	if !isNewUser {
		ctx.Error(httpError.ConflictError("Email/Username already taken!", map[string]string{
			"email": "Email/Username already taken!",
		}))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)

	if err != nil {
		ctx.Error(httpError.InternerServerError("", nil))
		return
	}

	newUserData := BuildNewUserObj(userData, hashedPassword)

	tx := config.DBInstance.Begin()

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		tx.Rollback()
	// 	}
	// }()

	if err := tx.Error; err != nil {
		ctx.Error(httpError.InternerServerError("", nil))
		return
	}

	err = CreateUser(&newUserData, tx)

	if err != nil {
		tx.Rollback()
		ctx.Error(httpError.InternerServerError("Failed to create new user", nil))
		return
	}

	otp, err := helpers.GenerateOTP(8)

	if err != nil {
		ctx.Error(httpError.InternerServerError("", nil))
		return
	}

	otpData := BuildOTPObj(otp, newUserData, "SIGNUP")

	err = CreateOTP(&otpData, tx)

	if err != nil {
		tx.Rollback()
		ctx.Error(httpError.InternerServerError("", nil))
		return
	}

	err = SendSignupMail(newUserData.Email, otp)

	if err != nil {
		tx.Rollback()
		ctx.Error(httpError.InternerServerError("", nil))
		return
	}

	tx.Commit()

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "An OTP sent to your mail. Please verify your account to continue",
		"data": gin.H{
			"expiresIn":    5,
			"timeUnit":     "minutes",
			"maximumRetry": 3,
			"email":        newUserData.Email,
			"username":     newUserData.Username,
		},
	})
}
