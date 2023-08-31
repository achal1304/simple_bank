package util

//constants for all supported currencies
const (
	USD = "USD"
	CAD = "CAD"
	EUR = "EUR"
)

//check valid currency
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
