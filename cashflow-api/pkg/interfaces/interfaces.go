package interfaces

import "cashflow/pkg/models"

// UserStore interface for user-related operations
type UserStore interface {
	GetByEmail(email string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
}

// CompanyStore interface for company-related operations
type CompanyStore interface {
	GetById(id int) (*models.Company, error)
	GetByIban(iban string) (*models.Company, error)
	GetAll(afterId, limit int) ([]models.Company, error)
}

type RateStore interface {
	GetOne(currency string) (*models.Rate, error)
	GetAll() ([]models.Rate, error)
}

// Validator interface for validation
type Validator interface {
	Validate(data interface{}) error
	Format(err error) string
}

type TransactionStore interface {
	GetSEPA(params map[string]string) ([]models.SEPA, error)
	GetSWIFT(params map[string]string) ([]models.SWIFT, error)
	GetBalance(iban string, rates []models.Rate, currency models.Currency) (float64, error)
	GetTransactionPerCountries(iban string, rate []models.Rate) ([]models.TransactionPerCountry, error)
	GetBalanceChange(iban string, rate []models.Rate) ([]models.TransactionChange, error)
}
