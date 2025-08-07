package ports

import domain "auto-trader-bot/internal/core/domain"

type AccountRetriever interface {
	GetAccountGeneralInfo() (domain.ExchangeAccount, error)
	GetAccountFundingBalances() (string, error)
}

type MarketRetriever interface {
	GetMarketCryptoInfo() (string, error)
}
