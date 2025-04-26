package models

import "time"

type Customer struct {
  	AccountNo    string    `json:"account_no" gorm:"primaryKey;not null;unique"`
    	Name         string    `json:"name" gorm:"type:varchar(255);not null"`
    	Nik          string    `json:"nik" gorm:"type:varchar(255);unique;not null"`
    	PhoneNumber  string    `json:"phone_number" gorm:"type:varchar(255);unique;not null"`
    	Balance      int64     `json:"balance" gorm:"type:bigint;not null;default:0"`
    	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
}
