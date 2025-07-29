package api

import (
	constants "auto-trader-bot/internal/adapters/constants"
	"net/http"
	"time"
)

type ExchangeAdapter interface {
	AuthenticatedRequest(method, endpoint string, body []byte) (*http.Response, error)
	GenerateNonce() string
}

type BaseExchangeAdapter struct {
	Client    *http.Client
	APIKey    string
	APISecret string
	BaseURL   string
}

func NewBaseAdapter(apiKey, apiSecret, baseURL string) *BaseExchangeAdapter {
	return &BaseExchangeAdapter{
		Client:    &http.Client{Timeout: constants.DefaultAPITimeoutSeconds * time.Second},
		APIKey:    apiKey,
		APISecret: apiSecret,
		BaseURL:   baseURL,
	}
}
