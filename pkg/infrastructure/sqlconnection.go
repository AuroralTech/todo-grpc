package infrastructure

import (
	"github.com/AuroralTech/todo-grpc/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewSQLConnection(cfg *config.AppConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.AppInfo.DatabaseURL), &gorm.Config{})
	return db, err
}
