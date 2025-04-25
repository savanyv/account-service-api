package models

import "time"

type Customer struct {
	ID uint `json:"id" gorm:"primary_key;auto_increment"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	NIK string `json:"nik" gorm:"type:varchar(255);not null"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(255);not null"`
	AccountNo string `json:"account_no" gorm:"type:varchar(255);not null"`
	Balance int64 `json:"balance" gorm:"type:bigint;not null;default:0"`
	CreatedAt time.Time
}
