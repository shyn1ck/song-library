package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"song-library/configs"
	"song-library/logger"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)
	r.GET("/ping", PingPong)

	if err := r.Run(fmt.Sprintf("%s:%s", configs.AppSettings.AppParams.ServerURL, configs.AppSettings.AppParams.PortRun)); err != nil {
		logger.Error.Fatalf("Error starting server: %v", err)
	}
	return r
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
