package models

import "time"

type Transaction struct {
	ID uint `json:"id" gorm:"primary_key;auto_increment"`
	AccountNo string `json:"account_no" gorm:"type:varchar(255);not null"`
	Type string `json:"type" gorm:"type:varchar(255);not null"`
	Amount int64 `json:"amount" gorm:"type:bigint;not null"`
	FinalBalance int64 `json:"final_balance" gorm:"type:bigint;not null"`
	CreatedAt time.Time
}
