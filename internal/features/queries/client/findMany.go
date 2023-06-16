package clientQuery

import (
	"context"
	"github.com/guilehm/solid-robot/internal/features/queries/filters"
	"github.com/guilehm/solid-robot/internal/features/types"
)

func (c *ClientQuery) FindMany(ctx context.Context, input *filters.ClientFilterInput) (*types.Client, error) {
	return nil, nil
}
