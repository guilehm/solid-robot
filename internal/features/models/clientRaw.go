package models

type ClientRaw struct {
	ID                 string
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
