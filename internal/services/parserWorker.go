package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/guilehm/solid-robot/internal/features/models"
	"strings"
	"time"
)

func getValue(start, end int, line string) string {
	if end == -1 {
		return strings.Trim(line[start:], " ")
	}
	return strings.Trim(line[start:end], " ")
}

func (service *ServiceGroup) parserWorker(ctx context.Context, rawMsgChannel <-chan string, channelClientRaw chan<- models.ClientRaw) {
	for line := range rawMsgChannel {
		now := time.Now().Format(time.RFC3339)
		channelClientRaw <- models.ClientRaw{
			ID:                 uuid.New(),
			Document:           getValue(DocumentIndexStart, DocumentIndexEnd, line),
			Private:            getValue(PrivateIndexStart, PrivateIndexEnd, line),
			Incomplete:         getValue(IncompleteIndexStart, IncompleteIndexEnd, line),
			LastPurchaseDate:   getValue(LastPurchaseDateIndexStart, LastPurchaseDateIndexEnd, line),
			TicketAverage:      getValue(TicketAverageIndexStart, TicketAverageIndexEnd, line),
			TicketLastPurchase: getValue(TicketLastPurchaseIndexStart, TicketLastPurchaseIndexEnd, line),
			StoreMostFrequent:  getValue(StoreMostFrequentIndexStart, StoreMostFrequentIndexEnd, line),
			StoreLastPurchase:  getValue(StoreLastPurchaseIndexStart, StoreLastPurchaseIndexEnd, line),
			Status:             models.ClientRawStatusWaiting,
			CreatedAt:          now,
			UpdatedAt:          now,
		}
	}

	close(channelClientRaw)

	service.logger.Info().Msg("finished parsing data")
}
