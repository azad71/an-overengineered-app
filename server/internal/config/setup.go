package config

import (
	"an-overengineered-app/internal/logger"
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Setup initialize the configuration instance
func SetupServerConfig() {
	logger.PrintInfo(context.TODO(), "Setting up server config...", nil)

	// Setup App Config
	AppConfig.AppUrl = os.Getenv("APP_URL")
	AppConfig.JwtSecret = os.Getenv("JWT_SECRET")
	AppConfig.AppEnv = os.Getenv("APP_ENV")
	AppConfig.HttpPort, _ = strconv.Atoi(os.Getenv("HTTP_PORT"))
	AppConfig.RunMode = os.Getenv("RUN_MODE")

	// Setup DB Config
	DBConfig.Dialect = os.Getenv("DIALECT")
	DBConfig.User = os.Getenv("DB_USER")
	DBConfig.Password = os.Getenv("DB_PASSWORD")
	DBConfig.Host = os.Getenv("DB_HOST")
	DBConfig.Name = os.Getenv("DB_NAME")
	DBConfig.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	DBConfig.ConnMaxLifeTime, _ = strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFE_TIME"))
	DBConfig.SSLMode = os.Getenv("SSL_MODE")

	// Setup Redis
	RedisConfig.Host = os.Getenv("REDIS_HOST")
	RedisConfig.Password = os.Getenv("REDIS_PASSWORD")

	// Setup SMTP
	EmailConfig.From = os.Getenv("EMAIL_FROM")
	EmailConfig.SMTPServer = os.Getenv("SMTP_SERVER")
	EmailConfig.Port, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	EmailConfig.Password = os.Getenv("SMTP_PASSWORD")

}

var DBInstance *gorm.DB

func SetupDB() {
	logger.PrintInfo(context.TODO(), "Connecting to Database...", nil)
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", DBConfig.Host, DBConfig.User, DBConfig.Password, DBConfig.Name, DBConfig.Port)
	DBInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger: logger.NewDBLogger(*logger.GetLogger(), gormLogger.Config{
			Colorful: true,
			LogLevel: gormLogger.Silent,
		}),
	})

	if err != nil {
		logger.PrintPanic(context.TODO(), "Failed to connect to db", err)
	}

	conn, err := DBInstance.DB()

	if err != nil {
		logger.PrintPanic(context.TODO(), "Failed to open connection", err)
	}

	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxLifetime(time.Hour * time.Duration(DBConfig.ConnMaxLifeTime))
	logger.PrintInfo(context.TODO(), "Server connected to database successfully", nil)
}
