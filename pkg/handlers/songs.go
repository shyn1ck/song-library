package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"song-library/logger"
	services "song-library/pkg/services"
	"song-library/utils"
	"strconv"
)

func GetSongs(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[GetSongs]: Client with IP=%s, requested to get songs", ip)
	group := c.Query("group")
	song := c.Query("song")

	pageParam := c.Query("page")
	limitParam := c.Query("limit")

	page := 1
	limit := 10

	if pageParam != "" {
		var err error
		page, err = strconv.Atoi(pageParam)
		if err != nil {
			handleError(c, utils.ErrInvalidPaginationParams)
			return
		}
	}

	if limitParam != "" {
		var err error
		limit, err = strconv.Atoi(limitParam)
		if err != nil {
			handleError(c, utils.ErrInvalidPaginationParams)
			return
		}
	}

	songs, err := services.GetSongs(group, song, page, limit)
	if err != nil {
		logger.Error.Printf("[handlers.GetSongs]: Error: %v", err)
		handleError(c, err)
		return
	}

	if songs == nil {
		logger.Info.Printf("[handlers.GetSongs]: Client with IP=%s, no songs found", ip)
		c.JSON(http.StatusOK, DefaultResponse{Message: "No songs found."})
		return
	}

	logger.Info.Printf("[handlers.GetSongs]: Client with IP=%s, successfully retrieved songs", ip)
	c.JSON(http.StatusOK, songs)
}

func GetSongByID(c *gin.Context) {
	ip := c.ClientIP()
	idParam := c.Param("id")

	logger.Info.Printf("[handlers.GetSongByID] Client IP: %s - Request to get song by id: %s", ip, idParam)

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		logger.Error.Printf("[handlers.GetSongByID] Invalid ID format: %s", err)
		handleError(c, utils.ErrInvalidID)
		return
	}

	song, err := services.GetSongByID(uint(id))
	if err != nil {
		logger.Error.Printf("[handlers.GetSongByID] Error getting song: %s", err)
		handleError(c, err)
		return
	}

	if song == nil {
		logger.Info.Printf("[handlers.GetSongByID] Song not found")
		handleError(c, utils.ErrSongNotFound)
		return
	}
	c.JSON(http.StatusOK, song)
}

func GetLyrics(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[handlers.GetLyrics]: Client with IP=%s, requested to get lyrics", ip)

	song := c.Param("title")
	pageParam := c.Query("page")
	limitParam := c.Query("limit")

	page := 1
	limit := 10
	if pageParam != "" {
		var err error
		page, err = strconv.Atoi(pageParam)
		if err != nil {
			handleError(c, utils.ErrInvalidPaginationParams)
			return
		}
	}

	if limitParam != "" {
		var err error
		limit, err = strconv.Atoi(limitParam)
		if err != nil {
			handleError(c, utils.ErrInvalidPaginationParams)
			return
		}
	}

	logger.Info.Printf("[handlers.GetLyrics]: Searching for song: %s", song)

	lyrics, err := services.GetLyrics(song, page, limit)
	if err != nil {
		logger.Error.Printf("[handlers.GetLyrics]: Error: %v", err)
		handleError(c, err)
		return
	}

	logger.Info.Printf("[handlers.GetLyrics]: Client with IP=%s, successfully retrieved lyrics", ip)
	c.JSON(http.StatusOK, lyrics)
}

func GetLyricsByText(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[handlers.GetLyricsByText]: Client with IP=%s, requested to get lyrics by text", ip)

	searchText := c.Query("search")
	if searchText == "" {
		handleError(c, utils.ErrInvalidText)
		return
	}

	pageParam := c.Query("page")
	limitParam := c.Query("limit")

	page := 1
	limit := 10

	if pageParam != "" {
		var err error
		page, err = strconv.Atoi(pageParam)
		if err != nil {
			handleError(c, utils.ErrInvalidPaginationParams)
			return
		}
	}

	if limitParam != "" {
		var err error
		limit, err = strconv.Atoi(limitParam)
		if err != nil {
			handleError(c, utils.ErrInvalidPaginationParams)
			return
		}
	}

	logger.Info.Printf("[handlers.GetLyricsByText]: Searching for lyrics containing: %s", searchText)

	lyrics, err := services.GetLyricsByText(searchText, page, limit)
	if err != nil {
		logger.Error.Printf("[handlers.GetLyricsByText]: Error: %v", err)
		handleError(c, err)
		return
	}

	if lyrics == nil {
		logger.Info.Printf("[handlers.GetLyricsByText]: Client with IP=%s, no lyrics found", ip)
		c.JSON(http.StatusOK, DefaultResponse{Message: "No lyrics found."})
		return
	}

	logger.Info.Printf("[handlers.GetLyricsByText]: Client with IP=%s, successfully retrieved lyrics by text", ip)
	c.JSON(http.StatusOK, lyrics)
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
