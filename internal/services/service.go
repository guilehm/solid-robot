package services

import (
	"github.com/guilehm/solid-robot/internal/features/commands"
	"github.com/guilehm/solid-robot/internal/features/models"
	"github.com/guilehm/solid-robot/internal/features/queries"
	"github.com/rs/zerolog"
)

type ServiceGroup struct {
	logger     *zerolog.Logger
	queries    *queries.QueryGroup
	commands   *commands.CommandGroup
	bulkAmount int

	channelRawMsg    chan string
	channelClientRaw chan models.ClientRaw
}

func newServiceGroup(logger *zerolog.Logger, queries *queries.QueryGroup, commands *commands.CommandGroup) *ServiceGroup {
	const bulkAmount = 1000
	const channelSize = bulkAmount * bulkAmount

	channelRawMsg := make(chan string, channelSize)
	channelClientRaw := make(chan models.ClientRaw, channelSize)

	return &ServiceGroup{
		logger:           logger,
		queries:          queries,
		commands:         commands,
		bulkAmount:       bulkAmount,
		channelRawMsg:    channelRawMsg,
		channelClientRaw: channelClientRaw,
	}
}
