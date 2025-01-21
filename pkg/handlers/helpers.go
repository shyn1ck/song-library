package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"song-library/logger"
	"song-library/utils"
)

func handleError(c *gin.Context, err error) {
	var statusCode int
	var errorResponse ErrorResponse

	switch {
	case errors.Is(err, utils.ErrSongNotFound),
		errors.Is(err, utils.ErrSongAlreadyExists),
		errors.Is(err, utils.ErrInvalidSongData),
		errors.Is(err, utils.ErrSongDeleteFailed),
		errors.Is(err, utils.ErrSongUpdateFailed),
		errors.Is(err, utils.ErrInvalidGroup),
		errors.Is(err, utils.ErrInvalidSongTitle),
		errors.Is(err, utils.ErrInvalidReleaseDate),
		errors.Is(err, utils.ErrInvalidText),
		errors.Is(err, utils.ErrInvalidLink),
		errors.Is(err, utils.ErrFailedToFetchSongInfoFromAPI),
		errors.Is(err, utils.ErrDatabaseConnectionFailed),
		errors.Is(err, utils.ErrSongNotFoundInDatabase),
		errors.Is(err, utils.ErrInvalidPaginationParams),
		errors.Is(err, utils.ErrFailedToParseJSON),
		errors.Is(err, utils.ErrUnexpectedError),
		errors.Is(err, utils.ErrInvalidID),
		errors.Is(err, utils.ErrFailedToGenerateSwagger),
		errors.Is(err, utils.ErrInvalidRequestBody),
		errors.Is(err, utils.ErrInvalidRequestParameter),
		errors.Is(err, utils.ErrMissingRequiredField):
		statusCode = http.StatusBadRequest
		errorResponse = NewErrorResponse(err.Error())

	case errors.Is(err, utils.ErrSongNotFoundInDatabase):
		statusCode = http.StatusNotFound
		errorResponse = NewErrorResponse(err.Error())

	case errors.Is(err, utils.ErrSongDeleteFailed),
		errors.Is(err, utils.ErrSongUpdateFailed):
		statusCode = http.StatusInternalServerError
		errorResponse = NewErrorResponse(err.Error())

	default:
		logger.Error.Printf("Standard error occurred: %v", err)
		statusCode = http.StatusInternalServerError
		errorResponse = NewErrorResponse(utils.ErrUnexpectedError.Error())
	}

	c.JSON(statusCode, errorResponse)
}
