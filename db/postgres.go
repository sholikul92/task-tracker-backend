package db

import (
	"fmt"
	"task-tracker/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB(dbConfig *models.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", dbConfig.HOST, dbConfig.USER, dbConfig.PASSWORD, dbConfig.DBNAME, dbConfig.PORT)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	return db, nil
}
