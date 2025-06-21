package utils

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var Logger = zerolog.Logger{}

func InitLogger() {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime},
	).Level(Config.LogLevel).With().Timestamp().Caller().Logger()

	Logger = logger

	logger.Debug().Msg("Logger setup complete")
}
