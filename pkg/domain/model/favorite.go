package model

import (
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserID uint `json:"user_id" gorm:"foreignKey:UserID"`
	PostID uint `json:"post_id" gorm:"foreignKey:PostID"`

	User *User `gorm:"foreignKey:UserID;references:ID"`
	Post *Post `gorm:"foreignKey:PostID;references:ID"`
}
