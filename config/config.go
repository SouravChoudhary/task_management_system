package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	Name       string `mapstructure:"name"`
	MaxAttempt int    `mapstructure:"max_attempt"`
}

func (db *DBConfig) Validate() error {
	if db.Host == "" {
		return errors.New("db.host is required")
	}
	if db.Port == "" {
		return errors.New("db.port is required")
	}
	if db.User == "" {
		return errors.New("db.user is required")
	}
	if db.Password == "" {
		return errors.New("db.password is required")
	}
	if db.Name == "" {
		return errors.New("db.name is required")
	}
	if db.MaxAttempt <= 0 {
		return errors.New("db.max_attempt must be greater than 0")
	}

	return nil
}

type AppConfig struct {
	ServerPort string   `mapstructure:"server_port"`
	DB         DBConfig `mapstructure:"db"`
}

func (cfg *AppConfig) Validate() error {
	if cfg.ServerPort == "" {
		return errors.New("server_port is required")
	}
	if err := cfg.DB.Validate(); err != nil {
		return err
	}
	return nil
}

func LoadConfig(path string) (*AppConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	log.Println("Loaded config.yaml successfully")
	return &config, nil
}
