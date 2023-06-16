package clientQuery

import (
	"context"
	"github.com/guilehm/solid-robot/internal/features/models"
	"github.com/guilehm/solid-robot/internal/features/queries/filters"
)

func (c *ClientQuery) FindMany(ctx context.Context, input *filters.ClientFilterInput) ([]*models.Client, error) {
	sql := `SELECT id FROM client`

	rows, err := c.postgres.DB.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	clients := make([]*models.Client, 0)
	for rows.Next() {
		var client models.Client
		err = rows.Scan(&client.ID)
		if err != nil {
			return nil, err
		}
		clients = append(clients, &client)
	}
	return clients, nil

}
