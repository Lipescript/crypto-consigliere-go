package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	api "auto-trader-bot/internal/adapters/api"
	constants "auto-trader-bot/internal/constants"
	ports "auto-trader-bot/internal/core/ports"

	binance_connector "github.com/binance/binance-connector-go"
)

// Interfaces
var _ ports.AccountRetriever = (*BinanceAdapter)(nil)
var _ ports.MarketRetriever = (*BinanceAdapter)(nil)

type BinanceAdapter struct {
	*api.BaseExchangeAdapter
	binanceClient *binance_connector.Client
}

func NewBinanceAdapter(apiKey, apiSecret string) *BinanceAdapter {
	return &BinanceAdapter{
		BaseExchangeAdapter: api.NewApiBaseAdapter(
			apiKey,
			apiSecret,
			constants.BinanceAPIBaseURL,
		),
		binanceClient: binance_connector.NewClient(apiKey, apiSecret, constants.BinanceAPIBaseURL),
	}
}

// GET /api/v3/account
func (adapter *BinanceAdapter) GetAccountGeneralInfo() {
	accountInfo, err := adapter.binanceClient.NewGetAccountService().Do(context.Background())
	if err != nil {
		log.Printf("%s: %s", constants.LogPrefix, fmt.Sprintf(constants.ErrorAccountInfo, err))
		return
	}

	fmt.Println(constants.AccountInfoHeader)
	fmt.Printf("%s\n", fmt.Sprintf(constants.CanTradeFormat, accountInfo.CanTrade))
	fmt.Printf("%s\n", fmt.Sprintf(constants.CanWithdrawFormat, accountInfo.CanWithdraw))
	fmt.Printf("%s\n", fmt.Sprintf(constants.CanDepositFormat, accountInfo.CanDeposit))
}

// POST /sapi/v1/asset/get-funding-asset
func (adapter *BinanceAdapter) GetAccountFundingBalances() (string, error) {
	res, err := adapter.binanceClient.NewFundingWalletService().NeedBtcValuation("YES").Do(context.Background())

	if err != nil {
		return "", fmt.Errorf("erro ao obter saldos: %v", err)
	}

	prettyJSON, err := json.MarshalIndent(res, "", "  ")

	if err != nil {
		return "", fmt.Errorf("erro na formatação JSON: %v", err)
	}

	return string(prettyJSON), err
}

// GET /api/v3/ticker/24hr
func (adapter *BinanceAdapter) GetMarketCryptoInfo() (string, error) {
	//usar filtro para pesquisa condicionada .Symbols()
	res, err := adapter.binanceClient.NewTicker24hrService().Symbol("BNBBTC").Do(context.Background())

	if err != nil {
		return "", fmt.Errorf("erro ao obter informações da moeda: %v", err)
	}

	prettyJSON, err := json.MarshalIndent(res, "", "  ")

	if err != nil {
		return "", fmt.Errorf("erro na formatação JSON: %v", err)
	}

	return string(prettyJSON), err
}
