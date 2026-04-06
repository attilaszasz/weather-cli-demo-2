package validation

import (
	"fmt"
	"strconv"
	"strings"

	"weather-cli/src/internal/weather"
)

// ErrorType identifies the failure category for lightweight CLI errors.
type ErrorType string

const (
	ErrorTypeValidation ErrorType = "validation_error"
)

// Error is a machine-readable lightweight command error.
type Error struct {
	Type    ErrorType
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

// ParseCoordinates validates and converts CLI inputs into domain coordinates.
func ParseCoordinates(latitudeInput, longitudeInput string) (weather.Coordinates, error) {
	lat, err := parseRequiredFloat("latitude", latitudeInput)
	if err != nil {
		return weather.Coordinates{}, err
	}

	lon, err := parseRequiredFloat("longitude", longitudeInput)
	if err != nil {
		return weather.Coordinates{}, err
	}

	if lat < -90 || lat > 90 {
		return weather.Coordinates{}, &Error{
			Type:    ErrorTypeValidation,
			Message: "latitude must be between -90 and 90",
		}
	}

	if lon < -180 || lon > 180 {
		return weather.Coordinates{}, &Error{
			Type:    ErrorTypeValidation,
			Message: "longitude must be between -180 and 180",
		}
	}

	return weather.Coordinates{
		Latitude:  lat,
		Longitude: lon,
	}, nil
}

func parseRequiredFloat(name, value string) (float64, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return 0, &Error{
			Type:    ErrorTypeValidation,
			Message: fmt.Sprintf("%s is required", name),
		}
	}

	parsed, err := strconv.ParseFloat(trimmed, 64)
	if err != nil {
		return 0, &Error{
			Type:    ErrorTypeValidation,
			Message: fmt.Sprintf("%s must be a valid decimal number", name),
		}
	}

	return parsed, nil
}
