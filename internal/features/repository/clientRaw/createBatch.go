package clientRawRepository

import (
	"github.com/guilehm/solid-robot/internal/features/models"
	"github.com/jackc/pgx/v5"
)

func CreateBatch(batch *pgx.Batch, clientRaw models.ClientRaw) {
	sql := `INSERT INTO clients_data_raw (id, document, private, incomplete, last_purchase_date, ticket_average, ticket_last_purchase, store_most_frequent, store_last_purchase, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			RETURNING id, document, private, incomplete, last_purchase_date, ticket_average, ticket_last_purchase, store_most_frequent, store_last_purchase`

	batch.Queue(
		sql,
		clientRaw.ID,
		clientRaw.Document,
		clientRaw.Private,
		clientRaw.Incomplete,
		clientRaw.LastPurchaseDate,
		clientRaw.TicketAverage,
		clientRaw.TicketLastPurchase,
		clientRaw.StoreMostFrequent,
		clientRaw.StoreLastPurchase,
		clientRaw.CreatedAt,
		clientRaw.UpdatedAt,
	)
}
