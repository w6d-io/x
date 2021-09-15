package errorx

import (
	"fmt"
	"github.com/pkg/errors"
)

type Interface interface {
	// Error type representing an error condition, with the nil value representing no error.
	Error() string

	// GetCause type representing an error condition, with the nil value representing no error.
	GetCause() error

	// GetMessage returns the Message
	GetMessage() string

	// GetStatusCode returns http status code.
	GetStatusCode() int

	// ShowStack prints the stack if exist
	ShowStack()
}

type Error struct {
	// Cause original error
	Cause error

	// Code for http response code
	Code int

	// Message details for this error
	Message string
}

var _ Interface = &Error{}

func (e *Error) GetMessage() string {
	return e.Message
}

func (e *Error) GetStatusCode() int {
	return e.Code
}

func (e *Error) GetCause() error {
	return errors.Cause(e.Cause)
}

func (e *Error) ShowStack() {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	if err, ok := errors.Cause(e.Cause).(stackTracer); ok {
		for _, f := range err.StackTrace() {
			fmt.Printf("%+s:%d\n", f, f)
		}
	}
}

func (e *Error) Error() string {
	if e.Cause == nil {
		return e.Message
	}
	if e.Message == "" {
		return fmt.Sprintf("%s", e.GetCause())
	}
	return fmt.Sprintf("%s: %s", e.GetMessage(), e.GetCause())
}

func New(err error, message string) error {
	return &Error{Cause: err, Message: message}
}

func NewHTTP(err error, code int, message string) error {
	return &Error{Cause: err, Code: code, Message: message}
}

func Wrap(err error, message string) error {
	return &Error{Cause: errors.Wrap(err, message)}
}

func GetError(err error) *Error {
	if err == nil {
		return nil
	}
	e, ok := err.(*Error)
	if ok {
		return e
	}
	return &Error{ Cause: err }
}