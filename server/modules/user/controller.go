package users

import (
	users "an-overengineered-social-media-app/modules/user/models"
	"an-overengineered-social-media-app/pkg/helpers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// Things to do
/**
1. Validate body
2. Check if email or username already exists
3. Hash the password
4. Save  user to DB
5. Send mail
6. Send response to client
*/
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

	err = CreateUser(&newUserData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create new user",
			"errors":  gin.H{},
			"data":    gin.H{},
		})
		return
	}

	userPayload := newUserData.Sanitize()

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully signed up",
		"data": gin.H{
			"user": userPayload,
		},
		"errors": gin.H{},
	})
}
