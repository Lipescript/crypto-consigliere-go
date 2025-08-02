package services

import (
	"auto-trader-bot/internal/core/ports"
)

type ExchangeUserService struct {
	userRetriever ports.UserRetriever
}

func NewExchangeUserService(userRetriever ports.UserRetriever) *ExchangeUserService {
	return &ExchangeUserService{userRetriever: userRetriever}
}

func (userWalletService *ExchangeUserService) GetUserInfo() {
}
