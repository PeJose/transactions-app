package handlers

import (
	"cashflow/pkg/db"
	"cashflow/pkg/interfaces"
	"cashflow/pkg/stores"
)

type Handler struct {
	userStore        interfaces.UserStore
	companyStore     interfaces.CompanyStore
	rateStore        interfaces.RateStore
	transactionStore interfaces.TransactionStore
	validator        interfaces.Validator
}

func NewHandler(validator interfaces.Validator) *Handler {
	return &Handler{
		userStore:        stores.CreateUserStore(db.DB),
		companyStore:     stores.CreateCompanyStore(),
		rateStore:        stores.CreateRateStore(),
		transactionStore: stores.CreateTransactionStore(),
		validator:        validator,
	}
}
