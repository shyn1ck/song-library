package repository

import (
	"song-library/db"
	"song-library/logger"
	"song-library/models"
	"song-library/utils"
)

func GetInfoByGroup(group string) (bool, error) {
	var count int64
	err := db.GetDBConn().Model(&models.SongDetail{}).Where("group = ? AND deleted_at IS NULL", group).Count(&count).Error
	if err != nil {
		logger.Error.Printf("[repository.GetInfoByGroup]: Error finding songs: %s\n", err.Error())
		return false, utils.ErrDatabaseConnectionFailed
	}

	return count > 0, nil
}

func GetInfoBySong(song string) (bool, error) {
	var count int64
	err := db.GetDBConn().Model(&models.SongDetail{}).Where("song = ? AND deleted_at IS NULL", song).Count(&count).Error
	if err != nil {
		logger.Error.Printf("[repository.GetInfoBySong]: Error finding songs: %s\n", err.Error())
		return false, utils.ErrDatabaseConnectionFailed
	}

	return count > 0, nil
}
