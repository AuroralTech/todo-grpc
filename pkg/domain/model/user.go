package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firebase_uid string `json:"firebase_uid"`
	Name         string `json:"name"`

	Profile   Profile    `json:"profile" gorm:"foreignKey:UserID"`
	Posts     []Post     `json:"posts" gorm:"foreignKey:UserID"`
	Favorites []Favorite `json:"favorites" gorm:"foreignKey:UserID"`
	Follows   []Follow   `json:"follows" gorm:"foreignKey:UserID"`
}
