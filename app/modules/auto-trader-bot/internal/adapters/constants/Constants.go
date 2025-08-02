package adapters

const (
	BinanceAPIBaseURL  = "https://api1.binance.com"
	BinanceAPI2BaseURL = "https://api2.binance.com"
	BinanceAPI3BaseURL = "https://api3.binance.com"

	GET_ACCOUNT_INFO = "/api/v3/account"

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
