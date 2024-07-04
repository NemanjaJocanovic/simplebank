package util

const (
	USD = "USD"
	EUR = "EUR"
	RSD = "RSD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, RSD:
		return true
	}
	return false
}
