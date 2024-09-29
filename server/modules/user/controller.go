package users

import (
	"an-overengineered-app/internal/config"
	"an-overengineered-app/internal/helpers"
	"an-overengineered-app/internal/httpResponse"
	"an-overengineered-app/internal/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func SignupUser(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()

	logger.PrintInfo(reqCtx, "SignupUser controller method invoked...", nil)

	var userData SignupBody

	// TODO move server side validation to middleware
	if err := ctx.ShouldBindJSON(&userData); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			logger.PrintErrorWithStack(reqCtx, "Failed to validate req body", err)
			ctx.Error(httpResponse.ValidationError("", errs))
			return
		}

		logger.PrintErrorWithStack(reqCtx, "Failed to parse req body", err)

		ctx.Error(httpResponse.BadRequestError("Failed to parse request data"))
		return
	}

	logger.PrintInfo(reqCtx, "Validated request body", userData)

	isNewUser, err := IsEmailUnique(reqCtx, userData.Email)

	if err != nil {
		ctx.Error(httpResponse.InternerServerError(""))
		return
	}

	if !isNewUser {
		logger.PrintError(reqCtx, "User with provided email already exists")
		ctx.Error(httpResponse.ConflictError("Email already taken!", map[string]string{
			"email": "Email already taken!",
		}))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)

	if err != nil {
		logger.PrintErrorWithStack(reqCtx, "Hashing password failed", err)
		ctx.Error(httpResponse.InternerServerError(""))
		return
	}

	newUserData := BuildNewUserObj(userData, hashedPassword)

	logger.PrintInfo(reqCtx, "Constructed new user object", newUserData)

	tx := config.DBInstance.Begin()

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		tx.Rollback()
	// 	}
	// }()

	if err := tx.Error; err != nil {
		logger.PrintErrorWithStack(reqCtx, "Initiating db transaction failed", err)
		ctx.Error(httpResponse.InternerServerError(""))
		return
	}

	err = CreateUser(&newUserData, tx, reqCtx)

	logger.PrintInfo(reqCtx, "User object after inserting into db", newUserData)

	if err != nil {
		logger.PrintErrorWithStack(reqCtx, "Failed insert new user in db", err)
		tx.Rollback()
		ctx.Error(httpResponse.InternerServerError("Failed to create new user"))
		return
	}

	otp, err := helpers.GenerateOTP(6)

	if err != nil {
		logger.PrintErrorWithStack(reqCtx, "Failed to generate OTP", err)
		ctx.Error(httpResponse.InternerServerError(""))
		return
	}

	logger.PrintInfo(reqCtx, "Generated OTP", map[string]string{"otp": otp})

	otpData := BuildOTPObj(otp, newUserData, config.OTP_TYPE_SIGNUP)

	logger.PrintInfo(reqCtx, "Generated OTP object to save into db", otpData)

	err = CreateOTP(&otpData, tx, reqCtx)
	logger.PrintInfo(reqCtx, "OTP object after inserting into db", otpData)

	if err != nil {
		logger.PrintErrorWithStack(reqCtx, "Failed to insert otp in db", err)
		tx.Rollback()
		ctx.Error(httpResponse.InternerServerError(""))
		return
	}

	err = SendSignupMail(reqCtx, newUserData.Email, otp)

	if err != nil {
		logger.PrintErrorWithStack(reqCtx, "Failed to send new signup mail with otp to user", err)

		tx.Rollback()
		ctx.Error(httpResponse.InternerServerError(""))
		return
	}

	tx.Commit()

	httpResponse.Created(ctx, "An OTP sent to your mail. Please verify your account to continue", gin.H{
		"email": newUserData.Email,
	})

}

/*
*
  - validate incoming request body
  - find if otp for the given email exists
  - check if otp expired
  - check if maximum retry count exceeds
  - mark user as verified
  - generate JWT
*/
func VerifySignupOTP(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()

	logger.PrintInfo(reqCtx, "Invoking ValidateSignupOTP controller func", nil)

	var body VerifyOTPBody

	// TODO should be moved to middleware
	if err := ctx.ShouldBindJSON(&body); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			logger.PrintErrorWithStack(reqCtx, "Failed to validate req body", err)
			ctx.Error(httpResponse.ValidationError("", errs))
		}

		logger.PrintErrorWithStack(reqCtx, "Failed to parse req body", err)
		ctx.Error(httpResponse.BadRequestError("Failed to parse request data"))
		return
	}

	logger.PrintInfo(reqCtx, "Parsed data from req body", body)

	foundOtp, err := FindOtp(reqCtx, body.Email, body.Otp, config.OTP_TYPE_SIGNUP)

	if err != nil {
		logger.PrintError(reqCtx, "No otp found with given data")
		ctx.Error(httpResponse.BadRequestError("Invalid OTP"))
		return
	}

	if time.Now().After(foundOtp.ExpiresAt) {
		logger.PrintError(reqCtx, "OTP expired!")
		ctx.Error(httpResponse.BadRequestError("OTP expired!"))
		return
	}

	if foundOtp.RetryCount > config.AppConfig.MaxOtpRetry {
		logger.PrintError(reqCtx, "Signup retry count exceeds maximum retry")
		ctx.Error(httpResponse.RetryExceeded("Too many retry. Please request for new OTP"))
		return
	}

	httpResponse.Success(ctx, "User successfully verified!", gin.H{"data": foundOtp})

}
