package commands

import (
	"fmt"
	"github.com/RealZone22/DiscordBot/internal/handlers"
	"github.com/RealZone22/DiscordBot/pkg/utils"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/json"
	"strings"
)

var warnPermission = json.NewNullable(discord.PermissionKickMembers)

var Warn = discord.SlashCommandCreate{
	Name:        "warn",
	Description: "Manage user warnings",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionSubCommand{
			Name:        "create",
			Description: "Create a warning for a user",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionUser{
					Name:        "member",
					Description: "The user to warn",
					Required:    true,
				},
				discord.ApplicationCommandOptionString{
					Name:        "reason",
					Description: "Reason for the warning",
					Required:    true,
				},
			},
		},
		discord.ApplicationCommandOptionSubCommand{
			Name:        "get",
			Description: "Get warnings for a user",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionUser{
					Name:        "member",
					Description: "The user to get warnings for",
					Required:    true,
				},
			},
		},
		discord.ApplicationCommandOptionSubCommand{
			Name:        "delete",
			Description: "Delete a specific warning",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionInt{
					Name:        "id",
					Description: "ID of the warning to delete",
					Required:    true,
				},
			},
		},
		discord.ApplicationCommandOptionSubCommand{
			Name:        "clear",
			Description: "Clear all warnings for a user",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionUser{
					Name:        "member",
					Description: "The user to clear warnings for",
					Required:    true,
				},
			},
		},
		discord.ApplicationCommandOptionSubCommand{
			Name:        "count",
			Description: "Get the number of warnings for a user",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionUser{
					Name:        "member",
					Description: "The user to get warning count for",
					Required:    true,
				},
			},
		},
	},
	DefaultMemberPermissions: &warnPermission,
}

func CreateWarnHandler(e *handler.CommandEvent) error {
	member := e.SlashCommandInteractionData().User("member")
	reason := e.SlashCommandInteractionData().String("reason")

	if err := e.DeferCreateMessage(false); err != nil {
		return err
	}

	err := handlers.CreateWarn(member.ID.String(), reason)
	if err != nil {
		_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			SetContent("Failed to create warn: " + err.Error()).
			Build())
		return err
	}

	_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
		AddEmbeds(discord.NewEmbedBuilder().
			SetTitle("Warn Created").
			SetDescription(fmt.Sprintf("Successfully warned **%s** for: %s", member.Username, reason)).
			SetColor(15548997).
			SetFooter(utils.Config.Embeds.FooterMessage, utils.Config.Embeds.FooterIcon).
			Build()).
		Build())

	return err
}

func GetWarnsHandler(e *handler.CommandEvent) error {
	member := e.SlashCommandInteractionData().User("member")

	if err := e.DeferCreateMessage(false); err != nil {
		return err
	}

	warns, err := handlers.GetWarns(member.ID.String())
	if err != nil {
		_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			SetContent("Failed to retrieve warns: " + err.Error()).
			Build())
		return err
	}

	if len(warns) == 0 {
		_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			SetContent(fmt.Sprintf("No warnings found for **%s**.", member.Username)).
			Build())
		return err
	}

	var warnList []string
	for _, warn := range warns {
		warnList = append(warnList, fmt.Sprintf("ID: %d, Created At: %s, Reason: %s", warn.ID, "<t:"+fmt.Sprintf("%d", warn.CreatedAt.Unix())+">", warn.Reason))
	}

	content := fmt.Sprintf("Warnings for **%s**:\n%s", member.Username, strings.Join(warnList, "\n"))

	_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
		AddEmbeds(discord.NewEmbedBuilder().
			SetTitlef("All Warnings from %s", member.Username).
			SetDescription(content).
			SetColor(15548997).
			SetFooter(utils.Config.Embeds.FooterMessage, utils.Config.Embeds.FooterIcon).
			Build()).
		Build())

	return err
}

func DeleteWarnHandler(e *handler.CommandEvent) error {
	warnID := e.SlashCommandInteractionData().Int("id")

	if err := e.DeferCreateMessage(false); err != nil {
		return err
	}

	err := handlers.DeleteWarn(int64(warnID))
	if err != nil {
		_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			SetContent("Failed to delete warn: " + err.Error()).
			Build())
		return err
	}

	_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
		AddEmbeds(discord.NewEmbedBuilder().
			SetTitle("Warn Deleted").
			SetDescription(fmt.Sprintf("Successfully deleted warn with ID: %d", warnID)).
			SetColor(15548997).
			SetFooter(utils.Config.Embeds.FooterMessage, utils.Config.Embeds.FooterIcon).
			Build()).
		Build())

	return err
}

func ClearWarnsHandler(e *handler.CommandEvent) error {
	member := e.SlashCommandInteractionData().User("member")

	if err := e.DeferCreateMessage(false); err != nil {
		return err
	}

	err := handlers.ClearWarns(member.ID.String())
	if err != nil {
		_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			SetContent("Failed to clear warns: " + err.Error()).
			Build())
		return err
	}

	_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
		AddEmbeds(discord.NewEmbedBuilder().
			SetTitle("Warns Cleared").
			SetDescription(fmt.Sprintf("Successfully cleared all warnings for **%s**.", member.Username)).
			SetColor(15548997).
			SetFooter(utils.Config.Embeds.FooterMessage, utils.Config.Embeds.FooterIcon).
			Build()).
		Build())

	return err
}

func CountWarnsHandler(e *handler.CommandEvent) error {
	member := e.SlashCommandInteractionData().User("member")

	if err := e.DeferCreateMessage(false); err != nil {
		return err
	}

	count, err := handlers.GetWarnCount(member.ID.String())
	if err != nil {
		_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
			SetContent("Failed to get warn count: " + err.Error()).
			Build())
		return err
	}

	_, err = e.UpdateInteractionResponse(discord.NewMessageUpdateBuilder().
		AddEmbeds(discord.NewEmbedBuilder().
			SetTitle("Warn Count").
			SetDescription(fmt.Sprintf("**%s** has %d warning(s).", member.Username, count)).
			SetColor(15548997).
			SetFooter(utils.Config.Embeds.FooterMessage, utils.Config.Embeds.FooterIcon).
			Build()).
		Build())

	return err
}
