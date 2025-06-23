package commands

import (
	"github.com/RealZone22/DiscordBot/internal/handlers"
	"github.com/RealZone22/DiscordBot/pkg/utils"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var Ticket = discord.SlashCommandCreate{
	Name:        "ticket",
	Description: "Manage tickets",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionSubCommand{
			Name:        "embed",
			Description: "Create a ticket embed in the current channel",
			Options:     nil,
		},
		discord.ApplicationCommandOptionSubCommand{
			Name:        "create",
			Description: "Create a new ticket",
			Options:     nil,
		},
		discord.ApplicationCommandOptionSubCommand{
			Name:        "close",
			Description: "Close the current ticket",
			Options:     nil,
		},
		discord.ApplicationCommandOptionSubCommand{
			Name:        "addmember",
			Description: "Add a member to the ticket",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionUser{
					Name:        "member",
					Description: "The member to add to the ticket",
					Required:    true,
				},
			},
		},
		discord.ApplicationCommandOptionSubCommand{
			Name:        "removemember",
			Description: "Remove a member from the ticket",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionUser{
					Name:        "member",
					Description: "The user to remove from the ticket",
					Required:    true,
				},
			},
		},
	},
	DefaultMemberPermissions: nil,
}

func TicketEmbedHandler(e *handler.CommandEvent) error {
	if !e.Member().Permissions.Has(discord.PermissionAdministrator) {
		return e.CreateMessage(discord.NewMessageCreateBuilder().
			SetContent("You do not have permission to use this command.").
			SetEphemeral(true).
			Build())
	}

	return e.CreateMessage(discord.NewMessageCreateBuilder().
		AddEmbeds(discord.NewEmbedBuilder().
			SetTitle("Ticket System").
			SetDescription("Click the button below to create a ticket.").
			SetColor(3447003).
			SetFooter(utils.Config.Embeds.FooterMessage, utils.Config.Embeds.FooterIcon).
			Build()).
		AddActionRow(discord.NewSuccessButton("Create Ticket", "create_ticket")).
		Build())
}

func CreateTicketHandler(e *handler.CommandEvent) error {
	channel, err := handlers.CreateTicket(e.User())
	if err != nil {
		return e.CreateMessage(discord.NewMessageCreateBuilder().
			SetContent("Failed to create ticket: " + err.Error()).
			SetEphemeral(true).
			Build())
	}

	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetContent("Your ticket has been created. You can view it here: <#" + channel.ID().String() + ">").
		SetEphemeral(true).
		Build())
}

func CloseTicketHandler(e *handler.CommandEvent) error {
	if err := handlers.CloseTicket(e.Channel().ID()); err != nil {
		return e.CreateMessage(discord.NewMessageCreateBuilder().
			SetContent("Failed to close ticket: " + err.Error()).
			Build())
	}

	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetContent("The ticket has been closed.").
		Build())
}

func AddMemberToTicketHandler(e *handler.CommandEvent) error {
	member := e.SlashCommandInteractionData().User("member")
	if err := handlers.AddMemberToTicket(e.Channel().ID(), member.ID); err != nil {
		return e.CreateMessage(discord.NewMessageCreateBuilder().
			SetContent("Failed to add member to ticket: " + err.Error()).
			Build())
	}

	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetContent("Successfully added <@" + member.ID.String() + "> to the ticket.").
		Build())
}

func RemoveMemberFromTicketHandler(e *handler.CommandEvent) error {
	member := e.SlashCommandInteractionData().User("member")
	if err := handlers.RemoveMemberFromTicket(e.Channel().ID(), member.ID); err != nil {
		return e.CreateMessage(discord.NewMessageCreateBuilder().
			SetContent("Failed to remove member from ticket: " + err.Error()).
			Build())
	}

	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetContent("Successfully removed <@" + member.ID.String() + "> from the ticket.").
		Build())
}
