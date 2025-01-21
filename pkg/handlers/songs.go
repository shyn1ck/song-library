package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"song-library/logger"
	services "song-library/pkg/services"
	"strconv"
)

func GetSongs(c *gin.Context) {
	group := c.Query("group")
	song := c.Query("song")

	pageParam := c.Query("page")
	limitParam := c.Query("limit")

	if pageParam == "" {
		pageParam = "1"
	}

	if limitParam == "" {
		limitParam = "10"
	}

	page, _ := strconv.Atoi(pageParam)
	limit, _ := strconv.Atoi(limitParam)

	songs, err := services.GetSongs(group, song, page, limit)
	if err != nil {
		handleError(c, err)
		return
	}

	if songs == nil {
		c.JSON(http.StatusOK, DefaultResponse{Message: "No songs found."})
		return
	}

	c.JSON(http.StatusOK, songs)
}

func GetSongByID(c *gin.Context) {
	id := c.Param("id")
	ip := c.ClientIP()
	logger.Info.Printf("[handlers.GetSongByID] Client IP: %s - Request to get song by id: %s\n]", ip, id)

}

func AddSong(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[handlers.AddSong] Client IP: %s - Request to add song\n", ip)

}

func UpdateSong(c *gin.Context) {
	id := c.Param("id")
	ip := c.ClientIP()

	logger.Info.Printf("[handlers.UpdateSong] Client IP: %s - Request to update song by id: %s\n", ip, id)
}

func DeleteSong(c *gin.Context) {
	id := c.Param("id")
	ip := c.ClientIP()

	logger.Info.Printf("[handlers.DeleteSong] Client IP: %s - Request to delete song by id: %s\n]", ip, id)

}
