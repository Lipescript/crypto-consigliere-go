package ports

type AccountRetriever interface {
	GetAccountGeneralInfo()
	GetAccountFundingBalances() (string, error)
}

type MarketRetriever interface {
	GetMarketCryptoInfo() (string, error)
}
