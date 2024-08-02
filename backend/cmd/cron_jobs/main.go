package main

import (
	"backend/internal/adapter/jobs"
	"backend/internal/adapter/logger"
	"backend/internal/adapter/repository"
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

	repo, err := repository.NewRepository(config, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Fail to initialize repository")
	}
	logger.Info().Msg("Initialed Repository")

	chanelLine := make(chan int)

	cron := jobs.NewJob(config, repo, logger, chanelLine)

	if err := cron.Start(); err != nil {
		logger.Fatal().Err(err).Msg("Fail to initialize cron job")
	}

	logger.Info().Msg("Initialed Cron Job")

	defer cron.Stop()

	<-chanelLine
	logger.Info().Msg("Finished sync result")
}
