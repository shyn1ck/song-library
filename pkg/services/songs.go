package service

import (
	"errors"
	"song-library/models"
	"song-library/pkg/repository"
)

func GetSongs(group, song string, page, limit int) ([]models.Song, error) {
	if group == "" && song == "" {
		return nil, errors.New("at least one of 'group' or 'song' must be provided")
	}

	if page <= 0 {
		return nil, errors.New("the 'page' parameter must be greater than 0")
	}
	if limit <= 0 || limit > 100 {
		return nil, errors.New("the 'limit' parameter must be between 1 and 100")
	}

	songs, err := repository.GetSongs(group, song, page, limit)
	if err != nil {
		return nil, err
	}

	return songs, nil
}
