package services

import (
	"github.com/guilehm/solid-robot/internal/features/commands"
	"github.com/guilehm/solid-robot/internal/features/queries"
	"github.com/rs/zerolog"
)

type ServiceGroup struct {
	logger     *zerolog.Logger
	queries    *queries.QueryGroup
	commands   *commands.CommandGroup
	bulkAmount int
}

func newServiceGroup(logger *zerolog.Logger, queries *queries.QueryGroup, commands *commands.CommandGroup) *ServiceGroup {
	const bulkAmount = 1000
	const channelSize = bulkAmount * bulkAmount

	return &ServiceGroup{
		logger:     logger,
		queries:    queries,
		commands:   commands,
		bulkAmount: bulkAmount,
	}
}
