package commands

import (
	"github.com/RealZone22/DiscordBot/pkg/utils"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/json"
	"github.com/disgoorg/snowflake/v2"
)

var purgePermission = json.NewNullable(discord.PermissionManageMessages)

var Purge = discord.SlashCommandCreate{
	Name:        "purge",
	Description: "Purge messages from the channel",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionInt{
			Name:        "amount",
			Description: "Amount of messages to purge",
			Required:    true,
		},
	},
	DefaultMemberPermissions: &purgePermission,
}

func PurgeHandler(e *handler.CommandEvent) error {
	amount := e.SlashCommandInteractionData().Int("amount")
	if amount < 1 || amount > 100 {
		return e.CreateMessage(discord.NewMessageCreateBuilder().
			SetContent("Please provide a valid amount of messages to purge (1-100)").
			SetEphemeral(true).Build())
	}

	if err := e.DeferCreateMessage(true); err != nil {
		return err
	}

	messages, err := e.Client().Rest().GetMessages(e.Channel().ID(), 0, 0, 0, amount)
	if err != nil {
		_, err := e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			SetContent("Failed to fetch messages.").Build())
		return err
	}

	var messageIDs []snowflake.ID
	for _, msg := range messages {
		messageIDs = append(messageIDs, msg.ID)
	}

	if err := e.Client().Rest().BulkDeleteMessages(e.Channel().ID(), messageIDs); err != nil {
		_, err := e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			SetContent("Failed to delete messages.").Build())
		return err
	}

	_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
		AddEmbeds(discord.NewEmbedBuilder().
			SetTitle("Messages Purged").
			SetDescriptionf("Successfully purged %d messages from this channel.", len(messageIDs)).
			SetColor(15548997).
			SetFooter(utils.Config.Embeds.FooterMessage, utils.Config.Embeds.FooterIcon).
			Build()).
		Build())

	return err
}
