package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/guilehm/solid-robot/internal/features/models"
	clientRawRepository "github.com/guilehm/solid-robot/internal/features/repository/clientRaw"
	"github.com/jackc/pgx/v5"
	"time"
)

type clientRawUpdateInput struct {
	ID     uuid.UUID
	Status models.ClientRawStatus
}

func (service *ServiceGroup) updateClientRaw(ctx context.Context, batch *pgx.Batch, channel <-chan clientRawUpdateInput) {

	for clientRaw := range channel {
		if batch.Len() >= service.bulkAmount {
			service.commands.SendUpdateBatch(ctx, batch)

			newBatch := &pgx.Batch{}
			*batch = *newBatch
		}
		clientRawRepository.UpdateBatch(batch, string(clientRaw.Status), clientRaw.ID)
	}

	if batch.Len() > 0 {
		service.commands.SendUpdateBatch(ctx, batch)
	}

	service.logger.Info().
		Str("now", time.Now().Format("15:04:05.999999999")).
		Msg("finished updating client raw")

}
