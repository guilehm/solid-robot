package commands

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (cg *CommandGroup) SendUpdateBatch(ctx context.Context, batch *pgx.Batch) {
	logger := cg.logger

	batchResults := cg.postgres.DB.SendBatch(ctx, batch)
	defer batchResults.Close()

	for i := 0; i < batch.Len(); i++ {
		row := batchResults.QueryRow()

		var id uuid.UUID
		err := row.Scan(
			&id,
		)
		if err != nil {
			logger.Error().Err(err).Msg("scanning update batch result")
			continue
		}

	}

}
