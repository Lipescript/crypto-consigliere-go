package ports

import (
	"math/big"
)

type FiatConverter interface {
	GetCurrencyInfo(fiatSymbol string) (string, string)
	ConvertCurrencyTo(assetValue *big.Int, userActualCurrency string, targetCurrency string) (*big.Int, error)
}

type ExchangeConverter interface {
	GetCurrencyInfo(assetSymbol string) (string, string)
	ConvertCurrencyTo(assetValue *big.Int, userActualCurrency string, targetCurrency string) (*big.Int, error)
}
