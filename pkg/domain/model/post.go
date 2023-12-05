package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID  uint   `json:"user_id" gorm:"foreignKey:UserID"`
	Content string `json:"content"`
	Image   string `json:"image"`

	User      *User      `json:"user" gorm:"foreignKey:UserID"`
	Favorites []Favorite `json:"favorite" gorm:"foreignKey:PostID"`
}
