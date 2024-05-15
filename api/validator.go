package api

import (
	"GoBank/util"
	"github.com/go-playground/validator/v10"
)

// 为货币绑定验证器（CNY USD EUR）
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// 检查是否支持该货币
		return util.IsSupportedCurrency(currency)
	}
	return false
}
