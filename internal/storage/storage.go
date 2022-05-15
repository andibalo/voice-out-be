package storage

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"voice-out-be/internal/config"
	"voice-out-be/internal/dto"
	"voice-out-be/internal/model"
	"voice-out-be/internal/storage/repositories"
)

var onceDb sync.Once

var instance *gorm.DB

type Store struct {
	logger         *zap.Logger
	userRepository UserRepository
}

func New(cfg *config.AppConfig) *Store {
	db := InitDB(cfg)

	migrateDB(db)

	userRepo := repositories.NewUserRepository(db)

	return &Store{
		logger:         cfg.Logger(),
		userRepository: userRepo,
	}
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
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

type Storage interface {
	CreateUser(in *dto.RegisterUser) (*model.User, error)
}

type UserRepository interface {
	SaveUser(user *model.User) error
}
