package database

import (
	"errors"
	"moon/pkg/config"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var (
	UnsupportDatabase         = errors.New("unsupport database")
	CouldNotConnectToDatabase = errors.New("could not connect to database")
)

func NewDatabase(cfg *config.Config, logger *zap.Logger) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	var gormConfig = gorm.Config{
		Logger: zapgorm2.New(logger),
	}

	switch cfg.DatabaseDriver {
	case "postgres":
		db, err = gorm.Open(postgres.Open(cfg.DSN), &gormConfig)
	default:
		return nil, UnsupportDatabase
	}
	if err != nil {
		return nil, CouldNotConnectToDatabase
	}
	return db, nil
}
