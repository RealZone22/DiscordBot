package main

import (
	"github.com/RealZone22/DiscordBot/cmd/bot"
	"github.com/RealZone22/DiscordBot/pkg/utils"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))

	err := utils.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	utils.InitLogger()

	utils.Logger.Debug().Msg("Pre-Initialization finished")

	bot.Init()
}
