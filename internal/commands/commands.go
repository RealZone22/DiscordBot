package commands

import (
	"github.com/RealZone22/DiscordBot/pkg/utils"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/snowflake/v2"
)

var Commands = []discord.ApplicationCommandCreate{
	Ping,
	UserInfo,
	Purge,
}

func Register() {
	h := handler.New()
	h.Command("/ping", PingHandler)
	h.Command("/userinfo", UserInfoHandler)
	h.Command("/purge", PurgeHandler)

	utils.Client.AddEventListeners(h)

	err := handler.SyncCommands(utils.Client, Commands, []snowflake.ID{utils.ConvertToSnowflake(utils.Config.DefaultGuildId)})
	if err != nil {
		utils.Logger.Error().Err(err).Msg("Failed to sync commands")
		return
	}
}
