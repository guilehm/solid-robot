package services

import (
	"github.com/guilehm/solid-robot/internal/features/models"
	"github.com/guilehm/solid-robot/internal/features/queries"
)

type ServiceGroup struct {
	queries    *queries.QueryGroup
	bulkAmount int

	channelRawMsg    chan string
	channelClientRaw chan models.ClientRaw
}

func newServiceGroup(queries *queries.QueryGroup) *ServiceGroup {
	const bulkAmount = 1000
	const channelSize = bulkAmount * bulkAmount

	channelRawMsg := make(chan string, channelSize)
	channelClientRaw := make(chan models.ClientRaw, channelSize)

	return &ServiceGroup{
		queries:          queries,
		bulkAmount:       bulkAmount,
		channelRawMsg:    channelRawMsg,
		channelClientRaw: channelClientRaw,
	}
}
