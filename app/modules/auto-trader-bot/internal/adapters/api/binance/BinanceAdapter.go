package binance

import (
	"context"
	"encoding/json"
	"fmt"

	api "auto-trader-bot/internal/adapters/api"
	constants "auto-trader-bot/internal/constants"
	domain "auto-trader-bot/internal/core/domain"
	ports "auto-trader-bot/internal/core/ports"

	binance_connector_lib "github.com/binance/binance-connector-go"
)

// Interfaces
var _ ports.AccountRetriever = (*BinanceAdapter)(nil)
var _ ports.MarketRetriever = (*BinanceAdapter)(nil)

type BinanceAdapter struct {
	*api.BaseAdapter
	binanceClient *binance_connector_lib.Client
}

func NewBinanceAdapter(apiKey, apiSecret string) *BinanceAdapter {
	return &BinanceAdapter{
		BaseAdapter: api.NewBaseAdapter(
			apiKey,
			apiSecret,
			constants.BinanceAPIBaseURL,
		),
		binanceClient: binance_connector_lib.NewClient(apiKey, apiSecret, constants.BinanceAPIBaseURL),
	}
}

// GET /api/v3/account
func (adapter *BinanceAdapter) GetAccountGeneralInfo() (domain.ExchangeAccount, error) {
	var err error

	result := domain.ExchangeAccount{}

	accountInfoRes, err := adapter.binanceClient.NewGetAccountService().Do(context.Background())

	if err != nil {
		return domain.ExchangeAccount{}, fmt.Errorf("erro ao obter informações da conta: %v", err)
	}

	result.CanTrade = accountInfoRes.CanTrade
	result.CanDeposit = accountInfoRes.CanDeposit
	result.CanWithdraw = accountInfoRes.CanWithdraw
	result.AccountType = accountInfoRes.AccountType
	result.UpdatedAt = api.ConvertTimestamp(accountInfoRes.UpdateTime)

	accountSnapshotRes, err := adapter.binanceClient.NewGetAccountSnapshotService().MarketType("SPOT").Do(context.Background())

	if err != nil {
		return domain.ExchangeAccount{}, fmt.Errorf("erro ao obter informações do snapshot da conta: %v", err)
	}

	if len(accountSnapshotRes.SnapshotVos) > 0 {
		// First (more recent) account snapshot
		result.TotalAssetOfBtc = accountSnapshotRes.SnapshotVos[0].Data.TotalAssetOfBtc
	}

	return result, nil
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
