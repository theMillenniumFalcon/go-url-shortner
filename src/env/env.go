package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func APIKey() string {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("No .env file found")
		return "undefined"
	}

	apiKey, exists := os.LookupEnv("API_KEY")

	if !exists {
		fmt.Println("no entry for API_KEY found in .env")
		return "undefined"
	}

	return apiKey
}
