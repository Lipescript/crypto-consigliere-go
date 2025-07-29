package binance

import (
	api "auto-trader-bot/internal/adapters/api"
	constants "auto-trader-bot/internal/adapters/constants"
)

type BinanceAdapter struct {
	*api.BaseExchangeAdapter
}

func NewBinanceAdapter(apiKey, apiSecret string) *BinanceAdapter {
	return &BinanceAdapter{
		BaseExchangeAdapter: api.NewBaseAdapter(
			apiKey,
			apiSecret,
			constants.BinanceAPIBaseURL,
		),
	}
}
