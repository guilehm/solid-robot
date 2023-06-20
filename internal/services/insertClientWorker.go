package services

import (
	"context"
	"github.com/guilehm/solid-robot/internal/features/models"
	clientRepository "github.com/guilehm/solid-robot/internal/features/repository/client"
	"github.com/jackc/pgx/v5"
	"time"
)

func (service *ServiceGroup) insertClientWorker(ctx context.Context, batch *pgx.Batch, clientChannel <-chan models.Client, quit chan<- bool) {

	for clientRaw := range clientChannel {
		if batch.Len() >= service.bulkAmount {
			service.commands.SendCreateBatch(ctx, batch)

			newBatch := &pgx.Batch{}
			*batch = *newBatch
		}
		clientRepository.CreateBatch(batch, clientRaw)
	}

	if batch.Len() > 0 {
		service.commands.SendCreateBatch(ctx, batch)
	}

	quit <- true

	service.logger.Info().
		Str("now", time.Now().Format(time.RFC3339Nano)).
		Msg("finished inserting client")

}
