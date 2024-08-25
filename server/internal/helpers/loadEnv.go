package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	fmt.Println("Initiating loading env...")
	appMode := GetAppMode()

	err := godotenv.Load(appMode)

	fmt.Printf("Env value loaded from %s\n", appMode)

	if err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}
}

func GetAppMode() string {
	appMode := os.Getenv("APP_ENV")

	switch appMode {
	case "production":
		return ".env"
	case "docker":
		return ".env.docker"
	default:
		return ".env.dev"
	}

}
