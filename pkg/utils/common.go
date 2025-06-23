package utils

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/snowflake/v2"
	"xorm.io/xorm"
)

var Client bot.Client
var DB *xorm.Engine

func ConvertToSnowflake(id string) snowflake.ID {
	snowflakeID, err := snowflake.Parse(id)
	if err != nil {
		Logger.Error().Err(err).Msg("Failed to convert string to snowflake ID")
		return 0
	}
	return snowflakeID
}
