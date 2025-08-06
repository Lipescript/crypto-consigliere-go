package config

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/host"
)

func GetLocalAPICredentials() (string, string, error) {
	clientID := os.Getenv("BINANCE_API_KEY_READ")
	clientSecret := os.Getenv("BINANC_API_SECRET_READ")

	if clientID == "" || clientSecret == "" {
		hostInfo, _ := host.Info()
		userSO := hostInfo.Platform
		return "", "", fmt.Errorf("credencials not found. Verify your '%s' system variables.", userSO)
	}

	return clientID, clientSecret, nil
}
