package repository

import (
	"errors"
	"gorm.io/gorm"
	"song-library/db"
	"song-library/models"
	"song-library/utils"
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
