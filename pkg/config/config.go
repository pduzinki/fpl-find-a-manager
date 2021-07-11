package config

import (
	"log"
	"os"
)

type Config struct {
	DBConfig DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Name     string
	User     string
	Password string
}

func Load() Config {
	if os.Getenv("POSTGRES_DB") == "" {
		log.Println("Env variable not found, loading default config...")
		return LoadDefault()
	}

	DBConfig := DatabaseConfig{
		Host:     "postgres",
		Name:     os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	return Config{
		DBConfig: DBConfig,
	}
}

func LoadDefault() Config {
	DBConfig := DatabaseConfig{
		Host:     "localhost",
		Name:     "fpl-find-a-manager-dev",
		User:     "postgres",
		Password: "123",
	}

	return Config{
		DBConfig: DBConfig,
	}
}
