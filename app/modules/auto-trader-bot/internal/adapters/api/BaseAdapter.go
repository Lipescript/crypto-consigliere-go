package api

import (
	constants "auto-trader-bot/internal/constants"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"
)

type BaseAdapter struct {
	Client    *http.Client
	APIKey    string
	APISecret string
	BaseURL   string
}

func NewBaseAdapter(apiKey, apiSecret, baseURL string) *BaseAdapter {
	return &BaseAdapter{
		Client:    &http.Client{Timeout: constants.DefaultAPITimeoutSeconds * time.Second},
		APIKey:    apiKey,
		APISecret: apiSecret,
		BaseURL:   baseURL,
	}
}

// security utils
func (b *BaseAdapter) generateHMACSHA256(data []byte) string {
	mac := hmac.New(sha256.New, []byte(b.APISecret))
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}

// timestamp utils
func ConvertTimestamp(timestamp uint64) time.Time {
	if timestamp == 0 {
		return time.Now()
	}

	if timestamp > 1e10 {
		return time.Unix(int64(timestamp/1000), int64((timestamp%1000)*1e6))
	}

	return time.Unix(int64(timestamp), 0)
}

func ConvertTimestampMillis(timestamp uint64) time.Time {
	if timestamp == 0 {
		return time.Now()
	}
	return time.Unix(int64(timestamp/1000), int64((timestamp%1000)*1e6))
}

func ConvertTimestampSeconds(timestamp uint64) time.Time {
	if timestamp == 0 {
		return time.Now()
	}
	return time.Unix(int64(timestamp), 0)
}
