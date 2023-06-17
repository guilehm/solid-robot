package services

import (
	"context"
	"github.com/guilehm/solid-robot/internal/features/models"
)

func (service *ServiceGroup) formatterWorker(ctx context.Context, formatChannel <-chan models.ClientRaw, clientChannel chan<- models.Client) {
	for clientRaw := range formatChannel {
		clientChannel <- service.format(ctx, clientRaw)
	}
	close(clientChannel)
}
