package interfaces

import "cashflow/pkg/models"

type UserStore interface {
	GetByEmail(email string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
}

type CompanyStore interface {
	GetByIban(iban string) (*models.Company, error)
	GetAll(afterId, limit int) ([]models.Company, error)
}

type RateStore interface {
	GetOne(currency string) (*models.Rate, error)
	GetAll() ([]models.Rate, error)
}

type Validator interface {
	Validate(data interface{}) error
	Format(err error) string
}

type TransactionStore interface {
	GetSEPA(params map[string]string) ([]models.SEPA, error)
	GetSWIFT(params map[string]string) ([]models.SWIFT, error)
	GetBalance(iban string, rates []models.Rate, currency models.Currency) (float64, error)
	GetPerCountries(string, []models.Rate) ([]models.TransactionPerCountry, error)
	GetBalanceChange(iban string, rates []models.Rate) ([]models.TransactionChange, error)
}
