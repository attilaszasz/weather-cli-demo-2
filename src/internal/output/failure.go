package output

import (
	"errors"
	"time"

	"weather-cli/src/internal/exitcode"
	"weather-cli/src/internal/provider/openmeteo"
	"weather-cli/src/internal/validation"
)

const errorStatus = "error"

// FailurePayload is the stable public failure JSON shape.
type FailurePayload struct {
	Status    string       `json:"status"`
	Timestamp string       `json:"timestamp"`
	Error     ErrorPayload `json:"error"`
}

type ErrorPayload struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Retryable bool   `json:"retryable"`
}

type FailureDescriptor struct {
	Code      string
	Message   string
	Retryable bool
	ExitCode  int
}

func DescribeFailure(err error) FailureDescriptor {
	descriptor := FailureDescriptor{
		Code:      "internal_error",
		Message:   err.Error(),
		Retryable: false,
		ExitCode:  exitcode.Internal,
	}

	var validationErr *validation.Error
	if errors.As(err, &validationErr) {
		return FailureDescriptor{
			Code:      "validation_error",
			Message:   validationErr.Message,
			Retryable: false,
			ExitCode:  exitcode.Validation,
		}
	}

	var providerErr *openmeteo.Error
	if errors.As(err, &providerErr) {
		switch providerErr.Type {
		case openmeteo.ErrorTypeTransport:
			return FailureDescriptor{
				Code:      "network_error",
				Message:   providerErr.Message,
				Retryable: true,
				ExitCode:  exitcode.Network,
			}
		case openmeteo.ErrorTypeData:
			return FailureDescriptor{
				Code:      "provider_error",
				Message:   providerErr.Message,
				Retryable: false,
				ExitCode:  exitcode.Provider,
			}
		}
	}

	return descriptor
}

func NewFailurePayload(descriptor FailureDescriptor, generatedAt time.Time) FailurePayload {
	return FailurePayload{
		Status:    errorStatus,
		Timestamp: generatedAt.UTC().Format(time.RFC3339),
		Error: ErrorPayload{
			Code:      descriptor.Code,
			Message:   descriptor.Message,
			Retryable: descriptor.Retryable,
		},
	}
}
