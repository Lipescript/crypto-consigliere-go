package main

import (
	binance "auto-trader-bot/internal/adapters/api/binance"
	"fmt"
	"os"
)

func main() {
	binanceAdapter := binance.NewBinanceAdapter(os.Getenv("BINANCE_API_KEY_READ"), os.Getenv("BINANCE_API_SECRET_READ"))

	binanceAdapter.GetAccountGeneralInfo()
	fmt.Printf(binanceAdapter.GetAccountFundingBalances())
	fmt.Printf(binanceAdapter.GetMarketCryptoInfo())

}
