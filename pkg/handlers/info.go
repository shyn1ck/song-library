package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"song-library/logger"
	services "song-library/pkg/services"
	"song-library/utils"
)

func ApiInfo(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[handlers.ApiInfo]: Client with ip: %s request to get InfoSong", ip)

	group := c.Query("group")
	song := c.Query("song")

	if group == "" || song == "" {
		logger.Error.Printf("[handlers.ApiInfo]: Error Invalid Request Parameter")
		handleError(c, utils.ErrInvalidRequestParameter)
		return
	}

	songDetail, err := services.GetSongDetail(group, song)
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[handlers.ApiInfo]: Client with ip: %s, successufly to get InfoSong", ip)
	c.JSON(http.StatusOK, songDetail)
}
