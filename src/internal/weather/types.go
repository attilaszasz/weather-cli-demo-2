package weather

// Coordinates identifies a single requested location.
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// CurrentWeather holds the approved MVP weather fields.
type CurrentWeather struct {
	Temperature          float64
	WindSpeed            float64
	WindDirection        float64
	WeatherCode          int
	ObservationTimestamp string
}
