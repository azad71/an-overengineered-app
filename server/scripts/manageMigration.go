package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	action := flag.String("action", "", "defines migration actions")
	name := flag.String("name", "", "migration name")

	flag.Parse()

	actionValue := *action

	switch actionValue {
	case "create":
		migrationName := *name
		if migrationName == "" {
			log.Fatal("Please provide migration name")
			return
		}

		cmd := exec.Command("migrate", "create", "-ext=sql", "-dir=migrations", migrationName)

		_, err := cmd.Output()

		if err != nil {
			fmt.Printf("Failed to run the migration, error: %v", err)
			return
		}

		fmt.Printf("%s migration successfully created", migrationName)

	case "up":
		err := runMigration(actionValue)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("Migration successfully synced up with database")
	case "down":
		err := runMigration(actionValue)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("Migration downed successfully")

	default:
		fmt.Println("Invalid action type")
		return

	}

}

func LoadEnv() error {

	appMode := GetAppMode()

	err := godotenv.Load(appMode)
	log.Printf("Env value loaded from %s", appMode)

	if err != nil {
		return err
	}

	return nil
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

func getDBConfig() (string, error) {

	err := LoadEnv()

	if err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSL_MODE")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, name, sslMode,
	)

	return dsn, nil
}

func runMigration(action string) error {
	dsn, err := getDBConfig()
	if err != nil {
		log.Fatalf("Failed to load db config, error: %v", err)
		return err
	}

	m, err := migrate.New("file://migrations", dsn)

	if err != nil {
		log.Fatalf("Migration failed, error: %v", err)
		return err
	}

	if action == "up" {
		if err := m.Up(); err != nil {
			log.Fatal(err)
			return err
		}
	} else if action == "down" {
		if err := m.Steps(-1); err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
