package main

import (
	voiceout "voice-out-be"
	"voice-out-be/internal/config"
)

func main() {

	cfg := config.InitConfig()

	server := voiceout.NewServer(cfg)

	err := server.Start(cfg.ServerAddress())

	if err != nil {
		cfg.Logger().Fatal("Port already used")
	}
}
