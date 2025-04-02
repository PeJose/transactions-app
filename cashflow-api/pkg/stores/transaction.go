package stores

import (
	"cashflow/pkg/config"
	"cashflow/pkg/helpers"
	"cashflow/pkg/interfaces"
	"cashflow/pkg/models"
	"cashflow/pkg/validators"
	"net/http"
	"net/url"
)

type transactionStore struct {
}
type Balance float64

func (ts *transactionStore) GetSEPA(params map[string]string) ([]models.SEPA, error) {
	baseUrl, err := url.JoinPath(config.ApiUrl, "transactions/sepa")
	if err != nil {
		return nil, err
	}

	u, err := helpers.BuildUrl(baseUrl, params)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var sepaList []models.SEPA
	if err := helpers.ReadBody(res, &sepaList); err != nil {
		return nil, err
	}

	filteredSepaList := []models.SEPA{}
	for _, tx := range sepaList {
		if validators.ValidateIban(tx.Payer, tx.Receiver) {
			filteredSepaList = append(filteredSepaList, tx)
		}
	}

	return filteredSepaList, nil
}

func (ts *transactionStore) GetSWIFT(params map[string]string) ([]models.SWIFT, error) {
	baseUrl, err := url.JoinPath(config.ApiUrl, "transactions/swift")
	if err != nil {
		return nil, err
	}

	u, err := helpers.BuildUrl(baseUrl, params)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var swiftList []models.SWIFT
	if err := helpers.ReadBody(res, &swiftList); err != nil {
		return nil, err
	}

	filteredSwiftList := []models.SWIFT{}
	for _, tx := range swiftList {
		if validators.ValidateIban(tx.Sender, tx.Beneficiary) {
			filteredSwiftList = append(filteredSwiftList, tx)
		}
	}

	return filteredSwiftList, nil
}

func (ts *transactionStore) GetSEPAOutgoing(iban string) ([]models.SEPA, error) {
	return ts.GetSEPA(map[string]string{"payer": iban})
}

func (ts *transactionStore) GetSEPAIncoming(iban string) ([]models.SEPA, error) {
	return ts.GetSEPA(map[string]string{"receiver": iban})
}

func (ts *transactionStore) GetSEPAAll(iban string) ([]models.SEPA, error) {
	sepaOutgoing, err := ts.GetSEPAOutgoing(iban)
	if err != nil {
		return nil, err
	}

	sepaIncoming, err := ts.GetSEPAIncoming(iban)
	if err != nil {
		return nil, err
	}

	return append(sepaOutgoing, sepaIncoming...), nil
}

func (ts *transactionStore) GetSWIFTOutgoing(iban string) ([]models.SWIFT, error) {
	return ts.GetSWIFT(map[string]string{"sender": iban})
}

func (ts *transactionStore) GetSWIFTIncoming(iban string) ([]models.SWIFT, error) {
	return ts.GetSWIFT(map[string]string{"beneficiary": iban})
}

func (ts *transactionStore) GetSwiftAll(iban string) ([]models.SWIFT, error) {
	sepaOutgoing, err := ts.GetSWIFTOutgoing(iban)
	if err != nil {
		return nil, err
	}

	sepaIncoming, err := ts.GetSWIFTIncoming(iban)
	if err != nil {
		return nil, err
	}

	return append(sepaOutgoing, sepaIncoming...), nil
}

func (ts *transactionStore) getAllTransactions(iban string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	sepaTransactions, err := ts.GetSEPAAll(iban)
	if err != nil {
		return nil, err
	}
	for _, tx := range sepaTransactions {
		transactions = append(transactions, tx)
	}

	swiftTransactions, err := ts.GetSwiftAll(iban)
	if err != nil {
		return nil, err
	}
	for _, tx := range swiftTransactions {
		transactions = append(transactions, tx)
	}

	return transactions, nil
}

