package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"song-library/configs"
	"song-library/logger"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", PingPong)

	songGroup := r.Group("/songs")
	{
		songGroup.GET("/", GetSongs)
		songGroup.GET("/:id", GetSongByID)
		songGroup.POST("/", AddSong)
		songGroup.PUT("/:id", UpdateSong)
		songGroup.DELETE("/:id", SoftDeleteSong)
		songGroup.DELETE("/hard/:id", HardDeleteSong)
	}

	lyricsGroup := r.Group("/lyrics")
	{
		lyricsGroup.GET("/:title", GetLyrics)
		lyricsGroup.GET("/", GetLyricsByText)
	}

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
