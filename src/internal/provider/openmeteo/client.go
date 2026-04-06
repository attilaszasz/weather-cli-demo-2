package openmeteo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://api.open-meteo.com"

// ErrorType identifies provider-related failures.
type ErrorType string

const (
	ErrorTypeTransport ErrorType = "provider_transport_error"
	ErrorTypeData      ErrorType = "provider_data_error"
)

// Error is a lightweight provider failure that the CLI can serialize.
type Error struct {
	Type    ErrorType
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

// HTTPClient captures the standard HTTP client contract.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client retrieves current weather from Open-Meteo.
type Client struct {
	baseURL    string
	httpClient HTTPClient
}

// NewClient builds a client with a sensible HTTP timeout for CLI usage.
func NewClient() *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// NewClientWithBaseURL is used by tests to target a local mock server.
func NewClientWithBaseURL(baseURL string, httpClient HTTPClient) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

// FetchCurrentWeather gets the provider payload for validated coordinates.
func (c *Client) FetchCurrentWeather(ctx context.Context, latitude, longitude float64) (Response, error) {
	endpoint, err := url.Parse(c.baseURL)
	if err != nil {
		return Response{}, &Error{Type: ErrorTypeTransport, Message: "failed to build provider URL"}
	}

	endpoint.Path = "/v1/forecast"
	query := endpoint.Query()
	query.Set("latitude", fmt.Sprintf("%f", latitude))
	query.Set("longitude", fmt.Sprintf("%f", longitude))
	query.Set("current", "temperature_2m,wind_speed_10m,wind_direction_10m,weather_code")
	endpoint.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return Response{}, &Error{Type: ErrorTypeTransport, Message: "failed to create provider request"}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Response{}, &Error{Type: ErrorTypeTransport, Message: "provider request failed"}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Response{}, &Error{
			Type:    ErrorTypeTransport,
			Message: fmt.Sprintf("provider returned status %d", resp.StatusCode),
		}
	}

	var payload Response
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return Response{}, &Error{Type: ErrorTypeData, Message: "provider returned invalid JSON"}
	}

	if err := validateResponse(payload); err != nil {
		return Response{}, err
	}

	return payload, nil
}

func validateResponse(payload Response) error {
	switch {
	case payload.Current.Time == "":
		return &Error{Type: ErrorTypeData, Message: "provider response missing current.time"}
	case payload.Latitude == 0 && payload.Longitude == 0:
		return &Error{Type: ErrorTypeData, Message: "provider response missing coordinates"}
	default:
		return nil
	}
}

// AsError extracts a provider error when possible.
func AsError(err error) (*Error, bool) {
	var providerErr *Error
	if errors.As(err, &providerErr) {
		return providerErr, true
	}

	return nil, false
}
