package binance

type AccountInfoResponse struct {
	MakerCommission int               `json:"makerCommission"`
	TakerCommission int               `json:"takerCommission"`
	CanTrade        bool              `json:"canTrade"`
	CanWithdraw     bool              `json:"canWithdraw"`
	CanDeposit      bool              `json:"canDeposit"`
	UpdateTime      int64             `json:"updateTime"`
	AccountType     string            `json:"accountType"`
	Balances        []BalanceResponse `json:"balances"`
	Permissions     []string          `json:"permissions"`
}

type BalanceResponse struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}