func (ts *transactionStore) GetBalanceChange(iban string, rates []models.Rate) ([]models.TransactionChange, error) {
	transactions, err := ts.getAllTransactions(iban)
	if err != nil {
		return nil, err
	}

	rateMap := helpers.TransformRates(rates, models.EUR)

	helpers.SortTransactionsByDate(transactions)

	transactionChanges := []models.TransactionChange{}
	currentBalance := 0.0

	initialDate := helpers.DefaultToCurrentDate(transactions, 0)
	transactionChanges = append(transactionChanges, models.TransactionChange{
		Amount: 0,
		Date:   initialDate,
	})

	for i, tx := range transactions {
		currentBalance = helpers.CalculateNewBalance(currentBalance, tx, rateMap, models.EUR, iban)
		nextDate := helpers.DefaultToCurrentDate(transactions, i+1)
		transactionChanges = append(transactionChanges, models.TransactionChange{
			Amount: currentBalance,
			Date:   nextDate,
		})
	}

	return transactionChanges, nil
}

func (ts *transactionStore) GetBalance(iban string, rates []models.Rate, currency models.Currency) (float64, error) {
	balance := 0.0

	rateMap := make(models.RateMap)
	for _, rate := range rates {
		switch currency {
		case models.EUR:
			rateMap[models.Currency(rate.Currency)] = rate.EurRate
		case models.USD:
			rateMap[models.Currency(rate.Currency)] = rate.UsdRate
		}
	}

	sepaTransactions, err := ts.GetSEPAAll(iban)
	if err != nil {
		return 0, err
	}
	for _, tx := range sepaTransactions {
		balance = helpers.CalculateNewBalance(balance, tx, rateMap, currency, iban)
	}

	swiftTransactions, err := ts.GetSwiftAll(iban)
	if err != nil {
		return 0, err
	}
	for _, tx := range swiftTransactions {
		balance = helpers.CalculateNewBalance(balance, tx, rateMap, currency, iban)
	}

	return balance, nil
}

func (ts *transactionStore) GetPerCountries(iban string, rates []models.Rate) ([]models.TransactionPerCountry, error) {
	countryTransactions := make(map[string]models.TransactionPerCountry)

	rateMap := make(models.RateMap)
	for _, rate := range rates {
		rateMap[models.Currency(rate.Currency)] = rate.EurRate
	}

	processTransaction := func(country string, tx models.Transaction) {
		if country == "" {
			return
		}

		txPerCountry := countryTransactions[country]
		txPerCountry.Amount = helpers.CalculateNewBalance(txPerCountry.Amount, tx, rateMap, models.EUR, iban)
		txPerCountry.Occurrances++
		countryTransactions[country] = txPerCountry
	}

	sepaOut, err := ts.GetSEPAOutgoing(iban)
	if err != nil {
		return nil, err
	}
	for _, tx := range sepaOut {
		processTransaction(helpers.ExtractCountryFromIBAN(tx.Receiver), tx)
	}

	sepaIn, err := ts.GetSEPAIncoming(iban)
	if err != nil {
		return nil, err
	}
	for _, tx := range sepaIn {
		processTransaction(helpers.ExtractCountryFromIBAN(tx.Payer), tx)
	}

	swiftOut, err := ts.GetSWIFTOutgoing(iban)
	if err != nil {
		return nil, err
	}
	for _, tx := range swiftOut {
		processTransaction(helpers.ExtractCountryFromIBAN(tx.Beneficiary), tx)
	}

	swiftIn, err := ts.GetSWIFTIncoming(iban)
	if err != nil {
		return nil, err
	}
	for _, tx := range swiftIn {
		processTransaction(helpers.ExtractCountryFromIBAN(tx.Sender), tx)
	}

	var result []models.TransactionPerCountry
	for country, txPerCountry := range countryTransactions {
		result = append(result, models.TransactionPerCountry{
			Amount:      txPerCountry.Amount,
			Occurrances: txPerCountry.Occurrances,
			Country:     country,
		})
	}

	return result, nil
}

func CreateTransactionStore() interfaces.TransactionStore {
	return &transactionStore{}
}
