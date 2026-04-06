package validation

import "testing"

func TestParseCoordinatesSuccess(t *testing.T) {
	coordinates, err := ParseCoordinates("44.4268", "26.1025")
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}

	if coordinates.Latitude != 44.4268 {
		t.Fatalf("unexpected latitude: %v", coordinates.Latitude)
	}

	if coordinates.Longitude != 26.1025 {
		t.Fatalf("unexpected longitude: %v", coordinates.Longitude)
	}
}

func TestParseCoordinatesFailures(t *testing.T) {
	testCases := []struct {
		name      string
		latitude  string
		longitude string
		message   string
	}{
		{name: "missing latitude", latitude: "", longitude: "10", message: "latitude is required"},
		{name: "missing longitude", latitude: "10", longitude: "", message: "longitude is required"},
		{name: "malformed latitude", latitude: "north", longitude: "10", message: "latitude must be a valid decimal number"},
		{name: "latitude out of range", latitude: "91", longitude: "10", message: "latitude must be between -90 and 90"},
		{name: "longitude out of range", latitude: "10", longitude: "181", message: "longitude must be between -180 and 180"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := ParseCoordinates(testCase.latitude, testCase.longitude)
			if err == nil {
				t.Fatal("expected validation error, got nil")
			}

			if err.Error() != testCase.message {
				t.Fatalf("unexpected error message: %v", err)
			}
		})
	}
}
