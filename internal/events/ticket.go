package events

import (
	"github.com/RealZone22/DiscordBot/internal/handlers"
	"github.com/RealZone22/DiscordBot/pkg/utils"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func TicketEmbedHandler() bot.EventListener {
	return bot.NewListenerFunc(func(e *events.ComponentInteractionCreate) {
		if e.ButtonInteractionData().CustomID() == "create_ticket" {
			if !utils.Config.Ticket.Enabled {
				return
			}

			channel, err := handlers.CreateTicket(e.User())
			if err != nil {
				err := e.CreateMessage(discord.NewMessageCreateBuilder().
					SetContent("Failed to create ticket: " + err.Error()).
					SetEphemeral(true).
					Build())
				if err != nil {
					return
				}
				return
			}

			err = e.CreateMessage(discord.NewMessageCreateBuilder().
				SetContent("Your ticket has been created. You can view it here: <#" + channel.ID().String() + ">").
				SetEphemeral(true).
				Build())
			if err != nil {
				return
			}
		}
	})
}
