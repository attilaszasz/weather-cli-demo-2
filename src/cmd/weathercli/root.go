package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/spf13/cobra"

	"weather-cli/src/internal/exitcode"
	"weather-cli/src/internal/output"
	"weather-cli/src/internal/provider/openmeteo"
	"weather-cli/src/internal/validation"
	"weather-cli/src/internal/weather"
)

type weatherService interface {
	GetCurrentWeather(ctx context.Context, coordinates weather.Coordinates) (weather.CurrentWeather, error)
}

var newWeatherService = func() weatherService {
	return weather.NewService(openmeteo.NewClient())
}

var now = time.Now

func newRootCommand() *cobra.Command {
	var latitudeInput string
	var longitudeInput string

	command := &cobra.Command{
		Use:   "weathercli",
		Short: "Retrieve current weather for explicit coordinates",
		Long:  "weathercli accepts one latitude and one longitude and returns current weather as machine-readable JSON.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context(), latitudeInput, longitudeInput, cmd.OutOrStdout())
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	command.Flags().StringVar(&latitudeInput, "latitude", "", "Latitude for the weather lookup")
	command.Flags().StringVar(&longitudeInput, "longitude", "", "Longitude for the weather lookup")

	return command
}

func run(ctx context.Context, latitudeInput, longitudeInput string, stdout io.Writer) error {
	coordinates, err := validation.ParseCoordinates(latitudeInput, longitudeInput)
	if err != nil {
		return writeFailure(stdout, err)
	}

	service := newWeatherService()
	currentWeather, err := service.GetCurrentWeather(ctx, coordinates)
	if err != nil {
		return writeFailure(stdout, err)
	}

	successPayload := output.NewSuccessPayload(coordinates, currentWeather, now())
	if err := output.WriteJSON(stdout, successPayload); err != nil {
		return exitcode.Wrap(exitcode.Internal, fmt.Errorf("write success payload: %w", err))
	}

	return nil
}

func writeFailure(stdout io.Writer, err error) error {
	descriptor := output.DescribeFailure(err)
	failure := output.NewFailurePayload(descriptor, now())

	if writeErr := output.WriteJSON(stdout, failure); writeErr != nil {
		return exitcode.Wrap(exitcode.Internal, fmt.Errorf("write failure payload: %w", writeErr))
	}

	return exitcode.Wrap(descriptor.ExitCode, err)
}
