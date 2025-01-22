package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"song-library/configs"
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

func AddSong(newSongRequest models.NewSongRequest) (*models.Song, error) {
	song := &models.Song{
		Group:       newSongRequest.Group,
		Song:        newSongRequest.Song,
		ReleaseDate: "",
		Text:        "",
		Link:        "",
	}

	apiURL := fmt.Sprintf(configs.AppSettings.AppParams.ApiURL, song.Group, song.Song)
	resp, err := http.Get(apiURL)
	if err != nil {
		logger.Error.Printf("[services.AddSong] Failed to fetch song info: %s", err)
		return nil, utils.ErrAPIRequestFailed
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error.Printf("[services.AddSong] Failed to close body reader: %s", err)
		}
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {
		var songDetail models.SongDetail
		if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
			logger.Error.Printf("[services.AddSong] Failed to decode response: %s", err)
			return nil, utils.ErrInvalidResponse
		}

		song.ReleaseDate = songDetail.ReleaseDate
		song.Text = songDetail.Text
		song.Link = songDetail.Link
	}

	if err := repository.AddSong(song); err != nil {
		return nil, err
	}

	return song, nil
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

func HardDeleteSong(id uint) (err error) {
	song, err := repository.GetSongByID(id)
	if err != nil {
		return err
	}
	if song == nil {
		logger.Error.Printf("[services.HardDeleteSong]: Song does not exist")
		return utils.ErrSongNotFound
	}
	return repository.HardDeleteSong(id)
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
