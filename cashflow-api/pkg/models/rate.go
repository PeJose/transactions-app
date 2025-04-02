package models

type Rate struct {
	Currency string
	EurRate  float32 `json:"eur_rate"`
	UsdRate  float32 `json:"usd_rate"`
}

type Currency string

const (
	EUR Currency = "EUR"
	USD Currency = "USD"
)

type RateMap map[Currency]float32

func IsValidCurrency(currency string) bool {
	switch Currency(currency) {
	case EUR, USD:
		return true
	default:
		return false
	}
}
