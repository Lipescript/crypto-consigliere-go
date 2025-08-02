package domain

import (
	"math/big"
	"time"
)

type UserAccount struct {
	Username       string    `json:"username"`
	AccountType    string    `json:"account_type"`
	CanTrade       bool      `json:"can_trade"`
	CanWithdraw    bool      `json:"can_withdraw"`
	CanDeposit     bool      `json:"can_deposit"`
	CommissionRate float64   `json:"commission_rate"`
	UpdatedAt      time.Time `json:"updated_at"`
	Permissions    []string  `json:"permissions"`
	PersonalWallet Wallet    `json:"wallet" validate:"required"` // Primary crypto wallet
}

type Wallet struct {
	CryptoAssets    []CryptoAsset `json:"assets" validate:"required,gt=0"` // Crypto portfolio
	TotalFiat       *big.Int      `json:"total_fiat"`                      // Sum in fiat (wei-like precision)
	DefaultFiatCode string        `json:"fiat_code" validate:"iso4217"`    // ISO 4217 (USD, EUR)
}

type CryptoAsset struct {
	Symbol    string     `json:"exchange_symbol" validate:"required"` // Exchange-compliant symbol (BTC, ETH)
	Name      string     `json:"exchange_name"`                       // Official exchange name
	Amount    *big.Int   `json:"amount" validate:"gt=0"`              // Precise amount (wei-like)
	FiatValue *FiatValue `json:"fiat_valuation"`                      // Current market value
}

type FiatValue struct {
	Amount   *big.Int `json:"value" validate:"required"`   // Precise amount (decimal * 10^18)
	Currency string   `json:"currency" validate:"iso4217"` // ISO 4217 code
}
