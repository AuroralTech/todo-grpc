package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Task        string `json:"task"`
	IsCompleted bool   `json:"is_completed"`
}
