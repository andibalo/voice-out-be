package storage

import (
	"sync"
	"voice-out-be/internal/config"
	"voice-out-be/internal/dto"
	"voice-out-be/internal/model"
	"voice-out-be/internal/storage/repositories"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var onceDb sync.Once

var instance *gorm.DB

type Store struct {
	logger         *zap.Logger
	userRepository UserRepository
	postRepository PostRepository
}

func New(cfg *config.AppConfig) *Store {
	db := InitDB(cfg)

	migrateDB(db)

	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepositroy(db)

	return &Store{
		logger:         cfg.Logger(),
		userRepository: userRepo,
		postRepository: postRepo,
	}
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Post{})
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
	FindUserByEmail(email string) (*model.User, error)
	CreatePost(in *dto.CreatePost) (*model.Post, error)
}

type UserRepository interface {
	SaveUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
}

type PostRepository interface {
	SavePost(post *model.Post) error
}
