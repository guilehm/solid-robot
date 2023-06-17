package logger

import (
	"github.com/rs/zerolog"
	"os"
)

func NewLogger() *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}

	logger := zerolog.New(output).
		With().
		Timestamp().
		Logger()

	return &logger
}
