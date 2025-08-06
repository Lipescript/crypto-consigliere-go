package main

import (
	constants "auto-trader-bot/internal/adapters/constants"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	binance "github.com/binance/binance-connector-go"
)

func main() {

	apiKey := os.Getenv("BINANCE_API_KEY_READ")
	apiSecret := os.Getenv("BINANCE_API_SECRET_READ")

	clientConnector := binance.NewClient(apiKey, apiSecret, constants.BinanceAPIBaseURL)

	result, err := getMarketCryptoInfo(clientConnector)
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}

	fmt.Println("Saldos formatados:")
	fmt.Println(result)

}

// POST /sapi/v1/asset/get-funding-asset
func getFundingBalances(client *binance.Client) (string, error) {
	res, err := client.NewFundingWalletService().NeedBtcValuation("YES").Do(context.Background())

	if err != nil {
		return "", fmt.Errorf("erro ao obter saldos: %v", err)
	}

	prettyJSON, err := json.MarshalIndent(res, "", "  ")

	if err != nil {
		return "", fmt.Errorf("erro na formatação JSON: %v", err)
	}

	return string(prettyJSON), nil
}

// GET /api/v3/ticker/24hr
func getMarketCryptoInfo(client *binance.Client) (string, error) {
	//usar filtro para pesquisa condicionada .Symbols()
	res, err := client.NewTicker24hrService().Symbol("BNBBTC").Do(context.Background())

	if err != nil {
		return "", fmt.Errorf("erro ao obter informações da moeda: %v", err)
	}

	prettyJSON, err := json.MarshalIndent(res, "", "  ")

	if err != nil {
		return "", fmt.Errorf("erro na formatação JSON: %v", err)
	}

	return string(prettyJSON), nil
}
