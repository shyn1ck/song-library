package service

import (
	"song-library/logger"
	"song-library/models"
	"song-library/pkg/repository"
	"song-library/utils"
)

func GetSongDetail(group, song string) (models.SongDetail, error) {
	groupExists, err := repository.GetInfoByGroup(group)
	if err != nil {
		logger.Error.Printf("[services.GetSongDetail]: Error checking song exists: %s", err.Error())
		return models.SongDetail{}, err
	}

	if !groupExists {
		return models.SongDetail{}, utils.ErrGroupNotFound
	}

	songExists, err := repository.GetInfoBySong(song)
	if err != nil {
		logger.Error.Printf("[services.GetSongDetail]: Error checking song: %s", err.Error())
		return models.SongDetail{}, err
	}

	if !songExists {
		return models.SongDetail{}, utils.ErrSongNotFound
	}

	songDetail, err := repository.GetSongDetail(group, song)
	if err != nil {
		logger.Error.Printf("[services.GetSongDetail]: Error getting song detail: %s", err.Error())
		return models.SongDetail{}, err
	}

	return songDetail, nil
}
