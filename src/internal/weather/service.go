package weather

import (
	"context"

	"weather-cli/src/internal/provider"
)

// Service orchestrates current-weather retrieval.
type Service struct {
	client provider.Client
}

// NewService builds a Service.
func NewService(client provider.Client) *Service {
	return &Service{client: client}
}

// GetCurrentWeather fetches and maps provider data into the domain model.
func (s *Service) GetCurrentWeather(ctx context.Context, coordinates Coordinates) (CurrentWeather, error) {
	response, err := s.client.FetchCurrentWeather(ctx, coordinates.Latitude, coordinates.Longitude)
	if err != nil {
		return CurrentWeather{}, err
	}

	return CurrentWeather{
		Temperature:          response.Temperature,
		WindSpeed:            response.WindSpeed,
		WindDirection:        response.WindDirection,
		WeatherCode:          response.WeatherCode,
		ObservationTimestamp: response.ObservationTimestamp,
	}, nil
}
