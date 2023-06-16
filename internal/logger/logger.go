package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

func NewLogger() *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	logger := zerolog.New(output).
		With().
		Timestamp().
		Logger()

	return &logger
}
