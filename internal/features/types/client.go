package types

import "time"

type Client struct {
	ID                 string
	Document           string
	DocumentType       string
	IsPrivate          bool
	IsIncomplete       bool
	LastPurchaseDate   *time.Time
	AverageTicket      float64
	LastPurchaseTicket float64
	MostFrequentStore  string
	LastPurchaseStore  string
	Status             string
	CreatedAt          string
}
