package provider

import "context"

// Client captures the provider dependency used by the weather service.
type Client interface {
	FetchCurrentWeather(ctx context.Context, latitude, longitude float64) (CurrentWeather, error)
}

// CurrentWeather is the provider-neutral result returned by adapters.
type CurrentWeather struct {
	Temperature          float64
	WindSpeed            float64
	WindDirection        float64
	WeatherCode          int
	ObservationTimestamp string
}

// ErrorType identifies provider-neutral failure categories.
type ErrorType string

const (
	ErrorTypeTransport ErrorType = "provider_transport_error"
	ErrorTypeData      ErrorType = "provider_data_error"
)

// Error is the provider-neutral failure type exposed outside concrete adapters.
type Error struct {
	Type    ErrorType
	Message string
}

func (e *Error) Error() string {
	return e.Message
}
