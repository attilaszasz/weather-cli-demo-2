package output

import (
	"bytes"
	"errors"
	"os"
	"testing"
	"time"

	"weather-cli/src/internal/exitcode"
	"weather-cli/src/internal/provider/openmeteo"
	"weather-cli/src/internal/validation"
	"weather-cli/src/internal/weather"
)

func TestNewSuccessPayload(t *testing.T) {
	generatedAt := time.Date(2026, 4, 6, 9, 15, 0, 0, time.UTC)
	payload := NewSuccessPayload(
		weather.Coordinates{Latitude: 44.4268, Longitude: 26.1025},
		weather.CurrentWeather{
			Temperature:          14.2,
			WindSpeed:            7.1,
			WindDirection:        220,
			WeatherCode:          3,
			ObservationTimestamp: "2026-04-06T10:00",
		},
		generatedAt,
	)

	if payload.Location.Latitude != 44.4268 {
		t.Fatalf("unexpected latitude: %v", payload.Location.Latitude)
	}

	if payload.Current.ObservationTimestamp != "2026-04-06T10:00" {
		t.Fatalf("unexpected timestamp: %s", payload.Current.ObservationTimestamp)
	}

	if payload.Timestamp != "2026-04-06T09:15:00Z" {
		t.Fatalf("unexpected generated timestamp: %s", payload.Timestamp)
	}

	if payload.Source.Name != "open-meteo" {
		t.Fatalf("unexpected source: %s", payload.Source.Name)
	}
}

func TestWriteJSON(t *testing.T) {
	var buffer bytes.Buffer

	err := WriteJSON(&buffer, NewFailurePayload(FailureDescriptor{
		Code:      "validation_error",
		Message:   "latitude must be a valid decimal number",
		Retryable: false,
		ExitCode:  exitcode.Validation,
	}, time.Date(2026, 4, 6, 9, 15, 0, 0, time.UTC)))
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}

	expected, readErr := os.ReadFile("../../tests/testdata/contract-failure-validation.json")
	if readErr != nil {
		t.Fatalf("read validation fixture: %v", readErr)
	}

	if buffer.String() != string(expected) {
		t.Fatalf("unexpected failure JSON:\nexpected:\n%s\ngot:\n%s", string(expected), buffer.String())
	}
}

func TestDescribeFailureValidation(t *testing.T) {
	payload := DescribeFailure(&validation.Error{Type: validation.ErrorTypeValidation, Message: "latitude is required"})

	if payload.Code != "validation_error" {
		t.Fatalf("unexpected error type: %s", payload.Code)
	}

	if payload.ExitCode != exitcode.Validation {
		t.Fatalf("unexpected exit code: %d", payload.ExitCode)
	}
}

func TestDescribeFailureNetwork(t *testing.T) {
	payload := DescribeFailure(&openmeteo.Error{Type: openmeteo.ErrorTypeTransport, Message: "provider request failed"})

	if payload.Code != "network_error" {
		t.Fatalf("unexpected error type: %s", payload.Code)
	}

	if !payload.Retryable {
		t.Fatal("expected network failures to be retryable")
	}
}

func TestDescribeFailureProvider(t *testing.T) {
	payload := DescribeFailure(&openmeteo.Error{Type: openmeteo.ErrorTypeData, Message: "provider response missing current.time"})

	if payload.Code != "provider_error" {
		t.Fatalf("unexpected error type: %s", payload.Code)
	}

	if payload.ExitCode != exitcode.Provider {
		t.Fatalf("unexpected exit code: %d", payload.ExitCode)
	}
}

func TestDescribeFailureInternal(t *testing.T) {
	payload := DescribeFailure(errors.New("boom"))

	if payload.Code != "internal_error" {
		t.Fatalf("unexpected error type: %s", payload.Code)
	}

	if payload.Message == "" {
		t.Fatal("expected failure message to be populated")
	}
}
