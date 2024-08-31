package helpers

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var GetAppModeFunc = GetAppMode

func LoadEnv() {
	log.Info().
		Str("source", "LoadEnv").
		Msg("Initiating loading env...")

	appMode := GetAppModeFunc()

	err := godotenv.Load(appMode)

	log.Info().
		Str("source", "LoadEnv").
		Msgf("Env value loaded from %s", appMode)

	if err != nil {
		log.Fatal().
			Str("source", "LoadEnv").
			Msgf("Failed to load env file: %v", err)
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
