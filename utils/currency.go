package utils 

const (
	USD = "USD"
	EURO = "EURO"
	INR = "INR"
)


func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EURO, INR :
		return true 
	}
	return false 
}