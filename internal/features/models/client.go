package models

import "time"

type Client struct {
	ID                  string
	Document            string
	DocumentType        string
	Private             bool
	Incomplete          bool
	LastPurchaseDate    *time.Time
	TicketAverage       int
	TicketLastPurchaset int
	StoreMostFrequent   string
	StoreLastPurchase   string
	Status              string
	CreatedAt           string
}
