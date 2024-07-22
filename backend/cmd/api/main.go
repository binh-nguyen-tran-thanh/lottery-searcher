package main

import (
	"backend/internal/adapter/handler"
	"backend/internal/adapter/logger"
	"backend/internal/adapter/repository"
	"backend/internal/core/service"
	"backend/internal/core/util"
	"fmt"
)

func main() {
	config, err := util.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Fail to load the config: %s", err))
	}

	logger := logger.NewLogger(config)

	logger.Info().Msg("Initialed Logger")

	service, err := service.NewService(config, logger)

	if err != nil {
		logger.Fatal().Err(err).Msg("Fail to initialize services")
	}
	logger.Info().Msg("Initialed Services")

	repo, err := repository.NewRepository(config, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Fail to initialize repository")
	}
	logger.Info().Msg("Initialed Repository")

	server := handler.NewServer(config, service, logger, repo)

	if err = server.Start(); err != nil {
		logger.Fatal().Err(err).Msg("failed to load service")
	}
}
