package helpers

import (
	"cashflow/pkg/middleware"
	"cashflow/pkg/models"
	"sort"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jbub/banking/iban"
)

func ExtractCountryFromIBAN(i string) string {
	ibanObj, err := iban.Parse(i)
	if err != nil {
		return ""
	}
	countryCode := ibanObj.CountryCode()
	if len(countryCode) > 2 {
		return ""
	}
	return strings.ToUpper(countryCode)
}

func CalculateNewBalance(balance float64, tx models.Transaction, rateMap models.RateMap, currency models.Currency, iban string) float64 {
	amount := tx.GetAmount()
	txCurr := tx.GetCurrency()
	if txCurr != currency {
		amount = amount * float64(rateMap[txCurr])
	}

	return balance + float64(tx.GetOperator(iban))*amount
}

func TransformRates(rates []models.Rate, currency models.Currency) models.RateMap {
	rateMap := make(models.RateMap)
	for _, rate := range rates {
		switch currency {
		case models.EUR:
			rateMap[models.Currency(rate.Currency)] = rate.EurRate
		case models.USD:
			rateMap[models.Currency(rate.Currency)] = rate.UsdRate
		}
	}
	return rateMap
}

func DefaultToCurrentDate(transactions []models.Transaction, index int) time.Time {
	if index >= 0 && index < len(transactions) {
		return transactions[index].GetDate()
	}
	return time.Now()
}

func SortTransactionsByDate(transactions []models.Transaction) {
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].GetDate().Before(transactions[j].GetDate())
	})
}

func GetIban(c *fiber.Ctx) (string, error) {
	userClaims := middleware.GetUser(c)
	if userClaims == nil {
		return "", ErrorResponse(c, fiber.ErrUnauthorized)
	}

	userIban := userClaims["iban"].(string)
	return c.Query("iban", userIban), nil
}
