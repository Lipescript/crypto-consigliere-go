package main

import (
	"context"
	"log"
	"os"

	binance "github.com/binance/binance-connector-go"
)

func main() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_API_SECRET")

	// Verificação básica de credenciais
	if apiKey == "" || apiSecret == "" {
		log.Fatal("Credenciais não configuradas")
	}

	client := binance.NewClient(apiKey, apiSecret)

	// Exemplo: Obter saldo da conta
	balances, err := getAccountBalances(client)
	if err != nil {
		log.Fatalf("Erro ao obter saldos: %v", err)
	}

	log.Printf("Saldos: %+v", balances)
}

func getAccountBalances(client *binance.Client) ([]binance.Balance, error) {
	res, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	return res.Balances, nil
}
