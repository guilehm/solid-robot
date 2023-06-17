package helpers

import (
	"github.com/guilehm/solid-robot/internal/features/models"
	"github.com/klassmann/cpfcnpj"
	"strconv"
	"strings"
	"time"
)

func clean(value string) string {
	if value == "NULL" {
		return ""
	}
	return value
}

func FormatDocument(document string) (string, models.ClientDocumentType) {
	document = clean(document)
	if cpfcnpj.ValidateCPF(document) {
		return cpfcnpj.Clean(document), models.ClientDocumentTypeCPF
	}
	if cpfcnpj.ValidateCPF(document) {
		return cpfcnpj.Clean(document), models.ClientDocumentTypeCNPJ
	}
	return document, models.ClientDocumentTypeUnknown
}

func FormatDate(date string) (*time.Time, error) {
	date = clean(date)
	if date == "" {
		return nil, nil
	}

	parsedDate, err := time.Parse("2006-01-02", clean(date))
	if err != nil {
		return nil, err
	}

	return &parsedDate, nil
}

func FormatFloat(value string) (int, error) {
	value = strings.Replace(clean(value), ",", ".", 1)
	if value == "" {
		return 0, nil
	}
	parsedFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return int(parsedFloat * 100), nil
}
