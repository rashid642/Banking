package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/rashid642/banking/utils"
)

var validCurrency validator.Func = func(FieldLevel validator.FieldLevel) bool {
	if currency, ok := FieldLevel.Field().Interface().(string); ok {
		return utils.IsSupportedCurrency(currency)
	}
	return false 
}