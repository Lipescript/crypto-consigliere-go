package ports

import (
	"math/big"
)

type FiatConverter interface {
	GetCurrencyInfo(fiatSymbol string) (string, string)
	ConvertToCurrency(assetValue *big.Int, currency string, targetCurrency string) (*big.Int, error)
}

type CryptoConverter interface {
	GetCurrencyInfo(fiatSymbol string) (string, string)
	ConvertToCurrency(assetValue *big.Int, currency string, targetCurrency string) (*big.Int, error)
}
