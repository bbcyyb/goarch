package entity

import (
	"time"
)

type User struct {
	ID              uint      `gorm:"primary_key"`
	RsaId           string    `gorm:"type:nvarchar(128)"`
	UserName        string    `gorm:"type:nvarchar(256)"`
	Email           string    `gorm:"type:nvarchar(256)"`
	LicenceAccepted bool      `gorm:"type:bit;not null"`
	LastLogin       time.Time `gorm:"type:datetime"`
	Role            int       `gorm:"type:int;not null"`
}

func (User) TableName() string {
	return "User2"
}
