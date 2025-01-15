package config

import (
	"an-overengineered-app/internal/logger"
	"context"
	"os"
	"strconv"
)

type App struct {
	JwtSecret   string
	AppUrl      string
	RunMode     string
	HttpPort    int
	AppEnv      string
	MaxOtpRetry int
	JWTExpiry   int
}

var AppConfig = &App{}

type Database struct {
	Dialect         string
	User            string
	Password        string
	Host            string
	Name            string
	Port            int
	ConnMaxLifeTime int
	SSLMode         string
}

var DBConfig = &Database{}

type Redis struct {
	Host     string
	Password string
}

var RedisConfig = &Redis{}

type Email struct {
	From       string
	SMTPServer string
	Port       int
	Password   string
}

var EmailConfig = &Email{}

// Setup initialize the configuration instance
func SetupServerConfig() {
	logger.Info(context.TODO(), "Setting up server config...", nil)

	// Setup App Config
	AppConfig.AppUrl = os.Getenv("APP_URL")
	AppConfig.JwtSecret = os.Getenv("JWT_SECRET")
	AppConfig.AppEnv = os.Getenv("APP_ENV")
	AppConfig.HttpPort, _ = strconv.Atoi(os.Getenv("HTTP_PORT"))
	AppConfig.MaxOtpRetry, _ = strconv.Atoi(os.Getenv("MAX_OTP_RETRY"))
	AppConfig.JWTExpiry, _ = strconv.Atoi(os.Getenv("JWT_EXPIRY"))
	if appMode := os.Getenv("APP_ENV"); appMode != "production" {
		AppConfig.RunMode = "debug"
	} else {
		AppConfig.RunMode = "release"
	}

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
