package models

import "github.com/google/uuid"

type ClientRaw struct {
	ID                 uuid.UUID
	Document           string
	Private            string
	Incomplete         string
	LastPurchaseDate   string
	TicketAverage      string
	TicketLastPurchase string
	StoreMostFrequent  string
	StoreLastPurchase  string
	// Status             string
	CreatedAt string
}
