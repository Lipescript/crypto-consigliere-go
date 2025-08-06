package api

import (
	constants "auto-trader-bot/internal/constants"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"
)

type BaseExchangeAdapter struct {
	Client    *http.Client
	APIKey    string
	APISecret string
	BaseURL   string
}

func NewApiBaseAdapter(apiKey, apiSecret, baseURL string) *BaseExchangeAdapter {
	return &BaseExchangeAdapter{
		Client:    &http.Client{Timeout: constants.DefaultAPITimeoutSeconds * time.Second},
		APIKey:    apiKey,
		APISecret: apiSecret,
		BaseURL:   baseURL,
	}
}

func (b *BaseExchangeAdapter) GenerateHMACSHA256(data []byte) string {
	mac := hmac.New(sha256.New, []byte(b.APISecret))
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}
