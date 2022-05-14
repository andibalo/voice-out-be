package config

import (
	"go.uber.org/zap"
	"os"
	logger2 "voice-out-be/internal/logger"
)

const envVarEnvironment = "ENV"

func InitConfig() *AppConfig {
	logger := logger2.NewMainLoggerSingleton()

	return &AppConfig{
		logger:      logger,
		environment: os.Getenv(envVarEnvironment),
	}
}

type AppConfig struct {
	logger      *zap.Logger
	environment string
}

func (a *AppConfig) Logger() *zap.Logger {
	return a.logger
}

func (a *AppConfig) ServerAddress() string {
	return "8080"
}
