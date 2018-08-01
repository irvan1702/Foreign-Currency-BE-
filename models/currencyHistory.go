package models

import (
	"time"
)

type ForexHistory struct {
	ID           uint64
	ForexDate    time.Time
	CurrencyFrom string
	CurrencyTo   string
	ExchangeRate float64
}
