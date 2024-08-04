package helpers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	fmt.Println("Initiating loading env...")
	err := godotenv.Load(".env.dev")

	if err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}
}
