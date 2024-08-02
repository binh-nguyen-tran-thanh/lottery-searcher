package jobs

import (
	"backend/internal/core/port"
	"backend/internal/core/util"

	"github.com/robfig/cron/v3"
)

type CronJob struct {
	config     util.Config
	repository port.Repository
	logger     port.Logger
	cron       *cron.Cron
	channel    chan int
}

func NewJob(config util.Config, repository port.Repository, logger port.Logger, channel chan int) port.CronJob {
	c := cron.New()

	cronJob := CronJob{
		config:     config,
		repository: repository,
		cron:       c,
		logger:     logger,
	}

	for _, timeSchedule := range config.CronJob.Schedule {
		if _, err := cronJob.cron.AddFunc(timeSchedule, cronJob.StartSyncResult); err != nil {
			cronJob.logger.Fatal().Msgf("Fail to initial Job. Reason: %v", err.Error())
		}
		cronJob.logger.Info().Msgf("Added job for %s", timeSchedule)
	}

	return &cronJob
}

func (c *CronJob) Start() error {
	c.cron.Start()

	c.logger.Info().Msg("Staring cron jobs")

	return nil
}

func (c *CronJob) Stop() error {
	c.cron.Stop()

	c.logger.Info().Msg("Stopping cron jobs")

	return nil
}
