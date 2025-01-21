package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"song-library/configs"
	"song-library/db"
	"song-library/logger"
	"song-library/pkg/handlers"
	"song-library/server"
	"syscall"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Error.Printf("Error loading .env file")
	} else {
		fmt.Println(".env file loaded successfully")
	}

	if err := configs.ReadSettings(); err != nil {
		logger.Error.Fatalf("Error reading settings: %s", err)
	} else {
		fmt.Printf("Settings loaded: %+v\n", configs.AppSettings.AppParams)
	}

	if err := logger.Init(); err != nil {
		logger.Error.Fatalf("Error initializing logger: %s", err)
	} else {
		fmt.Println("Logger initialized successfully")
	}

	if err := db.ConnectToDB(); err != nil {
		logger.Error.Fatalf("Error connecting to database: %s", err)
	} else {
		fmt.Println("Connected to the database successfully")
	}

	defer func() {
		if err := db.CloseDBConn(); err != nil {
			logger.Error.Printf("Error closing database connection: %v", err)
		}
	}()

	if err := db.Migrate(); err != nil {
		logger.Error.Fatalf("Failed to run database migrations: %v", err)
	} else {
		fmt.Println("Database migrations completed successfully")
	}

	mainServer := new(server.Server)
	go func() {
		fmt.Printf("Starting server on port %s\n", configs.AppSettings.AppParams.PortRun)
		if err := mainServer.Run(configs.AppSettings.AppParams.PortRun, handlers.InitRoutes()); err != nil {
			logger.Error.Fatalf("Error starting HTTP server: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if sqlDB, err := db.GetDBConn().DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			logger.Error.Fatalf("Error closing DB: %s", err)
		} else {
			fmt.Println("Database connection closed successfully")
		}
	} else {
		logger.Error.Fatalf("Error getting *sql.DB from GORM: %s", err)
	}

	if err := mainServer.Shutdown(context.Background()); err != nil {
		logger.Error.Fatalf("Error during server shutdown: %s", err)
	} else {
		fmt.Println("Server shut down gracefully")
	}
}
