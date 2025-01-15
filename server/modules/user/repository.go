package users

import (
	"an-overengineered-app/internal/db"
	"an-overengineered-app/internal/logger"
	users "an-overengineered-app/modules/user/models"
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func IsEmailUnique(ctx context.Context, email string) (bool, error) {
	dbInstance := db.DBInstance

	var count int64

	result := dbInstance.WithContext(ctx).Model(&users.User{}).
		Where("email = ?", email).
		Select("id").
		Count(&count)

	if result.Error != nil {
		logger.ErrorStack(ctx,
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

		logger.ErrorStack(ctx, "Error occurred while creating new user.", err)
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

	dbInstance := db.DBInstance.WithContext(ctx)

	var foundOtp users.OtpCodes

	err := dbInstance.Model(&users.OtpCodes{}).
		Where(&users.OtpCodes{
			Email:   email,
			Otp:     otp,
			OtpType: otpType,
		}).
		First(&foundOtp).Error

	if err != nil {
		logger.ErrorStack(ctx, "Failed to fetch otp data from db", err)
		return users.OtpCodes{}, err
	}

	return foundOtp, nil
}

func UpdateUser(ctx context.Context, updateData users.User, email string) (users.User, error) {

	logger.Info(ctx, "Invoking update user repository method", map[string]interface{}{
		"updatePayload": updateData,
		"email":         email,
	})

	dbInstance := db.DBInstance.WithContext(ctx).Model(users.User{})
	var verifiedUser users.User

	err := dbInstance.
		Clauses(clause.Returning{}).
		Where("email = ?", email).
		Updates(&users.User{AccountStatus: updateData.AccountStatus}).
		Scan(&verifiedUser).
		Error

	if err != nil {
		return verifiedUser, err
	}

	return verifiedUser, nil
}
