package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"song-library/configs"
	"song-library/logger"
)

var dbConn *gorm.DB

func ConnectToDB() error {
	if os.Getenv("DB_PASSWORD") == "" {
		logger.Error.Printf("[db.ConnectToDB] DB_PASSWORD is not set ")
		return fmt.Errorf("DB_PASSWORD environment variable is not set")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		configs.AppSettings.PostgresParams.Database,
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Error.Printf("[db.ConnectToDB] Failed to connect to DB: %v", err)
		fmt.Printf("Failed to connect to database: %v\n", err)
		return err
	}
	logger.Info.Printf("[db.ConnectToDB] Connected to DB")
	fmt.Println("Connected to database")
	dbConn = db
	return nil
}

func CloseDBConn() error {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}
