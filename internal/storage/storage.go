package storage

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"voice-out-be/internal/config"
)

var onceDb sync.Once

var instance *gorm.DB

type Store struct {
	logger *zap.Logger
}

func New(cfg *config.AppConfig) *Store {
	_ = InitDB(cfg)

	return &Store{
		logger: cfg.Logger(),
	}
}

func InitDB(cfg *config.AppConfig) *gorm.DB {
	onceDb.Do(func() {
		databaseConfig := cfg.PgConfig()

		db, err := gorm.Open(postgres.Open(databaseConfig.GetDBConnectionString()), &gorm.Config{})
		if err != nil {
			cfg.Logger().Fatal("Could not connect to database: %s", zap.Error(err))
		}

		cfg.Logger().Info("Successfully Connected to Database")

		instance = db
	})

	return instance
}
