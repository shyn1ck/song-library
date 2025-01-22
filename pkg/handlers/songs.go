package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"song-library/logger"
	"song-library/models"
	services "song-library/pkg/services"
	"song-library/utils"
	"strconv"
)

// GetSongs godoc
// @Summary      Get songs
// @Description  Retrieves a list of songs based on optional filters such as group name, song name, pagination, and limit.
// @Tags         Songs
// @Accept       json
// @Produce      json
// @Param        group    query   string  false  "Group name"
// @Param        song     query   string  false  "Song name"
// @Param        page     query   int     false  "Page number"  default(1)
// @Param        limit    query   int     false  "Number of results per page"  default(10)
// @Success      200      {array}  models.Song   "Success"  "List of songs"
// @Failure      400      {object}  ErrorResponse  "Invalid request"
// @Failure      500      {object}  ErrorResponse  "Internal server error"
// @Router       /songs [get]
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
		c.JSON(http.StatusNotFound, DefaultResponse{Message: "No songs found."})
		return
	}

	logger.Info.Printf("[handlers.GetSongs]: Client with IP=%s, successfully retrieved songs", ip)
	c.JSON(http.StatusOK, songs)
}

// GetSongByID godoc
// @Summary      Get song by ID
// @Description  Retrieves a song by its unique ID.
// @Tags         Songs
// @Accept       json
// @Produce      json
// @Param        id   path    int     true  "Song ID"
// @Success      200  {object}  models.Song   "Success"  "Song details"
// @Failure      400  {object}  ErrorResponse  "Invalid ID format"
// @Failure      404  {object}  ErrorResponse  "Song not found"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Router       /songs/{id} [get]
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

// AddSong godoc
// @Summary      Add a new song
// @Description  Adds a new song to the database with the provided details, such as title, artist, release date, and link.
// @Tags         Songs
// @Accept       json
// @Produce      json
// @Param        song  body    models.NewSongRequest  true  "New song details"
// @Success      200   {object}  DefaultResponse   "Success"  "Song added successfully with additional data."
// @Success      200   {object}  DefaultResponse   "Success"  "Song added successfully with provided data only."
// @Failure      400   {object}  ErrorResponse  "Invalid request body"
// @Failure      500   {object}  ErrorResponse  "Internal server error"
// @Router       /songs [post]
func AddSong(c *gin.Context) {
	ip := c.ClientIP()
	logger.Info.Printf("[handlers.AddSong] Client IP: %s - Request to add a new song", ip)
	var newSongRequest models.NewSongRequest
	if err := c.ShouldBindJSON(&newSongRequest); err != nil {
		logger.Error.Printf("[handlers.AddSong] Error binding JSON: %s", err)
		handleError(c, utils.ErrInvalidRequestBody)
		return
	}

	song, err := services.AddSong(newSongRequest)
	if err != nil {
		logger.Error.Printf("[handlers.AddSong] Error adding song: %s", err)
		handleError(c, err)
		return
	}

	if song.ReleaseDate != "" || song.Text != "" || song.Link != "" {
		response := DefaultResponse{Message: "Song added successfully with additional data."}
		c.JSON(http.StatusOK, response)
	} else {
		response := DefaultResponse{Message: "Song added successfully with provided data only."}
		c.JSON(http.StatusOK, response)
	}
}

// UpdateSong godoc
// @Summary      Update an existing song
// @Description  Updates an existing song by its unique ID with new details, such as title, artist, release date, and link.
// @Tags         Songs
// @Accept       json
// @Produce      json
// @Param        id   path    int     true  "Song ID"
// @Param        song body    	models.Song  true  "Updated song details"
// @Success      200  {object}  DefaultResponse   "Success"  "Song updated successfully"
// @Failure      400  {object}  ErrorResponse  "Invalid ID format or request body"
// @Failure      404  {object}  ErrorResponse  "Song not found"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Router       /songs/{id} [put]
func UpdateSong(c *gin.Context) {
	ip := c.ClientIP()
	idParam := c.Param("id")

	logger.Info.Printf("[handlers.UpdateSong] Client IP: %s - Request to update song by id: %s", ip, idParam)

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		logger.Error.Printf("[handlers.UpdateSong] Invalid ID format: %s", err)
		handleError(c, utils.ErrInvalidID)
		return
	}

	var songUpdate models.Song
	if err := c.ShouldBindJSON(&songUpdate); err != nil {
		logger.Error.Printf("[handlers.UpdateSong] Error binding JSON: %s", err)
		handleError(c, utils.ErrInvalidRequestBody)
		return
	}

	err = services.UpdateSong(uint(id), &songUpdate)
	if err != nil {
		logger.Error.Printf("[handlers.UpdateSong] Error updating song: %s", err)
		handleError(c, err)
		return
	}

	response := DefaultResponse{Message: fmt.Sprintf("Song with id: %d updated successfully.", id)}
	c.JSON(http.StatusOK, response)
}

// SoftDeleteSong godoc
// @Summary      Soft delete a song
// @Description  Soft deletes a song by its unique ID, marking it as deleted without actually removing it from the database.
// @Tags         Songs
// @Accept       json
// @Produce      json
// @Param        id   path    int     true  "Song ID"
// @Success      200  {object}  DefaultResponse   "Success"  "Song successfully soft deleted"
// @Failure      400  {object}  ErrorResponse  "Invalid ID format"
// @Failure      404  {object}  ErrorResponse  "Song not found"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /songs/{id} [delete]
func SoftDeleteSong(c *gin.Context) {
	ip := c.ClientIP()
	idParam := c.Param("id")

	logger.Info.Printf("[handlers.SoftDeleteSong] Client IP: %s - Request to soft delete song by id: %s", ip, idParam)

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		logger.Error.Printf("[handlers.SoftDeleteSong] Invalid ID format: %s", err)
		handleError(c, utils.ErrInvalidID)
		return
	}

	err = services.SoftDeleteSong(uint(id))
	if err != nil {
		logger.Error.Printf("[handlers.SoftDeleteSong] Error soft deleting song: %s", err)
		handleError(c, err)
		return
	}

	response := NewDefaultResponse("Song successfully deleted")
	c.JSON(http.StatusOK, response)
}

// HardDeleteSong godoc
// @Summary      Hard delete a song
// @Description  Permanently deletes a song by its unique ID from the database. This action cannot be undone.
// @Tags         Songs
// @Accept       json
// @Produce      json
// @Param        id   path    int     true  "Song ID"
// @Success      200  {object}  DefaultResponse   "Success"  "Song successfully hard deleted"
// @Failure      400  {object}  ErrorResponse  "Invalid ID format"
// @Failure      404  {object}  ErrorResponse  "Song not found"
// @Failure      500  {object}  ErrorResponse  "Internal server error"
// @Security     ApiKeyAuth
// @Router       /songs/hard/{id} [delete]
func HardDeleteSong(c *gin.Context) {
	ip := c.ClientIP()
	idParam := c.Param("id")

	logger.Info.Printf("[handlers.HardDeleteSong] Client IP: %s - Request to hard delete song by id: %s", ip, idParam)

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		logger.Error.Printf("[handlers.HardDeleteSong] Invalid ID format: %s", err)
		handleError(c, utils.ErrInvalidID)
		return
	}

	err = services.HardDeleteSong(uint(id))
	if err != nil {
		logger.Error.Printf("[handlers.HardDeleteSong] Error hard deleting song: %s", err)
		handleError(c, err)
		return
	}

	response := NewDefaultResponse("Song successfully hard deleted")
	c.JSON(http.StatusOK, response)
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
