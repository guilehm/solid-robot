package commands

import (
	"context"
	"github.com/google/uuid"
	"github.com/guilehm/solid-robot/internal/features/models"
	clientRawRepository "github.com/guilehm/solid-robot/internal/features/repository/clientRaw"
)

func (cg *CommandGroup) ClientRawCreate(ctx context.Context, clientRaw models.ClientRaw) (uuid.UUID, error) {
	logger := cg.logger.
		With().
		Str("id", clientRaw.ID.String()).
		Logger()

	logger.Debug().Msg("creating client raw")

	id, err := clientRawRepository.Create(ctx, cg.postgres.DB, clientRaw)
	if err != nil {
		return uuid.Nil, err
	}

	logger.Debug().Msg("client raw successfully created")
	return id, nil
}
