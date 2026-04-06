package main

import (
	"bytes"
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"weather-cli/src/internal/exitcode"
	"weather-cli/src/internal/provider"
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
	fixedNow := time.Date(2026, 4, 6, 9, 15, 0, 0, time.UTC)
	originalNow := now
	now = func() time.Time { return fixedNow }
	t.Cleanup(func() { now = originalNow })

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

	expected, readErr := os.ReadFile("../../tests/testdata/contract-success.json")
	if readErr != nil {
		t.Fatalf("read success fixture: %v", readErr)
	}

	if stdout.String() != string(expected) {
		t.Fatalf("unexpected success output:\nexpected:\n%s\ngot:\n%s", string(expected), stdout.String())
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
	fixedNow := time.Date(2026, 4, 6, 9, 15, 0, 0, time.UTC)
	originalNow := now
	now = func() time.Time { return fixedNow }
	t.Cleanup(func() { now = originalNow })

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

	if exitcode.FromError(err) != exitcode.Validation {
		t.Fatalf("expected validation exit code, got %d", exitcode.FromError(err))
	}

	expected, readErr := os.ReadFile("../../tests/testdata/contract-failure-validation.json")
	if readErr != nil {
		t.Fatalf("read validation fixture: %v", readErr)
	}

	if stdout.String() != string(expected) {
		t.Fatalf("unexpected validation output:\nexpected:\n%s\ngot:\n%s", string(expected), stdout.String())
	}
}

func TestRunProviderFailure(t *testing.T) {
	fixedNow := time.Date(2026, 4, 6, 9, 15, 0, 0, time.UTC)
	originalNow := now
	now = func() time.Time { return fixedNow }
	t.Cleanup(func() { now = originalNow })

	service := &stubWeatherService{
		err: &provider.Error{
			Type:    provider.ErrorTypeTransport,
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

	if exitcode.FromError(err) != exitcode.Network {
		t.Fatalf("expected network exit code, got %d", exitcode.FromError(err))
	}

	expected, readErr := os.ReadFile("../../tests/testdata/contract-failure-network.json")
	if readErr != nil {
		t.Fatalf("read network fixture: %v", readErr)
	}

	if stdout.String() != string(expected) {
		t.Fatalf("unexpected network output:\nexpected:\n%s\ngot:\n%s", string(expected), stdout.String())
	}
}

func TestRunProviderDataFailure(t *testing.T) {
	fixedNow := time.Date(2026, 4, 6, 9, 15, 0, 0, time.UTC)
	originalNow := now
	now = func() time.Time { return fixedNow }
	t.Cleanup(func() { now = originalNow })

	service := &stubWeatherService{
		err: &provider.Error{
			Type:    provider.ErrorTypeData,
			Message: "provider response missing current.time",
		},
	}

	originalFactory := newWeatherService
	newWeatherService = func() weatherService { return service }
	t.Cleanup(func() { newWeatherService = originalFactory })

	var stdout bytes.Buffer
	err := run(context.Background(), "44.4268", "26.1025", &stdout)
	if err == nil {
		t.Fatal("expected provider data failure, got nil")
	}

	if exitcode.FromError(err) != exitcode.Provider {
		t.Fatalf("expected provider exit code, got %d", exitcode.FromError(err))
	}

	expected, readErr := os.ReadFile("../../tests/testdata/contract-failure-provider.json")
	if readErr != nil {
		t.Fatalf("read provider fixture: %v", readErr)
	}

	if stdout.String() != string(expected) {
		t.Fatalf("unexpected provider output:\nexpected:\n%s\ngot:\n%s", string(expected), stdout.String())
	}
}

func TestWriteFailureUnexpectedErrorUsesInternalExitCode(t *testing.T) {
	fixedNow := time.Date(2026, 4, 6, 9, 15, 0, 0, time.UTC)
	originalNow := now
	now = func() time.Time { return fixedNow }
	t.Cleanup(func() { now = originalNow })

	var stdout bytes.Buffer
	err := writeFailure(&stdout, errors.New("boom"))
	if err == nil {
		t.Fatal("expected internal failure, got nil")
	}

	if exitcode.FromError(err) != exitcode.Internal {
		t.Fatalf("expected internal exit code, got %d", exitcode.FromError(err))
	}
}
