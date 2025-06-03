package database

import (
	"fmt"
	"os"
)

type Config struct {
	Host         string
	Port         string
	User         string
	Password     string
	databaseName string
	SSLMode      string
	MaxConns     string
	MaxIdleConns string
	MaxLifetime  string
}

func NewConfig() *Config {
	return &Config{
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		User:         os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		databaseName: os.Getenv("DB_NAME"),
		SSLMode:      os.Getenv("DB_SSL_MODE"),
		MaxConns:     os.Getenv("DB_MAX_CONNECTIONS"),
		MaxIdleConns: os.Getenv("DB_MAX_IDLE_CONNECTIONS"),
		MaxLifetime:  os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"),
	}
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.Host, c.Port, c.User, c.Password, c.databaseName, c.SSLMode)
}
