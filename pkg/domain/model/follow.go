package model

import (
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	UserID   uint `json:"user_id" gorm:"foreignKey:UserID"`
	TargetID uint `json:"target_id" gorm:"foreignKey:TargetID"`

	User   *User `gorm:"foreignKey:UserID;references:ID"`
	Target *User `gorm:"foreignKey:TargetID;references:ID"`
}
