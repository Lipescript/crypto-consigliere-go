package constants

import "testing"

func TestCurrency_String(t *testing.T) {
	cases := []struct {
		currency Currency
		expected string
	}{
		{BRL, "BRL"},
		{USD, "USD"},
		{BTC, "BTC"},
		{Unknown, "UNKNOWN"},
	}
	for _, c := range cases {
		if got := c.currency.String(); got != c.expected {
			t.Errorf("Currency(%d).String() = %q, want %q", c.currency, got, c.expected)
		}
	}
}

func TestCurrency_IsValid(t *testing.T) {
	for _, c := range []Currency{BRL, USD, EUR, GBP, JPY, CNY, BTC} {
		if !c.IsValid() {
			t.Errorf("expected %v to be valid", c)
		}
	}
	if Unknown.IsValid() {
		t.Error("expected Unknown to be invalid")
	}
}

func TestFromString(t *testing.T) {
	cases := []struct {
		input    string
		expected Currency
	}{
		{"BRL", BRL},
		{"usd", USD},
		{"₿", BTC},
		{"invalid", Unknown},
		{"", Unknown},
	}
	for _, c := range cases {
		if got := FromString(c.input); got != c.expected {
			t.Errorf("FromString(%q) = %v, want %v", c.input, got, c.expected)
		}
	}
}

func TestCurrency_Symbol(t *testing.T) {
	if BRL.Symbol() != "R$" {
		t.Errorf("expected R$, got %s", BRL.Symbol())
	}
	if Unknown.Symbol() != "" {
		t.Errorf("expected empty symbol for Unknown, got %s", Unknown.Symbol())
	}
}
