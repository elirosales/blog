package config

import "os"

type Config struct {
	API struct {
		Port string
	}

	Database struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
		DSN      string
	}
}

// Initialize config
func New() *Config {
	var config Config
	config.API.Port = getEnv("PORT")

	config.Database.Host = getEnv("DB_HOST")
	config.Database.Port = getEnv("DB_PORT")
	config.Database.Username = getEnv("DB_USERNAME")
	config.Database.Password = getEnv("DB_PASSWORD")
	config.Database.DBName = getEnv("DB_NAME")

	return &config
}

func getEnv(key string) string {
	return os.Getenv(key)
}
