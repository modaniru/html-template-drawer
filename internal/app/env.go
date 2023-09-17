package app

import (
	log "log/slog"

	"github.com/joho/godotenv"
)

func configureDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Warn(".env file not present, using already loaded(in system) envs")
	}
	log.Debug(".env file was successfully loaded")
}
