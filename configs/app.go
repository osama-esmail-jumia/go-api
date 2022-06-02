package configs

import (
	"go-api/pkg/config"
)

type App struct {
	Port int64 `mapstructure:"APP_PORT" validate:"required"`
	DB   DB    `mapstructure:",squash"`
}

type DB struct {
	Host     string `mapstructure:"DB_HOST" validate:"required"`
	Port     int64  `mapstructure:"DB_PORT" validate:"required,number"`
	Name     string `mapstructure:"DB_NAME" validate:"required"`
	User     string `mapstructure:"DB_USER" validate:"required"`
	Password string `mapstructure:"DB_PASSWORD" validate:"required"`
}

func NewAppConfig() (App, error) {
	var app App
	err := config.DefaultEnv().Bind(&app)
	return app, err
}
