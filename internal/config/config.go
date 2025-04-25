package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	return &Config{
		DBHost:     getEnv("DB_HOST"),
		DBPort:     getEnv("DB_PORT"),
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBName:     getEnv("DB_NAME"),
	}
}

func getEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return ""
	}

	return val
}
