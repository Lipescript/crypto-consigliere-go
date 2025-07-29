package services

import (
	"auto-trader-bot/internal/core/domain"
	"auto-trader-bot/internal/core/ports"
)

type UserWalletService struct {
	userRetriever ports.UserRetriever
}

func NewWalletService(userRetriever ports.UserRetriever) *UserWalletService {
	return &UserWalletService{userRetriever: userRetriever}
}

func (userWalletService *UserWalletService) GetUserWalletInfo(user domain.User) domain.Wallet {

	userWalletService.userRetriever.GetUserInfo()

	return user.PersonalWallet
}
