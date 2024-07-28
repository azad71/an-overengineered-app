package users

import (
	users "an-overengineered-social-media-app/modules/user/models"
	"an-overengineered-social-media-app/pkg/config"
	"fmt"
)

func IsEmailAndUsernameUnique(email string, username string) (bool, error) {
	db := config.DBInstance

	var count int64

	result := db.Model(&users.User{}).Where("email = ?", email).Or("username = ?", username).Select("id").Count(&count)

	if result.Error != nil {
		fmt.Printf("Error occurred while checking email and username uniqueness. Error: %v", result.Error)
		return false, result.Error
	}

	return count == 0, nil
}

func CreateUser(userData *users.User) error {
	db := config.DBInstance

	result := db.Model(&users.User{}).Create(&userData)

	if result.Error != nil {
		fmt.Printf("Error occurred while creating new user. Error: %v", result.Error)
		return result.Error
	}

	return nil
}
