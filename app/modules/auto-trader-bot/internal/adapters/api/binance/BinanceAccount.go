package binance

import (
	api "auto-trader-bot/internal/adapters/api"
	constants "auto-trader-bot/internal/adapters/constants"

	binance_connector "github.com/binance/binance-connector-go"
)

type BinanceAdapter struct {
	*api.BaseExchangeAdapter
	binanceClient *binance_connector.Client
}

func NewBinanceAdapter(apiKey, apiSecret string) *BinanceAdapter {
	return &BinanceAdapter{
		BaseExchangeAdapter: api.NewApiBaseAdapter(
			apiKey,
			apiSecret,
			constants.BinanceAPIBaseURL,
		),
		binanceClient: binance_connector.NewClient(apiKey, apiSecret, constants.BinanceAPIBaseURL),
	}
}
