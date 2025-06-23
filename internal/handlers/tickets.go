package handlers

import (
	"errors"
	"github.com/RealZone22/DiscordBot/pkg/utils"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
	"strings"
)

func CreateTicket(user discord.User) (discord.GuildChannel, error) {
	channels, err := utils.Client.Rest().GetGuildChannels(utils.ConvertToSnowflake(utils.Config.DefaultGuildId))
	if err != nil {
		return nil, err
	}
	for _, channel := range channels {
		if channel.Name() == "ticket-"+user.Username {
			return nil, errors.New("you already have an open ticket")
		}
	}

	channel, err := utils.Client.Rest().CreateGuildChannel(utils.ConvertToSnowflake(utils.Config.DefaultGuildId), discord.GuildTextChannelCreate{
		Name:     "ticket-" + user.Username,
		ParentID: utils.ConvertToSnowflake(utils.Config.Ticket.CategoryID),
		PermissionOverwrites: discord.PermissionOverwrites{
			discord.MemberPermissionOverwrite{
				UserID: user.ID,
				Allow:  discord.PermissionViewChannel | discord.PermissionSendMessages | discord.PermissionReadMessageHistory,
			},
			discord.RolePermissionOverwrite{
				RoleID: utils.ConvertToSnowflake(utils.Config.Ticket.SupportRole),
				Allow:  discord.PermissionViewChannel | discord.PermissionSendMessages | discord.PermissionReadMessageHistory | discord.PermissionManageMessages,
			},
			discord.RolePermissionOverwrite{
				RoleID: utils.ConvertToSnowflake(utils.Config.DefaultGuildId),
				Deny:   discord.PermissionViewChannel | discord.PermissionSendMessages | discord.PermissionReadMessageHistory,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	_, err = utils.Client.Rest().CreateMessage(channel.ID(), discord.NewMessageCreateBuilder().
		AddEmbeds(discord.NewEmbedBuilder().
			SetTitle("Ticket Information").
			SetDescription("Thank you for creating a ticket. Please provide details about your issue, and a support member will be with you shortly.").
			SetColor(3447003).
			SetFooter(utils.Config.Embeds.FooterMessage, utils.Config.Embeds.FooterIcon).
			Build()).
		Build())
	if err != nil {
		return nil, err
	}

	utils.Logger.Debug().Str("ticket_channel", channel.ID().String()).Str("user", user.Username).Msg("Ticket channel created")

	return channel, nil
}

func CloseTicket(channelID snowflake.ID) error {
	channel, err := utils.Client.Rest().GetChannel(channelID)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(channel.Name(), "ticket-") {
		return errors.New("this channel is not a ticket channel")
	}

	err = utils.Client.Rest().DeleteChannel(channelID)
	if err != nil {
		return err
	}

	utils.Logger.Debug().Str("ticket_channel", channelID.String()).Msg("Ticket channel closed")
	return nil
}

func AddMemberToTicket(channelID snowflake.ID, userID snowflake.ID) error {
	channel, err := utils.Client.Rest().GetChannel(channelID)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(channel.Name(), "ticket-") {
		return errors.New("this channel is not a ticket channel")
	}

	_, err = utils.Client.Rest().UpdateChannel(channelID, discord.GuildTextChannelUpdate{
		PermissionOverwrites: &[]discord.PermissionOverwrite{
			discord.MemberPermissionOverwrite{
				UserID: userID,
				Allow:  discord.PermissionViewChannel | discord.PermissionSendMessages | discord.PermissionReadMessageHistory,
			},
		},
	})

	utils.Logger.Debug().Str("ticket_channel", channelID.String()).Str("user", userID.String()).Msg("User added to ticket")
	return nil
}

func RemoveMemberFromTicket(channelID snowflake.ID, userID snowflake.ID) error {
	channel, err := utils.Client.Rest().GetChannel(channelID)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(channel.Name(), "ticket-") {
		return errors.New("this channel is not a ticket channel")
	}

	_, err = utils.Client.Rest().UpdateChannel(channelID, discord.GuildTextChannelUpdate{
		PermissionOverwrites: &[]discord.PermissionOverwrite{
			discord.MemberPermissionOverwrite{
				UserID: userID,
				Deny:   discord.PermissionViewChannel | discord.PermissionSendMessages | discord.PermissionReadMessageHistory,
			},
		},
	})

	utils.Logger.Debug().Str("ticket_channel", channelID.String()).Str("user", userID.String()).Msg("User removed from ticket")
	return nil
}
