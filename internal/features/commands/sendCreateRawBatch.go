package commands

import (
	"context"
	"github.com/guilehm/solid-robot/internal/features/models"
	"github.com/jackc/pgx/v5"
)

func (cg *CommandGroup) SendCreateRawBatch(ctx context.Context, batch *pgx.Batch, formatChannel chan<- models.ClientRaw) {
	logger := cg.logger

	batchResults := cg.postgres.DB.SendBatch(ctx, batch)
	defer batchResults.Close()

	for i := 0; i < batch.Len(); i++ {
		row := batchResults.QueryRow()

		var clientRaw models.ClientRaw
		err := row.Scan(
			&clientRaw.ID,
			&clientRaw.Document,
			&clientRaw.Private,
			&clientRaw.Incomplete,
			&clientRaw.LastPurchaseDate,
			&clientRaw.TicketAverage,
			&clientRaw.TicketLastPurchase,
			&clientRaw.StoreMostFrequent,
			&clientRaw.StoreLastPurchase,
		)
		if err != nil {
			logger.Error().Err(err).Msg("scanning create raw batch result")
			continue
		}

		formatChannel <- clientRaw
	}

}
