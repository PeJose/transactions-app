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

type companyStore struct {
}

func (cs *companyStore) GetById(id int) (*models.Company, error) {
	baseUrl, err := url.JoinPath(config.ApiUrl, "companies", fmt.Sprint(id))
	if err != nil {
		return nil, err
	}

	res, err := http.Get(baseUrl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var company *models.Company
	if err := helpers.ReadBody(res, &company); err != nil {
		return nil, err
	}

	return company, nil
}

func (cs *companyStore) GetAll(afterId, limit int) ([]models.Company, error) {
	baseUrl, err := url.JoinPath(config.ApiUrl, "companies")
	if err != nil {
		return nil, err
	}

	u, err := helpers.BuildUrl(baseUrl, map[string]string{
		"after-id": fmt.Sprint(afterId),
		"limit":    fmt.Sprint(limit),
	})
	if err != nil {
		return nil, err
	}

	res, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var companies []models.Company
	if err := helpers.ReadBody(res, &companies); err != nil {
		return nil, err
	}

	return companies, nil
}

func (cs *companyStore) GetByIban(iban string) (*models.Company, error) {
	companies, err := cs.GetAll(0, 5000)
	if err != nil {
		return nil, err
	}

	for _, company := range companies {
		for _, companyIban := range company.Ibans {
			if companyIban == iban {
				return &company, nil
			}
		}
	}
	
	return nil, fmt.Errorf("company with IBAN %s not found", iban)
}

func CreateCompanyStore() interfaces.CompanyStore {
	return &companyStore{}
}
