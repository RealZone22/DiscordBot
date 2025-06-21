package cron

import (
	"github.com/RealZone22/DiscordBot/internal/cron/jobs"
	"github.com/RealZone22/DiscordBot/pkg/utils"
	"github.com/robfig/cron"
)

func RunCronJobs() error {
	c := cron.New()

	err := c.AddFunc("@hourly", jobs.MemberStatsJob)
	if err != nil {
		utils.Logger.Error().Err(err).Msg("Failed to add cron job")
		return err
	}

	c.Start()

	utils.Logger.Info().Int("jobs", len(c.Entries())).Msg("Cron jobs started")

	return nil
}
