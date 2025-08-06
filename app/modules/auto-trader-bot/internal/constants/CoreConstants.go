package constants

import "strings"

type Currency int

const (
	Unknown Currency = iota
	BRL
	USD
	EUR
	GBP
	JPY
	CNY
	BTC
)

func (c Currency) String() string {
	switch c {
	case BRL:
		return "BRL"
	case USD:
		return "USD"
	case EUR:
		return "EUR"
	case GBP:
		return "GBP"
	case JPY:
		return "JPY"
	case CNY:
		return "CNY"
	case BTC:
		return "BTC"
	default:
		return "UNKNOWN"
	}
}

func (c Currency) Symbol() string {
	switch c {
	case BRL:
		return "R$"
	case USD:
		return "$"
	case EUR:
		return "€"
	case GBP:
		return "£"
	case JPY:
		return "¥"
	case CNY:
		return "¥"
	case BTC:
		return "₿"
	default:
		return ""
	}
}

func (c Currency) Name() string {
	switch c {
	case BRL:
		return "Brazilian Real"
	case USD:
		return "US Dollar"
	case EUR:
		return "Euro"
	case GBP:
		return "British Pound"
	case JPY:
		return "Japanese Yen"
	case CNY:
		return "Chinese Yuan"
	case BTC:
		return "Bitcoin"
	default:
		return "Unknown"
	}
}

func FromString(s string) Currency {
	s = strings.ToUpper(strings.TrimSpace(s))
	switch s {
	case "BRL", "REAL", "R$":
		return BRL
	case "USD", "DOLLAR", "$":
		return USD
	case "EUR", "EURO", "€":
		return EUR
	case "GBP", "POUND", "£":
		return GBP
	case "JPY", "YEN", "¥":
		return JPY
	case "CNY", "YUAN":
		return CNY
	case "BTC", "BITCOIN", "₿":
		return BTC
	default:
		return Unknown
	}
}

func (c Currency) IsValid() bool {
	return c > Unknown && c <= BTC
}
