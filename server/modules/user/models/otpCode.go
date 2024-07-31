package users

import (
	"gorm.io/gorm"
)

type OtpCodes struct {
	gorm.Model
	Otp        string `gorm:"NOT NULL;size:20"`
	Username   string `gorm:"NOT NULL;size:50"`
	Email      string `gorm:"NOT NULL;size:300"`
	OtpType    string `gorm:"column:otp_type;size:20"`
	RetryCount int    `gorm:"column:retry_count;max:3"`
}
