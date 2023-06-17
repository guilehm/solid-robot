package clientRawRepository

import (
	"context"
	"github.com/google/uuid"
	"github.com/guilehm/solid-robot/internal/features/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Create(ctx context.Context, db *pgxpool.Pool, clientRaw models.ClientRaw) (uuid.UUID, error) {
	sql := `INSERT INTO clients_data_raw (id, document, private, incomplete, last_purchase_date, ticket_average, ticket_last_purchase, store_most_frequent, store_last_purchase)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			RETURNING id`

	var id uuid.UUID
	err := db.QueryRow(
		ctx,
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
	).Scan(&id)

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
