package users

import (
	"an-overengineered-app/internal/logger"
	users "an-overengineered-app/modules/user/models"
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserImpl interface {
	IsEmailUnique(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, payload *users.User, tx *gorm.DB) error
	CreateOTP(ctx context.Context, payload *users.OtpCodes, tx *gorm.DB) error
	FindOTP(ctx context.Context, email, otp, otpType string) (users.OtpCodes, error)
	UpdateUser(ctx context.Context, payload users.User, email string) (users.User, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(conn *gorm.DB) UserImpl {
	return &UserRepo{
		DB: conn,
	}
}

func (u *UserRepo) IsEmailUnique(ctx context.Context, email string) (bool, error) {
	var count int64

	result := u.DB.WithContext(ctx).Model(&users.User{}).
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

func (u *UserRepo) CreateUser(ctx context.Context, payload *users.User, tx *gorm.DB) error {

	if err := tx.WithContext(ctx).Model(&users.User{}).Create(&payload).Error; err != nil {

		logger.ErrorStack(ctx, "Error occurred while creating new user.", err)
		return err
	}

	return nil
}

func (u *UserRepo) CreateOTP(ctx context.Context, payload *users.OtpCodes, tx *gorm.DB) error {

	if err := tx.WithContext(ctx).Model(&users.OtpCodes{}).Create(&payload).Error; err != nil {
		fmt.Printf("Error occurred while creating new otp record. %v\n", err)
		return err
	}

	return nil
}

func (u *UserRepo) FindOTP(ctx context.Context, email string, otp string, otpType string) (users.OtpCodes, error) {
	var foundOtp users.OtpCodes

	err := u.DB.WithContext(ctx).Model(&users.OtpCodes{}).
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

func (u *UserRepo) UpdateUser(ctx context.Context, payload users.User, email string) (users.User, error) {

	logger.Info(ctx, "Invoking update user repository method", map[string]interface{}{
		"updatePayload": payload,
		"email":         email,
	})

	dbInstance := u.DB.WithContext(ctx).Model(users.User{})
	var verifiedUser users.User

	err := dbInstance.
		Clauses(clause.Returning{}).
		Where("email = ?", email).
		Updates(&users.User{AccountStatus: payload.AccountStatus}).
		Scan(&verifiedUser).
		Error

	if err != nil {
		return verifiedUser, err
	}

	return verifiedUser, nil
}
