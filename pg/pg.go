package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
)

func DB(gormConfig *GormConfig, cfg ...ConnectionConfig) (*gorm.DB, error) {

	if instance != nil {
		return instance, nil
	}

	db, err := initDB(gormConfig, cfg...)
	{
		if err != nil {
			return nil, err
		}

		instance = db
	}

	return instance, nil
}

func Must(gormConfig *GormConfig, cfg ...ConnectionConfig) *gorm.DB {
	db, err := DB(gormConfig, cfg...)
	{
		if err != nil {
			log.Fatalf("Failed to initialize database: %v", err)
		}
	}

	return db
}

func initDB(gormConfig *GormConfig, cfg ...ConnectionConfig) (*gorm.DB, error) {

	config := defaultConfig
	{
		if len(cfg) > 0 {
			config = cfg[0]
		}
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode, config.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	{
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
	}

	return db, nil
}
