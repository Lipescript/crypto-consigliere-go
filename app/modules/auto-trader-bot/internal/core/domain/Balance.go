package domain

import (
	"time"
)

type User struct {
	Username    string  `json:"username"`
	AccountType string  `json:"account_type"`
	Sasvings    Savings `json:"savings"`
	Spot        Spot    `json:"spot"`
}
type Savings struct {
	SumBRL float64 `json:"sum_in_BRL"`
	SumDOL float64 `json:"sum_in_DOL"`

	UpdatedAt time.Time `json:"updated_at"`
}

type Spot struct {
	SumBRL float64 `json:"sum_in_BRL"`
	SumDOL float64 `json:"sum_in_DOL"`

	UpdatedAt time.Time `json:"updated_at"`
}
