package commands

import (
	"fmt"
	"github.com/RealZone22/DiscordBot/pkg/utils"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"strings"
)

var UserInfo = discord.SlashCommandCreate{
	Name:        "userinfo",
	Description: "Get information about a specific user",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionMentionable{
			Name:        "member",
			Description: "The user to get information about",
			Required:    true,
		},
	},
	DefaultMemberPermissions: nil,
}

func UserInfoHandler(e *handler.CommandEvent) error {
	member := e.SlashCommandInteractionData().Member("member")

	rolesList := make([]string, 0, len(member.RoleIDs))
	for _, roleID := range member.RoleIDs {
		rolesList = append(rolesList, fmt.Sprintf("<@&%s>", roleID))
	}
	rolesStr := "No roles"
	if len(rolesList) > 0 {
		rolesStr = strings.Join(rolesList, " ")
	}

	return e.CreateMessage(discord.NewMessageCreateBuilder().
		AddEmbeds(discord.NewEmbedBuilder().
			SetTitlef("User Information for %s", member.User.Username).
			AddField("Username", member.User.Username, true).
			AddField("ID", member.User.ID.String(), true).
			AddField("Roles", rolesStr, false).
			AddField("Joined At", fmt.Sprintf("<t:%d>", member.JoinedAt.Unix()), true).
			AddField("Created At", fmt.Sprintf("<t:%d>", member.CreatedAt().Unix()), true).
			SetThumbnail(*member.User.AvatarURL()).
			SetColor(3447003).
			SetFooter(utils.Config.Embeds.FooterMessage, utils.Config.Embeds.FooterIcon).
			Build()).
		Build())
}
