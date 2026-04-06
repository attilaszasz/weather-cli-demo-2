package output

import (
	"bytes"
	"strings"
	"testing"

	"weather-cli/src/internal/weather"
)

func TestNewSuccessPayload(t *testing.T) {
	payload := NewSuccessPayload(
		weather.Coordinates{Latitude: 44.4268, Longitude: 26.1025},
		weather.CurrentWeather{
			Temperature:          14.2,
			WindSpeed:            7.1,
			WindDirection:        220,
			WeatherCode:          3,
			ObservationTimestamp: "2026-04-06T10:00",
		},
	)

	if payload.Coordinates.Latitude != 44.4268 {
		t.Fatalf("unexpected latitude: %v", payload.Coordinates.Latitude)
	}

	if payload.Current.ObservationTimestamp != "2026-04-06T10:00" {
		t.Fatalf("unexpected timestamp: %s", payload.Current.ObservationTimestamp)
	}

	if payload.Source.Name != "open-meteo" {
		t.Fatalf("unexpected source: %s", payload.Source.Name)
	}
}

func TestWriteJSON(t *testing.T) {
	var buffer bytes.Buffer

	err := WriteJSON(&buffer, NewFailurePayload("validation_error", "latitude is required"))
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}

	output := buffer.String()
	if !strings.Contains(output, `"error_type": "validation_error"`) {
		t.Fatalf("expected error_type field, got %s", output)
	}
}

func TestNewFailurePayload(t *testing.T) {
	payload := NewFailurePayload("provider_data_error", "provider response missing current.time")

	if payload.ErrorType != "provider_data_error" {
		t.Fatalf("unexpected error type: %s", payload.ErrorType)
	}

	if payload.Message == "" {
		t.Fatal("expected failure message to be populated")
	}
}
