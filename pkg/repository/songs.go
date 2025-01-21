package repository

import (
	"errors"
	"gorm.io/gorm"
	"song-library/db"
	"song-library/models"
	"song-library/utils"
	"strings"
)

func GetSongs(group, song string, page, limit int) ([]models.Song, error) {
	var songs []models.Song
	offset := (page - 1) * limit

	query := db.GetDBConn().Model(&songs)

	if group != "" {
		query = query.Where("\"group\" = ?", group)
	}
	if song != "" {
		query = query.Where("song = ?", song)
	}

	err := query.Offset(offset).Limit(limit).Find(&songs).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, utils.ErrDatabaseConnectionFailed
	}

	if len(songs) == 0 {
		return nil, nil
	}

	return songs, nil
}

func GetLyrics(songName string, page, limit int) (verses []string, err error) {
	var song models.Song
	err = db.GetDBConn().Where("song = ?", songName).First(&song).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrSongNotFound
		}
		return nil, utils.ErrDatabaseConnectionFailed
	}

	verses = strings.Split(song.Text, "\n\n")

	start := (page - 1) * limit
	end := start + limit

	if start >= len(verses) {
		return nil, nil
	}

	if end > len(verses) {
		end = len(verses)
	}

	return verses[start:end], nil
}

func GetLyricsByText(searchText string, page, limit int) ([]string, error) {
	var songs []models.Song
	err := db.GetDBConn().Where("text LIKE ?", "%"+searchText+"%").Find(&songs).Error
	if err != nil {
		return nil, utils.ErrDatabaseConnectionFailed
	}

	if len(songs) == 0 {
		return nil, utils.ErrSongNotFound
	}

	verses := strings.Split(songs[0].Text, "\n\n")

	start := (page - 1) * limit
	end := start + limit

	if start >= len(verses) {
		return nil, nil
	}

	if end > len(verses) {
		end = len(verses)
	}

	return verses[start:end], nil
}
