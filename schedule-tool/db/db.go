package db

import (
	"github.com/hiroki90/goringorin/go-crud-sample/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnection() (*gorm.DB, error) {
	cfg := configs.NewConfig()

	return gorm.Open(mysql.Open(cfg.DBDial), &gorm.Config{})
}

