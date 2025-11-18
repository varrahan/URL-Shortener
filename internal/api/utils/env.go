package utils

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func LoadEnv(filepath string) {
	if err := godotenv.Load(filepath); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func GetAllEnv(filepath string) map[string]string {
	LoadEnv(filepath)

	envMap := make(map[string]string)
	for _, e := range os.Environ() {
		pair := splitOnce(e, "=")
		envMap[pair[0]] = pair[1]
	}
	return envMap
}

func splitOnce(s, sep string) [2]string {
	for i := 0; i < len(s); i++ {
		if string(s[i]) == sep {
			return [2]string{s[:i], s[i+1:]}
		}
	}
	return [2]string{s, ""}
}