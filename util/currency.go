package util

const (
	USD = "USD"
	EUR = "EUR"
	CNY = "CNY"
)

// IsSupportedCurrency 检查货币是否支持
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CNY:
		return true
	}
	return false
}
