package bot

import (
	"context"
	"github.com/RealZone22/DiscordBot/cmd/database"
	"github.com/RealZone22/DiscordBot/internal/commands"
	"github.com/RealZone22/DiscordBot/internal/cron"
	"github.com/RealZone22/DiscordBot/internal/events"
	"github.com/RealZone22/DiscordBot/pkg/utils"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	client, err := disgo.New(utils.Config.Token, bot.WithGatewayConfigOpts(gateway.WithIntents(gateway.IntentsAll)))
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("Error initializing Discord client")
	}

	utils.Client = client

	if err := client.OpenGateway(context.TODO()); err != nil {
		utils.Logger.Fatal().Err(err).Msg("Error opening Discord gateway")
	}

	database.Init()

	events.Register()
	commands.Register()
	err = cron.RunCronJobs()
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("Error running cron jobs")
		return
	}

	utils.Logger.Debug().Msg("Discord client initialized successfully")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	utils.Logger.Info().Msg("Shutting down...")
}
