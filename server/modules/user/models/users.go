package users

import (
	"time"

	"gorm.io/gorm"
)

type AccountStatus string
type Gender string

const (
	Active    AccountStatus = "ACTIVE"
	Inactive  AccountStatus = "INACTIVE"
	Pending   AccountStatus = "PENDING"
	Suspended AccountStatus = "SUSPENDED"

	Male   Gender = "MALE"
	Female Gender = "FEMALE"
	Others Gender = "OTHERS"
)

type UserModel struct {
	gorm.Model
	Username      string        `gorm:"unique;NOT NULL;size:50"`
	Password      string        `gorm:"size:250;NOT NULL"`
	Email         string        `gorm:"size:300;unique;NOT NULL"`
	FirstName     *string       `gorm:"size:100;column:first_name"`
	LastName      *string       `gorm:"size:100;column:last_name"`
	AccountStatus AccountStatus `gorm:"column:account_status;default:PENDING"`
	Avatar        *string       `gorm:"size:250"`
	BirthDate     *time.Time    `gorm:"type:date;column:birth_date"`
	Gender        *Gender       ``
	Address       *string       `gorm:"size:1000"`
	UserTimezone  *string       `gorm:"size:100;column:user_timezone"`
	Description   *string       `gorm:"type:text"`
}
