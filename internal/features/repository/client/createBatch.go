package clientRepository

import (
	"github.com/guilehm/solid-robot/internal/features/models"
	"github.com/jackc/pgx/v5"
)

func CreateBatch(batch *pgx.Batch, client models.Client) {
	sql := `INSERT INTO clients (id, document, document_type, private, incomplete, last_purchase_date, ticket_average, ticket_last_purchase, store_most_frequent, store_last_purchase, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
			RETURNING id`

	batch.Queue(
		sql,
		client.ID,
		client.Document,
		client.DocumentType,
		client.Private,
		client.Incomplete,
		client.LastPurchaseDate,
		client.TicketAverage,
		client.TicketLastPurchase,
		client.StoreMostFrequent,
		client.StoreLastPurchase,
		client.CreatedAt,
		client.UpdatedAt,
	)
}
