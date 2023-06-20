package services

import (
	"context"
	"github.com/guilehm/solid-robot/internal/features/models"
	clientRawRepository "github.com/guilehm/solid-robot/internal/features/repository/clientRaw"
	"github.com/jackc/pgx/v5"
	"time"
)

func (service *ServiceGroup) insertClientRawWorker(ctx context.Context, batch *pgx.Batch, channelClientRaw <-chan models.ClientRaw, formatChannel chan models.ClientRaw) {

	for clientRaw := range channelClientRaw {
		if batch.Len() >= service.bulkAmount {
			service.commands.SendCreateRawBatch(ctx, batch, formatChannel)

			newBatch := &pgx.Batch{}
			*batch = *newBatch
		}
		clientRawRepository.CreateBatch(batch, clientRaw)
	}

	if batch.Len() > 0 {
		service.commands.SendCreateRawBatch(ctx, batch, formatChannel)
	}

	close(formatChannel)

	service.logger.Info().
		Str("now", time.Now().Format("15:04:05.999999999")).
		Msg("finished inserting client raw")
}
