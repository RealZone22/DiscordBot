package commands

import (
	"github.com/RealZone22/DiscordBot/pkg/utils"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var Ping = discord.SlashCommandCreate{
	Name:                     "ping",
	Description:              "Check the bot's latency",
	Options:                  nil,
	DefaultMemberPermissions: nil,
}

func PingHandler(e *handler.CommandEvent) error {
	latency := e.Client().Gateway().Latency()

	return e.CreateMessage(discord.NewMessageCreateBuilder().
		AddEmbeds(discord.NewEmbedBuilder().
			SetTitle("Pong!").
			SetDescriptionf("Latency: %d ms", latency.Milliseconds()).
			SetColor(3447003).
			SetFooter(utils.Config.Embeds.FooterMessage, utils.Config.Embeds.FooterIcon).
			Build()).
		Build())
}
