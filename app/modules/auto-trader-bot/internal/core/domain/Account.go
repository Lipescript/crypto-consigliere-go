package domain

import (
	"math/big"
	"time"
)

type Account struct {
	Username    string `json:"username"`
	AccountType string `json:"account_type"`
	Wallet      Wallet `json:"wallet"`
}

type Wallet struct {
	Spot Spot `json:"spot"`
}

type Spot struct {
	Fiat             Fiat               `json:"fiat"`
	Cryptocurrencies []Cryptocurrencies `json:"Cryptocurrencies"`
	APISource        string             `json:"api"`
}

type Fiat struct {
	Currency  string    `json:"currency"`
	Value     float64   `json:"value"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Cryptocurrencies struct {
	Currency  string    `json:"currency"`
	Value     big.Float `json:"value"`
	UpdatedAt time.Time `json:"updated_at"`
}
