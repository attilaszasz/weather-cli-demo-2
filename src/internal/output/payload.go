package output

import (
	"encoding/json"
	"io"

	"weather-cli/src/internal/weather"
)

// SuccessPayload is the approved MVP success JSON shape.
type SuccessPayload struct {
	Coordinates CoordinatesPayload `json:"coordinates"`
	Current     CurrentPayload     `json:"current"`
	Source      SourcePayload      `json:"source"`
}

type CoordinatesPayload struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CurrentPayload struct {
	Temperature          float64 `json:"temperature"`
	WindSpeed            float64 `json:"wind_speed"`
	WindDirection        float64 `json:"wind_direction"`
	WeatherCode          int     `json:"weather_code"`
	ObservationTimestamp string  `json:"observation_timestamp"`
}

type SourcePayload struct {
	Name string `json:"name"`
}

// NewSuccessPayload maps domain values into the MVP JSON response.
func NewSuccessPayload(coordinates weather.Coordinates, current weather.CurrentWeather) SuccessPayload {
	return SuccessPayload{
		Coordinates: CoordinatesPayload{
			Latitude:  coordinates.Latitude,
			Longitude: coordinates.Longitude,
		},
		Current: CurrentPayload{
			Temperature:          current.Temperature,
			WindSpeed:            current.WindSpeed,
			WindDirection:        current.WindDirection,
			WeatherCode:          current.WeatherCode,
			ObservationTimestamp: current.ObservationTimestamp,
		},
		Source: SourcePayload{
			Name: "open-meteo",
		},
	}
}

// WriteJSON serializes any payload for command output.
func WriteJSON(writer io.Writer, payload any) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(payload)
}
