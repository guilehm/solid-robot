package services

import (
	"context"
	"fmt"
	"github.com/guilehm/solid-robot/internal/features/queries/filters"
	"github.com/rs/zerolog"
)

func (service *ServiceGroup) Process(logger *zerolog.Logger, ctx context.Context, filename string) {
	logger.Info().Msg("processing " + filename)
	clients, err := service.queries.Client.FindMany(ctx, &filters.ClientFilterInput{})
	if err != nil {
		return
	}

	fmt.Println("FOUND", clients)
}
