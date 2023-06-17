package commands

import (
	postgresStorage "github.com/guilehm/solid-robot/internal/storages/postgres"
	"github.com/rs/zerolog"
)

type CommandGroup struct {
	logger   *zerolog.Logger
	postgres *postgresStorage.Postgres
}

func newCommandGroup(logger *zerolog.Logger, postgres *postgresStorage.Postgres) *CommandGroup {
	return &CommandGroup{
		logger:   logger,
		postgres: postgres,
	}
}
