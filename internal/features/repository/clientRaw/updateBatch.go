package clientRawRepository

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func UpdateBatch(batch *pgx.Batch, status string, id uuid.UUID) {
	sql := `UPDATE clients_data_raw SET status = $1 WHERE ID = $2 RETURNING ID`

	batch.Queue(
		sql,
		status,
		id,
	)
}
