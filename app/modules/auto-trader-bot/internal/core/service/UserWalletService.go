package services

import (
	"auto-trader-bot/internal/core/ports"
)

type UserWalletService struct {
	converter ports.FiatConverter
}

func NewWalletService(converter ports.FiatConverter) *UserWalletService {
	return &UserWalletService{converter: converter}
}
