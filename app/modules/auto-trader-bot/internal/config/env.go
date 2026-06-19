package config

import (
	"fmt"
	"os"
	"runtime"
)

func GetLocalAPICredentials() (string, string, error) {
	clientID := os.Getenv("BINANCE_API_KEY_READ")
	clientSecret := os.Getenv("BINANCE_API_SECRET_READ")

	if clientID == "" || clientSecret == "" {
		return "", "", fmt.Errorf("credencials not found. Verify your '%s' system variables.", runtime.GOOS)
	}

	return clientID, clientSecret, nil
}
