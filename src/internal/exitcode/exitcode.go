package exitcode

import "errors"

const (
	Success    = 0
	Validation = 2
	Network    = 3
	Provider   = 4
	Internal   = 5
)

type CodedError interface {
	error
	ExitCode() int
}

type codedError struct {
	code int
	err  error
}

func (e *codedError) Error() string {
	return e.err.Error()
}

func (e *codedError) Unwrap() error {
	return e.err
}

func (e *codedError) ExitCode() int {
	return e.code
}

func Wrap(code int, err error) error {
	if err == nil {
		return nil
	}

	return &codedError{code: code, err: err}
}

func FromError(err error) int {
	if err == nil {
		return Success
	}

	var coded CodedError
	if errors.As(err, &coded) {
		return coded.ExitCode()
	}

	return Internal
}
