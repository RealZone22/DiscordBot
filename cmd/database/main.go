package database

import (
	"github.com/RealZone22/DiscordBot/internal/handlers"
	"github.com/RealZone22/DiscordBot/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

func Init() {
	engine, err := xorm.NewEngine("sqlite3", "./database.db")
	if err != nil {
		utils.Logger.Error().Err(err).Msg("Failed to connect to database. Retrying in 5 seconds")
		time.Sleep(5 * time.Second)
		Init()
		return
	}

	engine.SetMapper(names.GonicMapper{})

	utils.DB = engine

	err = utils.DB.Sync2(new(handlers.Warn))
	if err != nil {
		utils.Logger.Error().Err(err).Msg("Failed to sync database. Retrying in 5 seconds")
		time.Sleep(5 * time.Second)
		Init()
		return
	}

	utils.Logger.Info().Msg("Database setup complete")
}
