package service

import (
	"song-library/models"
	"song-library/pkg/repository"
	"song-library/utils"
)

func GetSongs(group, song string, page, limit int) (songs []models.Song, err error) {
	if page <= 0 || limit <= 0 || limit > 100 {
		return nil, utils.ErrInvalidPaginationParams
	}

	songs, err = repository.GetSongs(group, song, page, limit)
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func GetLyrics(song string, page int, limit int) ([]string, error) {
	if page <= 0 || limit <= 0 || limit > 100 {
		return nil, utils.ErrInvalidPaginationParams
	}
	verses, err := repository.GetLyrics(song, page, limit)
	if err != nil {
		return nil, err
	}

	return verses, nil
}

func GetLyricsByText(searchText string, page int, limit int) ([]string, error) {
	if page <= 0 || limit <= 0 || limit > 100 {
		return nil, utils.ErrInvalidPaginationParams
	}

	verses, err := repository.GetLyricsByText(searchText, page, limit)
	if err != nil {
		return nil, err
	}

	return verses, nil
}
