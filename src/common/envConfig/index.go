package envconfig

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(env string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return err.Error()
	}

	envVal := os.Getenv(env)
	if envVal == "" {
		return ""
	}
	return envVal
}
