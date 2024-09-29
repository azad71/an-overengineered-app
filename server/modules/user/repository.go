package users

import (
	"an-overengineered-app/internal/config"
	"an-overengineered-app/internal/logger"
	users "an-overengineered-app/modules/user/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func IsEmailUnique(ctx context.Context, email string) (bool, error) {
	db := config.DBInstance

	var count int64

	result := db.WithContext(ctx).Model(&users.User{}).
		Where("email = ?", email).
		Select("id").
		Count(&count)

	if result.Error != nil {
		logger.PrintErrorWithStack(ctx,
			"Error occurred while checking email uniqueness.",
			result.Error,
		)
		return false, result.Error
	}

	return count == 0, nil
}

// TODO need to set ctx as first func param
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

func FindOtp(ctx context.Context, email string, otp string, otpType string) (users.OtpCodes, error) {

	db := config.DBInstance.WithContext(ctx)

	var foundOtp users.OtpCodes

	err := db.Model(&users.OtpCodes{}).
		Where(&users.OtpCodes{
			Email:   email,
			Otp:     otp,
			OtpType: otpType,
		}).
		First(&foundOtp).Error

	if err != nil {
		logger.PrintErrorWithStack(ctx, "Failed to fetch otp data from db", err)
		return users.OtpCodes{}, err
	}

	return foundOtp, nil
}

func UpdateUser(ctx context.Context, updateData users.User, email string) {
	db := config.DBInstance.WithContext(ctx)

}
