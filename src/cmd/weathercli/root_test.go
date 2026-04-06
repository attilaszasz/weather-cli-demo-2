package main

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"weather-cli/src/internal/provider/openmeteo"
	"weather-cli/src/internal/weather"
)

type stubWeatherService struct {
	currentWeather weather.CurrentWeather
	err            error
	called         bool
}

func (s *stubWeatherService) GetCurrentWeather(ctx context.Context, coordinates weather.Coordinates) (weather.CurrentWeather, error) {
	s.called = true
	if s.err != nil {
		return weather.CurrentWeather{}, s.err
	}

	return s.currentWeather, nil
}

func TestRunSuccess(t *testing.T) {
	service := &stubWeatherService{
		currentWeather: weather.CurrentWeather{
			Temperature:          14.2,
			WindSpeed:            7.1,
			WindDirection:        220,
			WeatherCode:          3,
			ObservationTimestamp: "2026-04-06T10:00",
		},
	}

	originalFactory := newWeatherService
	newWeatherService = func() weatherService { return service }
	t.Cleanup(func() { newWeatherService = originalFactory })

	var stdout bytes.Buffer
	err := run(context.Background(), "44.4268", "26.1025", &stdout)
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}

	if !service.called {
		t.Fatal("expected weather service to be called")
	}

	output := stdout.String()
	for _, expected := range []string{
		`"latitude": 44.4268`,
		`"temperature": 14.2`,
		`"weather_code": 3`,
		`"name": "open-meteo"`,
	} {
		if !strings.Contains(output, expected) {
			t.Fatalf("expected output to contain %q, got %s", expected, output)
		}
	}
}

func TestNewRootCommand(t *testing.T) {
	command := newRootCommand()

	if command.Use != "weathercli" {
		t.Fatalf("unexpected command use: %s", command.Use)
	}

	if command.Flags().Lookup("latitude") == nil {
		t.Fatal("expected latitude flag to be registered")
	}

	if command.Flags().Lookup("longitude") == nil {
		t.Fatal("expected longitude flag to be registered")
	}
}

func TestRunValidationFailureSkipsService(t *testing.T) {
	service := &stubWeatherService{}

	originalFactory := newWeatherService
	newWeatherService = func() weatherService { return service }
	t.Cleanup(func() { newWeatherService = originalFactory })

	var stdout bytes.Buffer
	err := run(context.Background(), "north", "26.1025", &stdout)
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	if service.called {
		t.Fatal("expected validation failure to skip service call")
	}

	if !strings.Contains(stdout.String(), `"error_type": "validation_error"`) {
		t.Fatalf("expected validation error payload, got %s", stdout.String())
	}
}

func TestRunProviderFailure(t *testing.T) {
	service := &stubWeatherService{
		err: &openmeteo.Error{
			Type:    openmeteo.ErrorTypeTransport,
			Message: "provider request failed",
		},
	}

	originalFactory := newWeatherService
	newWeatherService = func() weatherService { return service }
	t.Cleanup(func() { newWeatherService = originalFactory })

	var stdout bytes.Buffer
	err := run(context.Background(), "44.4268", "26.1025", &stdout)
	if err == nil {
		t.Fatal("expected provider failure, got nil")
	}

	if !strings.Contains(stdout.String(), `"error_type": "provider_transport_error"`) {
		t.Fatalf("expected provider error payload, got %s", stdout.String())
	}
}
