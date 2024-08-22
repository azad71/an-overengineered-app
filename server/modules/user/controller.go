package users

import (
	"an-overengineered-social-media-app/internal/config"
	"an-overengineered-social-media-app/internal/helpers"
	"an-overengineered-social-media-app/internal/httpResponse"
	"an-overengineered-social-media-app/internal/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func SignupUser(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()

	logger.PrintInfo(reqCtx, "SignupUser controller method invoked...", nil)

	var userData SignupBody

	if err := ctx.ShouldBindJSON(&userData); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			logger.PrintErrorWithStack(reqCtx, "Failed to validate req body", err)
			ctx.Error(httpResponse.ValidationError("", errs))
			return
		}

		logger.PrintErrorWithStack(reqCtx, "Failed to parse req body", err)

		ctx.Error(httpResponse.BadRequestError("Failed to parse request data", nil))
		return
	}

	logger.PrintInfo(reqCtx, "Validated request body", userData)

	isNewUser, err := IsEmailAndUsernameUnique(userData.Email, userData.Username, reqCtx)

	if err != nil {
		ctx.Error(httpResponse.InternerServerError("", nil))
		return
	}

	if !isNewUser {
		logger.PrintError(reqCtx, "User with provided email or username already exists")
		ctx.Error(httpResponse.ConflictError("Email/Username already taken!", map[string]string{
			"email": "Email/Username already taken!",
		}))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)

	if err != nil {
		logger.PrintErrorWithStack(reqCtx, "Hashing password failed", err)
		ctx.Error(httpResponse.InternerServerError("", nil))
		return
	}

	newUserData := BuildNewUserObj(userData, hashedPassword)

	logger.PrintInfo(reqCtx, "Constructed new user object", newUserData)

	tx := config.DBInstance.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		logger.PrintErrorWithStack(reqCtx, "Initiating db transaction failed", err)
		ctx.Error(httpResponse.InternerServerError("", nil))
		return
	}

	err = CreateUser(&newUserData, tx, reqCtx)

	logger.PrintInfo(reqCtx, "User object after inserting into db", newUserData)

	if err != nil {
		logger.PrintErrorWithStack(reqCtx, "Failed insert new user in db", err)
		tx.Rollback()
		ctx.Error(httpResponse.InternerServerError("Failed to create new user", nil))
		return
	}

	otp, err := helpers.GenerateOTP(6)

	logger.PrintInfo(reqCtx, "Generated OTP", map[string]string{"otp": otp})

	if err != nil {
		logger.PrintErrorWithStack(reqCtx, "Failed to generate OTP", err)
		ctx.Error(httpResponse.InternerServerError("", nil))
		return
	}

	otpData := BuildOTPObj(otp, newUserData, "SIGNUP")

	logger.PrintInfo(reqCtx, "Generated OTP object to save into db", otpData)

	err = CreateOTP(&otpData, tx, reqCtx)
	logger.PrintInfo(reqCtx, "OTP object after inserting into db", otpData)

	if err != nil {
		logger.PrintErrorWithStack(reqCtx, "Failed to insert otp in db", err)
		tx.Rollback()
		ctx.Error(httpResponse.InternerServerError("", nil))
		return
	}

	err = SendSignupMail(newUserData.Email, otp)

	if err != nil {
		logger.PrintErrorWithStack(reqCtx, "Failed to new signup mail with otp to user", err)

		tx.Rollback()
		ctx.Error(httpResponse.InternerServerError("", nil))
		return
	}

	tx.Commit()

	resData := gin.H{
		"expiresIn":    5,
		"timeUnit":     "minutes",
		"maximumRetry": 3,
		"email":        newUserData.Email,
		"username":     newUserData.Username,
	}

	httpResponse.Created(ctx, "An OTP sent to your mail. Please verify your account to continue", resData)

}
