package constants

// HTTP constants
const (
	BinanceAPIBaseURL        = "https://api1.binance.com"
	BinanceAPI2BaseURL       = "https://api2.binance.com"
	BinanceAPI3BaseURL       = "https://api3.binance.com"
	BinanceAPITestnetBaseURL = "https://testnet.binance.vision"

	GET_ACCOUNT_INFO = "/api/v3/account"
)

// Headers & Params
const (
	XApiKeyHeader            = "X-API-KEY"
	XNonceHeader             = "X-NONCE"
	DefaultAPITimeoutSeconds = 30
	WalletTypeSpot           = "SPOT"
	WalletTypeFunding        = "FUNDING"
)

// Log constants
const (
	LogPrefix           = "BinanceAdapter"
	AccountInfoHeader   = "Account Information:"
	CanTradeFormat      = "Can Trade: %v"
	CanWithdrawFormat   = "Can Withdraw: %v"
	CanDepositFormat    = "Can Deposit: %v"
	SpotWalletHeader    = "Spot Wallet Balances:"
	AssetBalanceFormat  = "Asset: %s, Free: %s, Locked: %s"
	FundingWalletHeader = "Funding Wallet:"
	CoinBalanceFormat   = "Coin: %s, Free: %s, Frozen: %s"
	ErrorAccountInfo    = "Error getting account info: %v"
	ErrorSpotWallet     = "Error getting spot wallet: %v"
	ErrorFundingWallet  = "Error getting funding wallet: %v"
)
