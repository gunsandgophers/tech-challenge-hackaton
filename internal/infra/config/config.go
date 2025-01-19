package config

import (
	"log"
	"os"
)

func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	if key == "" {
		log.Printf("Environment variable '%s' has not been defined and has no default value\n", key)
		panic(1)
	}

	return defaultValue
}

func GetEnvOrNil(key string) *string {
	if value, exists := os.LookupEnv(key); exists {
		return &value
	}

	return nil
}
