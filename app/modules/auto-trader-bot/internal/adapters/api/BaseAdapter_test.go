package api

import (
	"testing"
	"time"
)

func TestConvertTimestamp_zero_returnsZeroTime(t *testing.T) {
	result := ConvertTimestamp(0)
	if !result.IsZero() {
		t.Errorf("expected zero time, got %v", result)
	}
}

func TestConvertTimestamp_seconds(t *testing.T) {
	ts := uint64(1_700_000_000)
	result := ConvertTimestamp(ts)
	expected := time.Unix(int64(ts), 0)
	if !result.Equal(expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestConvertTimestamp_millis(t *testing.T) {
	tsMillis := uint64(1_700_000_000_000)
	result := ConvertTimestamp(tsMillis)
	expected := time.Unix(1_700_000_000, 0)
	if !result.Equal(expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestConvertTimestampMillis_zero_returnsZeroTime(t *testing.T) {
	result := ConvertTimestampMillis(0)
	if !result.IsZero() {
		t.Errorf("expected zero time, got %v", result)
	}
}

func TestConvertTimestampSeconds_zero_returnsZeroTime(t *testing.T) {
	result := ConvertTimestampSeconds(0)
	if !result.IsZero() {
		t.Errorf("expected zero time, got %v", result)
	}
}
