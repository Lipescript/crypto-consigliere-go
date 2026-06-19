package config

import (
	"os"
	"testing"
)

func TestGetLocalAPICredentials_success(t *testing.T) {
	os.Setenv("BINANCE_API_KEY_READ", "test-key")
	os.Setenv("BINANCE_API_SECRET_READ", "test-secret")
	defer os.Unsetenv("BINANCE_API_KEY_READ")
	defer os.Unsetenv("BINANCE_API_SECRET_READ")

	key, secret, err := GetLocalAPICredentials()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if key != "test-key" || secret != "test-secret" {
		t.Errorf("unexpected credentials: key=%q secret=%q", key, secret)
	}
}

func TestGetLocalAPICredentials_missingKey(t *testing.T) {
	os.Unsetenv("BINANCE_API_KEY_READ")
	os.Setenv("BINANCE_API_SECRET_READ", "test-secret")
	defer os.Unsetenv("BINANCE_API_SECRET_READ")

	_, _, err := GetLocalAPICredentials()
	if err == nil {
		t.Fatal("expected error when key is missing")
	}
}

func TestGetLocalAPICredentials_missingSecret(t *testing.T) {
	os.Setenv("BINANCE_API_KEY_READ", "test-key")
	os.Unsetenv("BINANCE_API_SECRET_READ")
	defer os.Unsetenv("BINANCE_API_KEY_READ")

	_, _, err := GetLocalAPICredentials()
	if err == nil {
		t.Fatal("expected error when secret is missing")
	}
}
