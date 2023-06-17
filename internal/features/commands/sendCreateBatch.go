package commands

import (
	"context"
	"github.com/guilehm/solid-robot/internal/features/models"
	"github.com/jackc/pgx/v5"
)

func (cg *CommandGroup) SendCreateBatch(ctx context.Context, batch *pgx.Batch) {
	logger := cg.logger

	batchResults := cg.postgres.DB.SendBatch(ctx, batch)
	defer batchResults.Close()

	for i := 0; i < batch.Len(); i++ {
		row := batchResults.QueryRow()

		var clientRaw models.Client
		err := row.Scan(
			&clientRaw.ID,
		)
		if err != nil {
			logger.Error().Err(err).Msg("scanning create batch result")
			continue
		}

	}

}
