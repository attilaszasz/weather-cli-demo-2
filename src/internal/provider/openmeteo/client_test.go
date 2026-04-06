package openmeteo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"weather-cli/src/internal/provider"
)

func TestFetchCurrentWeatherSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("latitude") == "" || r.URL.Query().Get("longitude") == "" {
			t.Fatalf("expected latitude and longitude query params")
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
  "latitude": 44.4268,
  "longitude": 26.1025,
  "current": {
    "time": "2026-04-06T10:00",
    "temperature_2m": 14.2,
    "wind_speed_10m": 7.1,
    "wind_direction_10m": 220,
    "weather_code": 3
  }
}`))
	}))
	defer server.Close()

	client := NewClientWithBaseURL(server.URL, server.Client())
	response, err := client.FetchCurrentWeather(context.Background(), 44.4268, 26.1025)
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}

	if response.ObservationTimestamp == "" {
		t.Fatal("expected current time to be populated")
	}
}

func TestNewClientDefaults(t *testing.T) {
	client := NewClient()
	if client.baseURL != baseURL {
		t.Fatalf("unexpected base URL: %s", client.baseURL)
	}

	if client.httpClient == nil {
		t.Fatal("expected HTTP client to be configured")
	}
}

func TestProviderErrorString(t *testing.T) {
	err := &provider.Error{Type: provider.ErrorTypeData, Message: "provider response missing current.time"}
	if err.Error() != "provider response missing current.time" {
		t.Fatalf("unexpected error string: %s", err.Error())
	}
}

func TestFetchCurrentWeatherTransportFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "upstream unavailable", http.StatusBadGateway)
	}))
	server.Close()

	client := NewClientWithBaseURL(server.URL, http.DefaultClient)
	_, err := client.FetchCurrentWeather(context.Background(), 44.4268, 26.1025)
	if err == nil {
		t.Fatal("expected transport error, got nil")
	}

	providerErr, ok := err.(*provider.Error)
	if !ok {
		t.Fatalf("expected provider error, got %T", err)
	}

	if providerErr.Type != provider.ErrorTypeTransport {
		t.Fatalf("unexpected provider error type: %s", providerErr.Type)
	}
}

func TestFetchCurrentWeatherInvalidPayload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"latitude": 44.4, "longitude": 26.1, "current": {}}`))
	}))
	defer server.Close()

	client := NewClientWithBaseURL(server.URL, server.Client())
	_, err := client.FetchCurrentWeather(context.Background(), 44.4268, 26.1025)
	if err == nil {
		t.Fatal("expected provider data error, got nil")
	}

	providerErr, ok := err.(*provider.Error)
	if !ok {
		t.Fatalf("expected provider error, got %T", err)
	}

	if providerErr.Type != provider.ErrorTypeData {
		t.Fatalf("unexpected provider error type: %s", providerErr.Type)
	}
}
