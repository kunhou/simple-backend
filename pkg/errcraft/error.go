package errcraft

import "github/kunhou/simple-backend/pkg/reason"

// Error is a structured type for holding error details.
type Error struct {
	Code    int           // HTTP status code for the error.
	Reason  reason.Reason // Custom reason.
	Message string        // Error message.
	Err     error         // Original error, if any.
	Stack   string        // Stack trace, if needed.
}

// New creates a new Error with the provided HTTP status and custom error code.
func New(status int, reason reason.Reason) *Error {
	return &Error{Code: status, Reason: reason}
}

// Error implements the error interface, returning the error message.
func (e *Error) Error() string {
	return e.Message
}

// SetMessage sets a custom error message and returns the updated Error.
func (e *Error) SetMessage(message string) *Error {
	e.Message = message
	return e
}

// SetError sets the original error and returns the updated Error.
func (e *Error) SetError(err error) *Error {
	e.Err = err
	return e
}
