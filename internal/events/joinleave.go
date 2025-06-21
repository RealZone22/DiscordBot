package events

import (
	"github.com/RealZone22/DiscordBot/internal/handlers"
	"github.com/RealZone22/DiscordBot/pkg/utils"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func JoinEventHandler() bot.EventListener {
	return bot.NewListenerFunc(func(e *events.GuildMemberJoin) {
		if utils.Config.Events.Join.Enabled {
			_, err := e.Client().Rest().CreateMessage(utils.ConvertToSnowflake(utils.Config.Events.Join.ChannelID),
				discord.NewMessageCreateBuilder().
					AddEmbeds(discord.NewEmbedBuilder().
						SetTitlef("%s joined the server", e.Member.User.Username).
						SetDescriptionf("Welcome %s! We're glad to have you here.", e.Member.User.Mention()).
						SetColor(5763719).
						SetThumbnail(*e.Member.User.AvatarURL()).
						SetFooter(utils.Config.Embeds.FooterMessage, utils.Config.Embeds.FooterIcon).
						Build()).
					Build())
			if err != nil {
				utils.Logger.Error().Err(err).Msg("Failed to send join message")
				return
			}

			utils.Logger.Info().Str("user", e.Member.User.Username).Str("guild", e.GuildID.String()).Msg("joined the server")
		}

		handlers.HandleMemberStats(e.GuildID)
	})
}

func LeaveEventHandler() bot.EventListener {
	return bot.NewListenerFunc(func(e *events.GuildMemberLeave) {
		handlers.HandleMemberStats(e.GuildID)
	})
}
