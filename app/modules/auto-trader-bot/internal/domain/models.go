package domain

type Balance struct {
	Asset  string  `json:"asset"`
	Free   float64 `json:"free"`
	Locked float64 `json:"locked"`
}

type BinanceConfig struct {
	APIKey    string
	SecretKey string
}
