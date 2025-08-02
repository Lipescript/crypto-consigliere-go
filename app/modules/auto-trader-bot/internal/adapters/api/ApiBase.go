package api

import (
	constants "auto-trader-bot/internal/adapters/constants"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

// To perform any API adapter and handle request/response errors
func (base *BaseExchangeAdapter) ApiBaseRequest(
	method string,
	endpoint string,
	params url.Values,
	body []byte,
	headers map[string]string,
) ([]byte, error) {

	fullURL := base.BaseURL + endpoint
	if params != nil {
		fullURL += "?" + params.Encode()
	}

	req, err := http.NewRequest(method, fullURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("CREATE - request failed HTTP: %w", err)
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	if (method == "POST" || method == "PUT") && len(body) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := base.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("EXECUTE - request failed HTTP: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("READ RESPONSE - request failed HTTP: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error HTTP %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func (b *BaseExchangeAdapter) GenerateHMACSHA256(data []byte) string {
	mac := hmac.New(sha256.New, []byte(b.APISecret))
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}
