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
	Warn,
}

func Register() {
	h := handler.New()
	h.Command("/ping", PingHandler)
	h.Command("/userinfo", UserInfoHandler)
	h.Command("/purge", PurgeHandler)
	h.Route("/warn", func(r handler.Router) {
		r.Command("/create", CreateWarnHandler)
		r.Command("/get", GetWarnsHandler)
		r.Command("/delete", DeleteWarnHandler)
		r.Command("/clear", ClearWarnsHandler)
		r.Command("/count", CountWarnsHandler)
	})

	if utils.Config.Ticket.Enabled {
		h.Route("/ticket", func(r handler.Router) {
			r.Command("/embed", TicketEmbedHandler)
			r.Command("/create", CreateTicketHandler)
			r.Command("/close", CloseTicketHandler)
			r.Command("/addmember", AddMemberToTicketHandler)
			r.Command("/removemember", RemoveMemberFromTicketHandler)
		})
		Commands = append(Commands, Ticket)
	}

	utils.Client.AddEventListeners(h)

	err := handler.SyncCommands(utils.Client, Commands, []snowflake.ID{utils.ConvertToSnowflake(utils.Config.DefaultGuildId)})
	if err != nil {
		utils.Logger.Error().Err(err).Msg("Failed to sync commands")
		return
	}
}
