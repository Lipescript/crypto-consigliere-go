package services

import (
	"auto-trader-bot/internal/core/domain"
	"auto-trader-bot/internal/core/ports"
	"math/big"
)

type CurrencyService struct {
	converter ports.FiatConverter
}

func NewCurrencyService(converter ports.FiatConverter) *CurrencyService {
	return &CurrencyService{converter: converter}
}

func (currencyService *CurrencyService) CalculateWalletTotalFiat(wallet *domain.Wallet) error {

	if wallet == nil {
		return nil
	}

	calculatedTotal, err := currencyService.iterateAndSumAssets(wallet.CryptoAssets, wallet.DefaultFiatCurrency)

	if err != nil {
		return err
	}

	wallet.TotalFiat = calculatedTotal
	return nil
}

func (currencyService *CurrencyService) iterateAndSumAssets(assets []domain.CryptoAsset, targetCurrency string) (*big.Int, error) {
	total := big.NewInt(0)

	for _, asset := range assets {
		if asset.FiatValue == nil || asset.FiatValue.Amount == nil {
			continue
		}

		assetValue, err := currencyService.getAssetValueInTargetCurrency(asset, targetCurrency)

		if err != nil {
			return nil, err
		}
		total = total.Add(total, assetValue)
	}
	return total, nil
}

func (currencyService *CurrencyService) getAssetValueInTargetCurrency(asset domain.CryptoAsset, targetCurrency string) (*big.Int, error) {
	assetValue := new(big.Int).Set(asset.FiatValue.Amount)

	if asset.FiatValue.Currency != targetCurrency {
		convertedValue, err := currencyService.converter.ConvertToCurrency(
			assetValue,
			asset.FiatValue.Currency,
			targetCurrency,
		)
		if err != nil {
			return nil, err
		}
		return convertedValue, nil
	}
	return assetValue, nil
}
