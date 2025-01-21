package utils

import "errors"

var (
	ErrSongNotFound                 = errors.New("ErrSongNotFound")
	ErrSongAlreadyExists            = errors.New("ErrSongAlreadyExists")
	ErrInvalidSongData              = errors.New("ErrInvalidSongData")
	ErrSongDeleteFailed             = errors.New("ErrSongDeleteFailed")
	ErrSongUpdateFailed             = errors.New("ErrSongUpdateFailed")
	ErrInvalidGroup                 = errors.New("ErrInvalidGroup")
	ErrInvalidSongTitle             = errors.New("ErrInvalidSongTitle")
	ErrInvalidReleaseDate           = errors.New("ErrInvalidReleaseDate")
	ErrInvalidText                  = errors.New("ErrInvalidText")
	ErrInvalidLink                  = errors.New("ErrInvalidLink")
	ErrFailedToFetchSongInfoFromAPI = errors.New("ErrFailedToFetchSongInfoFromAPI")
	ErrDatabaseConnectionFailed     = errors.New("ErrDatabaseConnectionFailed")
	ErrSongNotFoundInDatabase       = errors.New("ErrSongNotFoundInDatabase")
	ErrInvalidPaginationParams      = errors.New("ErrInvalidPaginationParams")
	ErrFailedToParseJSON            = errors.New("ErrFailedToParseJSON")
	ErrUnexpectedError              = errors.New("ErrUnexpectedError")
	ErrFailedToGenerateSwagger      = errors.New("ErrFailedToGenerateSwagger")
	ErrMissingRequiredField         = errors.New("ErrMissingRequiredField")
	MissingParameters               = errors.New("MissingParameters")
	ErrInvalidID                    = errors.New("ErrInvalidID")
	ErrInvalidRequestBody           = errors.New("ErrInvalidRequestBody")
)
