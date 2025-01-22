package repository

import (
	"errors"
	"gorm.io/gorm"
	"song-library/db"
	"song-library/logger"
	"song-library/models"
	"song-library/utils"
)

func GetInfoByGroup(group string) ([]models.SongDetail, error) {
	var songDetails []models.SongDetail
	if err := db.GetDBConn().Where("\"group\" = ?", group).Find(&songDetails).Error; err != nil {
		logger.Error.Printf("[repository.GetInfoByGroup]: Error finding songs: %s\n", err.Error())
		return nil, err
	}
	return songDetails, nil
}

func GetInfoBySong(song string) (bool, error) {
	var count int64
	err := db.GetDBConn().Model(&models.SongDetail{}).Where("song = ?", song).Count(&count).Error
	if err != nil {
		logger.Error.Printf("[repository.GetInfoBySong]: Error finding songs: %s\n", err.Error())
		return false, utils.ErrDatabaseConnectionFailed
	}
	return count > 0, nil
}

func GetSongDetail(group, song string) (models.SongDetail, error) {
	var songDetail models.SongDetail
	err := db.GetDBConn().Model(&models.SongDetail{}).
		Where("\"group\" = ? AND song = ?", group, song).
		First(&songDetail).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.SongDetail{}, utils.ErrSongNotFound
		}
		logger.Error.Printf("[repository.GetSongDetail]: Error finding song details: %s\n", err.Error())
		return models.SongDetail{}, utils.ErrDatabaseConnectionFailed
	}
	return songDetail, nil
}
