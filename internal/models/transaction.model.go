package models

import "time"

type Transaction struct {
    ID           uint      `gorm:"primaryKey;autoIncrement"`
    AccountNo    string    `gorm:"type:varchar(255);not null"`
    Type         string    `gorm:"type:varchar(255);not null"`
    Amount       int64     `gorm:"type:bigint;not null"`
    FinalBalance int64     `gorm:"type:bigint;not null"`
    CreatedAt    time.Time `gorm:"autoCreateTime"`

    Customer Customer `gorm:"foreignKey:AccountNo;references:AccountNo"`
}
