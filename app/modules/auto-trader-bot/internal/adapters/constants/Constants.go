package adapters

const (
	BinanceAPIBaseURL          = "https://api.binance.com"
	APIVersion                 = "v3"
	ConvertGetQuoteEndpoint    = "/sapi/v1/convert/getQuote"
	ConvertAcceptQuoteEndpoint = "/sapi/v1/convert/acceptQuote"
	TradingPairsEndpoint       = "/api/v3/exchangeInfo"

	// Testnet (for developing purposes)
	BinanceAPITestnetBaseURL = "https://testnet.binance.vision"
)

const (
	XApiKeyHeader            = "X-API-KEY"
	XNonceHeader             = "X-NONCE"
	DefaultAPITimeoutSeconds = 30
	WalletTypeSpot           = "SPOT"
	WalletTypeFunding        = "FUNDING"
)
