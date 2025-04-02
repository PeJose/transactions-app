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

func (s SEPA) GetAmount() float64 {
	return s.Amount
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

func (s SEPA) GetId() string {
	return s.ID
}

func (s SEPA) GetFrom() string {
	return s.Payer
}

func (s SEPA) GetTo() string {
	return s.Receiver
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
	GetId() string
	GetFrom() string
	GetTo() string
}

func (s SWIFT) GetAmount() float64 {
	return s.Amount
}

func (s SWIFT) GetDate() time.Time {
	return s.Timestamp
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

func (s SWIFT) GetId() string {
	return s.ID
}

func (s SWIFT) GetFrom() string {
	return s.Sender
}

func (s SWIFT) GetTo() string {
	return s.Beneficiary
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

type TransactionStandarised struct {
	ID        string    `json:"id"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Amount    float64   `json:"amount"`
	Currency  Currency  `json:"currency"`
	Timestamp time.Time `json:"timestamp"`
}
