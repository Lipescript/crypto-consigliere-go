package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	binance "github.com/binance/binance-connector-go"
)

func main() {
	apiKey := os.Getenv("BINANCE_API_KEY_READ")
	apiSecret := os.Getenv("BINANCE_API_SECRET_READ")

	if apiKey == "" || apiSecret == "" {
		log.Fatal("Credenciais não configuradas")
	}

	client := binance.NewClient(apiKey, apiSecret)
	client.HTTPClient = &http.Client{Timeout: 10 * time.Second}

	getMyTrades(client)

}

func getAccountBalances(client *binance.Client) ([]binance.Balance, error) {
	res, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	return res.Balances, nil
}

func getAccountWalletInfo(client *binance.Client) {
	// AccountInfoService - /sapi/v1/account/apiTradingStatus
	accountInfo, err := client.NewAccountInfoService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(accountInfo))
}

func getMyTrades(client *binance.Client) {
	// AccountInfoService - /sapi/v1/account/apiTradingStatus
	accountInfo, err := client.NewGetMyTradesService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(accountInfo))
}
