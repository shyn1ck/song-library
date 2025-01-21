package service

import (
	"song-library/logger"
	"song-library/models"
	"song-library/pkg/repository"
	"song-library/utils"
	"time"
)

func GetSongs(group, song string, page, limit int) (songs []models.Song, err error) {
	if page <= 0 || limit <= 0 || limit > 100 {
		logger.Error.Printf("services.GetSongs: page %d or limit %d", page, limit)
		return nil, utils.ErrInvalidPaginationParams
	}

	songs, err = repository.GetSongs(group, song, page, limit)
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func GetSongByID(id uint) (song *models.Song, err error) {
	song, err = repository.GetSongByID(id)
	if err != nil {
		return nil, err
	}

	if song == nil {
		return nil, utils.ErrSongNotFound
	}

	return song, nil
}

func UpdateSong(id uint, songUpdate *models.Song) error {
	existingSong, err := repository.GetSongByID(id)
	if err != nil {
		logger.Error.Printf("[services.UpdateSong]: Error getting existing song: %v", err)
		return err
	}

	if existingSong == nil {
		logger.Error.Printf("[services.UpdateSong]: Song does not exist")
		return utils.ErrSongNotFound
	}

	existingSong.Group = songUpdate.Group
	existingSong.Song = songUpdate.Song
	existingSong.ReleaseDate = songUpdate.ReleaseDate
	existingSong.Text = songUpdate.Text
	existingSong.Link = songUpdate.Link
	existingSong.UpdatedAt = time.Now()
	if err := repository.UpdateSong(existingSong); err != nil {
		return err
	}
	return nil
}

func SoftDeleteSong(id uint) error {
	song, err := repository.GetSongByID(id)
	if err != nil {
		return err
	}

	if song == nil {
		logger.Error.Printf("[services.SoftDeleteSong]: Song does not exist")
		return utils.ErrSongNotFound
	}
	return repository.SoftDeleteSong(id)
}

func GetLyrics(song string, page int, limit int) ([]string, error) {
	if page <= 0 || limit <= 0 || limit > 100 {
		logger.Error.Printf("services.GetLyrics: page %d or limit %d", page, limit)
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
		logger.Error.Printf("services.GetLyrics: page %d or limit %d", page, limit)
		return nil, utils.ErrInvalidPaginationParams
	}

	verses, err := repository.GetLyricsByText(searchText, page, limit)
	if err != nil {
		return nil, err
	}

	return verses, nil
}
