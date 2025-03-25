package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBHost		string
	DBUser		string
	DBPassword	string
	DBName		string
	DBPort		string
}

func LoadConfig() *Config {
	return &Config{
		DBHost: 		os.Getenv("DB_HOST"),
		DBUser: 		os.Getenv("DB_USER"),
		DBPassword: 	os.Getenv("DB_PASSWORD"),
		DBName: 		os.Getenv("DB_NAME"),
		DBPort: 		os.Getenv("DB_PORT"),
	}
}

func getEnv(key, defualtValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defualtValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err ==nil {
		return value
	}
	return defaultValue
}