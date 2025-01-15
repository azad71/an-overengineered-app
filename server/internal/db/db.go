package db

import (
	"an-overengineered-app/internal/config"
	"an-overengineered-app/internal/logger"
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// TODO make sure it's initiated once
// var DBInstance *gorm.DB

// func SetupDB() error {
// 	ctx := context.TODO()
// 	logger.Info(ctx, "Connecting to Database...", nil)

// 	var err error
// 	dbConfig := config.DBConfig

// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)
// 	DBInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		SkipDefaultTransaction: true,
// 		Logger: logger.NewDBLogger(*logger.GetLogger(), gormLogger.Config{
// 			Colorful: true,
// 			LogLevel: gormLogger.Silent,
// 		}),
// 	})

// 	if err != nil {
// 		logger.ErrorStack(ctx, "Failed to connect to db", err)
// 		return err
// 	}

// 	conn, err := DBInstance.DB()

// 	if err != nil {
// 		logger.ErrorStack(ctx, "Failed to open connection", err)
// 		return err
// 	}

// 	conn.SetMaxIdleConns(10)
// 	conn.SetMaxOpenConns(100)
// 	conn.SetConnMaxLifetime(time.Hour * time.Duration(dbConfig.ConnMaxLifeTime))
// 	logger.Info(ctx, "Server connected to database successfully", nil)

// 	return nil
// }

type DB struct {
	DBInstance *gorm.DB
}

var dbConn = &DB{}

func ConnectDB() (*DB, error) {
	ctx := context.Background()
	logger.Info(ctx, "Connecting to Database...", nil)

	dbConfig := config.DBConfig

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger: logger.NewDBLogger(*logger.GetLogger(), gormLogger.Config{
			Colorful: true,
			LogLevel: gormLogger.Silent,
		}),
	})

	if err != nil {
		logger.ErrorStack(ctx, "Failed to connect to db", err)
		panic(err)
	}

	dbConn.DBInstance = conn

	return dbConn, nil
}
