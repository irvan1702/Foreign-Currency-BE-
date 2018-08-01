package models

import (
	"github.com/satori/go.uuid"
)

type ForexHistoryRequest struct {
	ForexDate    string
	CurrencyFrom string
	CurrencyTo   string
	ExchangeRate float64
	Group        uuid.UUID
}
