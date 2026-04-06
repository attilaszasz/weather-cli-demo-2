package openmeteo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"weather-cli/src/internal/provider"
)

const baseURL = "https://api.open-meteo.com"

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
func (c *Client) FetchCurrentWeather(ctx context.Context, latitude, longitude float64) (provider.CurrentWeather, error) {
	endpoint, err := url.Parse(c.baseURL)
	if err != nil {
		return provider.CurrentWeather{}, &provider.Error{Type: provider.ErrorTypeTransport, Message: "failed to build provider URL"}
	}

	endpoint.Path = "/v1/forecast"
	query := endpoint.Query()
	query.Set("latitude", fmt.Sprintf("%f", latitude))
	query.Set("longitude", fmt.Sprintf("%f", longitude))
	query.Set("current", "temperature_2m,wind_speed_10m,wind_direction_10m,weather_code")
	endpoint.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return provider.CurrentWeather{}, &provider.Error{Type: provider.ErrorTypeTransport, Message: "failed to create provider request"}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return provider.CurrentWeather{}, &provider.Error{Type: provider.ErrorTypeTransport, Message: "provider request failed"}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return provider.CurrentWeather{}, &provider.Error{
			Type:    provider.ErrorTypeTransport,
			Message: fmt.Sprintf("provider returned status %d", resp.StatusCode),
		}
	}

	var payload Response
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return provider.CurrentWeather{}, &provider.Error{Type: provider.ErrorTypeData, Message: "provider returned invalid JSON"}
	}

	if err := validateResponse(payload); err != nil {
		return provider.CurrentWeather{}, err
	}

	return provider.CurrentWeather{
		Temperature:          payload.Current.TimeSeries.Temperature2M,
		WindSpeed:            payload.Current.TimeSeries.WindSpeed10M,
		WindDirection:        payload.Current.TimeSeries.WindDirection10M,
		WeatherCode:          payload.Current.TimeSeries.WeatherCode,
		ObservationTimestamp: payload.Current.Time,
	}, nil
}

func validateResponse(payload Response) error {
	switch {
	case payload.Current.Time == "":
		return &provider.Error{Type: provider.ErrorTypeData, Message: "provider response missing current.time"}
	case payload.Latitude == 0 && payload.Longitude == 0:
		return &provider.Error{Type: provider.ErrorTypeData, Message: "provider response missing coordinates"}
	default:
		return nil
	}
}
