package config

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var cfg *ini.File

// Setup initialize the configuration instance
func SetupServerConfig() {
	var err error
	cfg, err = ini.Load("config.dev.ini")
	if err != nil {
		log.Panicf("setting.Setup, fail to parse 'config.ini': %v", err)
	}

	mapTo("app", AppConfig)
	mapTo("database", DBConfig)
	mapTo("redis", RedisConfig)

}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Panicf("Cfg.MapTo %s err: %v", section, err)
	}
}

var db *gorm.DB

func SetupDB() {
	fmt.Println("Connecting to Database...")
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", DBConfig.Host, DBConfig.User, DBConfig.Password, DBConfig.Name, DBConfig.Port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicf("Failed to connect to db, err: %v", err)
	}

	conn, err := db.DB()

	if err != nil {
		log.Panicf("Failed to open connection, err: %v", err)
	}

	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxLifetime(time.Hour * time.Duration(DBConfig.ConnMaxLifeTime))
	fmt.Println("Server connected to database successfully")
}
