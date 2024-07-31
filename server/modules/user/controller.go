package users

import (
	users "an-overengineered-social-media-app/modules/user/models"
	"an-overengineered-social-media-app/pkg/config"
	"an-overengineered-social-media-app/pkg/helpers"
	"an-overengineered-social-media-app/pkg/mailer"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func SignupUser(c *gin.Context) {

	var userData SignupBody

	if err := c.ShouldBindJSON(&userData); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "Invalid request data",
				"errors":  helpers.FormatValidationError(errs),
				"data":    gin.H{},
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Failed to parse request data",
			"errors":   err.Error(),
			"data":     gin.H{},
		})
		return
	}

	isNewUser, err := IsEmailAndUsernameUnique(userData.Email, userData.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"errors":  gin.H{},
			"data":    gin.H{},
		})
		return
	}

	if !isNewUser {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Email/Username already exists!",
			"errors": gin.H{
				"email": "Email/Username already exists!",
			},
			"data": gin.H{},
		})
		return
	}

	hasshedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"errors":  gin.H{},
			"data":    gin.H{},
		})
		return
	}

	birthDate, _ := time.Parse("2006-01-02", *userData.BirthDate)

	newUserData := users.User{
		Username:     userData.Username,
		Password:     string(hasshedPassword),
		Email:        userData.Email,
		FirstName:    userData.FirstName,
		LastName:     userData.LastName,
		BirthDate:    &birthDate,
		Gender:       userData.Gender,
		Address:      userData.Address,
		UserTimezone: userData.UserTimezone,
		Description:  userData.Description,
	}

	tx := config.DBInstance.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"errors":  gin.H{},
			"data":    gin.H{},
		})
		return
	}

	err = CreateUser(&newUserData, tx)

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create new user",
			"errors":  gin.H{},
			"data":    gin.H{},
		})
		return
	}

	otp, err := helpers.GenerateOTP(8)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"errors":  gin.H{},
			"data":    gin.H{},
		})
		return
	}

	newOtpData := users.OtpCodes{
		Otp:        otp,
		Username:   newUserData.Username,
		Email:      newUserData.Email,
		OtpType:    "SIGNUP",
		RetryCount: 0,
	}

	err = CreateOTP(&newOtpData, tx)

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"errors":  gin.H{},
			"data":    gin.H{},
		})
		return
	}

	mailContent := mailer.GetSignupContent(otp)

	err = mailer.SendMail([]string{newUserData.Email}, []byte(mailContent), "auth")

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"errors":  gin.H{},
			"data":    gin.H{},
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"message": "An OTP sent to your mail. Please verify your account to continue",
		"data": gin.H{
			"expiresIn":    5,
			"timeUnit":     "minutes",
			"maximumRetry": 3,
		},
		"errors": gin.H{},
	})
}
