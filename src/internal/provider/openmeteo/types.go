package openmeteo

// Response is the subset of Open-Meteo data required for the MVP flow.
type Response struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Current   struct {
		Time             string  `json:"time"`
		Temperature2M    float64 `json:"temperature_2m"`
		WindSpeed10M     float64 `json:"wind_speed_10m"`
		WindDirection10M float64 `json:"wind_direction_10m"`
		WeatherCode      int     `json:"weather_code"`
	} `json:"current"`
}
