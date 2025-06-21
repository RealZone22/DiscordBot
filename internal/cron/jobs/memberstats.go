package jobs

import (
	"github.com/RealZone22/DiscordBot/internal/handlers"
	"github.com/RealZone22/DiscordBot/pkg/utils"
)

func MemberStatsJob() {
	handlers.HandleMemberStats(utils.ConvertToSnowflake(utils.Config.DefaultGuildId))
}
