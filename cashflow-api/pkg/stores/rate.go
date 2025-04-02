package stores

import (
	"cashflow/pkg/config"
	"cashflow/pkg/helpers"
	"cashflow/pkg/interfaces"
	"cashflow/pkg/models"
	"fmt"
	"net/http"
	"net/url"
)

type rateStore struct {
}

// GetAll implements interfaces.RateStore.
func (r *rateStore) GetAll() ([]models.Rate, error) {
	u, err := url.JoinPath(config.ApiUrl, "exchange-rates")
	if err != nil {
		return nil, err
	}

	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var rates []models.Rate
	if err := helpers.ReadBody(res, &rates); err != nil {
		return nil, err
	}

	return rates, nil
}

// GetOne implements interfaces.RateStore.
func (r *rateStore) GetOne(currency string) (*models.Rate, error) {
	rates, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, rate := range rates {
		if rate.Currency == currency {
			return &rate, nil
		}
	}

	return nil, fmt.Errorf("rate not found for currency: %s", currency)
}

func CreateRateStore() interfaces.RateStore {
	return &rateStore{}
}
