package models

import "time"

type SEPA struct {
	ID        string    `json:"id"`
	Payer     string    `json:"payer"`
	Receiver  string    `json:"receiver"`
	Amount    float64   `json:"amount"`
	Currency  Currency  `json:"currency"`
	Timestamp time.Time `json:"timestamp"`
}

type SWIFT struct {
	ID          string    `json:"id"`
	Sender      string    `json:"sender"`
	Beneficiary string    `json:"beneficiary"`
	Amount      float64   `json:"amount"`
	Currency    Currency  `json:"currency"`
	Timestamp   time.Time `json:"timestamp"`
}

type Transaction interface {
	GetAmount() float64
	GetCurrency() Currency
	GetOperator(iban string) int
	GetDate() time.Time
}

func (s SWIFT) GetDate() time.Time {
	return s.Timestamp
}

func (s SEPA) GetCurrency() Currency {
	return s.Currency
}

func (s SEPA) GetOperator(iban string) int {
	if s.Payer == iban {
		return -1
	}
	return 1
}

func (s SEPA) GetDate() time.Time {
	return s.Timestamp
}

func (s SWIFT) GetAmount() float64 {
	return s.Amount
}

func (s SWIFT) GetCurrency() Currency {
	return s.Currency
}

func (s SWIFT) GetOperator(iban string) int {
	if s.Sender == iban {
		return -1
	}
	return 1
}

func (s SEPA) GetAmount() float64 {
	return s.Amount
}

type TransactionChange struct {
	Date   time.Time `json:"date"`
	Amount float64   `json:"amount"`
}

type TransactionPerCountry struct {
	Amount      float64 `json:"amount"`
	Occurrances int     `json:"occurrances"`
	Country     string  `json:"country"`
}
