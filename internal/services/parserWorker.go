package services

import (
	"github.com/google/uuid"
	"github.com/guilehm/solid-robot/internal/features/models"
	"runtime"
	"strings"
	"time"
)

func getValue(start, end int, line string) string {
	if end == -1 {
		return strings.Trim(line[start:], " ")
	}
	return strings.Trim(line[start:end], " ")
}

func (service *ServiceGroup) parserWorker(rawMsgChannel <-chan string, channelClientRaw chan<- models.ClientRaw) {
	done := make(chan bool, runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		for line := range rawMsgChannel {
			now := time.Now().Format(time.RFC3339Nano)
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
		done <- true
	}

	for i := 0; i < runtime.NumCPU(); i++ {
		<-done
	}
	close(channelClientRaw)

	service.logger.Info().
		Str("now", time.Now().Format(time.RFC3339Nano)).
		Msg("finished parsing data")
}
