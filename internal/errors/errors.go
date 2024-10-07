// Package errors contains helper functions for wrapping errors with stack traces, stack output, and panic recovery.
package errors

import (
	"fmt"

	goerrors "github.com/go-errors/errors"
)

// New creates a new instance of Error.
// If the given value does not contain an stack trace, it will be created.
func New(val any) error {
	if val == nil {
		return nil
	}

	skip := 2

	return newWithSkip(skip, val)
}

// Errorf creates a new error with the given format and values.
// It can be used as a drop-in replacement for fmt.Errorf() to provide descriptive errors in return values.
// If none of the given values contains an stack trace, it will be created.
func Errorf(format string, vals ...any) error {
	skip := 2

	return errorfWithSkip(skip, format, vals...)
}

func newWithSkip(skip int, val any) error {
	if err, ok := val.(error); ok && ContainsStackTrace(err) {
		return fmt.Errorf("%w", err)
	}

	return goerrors.Wrap(val, skip)
}

func errorfWithSkip(skip int, format string, vals ...any) error {
	err := fmt.Errorf(format, vals...)

	for _, val := range vals {
		if val, ok := val.(error); ok && val != nil && ContainsStackTrace(val) {
			return err
		}
	}

	return goerrors.Wrap(err, skip)
}
