package events

import "github.com/RealZone22/DiscordBot/pkg/utils"

func Register() {
	utils.Logger.Debug().Msg("Registering events")
	utils.Client.AddEventListeners(JoinEventHandler(), LeaveEventHandler())
}
