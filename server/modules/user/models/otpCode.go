package users

import (
	"time"

	"gorm.io/gorm"
)

type OtpCodes struct {
	gorm.Model
	Otp        string    `gorm:"NOT NULL;size:20"`
	Email      string    `gorm:"NOT NULL;size:300"`
	OtpType    string    `gorm:"column:otp_type;size:20"`
	RetryCount int       `gorm:"column:retry_count;default:0"`
	ExpiresAt  time.Time `gorm:"column:expires_at;"`
}
