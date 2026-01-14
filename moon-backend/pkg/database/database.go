package database

import (
	"errors"
	"moon/pkg/config"
	"moon/pkg/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var (
	UnsupportDatabase         = errors.New("unsupport database")
	CouldNotConnectToDatabase = errors.New("could not connect to database")
)

func New(cfg *config.Configuration) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	var gormConfig = gorm.Config{
		Logger: zapgorm2.New(log.Get()),
	}

	switch cfg.DatabaseConfiguration.Driver {
	case "postgres":
		db, err = gorm.Open(postgres.Open(cfg.DatabaseConfiguration.DSN), &gormConfig)
	default:
		return nil, UnsupportDatabase
	}

	if err != nil {
		return nil, CouldNotConnectToDatabase
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if cfg.DatabaseConfiguration.MaxOpenConnections > 0 {
		sqlDB.SetMaxOpenConns(cfg.DatabaseConfiguration.MaxOpenConnections)
	}
	if cfg.DatabaseConfiguration.MaxIdleConnections > 0 {
		sqlDB.SetMaxIdleConns(cfg.DatabaseConfiguration.MaxIdleConnections)
	}
	return db, nil
}
