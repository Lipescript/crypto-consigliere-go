package config

import (
	"errors"
	"os"
)

func GetLocalAPICredentials() (string, string, error) {
	clientID := os.Getenv("BINANCE_API_KEY_READ")
	clientSecret := os.Getenv("BINANC_API_SECRET_READ")

	if clientID == "" || clientSecret == "" {
		return "", "", errors.New("credencials not found. Verify your OS variables.")
	}

	return clientID, clientSecret, nil
}
