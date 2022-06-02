package config

import (
	"github.com/spf13/viper"
	"go-api/pkg/validation"
)

type Config interface {
	Bind(v interface{}) error
}

type config struct {
	Type string
	Path string
	Name string
}

func (c config) Bind(v interface{}) error {
	viper.SetConfigType(c.Type)
	viper.AddConfigPath(c.Path)
	viper.SetConfigName(c.Name)
	
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	
	err = viper.Unmarshal(&v)
	if err != nil {
		return err
	}
	
	return validation.Validate(v)
}

func DefaultEnv() Config {
	return config{
		Type: "env",
		Path: ".",
		Name: ".env",
	}
}
