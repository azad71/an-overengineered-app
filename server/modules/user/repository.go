package users

import (
	"an-overengineered-social-media-app/internal/config"
	users "an-overengineered-social-media-app/modules/user/models"
	"fmt"

	"gorm.io/gorm"
)

func IsEmailAndUsernameUnique(email string, username string) (bool, error) {
	db := config.DBInstance

	var count int64

	result := db.Model(&users.User{}).
		Where("email = ?", email).
		Or("username = ?", username).
		Select("id").
		Count(&count)

	if result.Error != nil {
		fmt.Printf("Error occurred while checking email and username uniqueness. %v\n", result.Error)
		return false, result.Error
	}

	return count == 0, nil
}

func CreateUser(userData *users.User, db *gorm.DB) error {

	if err := db.Model(&users.User{}).Create(&userData).Error; err != nil {

		fmt.Printf("Error occurred while creating new user. %v\n", err)
		return err
	}

	return nil
}

func CreateOTP(otpData *users.OtpCodes, db *gorm.DB) error {

	if err := db.Model(&users.OtpCodes{}).Create(&otpData).Error; err != nil {
		fmt.Printf("Error occurred while creating new otp record. %v\n", err)
		return err
	}

	return nil
}
