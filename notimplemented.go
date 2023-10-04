package terrors

import (
	"errors"
	"fmt"
)

var (
	// ErrNotImplemented is a simple error telling the caller
	// that something isn't implemented. No other data is provided.
	ErrNotImplemented = errors.New("not implemented")
)

// NotImplementedError is a more detailed variant of ErrNotImplemented,
// allowing a function to return more information via Msg.
//
// Note that calls to NotImplementedError.Error return msg prefixed with
// "not implemented: "
type NotImplementedError struct {
	Msg string
}

// Error implements the error interface.
func (e *NotImplementedError) Error() string {
	return fmt.Sprintf("not implemented: %s", e.Msg)
}

// Is implements errors.Is, allowing NotImplementedError
// to "be" ErrNotImplemented.
func (e *NotImplementedError) Is(err error) bool {
	return errors.Is(err, ErrNotImplemented)
}

// NewNotImplementedError returns a NotImplementedError
// with a given message, which can be a function name, details
// about why something isn't implemented, or anything else.
func NewNotImplementedError(msg string) error {
	return &NotImplementedError{msg}
}
