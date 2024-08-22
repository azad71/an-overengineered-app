package users

import (
	"an-overengineered-app/internal/config"
	"an-overengineered-app/internal/logger"
	users "an-overengineered-app/modules/user/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func IsEmailAndUsernameUnique(email string, username string, ctx context.Context) (bool, error) {
	db := config.DBInstance

	var count int64

	result := db.WithContext(ctx).Model(&users.User{}).
		Where("email = ?", email).
		Or("username = ?", username).
		Select("id").
		Count(&count)

	if result.Error != nil {
		logger.PrintErrorWithStack(ctx,
			"Error occurred while checking email and username uniqueness.",
			result.Error,
		)
		return false, result.Error
	}

	return count == 0, nil
}

func CreateUser(userData *users.User, db *gorm.DB, ctx context.Context) error {

	if err := db.WithContext(ctx).Model(&users.User{}).Create(&userData).Error; err != nil {

		logger.PrintErrorWithStack(ctx, "Error occurred while creating new user.", err)
		return err
	}

	return nil
}

func CreateOTP(otpData *users.OtpCodes, db *gorm.DB, ctx context.Context) error {

	if err := db.WithContext(ctx).Model(&users.OtpCodes{}).Create(&otpData).Error; err != nil {
		fmt.Printf("Error occurred while creating new otp record. %v\n", err)
		return err
	}

	return nil
}
