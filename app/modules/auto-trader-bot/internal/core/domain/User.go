package domain

import (
	"math/big"
)

type User struct {
	Username       string `json:"user_name"`
	Email          string `json:"email"`
	PersonalWallet Wallet `json:"personal_wallet"`
}

type Wallet struct {
	CryptoAssets        []CryptoAsset `json:"balance"`
	TotalFiat           *big.Int      `json:"total_fiat"`
	DefaultFiatCurrency string        `json:"fiat_currency"`        // (USD, EUR)
	FiatCurrencySymbol  string        `json:"fiat_currency_symbol"` // (US$, €)
}

type CryptoAsset struct {
	Symbol    string     `json:"symbol"` // (BTC, ETH)
	Name      string     `json:"name"`   // (Bitcoin)
	Amount    *big.Int   `json:"amount"`
	FiatValue *FiatValue `json:"fiat_value"`
}

type FiatValue struct {
	Amount   *big.Int `json:"amount"`
	Currency string   `json:"currency"` // (USD)
}
