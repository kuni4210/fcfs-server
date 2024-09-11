package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Jwt      JwtConfig
}

type ServerConfig struct {
	Mode string
	Port string
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

type JwtConfig struct {
	SecretKey string `json:"secret_key"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
