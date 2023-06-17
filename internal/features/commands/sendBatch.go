package commands

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func (cg *CommandGroup) SendBatch(ctx context.Context, batch *pgx.Batch) {
	logger := cg.logger

	results := cg.postgres.DB.SendBatch(ctx, batch)
	defer results.Close()

	for i := 0; i < batch.Len(); i++ {
		result, err := results.Exec()
		if err != nil {
			logger.Error().Err(err).Msg("sending batch command")
			continue
		}

		rowsAffected := result.RowsAffected()
		logger.Debug().
			Int64("rows_affected", rowsAffected).
			Msg("rows affected")
	}

}
