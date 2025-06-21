package utils

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"os"
)

type ConfigStruct struct {
	LogLevel       zerolog.Level `json:"log_level"` // 0: Debug, 1: Info, 2: Warn, 3: Error, 4: Fatal, 5: Panic, 6: NoLevel, 7: Disabled
	Token          string        `json:"token"`
	DefaultGuildId string        `json:"default_guild_id"` // Default guild ID for the bot to operate in
	Embeds         struct {
		FooterMessage string `json:"footer_message"` // Message to be displayed in the footer of embeds
		FooterIcon    string `json:"footer_icon"`    // URL of the icon to be displayed in the footer of embeds
	} `json:"embeds"`
	Activity struct {
		Type    string `json:"type"` // "PLAYING", "LISTENING", "WATCHING", etc.
		Message string `json:"message"`
	} `json:"activity"`
	Events struct {
		Join struct {
			Enabled   bool   `json:"enabled"`
			ChannelID string `json:"channel_id"`
		} `json:"join"`
	} `json:"events"`
	Commands struct {
		ModerationEnabled bool `json:"moderation_enabled"`
		GiveawayEnabled   bool `json:"giveaway_enabled"`
		HelpEnabled       bool `json:"help_enabled"`
		UserInfoEnabled   bool `json:"user_info_enabled"`
		PingEnabled       bool `json:"ping_enabled"`
	} `json:"commands"`
	Stats struct {
		Enabled        bool   `json:"enabled"`
		UsersChannelId string `json:"users_channel_id"`
	} `json:"stats"`
	Ticket struct {
		Enabled     bool   `json:"enabled"`
		CategoryID  string `json:"category_id"`
		SupportRole string `json:"ticket_role"`
	} `json:"ticket"`
}

var Config *ConfigStruct

func InitConfig() error {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		if err = createConfig(); err != nil {
			return err
		}
	}

	if err := readConfig(); err != nil {
		return err
	}

	return nil
}

func readConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	Config = &ConfigStruct{}
	err = json.NewDecoder(file).Decode(Config)
	if err != nil {
		return err
	}

	return nil
}

func createConfig() error {
	file, err := os.Create("config.json")
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	data, err := json.MarshalIndent(&ConfigStruct{}, "", "   ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
