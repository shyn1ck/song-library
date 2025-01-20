package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"song-library/logger"
	service "song-library/pkg/services"
	"strconv"
)

func GetSongs(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[controllers.GetSongs] Client IP: %s - Request to get list of songs\n", ip)

	group := c.DefaultQuery("group", "")
	song := c.DefaultQuery("song", "")
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		logger.Error.Printf("[controllers.GetSongs] Client IP: %s - Invalid page parameter: %v\n", ip, err)
		handleError(c, err)
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		logger.Error.Printf("[controllers.GetSongs] Client IP: %s - Invalid limit parameter: %v\n", ip, err)
		handleError(c, err)
		return
	}

	if group == "" && song == "" {
		err := errors.New("at least one of 'group' or 'song' must be provided")
		logger.Error.Printf("[controllers.GetSongs] Client IP: %s - %v\n", ip, err)
		handleError(c, err)
		return
	}

	songs, err := service.GetSongs(group, song, page, limit)
	if err != nil {
		handleError(c, err)
		return
	}
	logger.Info.Printf("[controllers.GetSongs] Client IP: %s - Successfully retrieved list of songs.\n", ip)
	c.JSON(http.StatusOK, songs)
}
