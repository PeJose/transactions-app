package validators

import "github.com/jbub/banking/iban"

func ValidateIban(ibans ...string) bool {
	for _, i := range ibans {
		err := iban.Validate(i)
		if err != nil {
			return false
		}
	}
	return true
}
