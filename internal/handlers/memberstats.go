package handlers

import (
	"github.com/RealZone22/DiscordBot/pkg/utils"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
	"strconv"
)

func HandleMemberStats(guildId snowflake.ID) {
	if !utils.Config.Stats.Enabled {
		return
	}

	guild, err := utils.Client.Rest().GetGuild(guildId, true)
	if err != nil {
		utils.Logger.Error().Str("guild", guildId.String()).Msg("Failed to get guild")
		return
	}

	name := "Users: " + strconv.Itoa(guild.ApproximateMemberCount)

	_, err = utils.Client.Rest().UpdateChannel(utils.ConvertToSnowflake(utils.Config.Stats.UsersChannelId), discord.GuildVoiceChannelUpdate{Name: &name})
	if err != nil {
		utils.Logger.Error().Err(err).Msg("Failed to update users channel")
		return
	}

	utils.Logger.Info().Str("guild", guildId.String()).Int("member_count", guild.ApproximateMemberCount).Msg("Updated users channel")
}
