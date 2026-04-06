package weather

import (
	"context"
	"errors"
	"testing"

	"weather-cli/src/internal/provider/openmeteo"
)

type fakeProviderClient struct {
	response openmeteo.Response
	err      error
}

func (f fakeProviderClient) FetchCurrentWeather(ctx context.Context, latitude, longitude float64) (openmeteo.Response, error) {
	if f.err != nil {
		return openmeteo.Response{}, f.err
	}

	return f.response, nil
}

func TestGetCurrentWeatherSuccess(t *testing.T) {
	var response openmeteo.Response
	response.Current.Time = "2026-04-06T10:00"
	response.Current.Temperature2M = 14.2
	response.Current.WindSpeed10M = 7.1
	response.Current.WindDirection10M = 220
	response.Current.WeatherCode = 3

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
	expectedErr := &openmeteo.Error{Type: openmeteo.ErrorTypeTransport, Message: "provider request failed"}
	service := NewService(fakeProviderClient{err: expectedErr})

	_, err := service.GetCurrentWeather(context.Background(), Coordinates{
		Latitude:  44.4268,
		Longitude: 26.1025,
	})
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected provider error, got %v", err)
	}
}
