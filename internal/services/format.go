package services

import (
	"context"
	"github.com/guilehm/solid-robot/internal/features/models"
	"github.com/guilehm/solid-robot/internal/services/helpers"
	"strconv"
	"time"
)

func (service *ServiceGroup) format(ctx context.Context, clientRaw models.ClientRaw) models.Client {

	logger := service.logger.
		With().
		Str("client_id", clientRaw.ID.String()).
		Logger()

	doc, docType := helpers.FormatDocument(clientRaw.Document)

	private, err := strconv.ParseBool(clientRaw.Private)
	if err != nil {
		logger.Error().Err(err).Msg("parsing private value")
	}
	incomplete, err := strconv.ParseBool(clientRaw.Private)
	if err != nil {
		logger.Error().Err(err).Msg("parsing incomplete value")
	}

	lpd, err := helpers.FormatDate(clientRaw.LastPurchaseDate)
	if err != nil {
		logger.Error().Err(err).Msg("parsing last purchase date")
	}

	ticketAverage, err := helpers.FormatFloat(clientRaw.TicketAverage)
	if err != nil {
		logger.Error().Err(err).Msg("parsing ticket average")
	}

	ticketLastPurchase, err := helpers.FormatFloat(clientRaw.TicketAverage)
	if err != nil {
		logger.Error().Err(err).Msg("parsing ticket last purchase")
	}

	now := time.Now()
	return models.Client{
		ID:                 clientRaw.ID,
		Document:           doc,
		DocumentType:       docType,
		Private:            private,
		Incomplete:         incomplete,
		LastPurchaseDate:   lpd,
		TicketAverage:      ticketAverage,
		TicketLastPurchase: ticketLastPurchase,
		StoreMostFrequent:  clientRaw.StoreMostFrequent,
		StoreLastPurchase:  clientRaw.StoreLastPurchase,
		CreatedAt:          &now,
		UpdatedAt:          &now,
	}

}
