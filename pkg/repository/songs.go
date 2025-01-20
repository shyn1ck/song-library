package repository

import (
	"song-library/db"
	"song-library/models"
)

func GetSongs(group string, song string, page int, limit int) (songs []models.Song, err error) {

	err = db.GetDBConn().Model(&songs).Find(&songs).Error

	return songs, err
}
