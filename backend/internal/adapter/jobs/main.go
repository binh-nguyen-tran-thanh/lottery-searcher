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
	isPureJob  bool
}

func NewJob(config util.Config, repository port.Repository, logger port.Logger) port.CronJob {
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

func NewPureJob(config util.Config, repository port.Repository, logger port.Logger) port.CronJob {
	c := cron.New()

	cronJob := CronJob{
		config:     config,
		repository: repository,
		cron:       c,
		logger:     logger,
		isPureJob:  true,
	}

	return &cronJob

}

func (c *CronJob) Start() error {
	c.logger.Info().Msg("Staring cron jobs")

	if c.isPureJob {
		c.StartSyncResult()
	} else {
		c.cron.Start()
	}

	return nil
}

func (c *CronJob) Stop() error {
	c.logger.Info().Msg("Stopping cron jobs")

	if !c.isPureJob {
		c.cron.Stop()
	}

	return nil
}
