package models

import (
	"github.com/google/uuid"
)

type ClientRawStatus string

const (
	ClientRawStatusWaiting   ClientRawStatus = "waiting"
	ClientRawStatusProcessed ClientRawStatus = "processed"
	ClientRawStatusFailed    ClientRawStatus = "failed"
)

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
	Status             ClientRawStatus
	CreatedAt          string
	UpdatedAt          string
}
