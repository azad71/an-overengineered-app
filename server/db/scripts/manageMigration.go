package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os/exec"

	"github.com/go-ini/ini"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

		cmd := exec.Command("migrate", "create", "-ext=sql", "-dir=db/migrations", migrationName)

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

func getDBConfig() (string, error) {
	var cfg *ini.File
	var err error

	cfg, err = ini.Load("config.dev.ini")

	if err != nil {
		return "", err
	}

	dbConfig, err := cfg.GetSection("database")

	if err != nil {
		return "", err
	}

	user, _ := dbConfig.GetKey("User")
	password, _ := dbConfig.GetKey("Password")
	host, _ := dbConfig.GetKey("Host")
	name, _ := dbConfig.GetKey("Name")
	port, _ := dbConfig.GetKey("Port")
	sslMode, _ := dbConfig.GetKey("SSLMode")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", user.Value(), url.QueryEscape(password.Value()), host.Value(), port.Value(), name.Value(), sslMode.Value())

	return dsn, nil
}

func runMigration(action string) error {
	dsn, err := getDBConfig()
	if err != nil {
		log.Fatalf("Failed to load db config, error: %v", err)
		return err
	}

	m, err := migrate.New("file://db/migrations", dsn)
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
