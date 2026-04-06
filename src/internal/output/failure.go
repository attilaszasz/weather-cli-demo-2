package output

// FailurePayload is the lightweight machine-readable failure result for E001.
type FailurePayload struct {
	ErrorType string `json:"error_type"`
	Message   string `json:"message"`
}

// NewFailurePayload builds a non-success CLI payload.
func NewFailurePayload(errorType, message string) FailurePayload {
	return FailurePayload{
		ErrorType: errorType,
		Message:   message,
	}
}
