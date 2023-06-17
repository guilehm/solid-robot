package services

import (
	"context"
	"github.com/guilehm/solid-robot/internal/features/models"
	clientRawRepository "github.com/guilehm/solid-robot/internal/features/repository/clientRaw"
	"github.com/jackc/pgx/v5"
)

func (service *ServiceGroup) insertClientRaw(ctx context.Context, batch *pgx.Batch, channel <-chan models.ClientRaw) {

	for clientRaw := range channel {
		if batch.Len() >= service.bulkAmount {
			service.commands.SendBatch(ctx, batch)

			newBatch := &pgx.Batch{}
			*batch = *newBatch
		}
		clientRawRepository.CreateBatch(batch, clientRaw)
	}

	if batch.Len() > 0 {
		service.commands.SendBatch(ctx, batch)
	}

	service.logger.Info().Msg("finished inserting client raw")

}
