package database

import (
	"fmt"
	"go-api/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMysql(config configs.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
	return gorm.Open(
		mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
}
