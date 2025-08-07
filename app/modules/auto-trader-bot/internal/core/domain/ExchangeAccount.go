package api

import (
	"math/big"
	"time"
)

type ExchangeAccount struct {
	AccountType     string    `json:"account_type"`
	CanTrade        bool      `json:"can_trade"`
	CanDeposit      bool      `json:"can_deposit"`
	CanWithdraw     bool      `json:"can_withdraw"`
	UpdatedAt       time.Time `json:"updated_at"`
	Wallet          Wallet    `json:"wallet"`
	TotalAssetOfBtc string    `json:"total_asset_btc"`
}

type Wallet struct {
	Spot Spot `json:"spot"`
}

type Spot struct {
	Fiat             *Fiat             `json:"fiat"`
	Cryptocurrencies *[]Cryptocurrency `json:"cryptocurrencies"`
	APISource        *string           `json:"api_source"`
}

type Fiat struct {
	Currency  *string    `json:"currency"`
	Value     *float64   `json:"value"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Cryptocurrency struct {
	Asset     *string    `json:"asset"`
	Value     *big.Float `json:"value"`
	UpdatedAt *time.Time `json:"updated_at"`
	Free      *big.Float `json:"free"`
	Locked    *big.Float `json:"locked"`
}
