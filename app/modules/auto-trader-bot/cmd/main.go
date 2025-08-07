package main

import (
	binance "auto-trader-bot/internal/adapters/api/binance"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	binanceAdapter := binance.NewBinanceAdapter(os.Getenv("BINANCE_API_KEY_READ"), os.Getenv("BINANCE_API_SECRET_READ"))
	binanceAdapter.GetAccountGeneralInfo()

	accountInfo, err := binanceAdapter.GetAccountGeneralInfo()
	if err != nil {
		log.Fatalf("Error getting account info: %v", err)
	}

	// Imprimir como JSON formatado
	jsonData, err := json.MarshalIndent(accountInfo, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}

	fmt.Println(string(jsonData))

}
