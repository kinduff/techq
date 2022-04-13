package config

import "os"

type Config struct {
	Port string
}

func New() *Config {
	return &Config{
		Port: getEnv("PORT", "3000"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
