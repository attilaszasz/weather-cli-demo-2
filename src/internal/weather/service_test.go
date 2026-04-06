package weather

import (
	"context"
	"errors"
	"testing"

	"weather-cli/src/internal/provider"
)

type fakeProviderClient struct {
	response provider.CurrentWeather
	err      error
}

func (f fakeProviderClient) FetchCurrentWeather(ctx context.Context, latitude, longitude float64) (provider.CurrentWeather, error) {
	if f.err != nil {
		return provider.CurrentWeather{}, f.err
	}

	return f.response, nil
}

func TestGetCurrentWeatherSuccess(t *testing.T) {
	response := provider.CurrentWeather{
		ObservationTimestamp: "2026-04-06T10:00",
		Temperature:          14.2,
		WindSpeed:            7.1,
		WindDirection:        220,
		WeatherCode:          3,
	}

	service := NewService(fakeProviderClient{response: response})
	currentWeather, err := service.GetCurrentWeather(context.Background(), Coordinates{
		Latitude:  44.4268,
		Longitude: 26.1025,
	})
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}

	if currentWeather.WeatherCode != 3 {
		t.Fatalf("unexpected weather code: %d", currentWeather.WeatherCode)
	}
}

func TestGetCurrentWeatherFailure(t *testing.T) {
	expectedErr := &provider.Error{Type: provider.ErrorTypeTransport, Message: "provider request failed"}
	service := NewService(fakeProviderClient{err: expectedErr})

	_, err := service.GetCurrentWeather(context.Background(), Coordinates{
		Latitude:  44.4268,
		Longitude: 26.1025,
	})
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected provider error, got %v", err)
	}
}
