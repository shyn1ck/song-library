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

// @title Song Library API 🎶
// @version 1.0
// @description API for managing a song library.

// @host localhost:8181
// @BasePath /

// @contact.name   Personal tg
// @contact.url    https://t.me/parvizjon_hasanov
// @contact.email  hy.parvizjon@outlook.com

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	} else {
		fmt.Println(".env file loaded successfully")
	}

	if err := configs.ReadSettings(); err != nil {
		fmt.Printf("Error reading settings: %v\n", err)
		return
	} else {
		fmt.Printf("Settings loaded: %+v\n", configs.AppSettings.AppParams)
	}

	if err := logger.Init(); err != nil {
		fmt.Printf("Error initializing logger: %v\n", err)
	}
	fmt.Println("Logger initialized successfully")

	if err := db.ConnectToDB(); err != nil {
		fmt.Printf("Error connecting to DB: %v\n", err)
		return
	}
	defer func() {
		if err := db.CloseDBConn(); err != nil {
			fmt.Printf("Error closing database connection: %v\n", err)
		} else {
			fmt.Println("Database connection closed successfully")
		}
	}()
	fmt.Println("Connected to the database successfully")

	if err := db.Migrate(); err != nil {
		fmt.Printf("Error initializing database migrations: %v\n", err)
		return
	}
	fmt.Println("Database migrations completed successfully")

	mainServer := new(server.Server)
	secondServer := new(server.Server)

	go func() {
		appPort := configs.AppSettings.AppParams.PortRun
		fmt.Printf("Starting application server on port %s\n", appPort)
		if err := mainServer.Run(appPort, handlers.InitRoutes()); err != nil {
			fmt.Printf("Error starting application server: %s\n", err)
		}
	}()

	go func() {
		apiPort := configs.AppSettings.AppParams.ApiPortRun
		fmt.Printf("Starting API server on port %s\n", apiPort)
		if err := mainServer.Run(apiPort, handlers.InitRoutes()); err != nil {
			fmt.Printf("Error starting API server: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	fmt.Println("Shutting down servers...")

	if err := mainServer.Shutdown(context.Background()); err != nil {
		fmt.Printf("Error during application server shutdown: %s\n", err)
	}
	fmt.Println("Application server shut down gracefully")

	if err := secondServer.Shutdown(context.Background()); err != nil {
		fmt.Printf("Error during application server shutdown: %s\n", err)
	}
	fmt.Println("Application server shut down gracefully")
}
