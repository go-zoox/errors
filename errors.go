package errors

import (
	stderrors "errors"
	"fmt"
)

type err struct {
	cause   error
	message string
}

// New returns an error that formats as the given text.
func New(message string) error {
	return err{
		message: message,
	}
}

// Error implements the error interface.
func (e err) Error() string {
	if e.message == "" {
		return e.cause.Error()
	}

	if e.cause == nil {
		return e.message
	}

	return e.cause.Error() + ": " + e.message
}

// Is reports whether any error in err's chain matches target.
func (e err) Is(target error) bool {
	if e == target {
		return true
	}

	if stderrors.Is(e.cause, target) {
		return true
	}

	// @TODO
	return target.Error() == e.message
}

// Cause returns the underlying cause of the error, if possible.
func (e err) Cause() error {
	return e.cause
}

// Unwrap returns the result of calling the Unwrap method on the error, if any.
func (e err) Unwrap() error {
	return e.Cause()
}

// Wrap returns an error annotating err with message.
func Wrap(origin error, message string) error {
	return err{
		cause:   origin,
		message: message,
	}
}

// Wrapf returns an error annotating err with the format specifier.
func Wrapf(cause error, format string, args ...interface{}) error {
	return Wrap(cause, fmt.Sprintf(format, args...))
}

// Errorf formats according to a format specifier and returns the string as a
// value that satisfies error.
func Errorf(format string, args ...interface{}) error {
	return New(fmt.Sprintf(format, args...))
}
