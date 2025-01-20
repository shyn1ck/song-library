package db

import (
	"errors"
	"fmt"
	"song-library/logger"
)

func Migrate() (err error) {
	if dbConn == nil {

		logger.Error.Printf("[db.Migrate]: database connection is not initialized: %v", err)
		return errors.New("database connection is not initialized")

	}

	migrateModels := []interface{}{}

	for _, model := range migrateModels {
		err := dbConn.AutoMigrate(model)
		if err != nil {
			return fmt.Errorf("failed to migrate %T: %v", model, err)
		}
	}

	return nil
}
