package models

import (
	"github.com/google/uuid"
	"time"
)

type ClientDocumentType string

const (
	ClientDocumentTypeCPF     ClientDocumentType = "cpf"
	ClientDocumentTypeCNPJ    ClientDocumentType = "cnpj"
	ClientDocumentTypeUnknown ClientDocumentType = "unknown"
)

type Client struct {
	ID                 uuid.UUID
	Document           string
	DocumentType       ClientDocumentType
	Private            bool
	Incomplete         bool
	LastPurchaseDate   *time.Time
	TicketAverage      int
	TicketLastPurchase int
	StoreMostFrequent  string
	StoreLastPurchase  string
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}
