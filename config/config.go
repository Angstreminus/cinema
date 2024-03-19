package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	ServerAddr string
	DBPort     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
	SSLMode    string
	AccSecr    string
	AccExp     string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	once.Do(
		func() {
			config = &Config{
				ServerAddr: os.Getenv("SERVER_ADDR"),
				DBPort:     os.Getenv("POSTGRES_PORT"),
				DBUser:     os.Getenv("POSTGRES_USER"),
				DBPassword: os.Getenv("POSTGRES_PASSWORD"),
				DBHost:     os.Getenv("POSTGRES_HOST"),
				DBName:     os.Getenv("POSTGRES_DB"),
				SSLMode:    os.Getenv("SSLMode"),
				AccSecr:    os.Getenv("ACCESS_TOKEN_SECRET"),
				AccExp:     os.Getenv("ACCESS_TOKEN_EXPIRY_TIME"),
			}
		})
	return config, nil
}
