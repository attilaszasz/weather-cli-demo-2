package weather

import (
	"context"

	"weather-cli/src/internal/provider/openmeteo"
)

// ProviderClient captures the provider dependency needed by the service.
type ProviderClient interface {
	FetchCurrentWeather(ctx context.Context, latitude, longitude float64) (openmeteo.Response, error)
}

// Service orchestrates current-weather retrieval.
type Service struct {
	client ProviderClient
}

// NewService builds a Service.
func NewService(client ProviderClient) *Service {
	return &Service{client: client}
}

// GetCurrentWeather fetches and maps provider data into the domain model.
func (s *Service) GetCurrentWeather(ctx context.Context, coordinates Coordinates) (CurrentWeather, error) {
	response, err := s.client.FetchCurrentWeather(ctx, coordinates.Latitude, coordinates.Longitude)
	if err != nil {
		return CurrentWeather{}, err
	}

	return CurrentWeather{
		Temperature:          response.Current.Temperature2M,
		WindSpeed:            response.Current.WindSpeed10M,
		WindDirection:        response.Current.WindDirection10M,
		WeatherCode:          response.Current.WeatherCode,
		ObservationTimestamp: response.Current.Time,
	}, nil
}
