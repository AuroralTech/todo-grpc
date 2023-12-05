package model

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID      uint   `json:"user_id" gorm:"foreignKey:UserID"`
	Description string `json:"description"`

	User *User `gorm:"foreignKey:UserID;references:ID"`
}
