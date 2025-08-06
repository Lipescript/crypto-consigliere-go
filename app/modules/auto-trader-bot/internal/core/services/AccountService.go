package services

import (
	"auto-trader-bot/internal/core/ports"
)

type AccountService struct {
	accountRetriever ports.AccountRetriever
}

func ExchangeAccountService(accountRetriever ports.AccountRetriever) *AccountService {
	return &AccountService{accountRetriever: accountRetriever}
}

// func (AccountService *AccountService) GetAccountInfo() domain.Account {

// }
